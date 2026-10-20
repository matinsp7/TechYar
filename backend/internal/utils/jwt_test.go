package utility

import (
	"strings"
	"testing"
)

func TestNewJWTManagerRejectsShortSecret(t *testing.T) {
	_, err := NewJWTManager("short")

	if err == nil {
		t.Fatal("expected an error for a short JWT secret")
	}
}

func TestJWTManagerGenerateAndParseToken(t *testing.T) {
	manager, err := NewJWTManager(strings.Repeat("a", 32))
	if err != nil {
		t.Fatalf("create JWT manager: %v", err)
	}

	tokenString, err := manager.GenerateToken(
		42,
		"09120000000",
		"Matin",
		"Parian",
	)
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}

	claims, err := manager.ParseToken(tokenString)
	if err != nil {
		t.Fatalf("parse token: %v", err)
	}

	if claims.UserID != 42 {
		t.Errorf("expected user ID 42, got %d", claims.UserID)
	}

	if claims.Phone != "09120000000" {
		t.Errorf("unexpected phone: %s", claims.Phone)
	}

	if claims.Name != "Matin" {
		t.Errorf("unexpected name: %s", claims.Name)
	}

	if claims.FamilyName != "Parian" {
		t.Errorf("unexpected family name: %s", claims.FamilyName)
	}
}

func TestJWTManagerRejectsTokenSignedWithDifferentSecret(t *testing.T) {
	signer, err := NewJWTManager(strings.Repeat("a", 32))
	if err != nil {
		t.Fatalf("create signer: %v", err)
	}

	verifier, err := NewJWTManager(strings.Repeat("b", 32))
	if err != nil {
		t.Fatalf("create verifier: %v", err)
	}

	tokenString, err := signer.GenerateToken(
		42,
		"09120000000",
		"Matin",
		"Parian",
	)
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}

	_, err = verifier.ParseToken(tokenString)
	if err == nil {
		t.Fatal("expected token signed with a different secret to be rejected")
	}
}
