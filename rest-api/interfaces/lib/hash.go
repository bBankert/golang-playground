package interfaces

type IHasher interface {
	HashPassword(password string) (string, error)
	ValidatePasswordHash(password, hashedPassword string) bool
}
