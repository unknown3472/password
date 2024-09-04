package password

import "golang.org/x/crypto/bcrypt"


func GeneratePassword(password string)(string, error){
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func CheckPassword(hashedPassword, password string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err==nil
}