package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/Uttamnath64/quixzap/app/common/types"
	"github.com/Uttamnath64/quixzap/app/config"
	"github.com/Uttamnath64/quixzap/app/models"
	"github.com/Uttamnath64/quixzap/app/repositories"
	"github.com/Uttamnath64/quixzap/app/utils/requests"
	"github.com/golang-jwt/jwt"
)

type JWT struct {
	env      *config.AppEnv
	authRepo repositories.AuthRepository
}

type JWTClaim struct {
	UserId    uint           `json:"user_id"`
	UserType  types.UserType `json:"user_type"`
	SessionID uint           `json:"session_id"`
	jwt.StandardClaims
}

func NewJWT(env *config.AppEnv, authRepo repositories.AuthRepository) *JWT {
	return &JWT{
		env:      env,
		authRepo: authRepo,
	}
}

func (auth *JWT) GenerateToken(rctx *requests.RequestContext, userId uint, userType types.UserType, deviceInfo, ipAddress string) (string, string, error) {

	var accessExpiresAt = time.Now().Add(auth.env.Auth.AccessTokenExpired).Unix()
	var refreshExpiresAt = time.Now().Add(auth.env.Auth.RefreshTokenExpired).Unix()

	// create settion
	sessionId, err := auth.authRepo.CreateSession(rctx, &models.Session{
		UserId:     userId,
		UserType:   userType,
		DeviceInfo: deviceInfo,
		IPAddress:  ipAddress,
	})
	if err != nil {
		return "", "", err
	}

	accessTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &JWTClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpiresAt,
		},
	})

	accessToken, err := accessTokenJWT.SignedString(auth.env.Auth.AccessPrivateKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenJWT := jwt.NewWithClaims(jwt.SigningMethodRS256, &JWTClaim{
		UserId:    userId,
		UserType:  userType,
		SessionID: sessionId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpiresAt,
		},
	})

	refreshToken, err := refreshTokenJWT.SignedString(auth.env.Auth.RefreshPrivateKey)
	if err != nil {
		return "", "", err
	}

	if err := auth.authRepo.UpdateSession(rctx, sessionId, refreshToken, refreshExpiresAt); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (auth *JWT) VerifyRefreshToken(rctx *requests.RequestContext, refreshToken string) (interface{}, error) {

	token, err := jwt.ParseWithClaims(
		refreshToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return auth.env.Auth.RefreshPublicKey, nil
		},
	)
	if err != nil || !token.Valid {
		return nil, errors.New("refresh token is invalid")
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || claims.SessionID == 0 {
		return nil, errors.New("invalid refresh token claims")
	}

	if err := auth.isValidRefreshToken(rctx, claims.UserType, refreshToken); err != nil {
		return nil, err
	}

	return claims, nil
}

func (auth *JWT) isValidRefreshToken(rctx *requests.RequestContext, userType types.UserType, refreshToken string) error {
	session, err := auth.authRepo.GetSessionByRefreshToken(rctx, refreshToken, userType)
	if err != nil {
		return err
	}

	// Check if token exists
	if session == nil {
		return errors.New("refresh token not found")
	}

	if session.ExpiresAt < time.Now().Unix() {
		return errors.New("refresh token is expired")
	}

	return nil
}
