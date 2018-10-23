package main

import (
  "fmt"
  "github.com/kelsier27/time-based-otp/totp"
)

func main() {
  fmt.Println("Generating new Time-based OTP")

  T := totp.CalculateTimeSteps("2033-05-18T03:33:20Z")
  fmt.Println(T)
}
