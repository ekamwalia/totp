package totp

import (
  "math"
  "time"
)

var (
  X  = 30.0
  T0 = "1970-01-01T00:00:00Z"
)

func SetTimeStep(step float64) {
  X = step
}

func SetEpochTime(epoch string) {
  T0 = epoch
}

func CalculateTimeSteps(curr_time_str string) float64 {
  epoch_time, _ := time.Parse(time.RFC3339, T0)
  curr_time, _ := time.Parse(time.RFC3339, curr_time_str)
  duration := curr_time.Sub(epoch_time)
  
  T := math.Floor(duration.Seconds()/ X)
  return T
}
