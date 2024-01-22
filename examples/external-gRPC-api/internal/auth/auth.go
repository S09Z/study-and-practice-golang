package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// Check if the "authorization" header exists
	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// Validate the authorization token (e.g., using JWT)
	valid := validateToken(authHeader[0])
	if !valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token")
	}

	// Proceed with the RPC call if the token is valid
	return handler(ctx, req)
}

func validateToken(token string) bool {
	// Implement token validation logic here (e.g., verify JWT signature, check expiration)
	// Return true if the token is valid, false otherwise
	return true
}
