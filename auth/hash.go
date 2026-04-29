package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err

}

func CheckPassword(password, matchPass string) bool {
	isValiderr := bcrypt.CompareHashAndPassword([]byte(matchPass), []byte(password))
	return isValiderr == nil
}
