package service

import (
	"context"
	"github.com/webook-project-go/webook-user/domain"
)

type UserService interface {
	SignUp(ctx context.Context, user domain.User) error
	Edit(ctx context.Context, userInfo domain.User) error
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	FindUserById(ctx context.Context, id int64) (domain.User, error)
	FindUser(ctx context.Context, filed string, value any) (domain.User, error)
	FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error)
	Login(ctx context.Context, email string, password string) (domain.User, error)
}
type AuthBindingService interface {
	FindOrCreateOauth2Binding(ctx context.Context, binding domain.Oauth2Binding) (domain.Oauth2Binding, error)
}
