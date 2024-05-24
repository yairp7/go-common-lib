package debug

import "time"

/*
*
Use like this: defer fmt.Println(MethodDuration(time.Now()))
*/
func MethodDuration(startTime time.Time) time.Duration {
	return time.Since(startTime)
}
