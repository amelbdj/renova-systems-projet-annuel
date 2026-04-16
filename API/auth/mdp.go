package auth // ou le nom de ton dossier

import "golang.org/x/crypto/bcrypt"


func HashPassword(password string) (string, error) {
	
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}