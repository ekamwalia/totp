package totp_test

import (
	"github.com/kelsier27/totp"
	"testing"
)

/**
All the tests in the following functions use Time and TOTP values
from Appendix B (Test Vectors) of RFC6238.
The sharedSecrets used are taken from Appendix A as
indicated in Errata.

Reference - https://tools.ietf.org/html/rfc6238#appendix-B
*/
func TestDefaultGeneratorSHA1(t *testing.T) {
	var sharedSecret = "3132333435363738393031323334353637383930"
	var inputs = []struct {
		Time string
		OTP  string
	}{
		{"1970-01-01T00:00:59Z", "94287082"},
		{"2005-03-18T01:58:29Z", "07081804"},
		{"2005-03-18T01:58:31Z", "14050471"},
		{"2009-02-13T23:31:30Z", "89005924"},
		{"2033-05-18T03:33:20Z", "69279037"},
		{"2603-10-11T11:33:20Z", "65353130"},
	}

	otpSHA1 := totp.NewDefaultSHA1(sharedSecret)
	for _, input := range inputs {
		otp, _ := otpSHA1.GenerateTOTP(input.Time)

		if otp != input.OTP {
			t.Error(
				"For", input.Time,
				"Expected", input.OTP,
				"Got", otp,
			)
		}
	}
}

func TestDefaultGeneratorSHA256(t *testing.T) {
	var sharedSecret = "3132333435363738393031323334353637383930" + "313233343536373839303132"
	var tests = []struct {
		Time string
		OTP  string
	}{
		{"1970-01-01T00:00:59Z", "46119246"},
		{"2005-03-18T01:58:29Z", "68084774"},
		{"2005-03-18T01:58:31Z", "67062674"},
		{"2009-02-13T23:31:30Z", "91819424"},
		{"2033-05-18T03:33:20Z", "90698825"},
		{"2603-10-11T11:33:20Z", "77737706"},
	}

	otpGen := totp.NewDefaultSHA256(sharedSecret)
	for _, test := range tests {
		otp, _ := otpGen.GenerateTOTP(test.Time)

		if otp != test.OTP {
			t.Error(
				"For", test.Time,
				"Expected", test.OTP,
				"Got", otp,
			)
		}
	}
}

func TestDefaultGeneratorSHA512(t *testing.T) {
	var sharedSecret = "3132333435363738393031323334353637383930" + "3132333435363738393031323334353637383930" + "3132333435363738393031323334353637383930" + "31323334"
	var tests = []struct {
		Time string
		OTP  string
	}{
		{"1970-01-01T00:00:59Z", "90693936"},
		{"2005-03-18T01:58:29Z", "25091201"},
		{"2005-03-18T01:58:31Z", "99943326"},
		{"2009-02-13T23:31:30Z", "93441116"},
		{"2033-05-18T03:33:20Z", "38618901"},
		{"2603-10-11T11:33:20Z", "47863826"},
	}

	otpSHA512 := totp.NewDefaultSHA512(sharedSecret)
	for _, test := range tests {
		otp, _ := otpSHA512.GenerateTOTP(test.Time)

		if otp != test.OTP {
			t.Error(
				"For", test.Time,
				"Expected", test.OTP,
				"Got", otp,
			)
		}
	}
}
