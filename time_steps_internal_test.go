package totp

import (
	"testing"
)

/**
All the tests in the following functions use Time and T (hex) values
from Appendix B (Test Vectors) of RFC6238.
The sharedSecrets used are taken from Appendix A as
indicated in Errata.

Reference - https://tools.ietf.org/html/rfc6238#appendix-B
*/
func TestCalculateTimeSteps(t *testing.T) {
	var sharedSecret = "3132333435363738393031323334353637383930"
	var inputs = []struct {
		Time string
		OTP  string
	}{
		{"1970-01-01T00:00:59Z", "0000000000000001"},
		{"2005-03-18T01:58:29Z", "00000000023523EC"},
		{"2005-03-18T01:58:31Z", "00000000023523ED"},
		{"2009-02-13T23:31:30Z", "000000000273EF07"},
		{"2033-05-18T03:33:20Z", "0000000003F940AA"},
		{"2603-10-11T11:33:20Z", "0000000027BC86AA"},
	}

	otpSHA1 := NewDefaultSHA1(sharedSecret)
	for _, input := range inputs {
		T, _ := otpSHA1.calculateTimeSteps(input.Time)

		if T != input.OTP {
			t.Error(
				"For", input.Time,
				"Expected", input.OTP,
				"Got", T,
			)
		}
	}
}
