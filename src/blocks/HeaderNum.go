package blocks

import (
	"strings"
)

func HeaderNum(block string) (int, int) {

	trimmed := strings.TrimLeft(block, " ")
	max := 6

	for i, char := range trimmed {
		if string(char) != "#" {
			if i > 6 {
				return max, i
			}
			return i, i
		}
	}
	return 0, 0
}
