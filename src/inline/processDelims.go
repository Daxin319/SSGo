package inline

import (
	"slices"
	"strings"

	"github.com/Daxin319/SSGo/src/nodes"
)

// processDelims handles the processing of delimiter markers like *, **, ***, etc.
// It manages the stack of delimiters and creates appropriate text nodes.
func processDelims(m string, stack []delimRun, newNodes []nodes.TextNode) ([]delimRun, []nodes.TextNode) {
	length := len(m)
	char := m[0]
	if length <= 2 {
		for i := len(stack) - 1; i >= 0; i-- {
			if stack[i].marker == m {
				op := stack[i].pos                            // set original position to current token
				content := slices.Clone(newNodes[op:])        // initialize content stack
				newNodes = newNodes[:op]                      // set nodes stack equal to current nodes stack up to the current position
				newNodes = append(newNodes, wrap(m, content)) // append content stack wrapped in parent delim
				stack = slices.Delete(stack, i, i+1)          // pop token from stack
				return stack, newNodes
			}
			if len(stack[i].marker) == 3 && stack[i].marker[0] == char && length < 3 {
				op := stack[i].pos
				if len(newNodes) > op {
					content := slices.Clone(newNodes[op:])
					newNodes = newNodes[:op]
					newNodes = append(newNodes, wrap(m, content))

					stack[i].marker = strings.Repeat(string(char), 3-length)
					return stack, newNodes
				}
			}
		}
		stack = append(stack, delimRun{marker: m, pos: len(newNodes)}) // append to be treated as plaintext
		return stack, newNodes
	}
	remaining := length
	for remaining > 0 { // as long as there are still characters in the delim
		idx := -1
		for j := len(stack) - 1; j >= 0; j-- {
			if stack[j].marker[0] == char && len(stack[j].marker) <= remaining {
				idx = j // set index to current stack position
				break
			}
		}
		if idx >= 0 {
			mrk := stack[idx].marker                        // set marker2 to the marker at current stack index
			op := stack[idx].pos                            // set original position to current stack position
			content := slices.Clone(newNodes[op:])          // initialize content stack of all nodes above this one in the nodes stack
			newNodes = newNodes[:op]                        // set nodes stack equal to all nodes below this one in the stack
			newNodes = append(newNodes, wrap(mrk, content)) // append content stack wrapped in parent delim to nodes stack
			stack = slices.Delete(stack, idx, idx+1)        // pop this token off the stack
			remaining -= len(mrk)                           // subtract this marker's length from remaining to ensure triple delims are properly split and matched
		} else {
			marker := strings.Repeat(string(char), remaining) // treat as unmatched, append all to stack
			stack = append(stack, delimRun{marker: marker, pos: len(newNodes)})
			break
		}
	}
	return stack, newNodes
}
