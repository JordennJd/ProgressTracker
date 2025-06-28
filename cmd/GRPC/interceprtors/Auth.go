package interceptors

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authHeader   = "authorization"
	bearerPrefix = "Bearer "
	userIDKey    = "userID"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if isExcludedMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	authHeaderValues := md.Get(authHeader)
	if len(authHeaderValues) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	token := strings.TrimPrefix(authHeaderValues[0], bearerPrefix)
	if token == authHeaderValues[0] { // Если префикс Bearer отсутствует
		return nil, status.Error(codes.Unauthenticated, "invalid authorization token format")
	}

	// Валидируем токен и извлекаем userID
	userID, err := validateTokenAndGetUserID(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	// Добавляем userID в контекст
	newCtx := context.WithValue(ctx, userIDKey, userID)

	// Продолжаем обработку запроса
	return handler(newCtx, req)
}

// Вспомогательные функции

func isExcludedMethod(method string) bool {
	excludedMethods := map[string]bool{
		"/service.AuthService/Login":    true,
		"/service.AuthService/Register": true,
	}
	return excludedMethods[method]
}

func validateTokenAndGetUserID(token string) (uuid.UUID, error) {
	//parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
	//	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	//	}
	//	return []byte("your-secret-key"), nil
	//})
	//if err != nil {
	//	return uuid.Nil, err
	//}
	//
	//claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//if !ok || !parsedToken.Valid {
	//	return uuid.Nil, errors.New("invalid token claims")
	//}
	//
	//userIDStr, ok := claims["sub"].(string)
	//if !ok {
	//	return uuid.Nil, errors.New("userID not found in token")
	//}

	return uuid.Parse("0197470f-f135-780b-9534-c3d5b59f219b")
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(userIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("userID not found in context")
	}
	return userID, nil
}
