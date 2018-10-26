// Package totp generates a Timed One Time Password conforming to RFC6238
package totp

const (
	DefaultT0         = "1970-01-01T00:00:00Z"
	DefaultX          = 30
	DefaultCodeDigits = 8

	CryptoSHA1   = "HMACSHA1"
	CryptoSHA256 = "HMACSHA256"
	CryptoSHA512 = "HMACSHA512"
)

type TOTPConfig struct {
	// Epoc time in standard format
	// eg: "1970-01-01T00:00:00Z"
	T0 string

	// Time step
	// eg 30.0
	X int64

	// shared secret K
	SecretK string

	// number of digits in OTP
	CodeDigits int

	// HMAC algorithm
	// enum{'HMACSHA1', 'HMACSHA256', 'HMACSHA512'}
	Crypto string
}

/**
Returns a TOTP config object that uses SHA1 for HMAC hashing
and default values for `T0`, `X`, and `CodeDigits`.

The secret key still needs supplied by the user.
*/
func NewDefaultSHA1(seed string) *TOTPConfig {
	totp := DefaultConfig(seed)
	totp.Crypto = CryptoSHA1
	return totp
}

/**
Returns a TOTP config object that uses SHA256 for HMAC hashing
and default values for `T0`, `X`, and `CodeDigits`.

The secret key still needs supplied by the user.
*/
func NewDefaultSHA256(seed string) *TOTPConfig {
	totp := DefaultConfig(seed)
	totp.Crypto = CryptoSHA256
	return totp
}

/**
Returns a TOTP config object that uses SHA512 for HMAC hashing
and default values for `T0`, `X`, and `CodeDigits`.

The secret key still needs supplied by the user.
*/
func NewDefaultSHA512(seed string) *TOTPConfig {
	totp := DefaultConfig(seed)
	totp.Crypto = CryptoSHA512
	return totp
}

// Returns TOTP config object with defaults
func DefaultConfig(seed string) *TOTPConfig {
	return &TOTPConfig{
		T0:         DefaultT0,
		X:          DefaultX,
		SecretK:    seed,
		CodeDigits: DefaultCodeDigits,
	}
}
