package store

import "io"

// CloseAll 关闭列表中的所有项目, 如果 reverse==true, 就从后往前执行
func CloseAll(all []io.Closer, reverse bool) error {
	if all == nil {
		return nil
	}
	var result error = nil
	size := len(all)
	if !reverse {
		// forward
		for i := 0; i < size; i++ {
			c := all[i]
			err := c.Close()
			if err != nil {
				result = err
			}
		}
	} else {
		// backward
		for i := size - 1; i >= 0; i-- {
			c := all[i]
			err := c.Close()
			if err != nil {
				result = err
			}
		}
	}
	return result
}
