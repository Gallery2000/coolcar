package auth

import (
	"context"
	authbp "coolcar/auth/api/gen/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	OpenIdResolver OpenIdResolver
	TokenGenerator TokenGenerator
}

type OpenIdResolver interface {
	Resolve(code string) (string error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string error)
}

func (s *Service) Login(c context.Context, req *authbp.LoginRequest) (*authbp.LoginResponse, error) {
	openId, err := s.OpenIdResolver(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot resolve openid: %v", err)
	}

}
