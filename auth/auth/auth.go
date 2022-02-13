package auth

import (
	"context"
	authbp "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	OpenIdResolver OpenIdResolver
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Mongo          *dao.Mongo
	Logger         *zap.Logger
}

type OpenIdResolver interface {
	Resolve(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authbp.LoginRequest) (*authbp.LoginResponse, error) {
	openId, err := s.OpenIdResolver.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "cannot resolve openid: %v", err)
	}
	accountID, err := s.Mongo.ResolveAccountId(c, openId)
	if err != nil {
		s.Logger.Error("cannot resolve account id: %v", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	tkn, err := s.TokenGenerator.GenerateToken(accountID.String(), s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &authbp.LoginResponse{
		AccessToken: tkn,
		ExpiresIn:   int32(s.TokenExpire.Seconds()),
	}, nil
}
