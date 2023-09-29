package helpers

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// JWT_KEY merupakan variabel untuk penandatanganan token JWT
var JWT_KEY = "rizkirobawa"

// Validation merupakan fungsi untuk memvalidasi data menggunakan govalidator
func Validation(c *gin.Context, data interface{}) error {
	// Melakukan validasi terhadap struct menggunakan govalidator.ValidateStruct
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		// Jika terdapat error validasi, mengembalikan respons JSON dengan pesan error
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return err
	}
	return nil
}

// fungsi ini untuk mengenkripsi password menggunakan bcrypt
func EncryptPassword(c *gin.Context, password string) string {
	// hash pass menggunakan bcrypt
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal mengenkripsi password"})
		return ""
	}
	return string(hash)

}

// fungsi ini untuk mengecek apakah password yang terenkripsi cocok dengan kata sandi asli
func CheckPassword(pass_1 string, pass_2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pass_1), []byte(pass_2))
	if err != nil {
		return err
	}
	return err
}

// fungsi ini untuk menginisialisasi token JWT dengan ID Pengguna
func InitToken(userid string) (string, error) {
	// mengatur waktu expired time Token JWT menjadi 24 jam
	expTime := time.Now().Add(time.Hour * 24)

	// Membuat Token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userid,
		"exp": expTime.Unix(),
	})

	// Menandatangani token dengan JWT_KEY
	tokenStr, err := token.SignedString([]byte(JWT_KEY))
	return tokenStr, err
}
