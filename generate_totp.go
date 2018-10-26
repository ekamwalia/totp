package totp

import "encoding/hex"

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
