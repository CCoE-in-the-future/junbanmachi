package jwtMiddleware

import (
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware creates a middleware function for verifying JWT tokens
func JwtMiddleware(verifier *oidc.IDTokenVerifier) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Cookie からトークンを取得
			cookie, err := c.Cookie("id_token")
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing id_token cookie"})
			}

			// トークンの検証
			idToken, err := verifier.Verify(c.Request().Context(), cookie.Value)
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
}
