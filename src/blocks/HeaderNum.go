package blocks

import (
	"strings"
)

func HeaderNum(block string) (int, int) {

	trimmed := strings.TrimLeft(block, " ")

	for i, char := range trimmed {
		if string(char) != "#" {
			if i <= 6 {
				return i, i
			}
			return 6, i
		}
	}
	return 0, 0
}
