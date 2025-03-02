package inputport

type IUserUsecase interface {
	Authenticate(token string) (string, error) 
}