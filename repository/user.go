package repository

import (
	"context"
	"errors"
	"github.com/webook-project-go/webook-user/domain"
	"github.com/webook-project-go/webook-user/repository/cache"
	"github.com/webook-project-go/webook-user/repository/dao"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	Edit(ctx context.Context, info domain.User) error
	FindById(ctx context.Context, id int64) (domain.User, error)
	FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error)
	FindUser(ctx context.Context, filed string, value any) (domain.User, error)
}

var (
	ErrEmailDuplicate         = dao.ErrDuplicate
	ErrInvalidEmailOrPassword = dao.ErrInvalidEmailOrPassword
	ErrRecordNotExist         = dao.ErrRecordNotFound
)

type userRepository struct {
	db    dao.Dao
	cache cache.Cache
}

func New(d dao.Dao, cache cache.Cache) UserRepository {
	return &userRepository{
		db:    d,
		cache: cache,
	}
}

func (u *userRepository) Create(ctx context.Context, user domain.User) error {
	_, err := u.db.InsertUser(ctx, user)
	return err
}
func (u *userRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return u.db.FindByEmail(ctx, email)
}

func (u *userRepository) Edit(ctx context.Context, info domain.User) error {
	err := u.db.Edit(ctx, info)
	if err != nil {
		return err
	}
	go func() {
		newUser, _ := u.db.FindById(ctx, info.Id)
		err = u.cache.Set(ctx, newUser)
		if err != nil {
			// log
		}
	}()
	return nil
}

func (u *userRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	if user, err := u.cache.Get(ctx, id); err == nil {
		return user, nil
	} else if errors.Is(err, cache.ErrKeyNotFound) {
		du, err := u.db.FindById(ctx, id)
		if err != nil {
			return domain.User{}, err
		}
		//go func() {
		err = u.cache.Set(ctx, du)
		if err != nil {
			// log
		}
		//}()
		return du, nil
	}

	return domain.User{}, errors.New("redisCache error")

}

func (u *userRepository) FindUser(ctx context.Context, filed string, value any) (domain.User, error) {
	return u.db.FindUser(ctx, filed, value)
}

func (u *userRepository) FindOrCreateByPhone(ctx context.Context, phone string) (domain.User, error) {
	ud, err := u.FindUser(ctx, "Phone", phone)
	if err == nil {
		return ud, nil
	} else if errors.Is(err, dao.ErrRecordNotFound) {
		ud, err = u.db.InsertUser(ctx, domain.User{
			Phone: phone,
		})
		if err != nil {
			return domain.User{}, err
		}
		return ud, nil
	}
	return domain.User{}, err
}
