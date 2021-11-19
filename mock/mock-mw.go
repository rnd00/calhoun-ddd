package mock

import "lrn/calhoun-ddd/mw"

type UserService struct {
	UserFn func(token string) (mw.User, error)
}

func (us *UserService) Authenticate(token string) (mw.User, error) {
	return us.UserFn(token)
}
