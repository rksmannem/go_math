package geometry

import "fmt"

//CubeVolume ...
func CubeVolume(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("zero/negetive edge is not allowed: %v", n)
	}
	return n * n * n, nil
}
