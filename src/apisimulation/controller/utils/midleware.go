package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Secret key untuk JWT
var jwtSecret = []byte(viper.GetString("token")) // Ganti dengan secret key yang aman

// Fungsi untuk membuat token JWT
func GenerateToken(phone string) (string, error) {
    claims := jwt.MapClaims{
        "sub": phone, // Menggunakan phone sebagai ID pengguna
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

        // Memeriksa klaim untuk memastikan token milik pengguna yang tepat
        userID := claims["sub"].(string) // Mengambil ID pengguna dari klaim
        requestUserID := c.Param("userID") // Mengambil ID pengguna dari URL atau parameter

        if userID != requestUserID {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Token does not belong to this user"})
            c.Abort()
            return
        }

        // Simpan klaim ke dalam konteks
        c.Set("userID", userID)
        c.Next()
    }
}