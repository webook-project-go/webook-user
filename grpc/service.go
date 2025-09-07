package grpc

import (
	"context"
	v1 "github.com/webook-project-go/webook-apis/gen/go/apis/user/v1"
	"github.com/webook-project-go/webook-user/domain"
	"github.com/webook-project-go/webook-user/service"
)

type Service struct {
	user service.UserService
	auth service.AuthBindingService
	v1.UnimplementedUserAuthServiceServer
}

func NewService(user service.UserService, auth service.AuthBindingService) *Service {
	return &Service{
		user: user,
		auth: auth,
	}
}

func (s *Service) SignUp(ctx context.Context, request *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	u := protoToDomainUser(request.User)
	if err := s.user.SignUp(ctx, u); err != nil {
		return nil, err
	}
	return &v1.SignUpResponse{}, nil
}

func (s *Service) Edit(ctx context.Context, request *v1.EditRequest) (*v1.EditResponse, error) {
	u := protoToDomainUser(request.User)
	if err := s.user.Edit(ctx, u); err != nil {
		return nil, err
	}
	return &v1.EditResponse{}, nil
}

func (s *Service) FindUserByEmail(ctx context.Context, request *v1.FindUserByEmailRequest) (*v1.FindUserResponse, error) {
	u, err := s.user.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	return &v1.FindUserResponse{User: domainToProtoUser(u)}, nil
}

func (s *Service) FindUserById(ctx context.Context, request *v1.FindUserByIdRequest) (*v1.FindUserResponse, error) {
	u, err := s.user.FindUserById(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &v1.FindUserResponse{User: domainToProtoUser(u)}, nil
}

func (s *Service) FindUser(ctx context.Context, request *v1.FindUserRequest) (*v1.FindUserResponse, error) {
	u, err := s.user.FindUser(ctx, request.Field, request.Value)
	if err != nil {
		return nil, err
	}
	return &v1.FindUserResponse{User: domainToProtoUser(u)}, nil
}

func (s *Service) FindOrCreateByPhone(ctx context.Context, request *v1.FindOrCreateByPhoneRequest) (*v1.FindUserResponse, error) {
	u, err := s.user.FindOrCreateByPhone(ctx, request.Phone)
	if err != nil {
		return nil, err
	}
	return &v1.FindUserResponse{User: domainToProtoUser(u)}, nil
}

func (s *Service) Login(ctx context.Context, request *v1.LoginRequest) (*v1.FindUserResponse, error) {
	u, err := s.user.Login(ctx, request.Email, request.Password)
	if err != nil {
		return nil, err
	}
	return &v1.FindUserResponse{User: domainToProtoUser(u)}, nil
}

func (s *Service) FindOrCreateOauth2Binding(ctx context.Context, request *v1.FindOrCreateOauth2BindingRequest) (*v1.FindOrCreateOauth2BindingResponse, error) {
	binding := protoToDomainBinding(request.Binding)
	b, err := s.auth.FindOrCreateOauth2Binding(ctx, binding)
	if err != nil {
		return nil, err
	}
	return &v1.FindOrCreateOauth2BindingResponse{Binding: domainToProtoBinding(b)}, nil
}

func protoToDomainUser(u *v1.User) domain.User {
	if u == nil {
		return domain.User{}
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Name:     u.Name,
		Birthday: u.Birthday,
		AboutMe:  u.AboutMe,
	}
}

func domainToProtoUser(u domain.User) *v1.User {
	return &v1.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Name:     u.Name,
		Birthday: u.Birthday,
		AboutMe:  u.AboutMe,
	}
}

func protoToDomainBinding(b *v1.Oauth2Binding) domain.Oauth2Binding {
	if b == nil {
		return domain.Oauth2Binding{}
	}
	return domain.Oauth2Binding{
		UserID:      b.UserId,
		ID:          b.Id,
		Provider:    b.Provider,
		ExternalID:  b.ExternalId,
		AccessToken: b.AccessToken,
	}
}

func domainToProtoBinding(b domain.Oauth2Binding) *v1.Oauth2Binding {
	return &v1.Oauth2Binding{
		UserId:      b.UserID,
		Id:          b.ID,
		Provider:    b.Provider,
		ExternalId:  b.ExternalID,
		AccessToken: b.AccessToken,
	}
}
