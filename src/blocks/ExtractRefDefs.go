package blocks

import (
	"regexp"
	"strings"

	"github.com/Daxin319/SSGo/src/nodes"
)

var refDefs = map[string]struct{ URL, Title string }{}

var defRe = regexp.MustCompile(`^\[([^\]]+)\]:\s*(\S+)(?:\s+"([^"]+)")?\s*$`)

// func to extract reference links, the rest of the architecture hasn't been started yet
func extractRefDefs(lines []string) []string {
	var out []string
	for _, line := range lines {
		if m := defRe.FindStringSubmatch(line); m != nil {
			label := strings.ToLower(nodes.UnescapeString(m[1]))
			url := nodes.UnescapeString(m[2])
			title := ""
			if len(m) >= 4 {
				title = m[3]
			}
			refDefs[label] = struct{ URL, Title string }{url, title}
		} else {
			out = append(out, line)
		}
	}
	return out
}
