package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rizkirobawa/task-5-pbi-btpns-rizki/helpers"
)

// Fungsi Require_Auth adalah middleware yang memeriksa keberadaan dan kevalidan token otentikasi JWT pada permintaan.
func Req_Auth(c *gin.Context) {
	// Mengekstrak token dari cookie "Authorization".
	tokenStr, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Melakukan parsing token JWT.
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode penandatanganan yang tidak diharapkan: %v", token.Header["alg"])
		}
		return []byte(helpers.JWT_KEY), nil
	})

	// Memeriksa validitas token.
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Memeriksa apakah token sudah kedaluwarsa.
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Menambahkan UserID ke konteks Gin untuk digunakan oleh handler selanjutnya.
		c.Set("userid", claims["sub"])
		c.Next()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token JWT tidak valid",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
