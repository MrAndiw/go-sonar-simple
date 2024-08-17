package pkg

import "fmt"

type (
	UserModule interface {
		GetUser(email string) (string, error)
		GetProfile(email string) (string, error)
	}

	userModule struct {
	}
)

func NewUserModule() UserModule {
	return &userModule{}
}

func (u *userModule) GetUser(email string) (string, error) {

	user := fmt.Sprintf("The User Email is %s", email)
	return user, nil
}

func (u *userModule) GetProfile(email string) (string, error) {

	user := fmt.Sprintf("The User Profile is %s", email)
	return user, nil
}
