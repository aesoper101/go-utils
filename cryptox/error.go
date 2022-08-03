package cryptox

import "errors"

var (
	ErrAesKeySize          = errors.New("key size must be 16, 24, or 32 bytes")
	ErrDesKeySize          = errors.New("key size must be 8")
	ErrAesNotSupportedGcm  = errors.New("aes-gcm not supported")
	ErrDesNotSupportedGcm  = errors.New("des-gcm not supported")
	ErrDes3NotSupportedGcm = errors.New("des3-gcm not supported")
)
