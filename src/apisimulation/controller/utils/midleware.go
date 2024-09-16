package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret key untuk JWT
var jwtSecret = []byte("your_secret_key") // Ganti dengan secret key yang aman

// Fungsi untuk membuat token JWT
func GenerateToken(phone string) (string, error) {
    claims := jwt.MapClaims{}
    claims["phone"] = phone
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix() // Token berlaku selama 72 jam

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// Middleware untuk memverifikasi JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.Request.Header.Get("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is required"})
            c.Abort()
            return
        }

        // Menghapus prefix "Bearer " jika ada
        if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
            tokenString = tokenString[7:]
        }

        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
            c.Abort()
            return
        }

        // Jika token valid, simpan klaim ke dalam konteks
        c.Set("phone", claims["phone"])
        c.Next()
    }
}