package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"

	"github.com/joho/godotenv"

	"back/handler"
	"back/repository"
	"back/service"
)

type ClaimsPage struct {
    AccessToken string
    Claims      jwt.MapClaims
}

var (
	clientID     string
	clientSecret string
	redirectURL  string
	issuerURL    string
	provider     *oidc.Provider
	verifier     *oidc.IDTokenVerifier
	oauth2Config oauth2.Config
)

// 初期化処理
func init() {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

	// 環境変数から値を取得
	clientID = os.Getenv("COGNITO_CLIENT_ID")
	clientSecret = os.Getenv("COGNITO_CLIENT_SECRET")
	redirectURL = os.Getenv("COGNITO_REDIRECT_URL")
	issuerURL = os.Getenv("COGNITO_ISSUER_URL")

	// 必須環境変数のチェック
	if clientID == "" || clientSecret == "" || redirectURL == "" || issuerURL == "" {
		log.Fatalf("One or more required environment variables are missing")
	}
	
	var err error

	provider, err = oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		log.Fatalf("Failed to create OIDC provider: %v", err)
	}

	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "openid"},
	}
	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
}

// /login: Cognito 認証ページにリダイレクト
func handleLogin(c echo.Context) error {
	state := "random_state"
	url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

// /callback: 認証後の処理
func handleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing code parameter"})
	}

	// トークン取得
	token, err := oauth2Config.Exchange(c.Request().Context(), code) // echo.Contextを使う
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to exchange token", "details": err.Error()})
	}

	// ID トークンの検証
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No id_token in token response"})
	}
	idToken, err := verifier.Verify(c.Request().Context(), rawIDToken) // echo.Contextを使う
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid ID token"})
	}

	// クレームを取得
	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse claims"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken": token.AccessToken,
		"idToken":     rawIDToken,
		"claims":      claims,
	})
}

// /logout: ログアウト処理
func handleLogout(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/")
}

// JWT ミドルウェア
func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing Authorization header"})
		}

		// Bearer トークンを抽出
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid Authorization header format"})
		}

		// トークンの検証
		idToken, err := verifier.Verify(c.Request().Context(), tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token", "details": err.Error()})
		}

		// クレームを取得
		var claims map[string]interface{}
		if err := idToken.Claims(&claims); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse claims"})
		}

		// コンテキストにユーザー情報をセット
		c.Set("user", claims)
		return next(c)
	}
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	db := dynamodb.New(sess)
	
	var userRepo service.UserRepositoryInterface = repository.NewUserRepository(db, "junbanmachi-table")
	var userService service.UserServiceInterface = service.NewUserService(userRepo) 
	var userHandler handler.UserHandlerInterface = handler.NewUserHandler(userService) 

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "hello!")
	})

	e.GET("/login", handleLogin)
	e.GET("/callback", handleCallback)
	e.GET("/logout", handleLogout)

	api := e.Group("/api")

	api.POST("/users", userHandler.CreateUser, jwtMiddleware)
	api.DELETE("/users", userHandler.DeleteUser, jwtMiddleware)
	api.PUT("/users", userHandler.UpdateUserWaitStatus, jwtMiddleware)

	api.GET("/users", userHandler.GetAllUsers)
	api.GET("/wait-time", userHandler.GetEstimatedWaitTime)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := ":" + port
	
	log.Fatal(e.Start(address))
}
