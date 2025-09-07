package service

import (
	"context"
	"github.com/webook-project-go/webook-user/domain"
	"github.com/webook-project-go/webook-user/repository"
)

type authBindingService struct {
	repo repository.AuthBindingRepository
}

func NewAuthBindingService(repo repository.AuthBindingRepository) AuthBindingService {
	return &authBindingService{
		repo: repo,
	}
}

func (s *authBindingService) FindOrCreateOauth2Binding(ctx context.Context, binding domain.Oauth2Binding) (domain.Oauth2Binding, error) {
	return s.repo.FindOrCreateOauth2Binding(ctx, binding)
}
