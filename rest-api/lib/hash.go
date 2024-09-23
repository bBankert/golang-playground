package lib

import "golang.org/x/crypto/bcrypt"

type Hasher struct{}

func (h *Hasher) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", nil
	}

	return string(hashedBytes), nil
}

func (h *Hasher) ValidatePasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewHasher() *Hasher {
	return &Hasher{}
}
