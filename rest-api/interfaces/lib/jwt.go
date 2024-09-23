package interfaces

type IJwtAuthorizer interface {
	GenerateToken(email, userId string) (string, error)
	ValidateToken(token string) (int64, error)
}
