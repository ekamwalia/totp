package totp

import "encoding/hex"

var (
	SeedK    = "3132333435363738393031323334353637383930"
	SeedK256 = "3132333435363738393031323334353637383930" + "313233343536373839303132"
	SeedK512 = "3132333435363738393031323334353637383930" + "3132333435363738393031323334353637383930" + "3132333435363738393031323334353637383930" + "31323334"

	DIGIT_POWERS = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000}
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
