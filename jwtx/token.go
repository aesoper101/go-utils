package jwtx

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type registeredClaims[T any] struct {
	jwt.RegisteredClaims
	Payload T
}

type jwtToken[V any] struct {
	opt *options
	_   V
}

// NewJwtToken creates a new JWT token
func NewJwtToken[V any](opts ...JwtTokenOptions) *jwtToken[V] {
	opt := &options{
		secretKey:     "",
		signingMethod: jwt.SigningMethodHS256,
	}
	for _, o := range opts {
		o(opt)
	}
	return &jwtToken[V]{opt: opt}
}

// GenerateToken generates a new token with the given payload
// an error if the token could not be generated
func (t *jwtToken[V]) GenerateToken(claims V) (string, error) {
	exp := time.Now().Add(t.opt.expired)
	expiresAt := jwt.NewNumericDate(exp)
	issuedAt := jwt.NewNumericDate(time.Now())

	jwtClaims := registeredClaims[V]{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    t.opt.issuer,
			Audience:  []string{t.opt.audience},
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
		Payload: claims,
	}
	token := jwt.NewWithClaims(t.opt.signingMethod, jwtClaims)
	return token.SignedString([]byte(t.opt.secretKey))
}

// ParseToken parses the given token and returns the payload. an error is returned if the token is invalid
func (t *jwtToken[V]) ParseToken(token string) (V, error) {
	claims := &registeredClaims[V]{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.opt.secretKey), nil
	})

	return claims.Payload, err
}
