package totp

import "encoding/hex"

var (
	DIGIT_POWERS = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000}
)

const (
	DefaultT0         = "1970-01-01T00:00:00Z"
	DefaultX          = 30
	CryptoSHA1        = "HMACSHA1"
	CryptoSHA256      = "HMACSHA256"
	CryptoSHA512      = "HMACSHA512"
	DefaultCodeDigits = 8
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

func NewDefaultSHA1(seed string) *TOTPConfig {
	totp := newDefault(seed)
	totp.Crypto = CryptoSHA1
	return totp
}

func NewDefaultSHA256(seed string) *TOTPConfig {
	totp := newDefault(seed)
	totp.Crypto = CryptoSHA256
	return totp
}

func NewDefaultSHA512(seed string) *TOTPConfig {
	totp := newDefault(seed)
	totp.Crypto = CryptoSHA512
	return totp
}

func newDefault(seed string) *TOTPConfig {
	return &TOTPConfig{
		T0:         DefaultT0,
		X:          DefaultX,
		SecretK:    seed,
		CodeDigits: DefaultCodeDigits,
	}
}

func (otpConf *TOTPConfig) GenerateTOTP(date string) (string, error) {
	T, err := otpConf.calculateTimeSteps(date)
	if err != nil {
		return "", err
	}

	msg, _ := hex.DecodeString(T)
	key, _ := hex.DecodeString(otpConf.SecretK)

	hash := getHmacHash(msg, key, otpConf.Crypto)
	newTotp := truncate(hash, otpConf.CodeDigits)

	for len(newTotp) < otpConf.CodeDigits {
		newTotp = "0" + newTotp
	}
	return newTotp, nil
}
