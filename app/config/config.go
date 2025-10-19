package config

import (
	"encoding/base64"

	"github.com/golang-jwt/jwt"
)

func LoadKeys(env *AppEnv) error {
	// AccessPublicKey
	decodedAccessPublicKey, err := base64.StdEncoding.DecodeString(env.Auth.Tokens.AccessPublicKey)
	if err != nil {
		return err
	}
	env.Auth.AccessPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(decodedAccessPublicKey)
	if err != nil {
		return err
	}

	// AccessPrivateKey
	decodedAccessPrivateKey, err := base64.StdEncoding.DecodeString(env.Auth.Tokens.AccessPrivateKey)
	if err != nil {
		return err
	}
	env.Auth.AccessPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(decodedAccessPrivateKey)
	if err != nil {
		return err
	}

	// RefreshPublicKey
	decodedRefreshPublicKey, err := base64.StdEncoding.DecodeString(env.Auth.Tokens.RefreshPublicKey)
	if err != nil {
		return err
	}
	env.Auth.RefreshPublicKey, err = jwt.ParseRSAPublicKeyFromPEM(decodedRefreshPublicKey)
	if err != nil {
		return err
	}

	// RefreshPrivateKey
	decodedRefreshPrivateKey, err := base64.StdEncoding.DecodeString(env.Auth.Tokens.RefreshPrivateKey)
	if err != nil {
		return err
	}
	env.Auth.RefreshPrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(decodedRefreshPrivateKey)
	if err != nil {
		return err
	}
	return nil
}
