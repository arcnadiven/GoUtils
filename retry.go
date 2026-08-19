package utils

import "time"

func RetryInterval(num int, interval time.Duration, f func() error) error {
	var err error
	for {
		if err = f(); err == nil {
			break
		}
		num -= 1
		if num == 0 {
			break
		}
		time.Sleep(interval)
	}
	return err
}
