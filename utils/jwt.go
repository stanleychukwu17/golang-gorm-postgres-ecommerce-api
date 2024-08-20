package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// Your function to sign JWT
// "github.com/dgrijalva/jwt-go/v4"
func SignJWT(payload map[string]interface{}, days int) (string, error) {
	// ###--START--
	// decode my private key
	// pemData := os.Getenv("PRIVATE_KEY")
	// fmt.Println(pemData)
	privateKey := `-----BEGIN RSA PRIVATE KEY-----
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
f/Z6bT7R6eFnCkjRXcc=
-----END RSA PRIVATE KEY-----
`
	//###--END--

	// Convert the duration string "7d" to time.Duration
	expiresIn := time.Duration(days*24) * time.Hour

	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":  time.Now().Add(expiresIn).Unix(),
		"data": payload,
	})

	// Sign the token with the private key
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}
	return signedToken, nil
}
