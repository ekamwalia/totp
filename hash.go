package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strconv"
)

var (
	DIGIT_POWERS = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000}
)

func getHmacHash(msg, key []byte, crypto string) []byte {

	var mac hash.Hash
	if crypto == "HMACSHA1" {
		mac = hmac.New(sha1.New, key)
	} else if crypto == "HMACSHA256" {
		mac = hmac.New(sha256.New, key)
	} else {
		mac = hmac.New(sha512.New, key)
	}

	mac.Write(msg)
	return mac.Sum(nil)
}

func truncate(hash []byte, digits int) string {
	offset := hash[len(hash)-1] & 0xf
	binCode := (int64(hash[offset])&0x7f)<<24 |
		(int64(hash[offset+1])&0xff)<<16 |
		(int64(hash[offset+2])&0xff)<<8 |
		(int64(hash[offset+3]) & 0xff)

	otp := binCode % DIGIT_POWERS[digits]
	return strconv.FormatInt(otp, 10)
}
