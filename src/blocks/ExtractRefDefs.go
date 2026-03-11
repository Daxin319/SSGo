package blocks

import (
	"regexp"
	"strings"

	"github.com/Daxin319/SSGo/src/nodes"
)

var defRe = regexp.MustCompile(`^\[([^]]+)]:\s*(\S+)(?:\s+"([^"]+)")?\s*$`)

func ExtractRefDefs(block string) string {
	lines := strings.Split(block, "\n")
	var out []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if m := defRe.FindStringSubmatch(trimmed); m != nil {
			label := nodes.UnescapeString(m[1])
			url := nodes.UnescapeString(m[2])
			title := ""
			if len(m) >= 4 {
				title = m[3]
			}
			nodes.SetRef(label, url, title)
		} else {
			out = append(out, line)
		}
	}
	return strings.Join(out, "\n")
}

func GetRef(label string) (string, string, bool) {
	return nodes.GetRef(label)
}
