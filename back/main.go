package main

import (
	"context"
	"log"
	"net/http"
	"os"

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
	"back/jwtMiddleware"
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
	allowFrontURL string
	provider     *oidc.Provider
	verifier     *oidc.IDTokenVerifier
	oauth2Config oauth2.Config
)

// 初期化処理
func init() {

	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

	env := os.Getenv("GO_ENV") // 環境変数 GO_ENV を取得
	if env == "" {
		env = "development" // デフォルトは開発環境
	}

	// 環境ごとの .env ファイルをロード
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No %s file found, using default .env", envFile)
		godotenv.Load()
	}
	
	// 環境変数から値を取得
	clientID = os.Getenv("COGNITO_CLIENT_ID")
	clientSecret = os.Getenv("COGNITO_CLIENT_SECRET")
	redirectURL = os.Getenv("COGNITO_REDIRECT_URL")
	issuerURL = os.Getenv("COGNITO_ISSUER_URL")
	allowFrontURL = os.Getenv("ALLOW_FRONT_URL")

	// 必須環境変数のチェック
	if clientID == "" || clientSecret == "" || redirectURL == "" || issuerURL == "" || allowFrontURL == ""{
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


func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	db := dynamodb.New(sess)
	
	var userRepo service.UserRepositoryInterface = repository.NewUserRepository(db, "junbanmachi-table")
	var userService service.UserServiceInterface = service.NewUserService(userRepo) 
	var userHandler handler.UserHandlerInterface = handler.NewUserHandler(userService) 
	var authHandler = handler.NewAuthHandler(&oauth2Config, verifier)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{allowFrontURL}, // フロントエンドのURL
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // 必要なヘッダーを指定
		AllowCredentials: true, // Cookieを含める場合はtrue
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "hello!")
	})

	api := e.Group("/api")

	api.GET("/login", authHandler.HandleLogin)
	api.GET("/callback", authHandler.HandleCallback)
	api.GET("/logout", authHandler.HandleLogout)
	api.GET("/auth-status", authHandler.HandleAuthStatus)

	api.POST("/users", userHandler.CreateUser, jwtMiddleware.JwtMiddleware(verifier))
	api.DELETE("/users", userHandler.DeleteUser, jwtMiddleware.JwtMiddleware(verifier))
	api.PUT("/users", userHandler.UpdateUserWaitStatus, jwtMiddleware.JwtMiddleware(verifier))

	api.GET("/users", userHandler.GetAllUsers)
	api.GET("/wait-time", userHandler.GetEstimatedWaitTime)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := ":" + port
	
	log.Fatal(e.Start(address))
}
