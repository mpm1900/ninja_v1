package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"

	"ninja_v1/internal/db"
)

const COOKIE_NAME = "session_id"
const SESSION_DURATION = 24 * time.Hour

func HashPassword(password string) (string, string, error) {
	salt := uuid.New().String()
	salted := fmt.Sprintf("%s$%s", password, salt)
	hashed, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashed), salt, nil
}

func CheckPasswords(a, b, salt string) error {
	salted := fmt.Sprintf("%s$%s", a, salt)
	return bcrypt.CompareHashAndPassword([]byte(b), []byte(salted))
}

func WithSession(next http.HandlerFunc, queries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(COOKIE_NAME)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		sessionID, err := uuid.Parse(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user, err := queries.GetUserBySessionID(r.Context(), sessionID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), "user", user)
		next(w, r.Clone(ctxWithUser))
	}
}

func CreateSession(ctx context.Context, queries *db.Queries, userID uuid.UUID) (*http.Cookie, error) {
	expiresAt := time.Now().Add(SESSION_DURATION)
	session, err := queries.CreateSession(ctx, db.CreateSessionParams{
		UserID: userID,
		ExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    session.ID.String(),
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}, nil
}
