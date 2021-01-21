package auth

import (
	"context"
	"encoding/base64"
	"errors"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func userClaimFromToken(struct{}) string {
	return "foobar"
}

func parseToken(token string) (struct{}, error) {
	if token == base64.StdEncoding.EncodeToString([]byte("admin:admin")) {
		return struct{}{}, nil
	}
	return struct{}{}, errors.New("invalid credential")
}

// authFunc is used by a middleware to authenticate requests
func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	log.Print(token, tokenInfo)

	grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	// WARNING: in production define your own type to avoid context collisions
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)

	return newCtx, nil
}
