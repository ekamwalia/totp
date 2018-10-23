package totp

import (
  "math"
  "time"
)

var (
  time_step = 30.0
  time0  = "1970-01-01T00:00:00Z"
)

func SetTimeStep(step float64) {
  time_step = step
}

func SetEpochTime(epoch string) {
  time0 = epoch
}

func CalculateTimeSteps(curr_time_str string) float64 {
  epoch_time, _ := time.Parse(time.RFC3339, time0)
  curr_time, _ := time.Parse(time.RFC3339, curr_time_str)
  duration := curr_time.Sub(epoch_time)
  
  T := math.Floor(duration.Seconds()/time_step)
  return T
}
