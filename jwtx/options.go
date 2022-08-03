package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type options struct {
	secretKey     string
	signingMethod jwt.SigningMethod
	issuer        string
	audience      string
	expired       time.Duration
}

type JwtTokenOptions func(*options)

// WithSecretKey sets the secret key for the jwt token
// The functionx can not be used with WithSecretKeyCallback together
func WithSecretKey(secretKey string) JwtTokenOptions {
	return func(o *options) {
		o.secretKey = secretKey
	}
}

// WithSecretKeyCallback sets the secret key for the jwt token by callback.
// The callback will be called when the secret key is needed.
// WithSecretKeyCallback can not be used with WithSecretKey together.
func WithSecretKeyCallback(callback func() string) JwtTokenOptions {
	return func(o *options) {
		o.secretKey = callback()
	}
}

// WithSigningMethod sets the signing method for the jwt token
func WithSigningMethod(signingMethod jwt.SigningMethod) JwtTokenOptions {
	return func(o *options) {
		o.signingMethod = signingMethod
	}
}

// WithIssuer sets the issuer for the jwt token
func WithIssuer(issuer string) JwtTokenOptions {
	return func(o *options) {
		o.issuer = issuer
	}
}

// WithAudience sets the audience for the jwt token
func WithAudience(audience string) JwtTokenOptions {
	return func(o *options) {
		o.audience = audience
	}
}

// WithExpired sets the expired for the jwt token
func WithExpired(expired time.Duration) JwtTokenOptions {
	return func(o *options) {
		o.expired = expired
	}
}
