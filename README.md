# TOTP (Time-based One Time Password)
 
This package provides a convenient way of generarting Time-based One-Time Passwords in accordance with IETF's RFC6238.

Hash-based One-Time Password Algorithm(HOTP)uses a counter as a moving factor while generating the OTP. This counter is hashed using an HMAC hashing algorithms (HAMC-SHA1, HMAC-SHA256 or HMAC-512). TOTP builds on on HOTP by deriving the moving factor from UNIX time instead of a counter. 
 
TOTP can be used to create OTPs containing more than 6 digits. Users can choose between the HMAC-SHA1, HMAC-SHA256 and HMAC-SHA512 for generating TOTPs.

## Example
```go
package main

import (
 "github.com/kelsier27/totp"
 "fmt"
)

func main() {
  var sharedSecret = "3132333435363738393031323334353637383930"
  otpgen := totp.NewDefaultTOTP(sharedSecret)
  totp := otpgen.GenerateTOTP("1970-01-01T00:00:59Z")
  fmt.Println(totp) // output - 94287082
}
```

## Install
TOTP can be installed using `go get`

`go get github.com/kelsier27/totp`

## Usage

1. _Obtain TOTPConfig object_

  Before generating the OTP, we need to obtain a TOTPConfig object. The default objects generate a 8-digit TOTP using T0 as Epoch time 1970-01-01T00:00:00Z and _Time Step X_ as 30s. To get the default TOTPConfig, use one of the following functions depending on the choice of HMAC algorithm
  
 - NewDefaultSHA1
 - NewDefaultSHA256
 - NewDefaultSHA512
  
2. _Generate TOTP_

  Once we obtain a config object, we can call the GenerateTOTP method. This method takes a time string as the input and returns a TOTP string.
 
  
