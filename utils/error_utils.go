package utils

import "fmt"

// RecoverError 把 recover() 捕捉到的 panic 信息，转换成 error
func RecoverError(x any) error {
	if x == nil {
		return nil
	}
	err, ok := x.(error)
	if ok {
		return err
	}
	str, ok := x.(string)
	if ok {
		return fmt.Errorf(str)
	}
	return fmt.Errorf("unknown panic")
}
