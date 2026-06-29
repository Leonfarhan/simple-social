package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Leonfarhan/simple-social/internal/store"
)

const authTokenTTL = 24 * time.Hour

type authTokenClaims struct {
	Subject   int64 `json:"sub"`
	ExpiresAt int64 `json:"exp"`
	IssuedAt  int64 `json:"iat"`
}

func (app *application) AuthTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			app.unauthorizedErrorResponse(w, r, errors.New("authorization header is missing"))
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedErrorResponse(w, r, errors.New("authorization header is malformed"))
			return
		}

		userID, err := app.parseAuthToken(parts[1])
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		user, err := app.store.Users.GetByID(r.Context(), userID)
		if err != nil {
			app.unauthorizedErrorResponse(w, r, err)
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) getUser(ctx context.Context, userID int64) (*store.User, error) {
	return app.store.Users.GetByID(ctx, userID)
}

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("unauthorized", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *application) generateAuthToken(userID int64) (string, error) {
	now := time.Now()
	claims := authTokenClaims{
		Subject:   userID,
		ExpiresAt: now.Add(authTokenTTL).Unix(),
		IssuedAt:  now.Unix(),
	}

	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodedPayload := base64.RawURLEncoding.EncodeToString(payload)
	signature := app.signAuthToken(encodedPayload)

	return encodedPayload + "." + signature, nil
}

func (app *application) parseAuthToken(token string) (int64, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return 0, errors.New("token is malformed")
	}

	expectedSignature := app.signAuthToken(parts[0])
	if !hmac.Equal([]byte(parts[1]), []byte(expectedSignature)) {
		return 0, errors.New("token signature is invalid")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return 0, err
	}

	var claims authTokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return 0, err
	}

	if claims.Subject <= 0 {
		return 0, errors.New("token subject is invalid")
	}

	if claims.ExpiresAt <= time.Now().Unix() {
		return 0, errors.New("token has expired")
	}

	return claims.Subject, nil
}

func (app *application) signAuthToken(payload string) string {
	mac := hmac.New(sha256.New, []byte(app.authTokenSecret()))
	_, _ = mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func (app *application) authTokenSecret() string {
	if secret := os.Getenv("AUTH_TOKEN_SECRET"); secret != "" {
		return secret
	}

	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		return secret
	}

	return fmt.Sprintf("%s:%s", app.config.addr, app.config.env)
}
