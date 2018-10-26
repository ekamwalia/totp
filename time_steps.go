package totp

import (
	"errors"
	"fmt"
	"math"
	"time"
)

/**
	@param: curr_time_str: input time in standard format string
 	@returns: time steps in Hex string
*/
func (otpConf *TOTPConfig) calculateTimeSteps(curr_time_str string) (string, error) {
	epoch_time, _ := time.Parse(time.RFC3339, otpConf.T0)
	curr_time, _ := time.Parse(time.RFC3339, curr_time_str)

	steps := math.Floor(float64((curr_time.Unix() - epoch_time.Unix()) / otpConf.X))
	if steps >= math.MaxInt64 || steps <= math.MinInt64 {
		err := errors.New("f64 is out of int range")
		return "0", err
	}

	decT := int64(steps)
	T := fmt.Sprintf("%X", decT)
	for len(T) < 16 {
		T = "0" + T
	}
	return T, nil
}
