package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Auth middleware handles authentication using environment variables
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for health endpoint
		if r.URL.Path == "/health" || r.URL.Path == "/" {
			next.ServeHTTP(w, r)
			return
		}

		// Get auth method from environment
		authMethod := os.Getenv("PINGME_AUTH_METHOD")

		var authenticated bool
		switch authMethod {
		case "apikey":
			authenticated = validateAPIKey(r)
		case "hmac":
			authenticated = validateHMAC(r)
		case "basic":
			authenticated = validateBasicAuth(r)
		default:
			authenticated = true // no auth
		}

		if !authenticated {
			log.Printf("Authentication failed for %s from %s", r.URL.Path, r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// validateAPIKey checks if the request has a valid API key
// Set PINGME_API_KEYS="key1,key2,key3" (comma-separated)
func validateAPIKey(r *http.Request) bool {
	// Check Authorization header: "Bearer <api_key>"
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}

	providedKey := parts[1]

	// Get valid keys from environment
	validKeysStr := os.Getenv("PINGME_API_KEYS")
	if validKeysStr == "" {
		log.Println("PINGME_API_KEYS not set")
		return false
	}

	validKeys := strings.Split(validKeysStr, ",")
	for _, validKey := range validKeys {
		if strings.TrimSpace(validKey) == providedKey {
			return true
		}
	}

	return false
}

// validateHMAC validates HMAC-SHA256 signature
// Expects X-Signature header with hex-encoded HMAC
// Set PINGME_HMAC_SECRET="your-secret"
func validateHMAC(r *http.Request) bool {
	secret := os.Getenv("PINGME_HMAC_SECRET")
	if secret == "" {
		log.Println("PINGME_HMAC_SECRET not set")
		return false
	}

	// Get signature from header
	providedSignature := r.Header.Get("X-Signature")
	if providedSignature == "" {
		return false
	}

	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	// Important: Restore body for next handler
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	// Calculate expected signature
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	// Compare signatures
	return hmac.Equal([]byte(providedSignature), []byte(expectedSignature))
}

// validateBasicAuth validates HTTP Basic Authentication
// Set PINGME_BASIC_USER="username" and PINGME_BASIC_PASS="password"
func validateBasicAuth(r *http.Request) bool {
	expectedUser := os.Getenv("PINGME_BASIC_USER")
	expectedPass := os.Getenv("PINGME_BASIC_PASS")

	if expectedUser == "" || expectedPass == "" {
		log.Println("PINGME_BASIC_USER or PINGME_BASIC_PASS not set")
		return false
	}

	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return user == expectedUser && pass == expectedPass
}
