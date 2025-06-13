package inline

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Daxin319/SSGo/src/nodes"
	"github.com/Daxin319/SSGo/src/tokenizer"
	"golang.org/x/net/publicsuffix"
)

// Combined regex for finding potential autolinks (URLs and emails).
var autolinkFindRegex = regexp.MustCompile(
	fmt.Sprintf(`(?i)\b(%s|%s|%s)\b`,
		// Strip anchors from the regexes since we're not matching whole strings anymore
		strings.Trim(tokenizer.EmailRE.String(), "^$"),
		strings.Trim(tokenizer.ProtocolRE.String(), "^$"),
		strings.Trim(tokenizer.GfmDomainRE.String(), "^$"),
	),
)

func autolink(text string) []nodes.TextNode {
	matches := autolinkFindRegex.FindAllStringIndex(text, -1)

	if len(matches) == 0 {
		return []nodes.TextNode{{TextType: nodes.Text, Text: text}}
	}

	var newNodes []nodes.TextNode
	lastIndex := 0

	for _, match := range matches {
		startIndex, endIndex := match[0], match[1]
		potentialLink := text[startIndex:endIndex]

		// Add preceding text if any
		if startIndex > lastIndex {
			newNodes = append(newNodes, nodes.TextNode{
				TextType: nodes.Text,
				Text:     text[lastIndex:startIndex],
			})
		}

		isLink := false
		href := potentialLink

		if tokenizer.EmailRE.MatchString(potentialLink) {
			isLink = true
			href = "mailto:" + potentialLink
		} else if tokenizer.ProtocolRE.MatchString(potentialLink) {
			isLink = true
		} else if tokenizer.GfmDomainRE.MatchString(potentialLink) {
			domain := potentialLink
			if at := strings.LastIndex(domain, "@"); at != -1 {
				domain = domain[at+1:]
			}
			if _, icann := publicsuffix.PublicSuffix(domain); icann {
				isLink = true
				href = "https://" + potentialLink
			}
		}

		if isLink {
			linkNode := nodes.TextNode{
				TextType: nodes.Link,
				Url:      href,
				Children: []nodes.TextNode{{TextType: nodes.Text, Text: potentialLink}},
			}
			newNodes = append(newNodes, linkNode)
		} else {
			// If not a valid link, treat as plain text
			newNodes = append(newNodes, nodes.TextNode{
				TextType: nodes.Text,
				Text:     potentialLink,
			})
		}

		lastIndex = endIndex
	}

	// Add any remaining text after the last match
	if lastIndex < len(text) {
		newNodes = append(newNodes, nodes.TextNode{
			TextType: nodes.Text,
			Text:     text[lastIndex:],
		})
	}

	return newNodes
}
