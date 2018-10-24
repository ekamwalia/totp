package main

import (
	"fmt"
	"github.com/kelsier27/time-based-otp/totp"
)

func main() {
	fmt.Println("Generating New Time-based OTP")
	// Get T (time factor) hex string
	newOTP1 := totp.TOTPConfig{
		T0:         "1970-01-01T00:00:00Z",
		X:          30,
		SecretK:    totp.SeedK,
		CodeDigits: 8,
		Crypto:     "HMACSHA1",
	}
	newOTP256 := totp.TOTPConfig{
		T0:         "1970-01-01T00:00:00Z",
		X:          30,
		SecretK:    totp.SeedK256,
		CodeDigits: 8,
		Crypto:     "HMACSHA256",
	}
	newOTP512 := totp.TOTPConfig{
		T0:         "1970-01-01T00:00:00Z",
		X:          30,
		SecretK:    totp.SeedK512,
		CodeDigits: 8,
		Crypto:     "HMACSHA512",
	}

	testArray := []string{
      "1970-01-01T00:00:59Z",
      "2005-03-18T01:58:29Z",
      "2005-03-18T01:58:31Z",
      "2009-02-13T23:31:30Z",
      "2033-05-18T03:33:20Z",
      "2603-10-11T11:33:20Z",
    }

	for _, date := range testArray {
      otp, _ := newOTP1.GenerateTOTP(date)
      fmt.Println(date, otp)
      otp, _ = newOTP256.GenerateTOTP(date)
      fmt.Println(date, otp)
      otp, _ = newOTP512.GenerateTOTP(date)
	  fmt.Println(date, otp)

	  fmt.Println("\n")
    }
}






