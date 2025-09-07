package repository

import (
	"context"
	"errors"
	"github.com/webook-project-go/webook-user/domain"
	"github.com/webook-project-go/webook-user/repository/dao"
)

type AuthBindingRepository interface {
	FindOrCreateOauth2Binding(ctx context.Context, binding domain.Oauth2Binding) (domain.Oauth2Binding, error)
}
type authBindingRepository struct {
	d dao.Dao
}

func NewRepository(d dao.Dao) AuthBindingRepository {
	return &authBindingRepository{
		d: d,
	}
}

func (r *authBindingRepository) FindOrCreateOauth2Binding(ctx context.Context, binding domain.Oauth2Binding) (domain.Oauth2Binding, error) {
	res, err := r.d.FindBinding(ctx, binding)
	if err != nil {
		if errors.Is(err, dao.ErrRecordNotFound) {
			user, err := r.d.InsertUser(ctx, domain.User{})
			if err != nil {
				return domain.Oauth2Binding{}, err
			}
			binding.UserID = user.Id
			res, err = r.d.InsertOauth2Binding(ctx, binding)
			if err != nil {
				return domain.Oauth2Binding{}, err
			}
			return res, nil
		} else {
			return domain.Oauth2Binding{}, err
		}
	}
	return res, nil
}
