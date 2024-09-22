package utils

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Secret key untuk JWT
var jwtSecret = []byte(viper.GetString("token")) // Ganti dengan secret key yang aman

// Fungsi untuk membuat token JWT
func GenerateToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "sub": strconv.Itoa(userID), // Mengonversi userID ke string
        "exp": time.Now().Add(time.Hour * 1).Unix(), // Token berlaku selama 1 jam
    }
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

        // Menghapus prefix "Bearer "
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

        // Mengambil ID pengguna dari klaim
        userIDStr, ok := claims["sub"].(string)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
            c.Abort()
            return
        }

        // Mengonversi kembali ke integer
        userID, err := strconv.Atoi(userIDStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
            c.Abort()
            return
        }

        // Simpan userID ke dalam konteks
        c.Set("userID", userID)
        c.Next()
    }
}