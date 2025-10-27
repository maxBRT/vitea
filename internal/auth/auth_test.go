package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	password := "test"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
		return
	}
	if hash == "" {
		t.Errorf("HashPassword() hash is empty")
		return
	}
	err = CheckPassword(password, hash)
	if err != nil {
		t.Errorf("CheckPassword() error = %v", err)
		return
	}
}

func TestGenerateAndValidateToken(t *testing.T) {
	userID := uuid.New()
	token, err := GenerateToken(userID, "test", time.Hour)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}
	if token == "" {
		t.Errorf("GenerateToken() token is empty")
		return
	}
	id, err := ValidateJWT(token, "test")
	if err != nil {
		t.Errorf("ValidateJWT() error = %v", err)
		return
	}
	if id != userID {
		t.Errorf("ValidateJWT() user ID mismatch")
	}
	_, err = ValidateJWT(token, "wrong")
	if err == nil {
		t.Errorf("ValidateJWT() error = %v", err)
	}
}

func TestGetBearerToken(t *testing.T) {
	t.Run("NoAuthHeader", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetBearerToken(headers)
		if err == nil {
			t.Errorf("GetBearerToken() error = %v", err)
			return
		}
	})

	t.Run("MalformedAuthHeader", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"Bearer"},
		}
		_, err := GetBearerToken(headers)
		if err == nil {
			t.Errorf("GetBearerToken() error = %v", err)
			return
		}

	})

	t.Run("ValidAuthHeader", func(t *testing.T) {
		headers := http.Header{
			"Authorization": []string{"Bearer test"},
		}
		token, err := GetBearerToken(headers)
		if err != nil {
			t.Errorf("GetBearerToken() error = %v", err)
			return
		}
		if token != "test" {
			t.Errorf("GetBearerToken() token mismatch")
			return
		}
	})
}
