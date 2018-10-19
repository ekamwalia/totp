package main

import (
  "fmt"
  "math"
  "time"
)

const (
  time_step = 30
  time0  = "1970-01-01T00:00:00Z"
)

func CalculateTimeSteps(curr_time_str string) float64 {
  epoch_time, _ := time.Parse(time.RFC3339, time0)
  curr_time, _ := time.Parse(time.RFC3339, curr_time_str)
  duration := curr_time.Sub(epoch_time)
  
  T := math.Floor(duration.Seconds()/time_step)
  return T
}

func main() {
  fmt.Println("Generating new Time-based OTP")

  T := CalculateTimeSteps("2033-05-18T03:33:20Z")
  fmt.Println(T)
}
