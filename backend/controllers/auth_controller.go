package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"Rp/backend/config"

	"github.com/golang-jwt/jwt"
	"google.golang.org/api/idtoken"
)

// GoogleUser represents the user data we get from Google
type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

// Claims structure for JWT
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte("123@dummy") // Change this to a secure secret key

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		SendErrorResponse(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		SendErrorResponse(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Verify Google token and get user info
	googleUser, err := verifyGoogleToken(requestData.Token)
	if err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to verify Google token: %v", err), http.StatusUnauthorized)
		return
	}

	// Check if this is a valid student email (you can add your domain validation here)
	if !strings.HasSuffix(googleUser.Email, "@gmail.com") {
		SendErrorResponse(w, "Only university email addresses are allowed", http.StatusForbidden)
		return
	}

	// Check if student exists in database or create new record
	studentID := strings.Split(googleUser.Email, "@")[0] // Use email prefix as student ID
	err = createOrUpdateStudent(studentID, googleUser.Email, googleUser.Name)
	if err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to process student record: %v", err), http.StatusInternalServerError)
		return
	}

	// Create JWT token
	token, err := createJWTToken(googleUser.Email)
	if err != nil {
		SendErrorResponse(w, fmt.Sprintf("Failed to create session token: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"token":   token,
		"user": map[string]string{
			"id":    studentID, // Make sure this is the correct student ID
			"email": googleUser.Email,
			"name":  googleUser.Name,
		},
	})
}

func verifyGoogleToken(token string) (*GoogleUser, error) {
	ctx := context.Background()

	// Validate the ID token
	payload, err := idtoken.Validate(ctx, token, "383142498918-a9g41rq6stqn8284nf9d9lc5lh67ddcg.apps.googleusercontent.com")
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Extract user information
	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("email claim not found in token")
	}

	name, _ := payload.Claims["name"].(string) // Name is optional

	return &GoogleUser{
		ID:            payload.Subject,
		Email:         email,
		VerifiedEmail: true,
		Name:          name,
		GivenName:     payload.Claims["given_name"].(string),
		FamilyName:    payload.Claims["family_name"].(string),
		Picture:       payload.Claims["picture"].(string),
	}, nil
}
func createOrUpdateStudent(studentID, email, name string) error {
	// Check if student exists
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM student_rewards WHERE student_id = ?", studentID).Scan(&count)
	if err != nil {
		return fmt.Errorf("database error checking student: %v", err)
	}

	if count == 0 {
		// Create new student
		_, err = config.DB.Exec(`
            INSERT INTO student_rewards (student_id, student_email, student_name, points, total_marks, points_remaining)
            VALUES (?, ?, ?, 0, 0, 0)
        `, studentID, email, name)
		if err != nil {
			return fmt.Errorf("database error creating student: %v", err)
		}
	} else {
		// Update existing student's email/name if needed
		_, err = config.DB.Exec(`
            UPDATE student_rewards 
            SET student_email = ?, student_name = ?
            WHERE student_id = ?
        `, email, name, studentID)
		if err != nil {
			return fmt.Errorf("database error updating student: %v", err)
		}
	}

	return nil
}

func createJWTToken(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			SendErrorResponse(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			SendErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed with the request
		next.ServeHTTP(w, r)
	})
}
