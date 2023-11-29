package main

import (
    "fmt"
    "net/http"
    "time"
    //"github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v5"
	"encoding/json"
	"strings"
)

// Secret key for signing and verifying JWT tokens
var secretKey = []byte("your-secret-key")

// CustomClaims represents the JWT claims you can include
type CustomClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func main() {
    http.HandleFunc("/login", login)
    http.HandleFunc("/api", withJWTAuthentication(api))

    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
    // Simulate user authentication, e.g., by checking a username and password
    username := "exampleUser"
    password := "password"

    // Check username and password (authentication logic)

    if username != "exampleUser" || password != "password" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Create an access token and an ID token
    accessToken, err := createToken(username, time.Hour*24)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    idToken, err := createToken(username, time.Hour*1) // ID token expires in 1 hour
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Send the access token and ID token as a response
    tokens := map[string]string{
        "access_token": accessToken,
        "id_token":     idToken,
    }

    // Convert the tokens to JSON
    jsonTokens, err := json.Marshal(tokens)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonTokens)
}

func createToken(username string, duration time.Duration) (string, error) {
    claims := CustomClaims{
        username,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(duration).Unix(),
            Issuer:    "your-app-name",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}



func withJWTAuthentication(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
		fmt.Println("Authorization: "+tokenString)
        if tokenString == "" {
            http.Error(w, "Unauthorized 0x0", http.StatusUnauthorized)
            return
        }

        claims := &CustomClaims{}

		bearerToken := strings.Split(tokenString, "token ") // This seems to be correct
		tokenExtracted := strings.Replace(bearerToken[1], "\n", "", -1) // Tried this as a last resort
		fmt.Println("TokenExtracted: "+tokenExtracted) // This does in fact produce the token string as I would expect

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })

        if err != nil {
			fmt.Println(err)
            http.Error(w, "Unauthorized 0x1", http.StatusUnauthorized)
        	return
        }

        if !token.Valid {
            http.Error(w, "Unauthorized 0x2", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}

func api(w http.ResponseWriter, r *http.Request) {
    // This is the protected API endpoint
    username := r.Context().Value("username").(string)
    fmt.Fprintf(w, "Hello, %s! This is a protected API endpoint.", username)
}
