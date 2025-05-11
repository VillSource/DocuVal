package identity

type Repository interface {
	AddAuthenticationCodeFlow(code string, challenge string, method string) error
	GetAuthenticationCodeFlow(code string) 
}