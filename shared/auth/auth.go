package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer"
)

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	aid := impersonationFromContext(ctx)
	if aid != nil {
		return handler(ContextWithAccountId(ctx, id.AccountID(aid)), req)
	}
	//tkn,err :=
}

func tokenFromContext(c context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", unauthenticated
	}
	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}

func impersonationFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}
	return imp[0]
}

type AccountIDKey struct {
}

func ContextWithAccountId(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, AccountIDKey{}, aid)
}

func Interceptor(PublicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(PublicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key:%v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key:%v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key:%v", err)
	}

	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandleReq, nil
}
