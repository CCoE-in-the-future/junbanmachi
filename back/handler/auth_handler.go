package handler

import (
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	OAuth2Config *oauth2.Config
	Verifier     *oidc.IDTokenVerifier
	AllowFrontURL string
}

func NewAuthHandler(oauth2Config *oauth2.Config, verifier *oidc.IDTokenVerifier, allowFrontURL string) *AuthHandler {
	return &AuthHandler{
		OAuth2Config: oauth2Config,
		Verifier:     verifier,
		AllowFrontURL: allowFrontURL,
	}
}

func (h *AuthHandler) HandleLogin(c echo.Context) error {
	state := "random_state"
	url := h.OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusFound, url)
}

func (h *AuthHandler) HandleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing code parameter"})
	}

	redirectURI := c.QueryParam("redirect_uri")
	if redirectURI == "" {
		redirectURI = h.AllowFrontURL + "/admin"
	}

	token, err := h.OAuth2Config.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to exchange token", "details": err.Error()})
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No id_token in token response"})
	}

	idToken, err := h.Verifier.Verify(c.Request().Context(), rawIDToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid ID token"})
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse claims"})
	}

	cookie := &http.Cookie{
		Name:     "id_token",
		Value:    rawIDToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(c.Response().Writer, cookie)

	return c.Redirect(http.StatusFound, redirectURI)
}

func (h *AuthHandler) HandleLogout(c echo.Context) error {
	redirectURI := c.QueryParam("redirect_uri")
	if redirectURI == "" {
		redirectURI = "http://localhost:3000"
	}
	return c.Redirect(http.StatusFound, redirectURI)
}

func (h *AuthHandler) HandleAuthStatus(c echo.Context) error {
	cookie, err := c.Cookie("id_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Missing id_token cookie"})
	}

	idToken, err := h.Verifier.Verify(c.Request().Context(), cookie.Value)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token", "details": err.Error()})
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse claims"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "authenticated",
		"claims": claims,
	})
}
