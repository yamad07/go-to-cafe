package value

import "golang.org/x/crypto/bcrypt"

type Password string

func (p Password) Encrypt() (string, error) {
	password := []byte(p)

	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return "", nil
	}
	return string(hashed), nil
}

func (p Password) Compare(hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(p))
}
