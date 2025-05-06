package blocks

import (
	"strings"
)

func HeaderNum(block string) int {

	trimmed := strings.TrimLeft(block, " ")

	for i, char := range trimmed {
		if i > 6 {
			return 6
		}
		if string(char) != "#" {
			return i
		}
		continue
	}
	return 0
}
