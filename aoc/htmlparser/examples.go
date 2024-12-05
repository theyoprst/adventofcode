package htmlparser

import (
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
	"golang.org/x/net/html"
)

// ExtractExamples extracts examples from the article.
// example is a <pre><code>...</code></pre> block followed by a <p>...</p> paragraph with "For example" words.
func ExtractExamples(paragraphs []soup.Root) ([]string, error) {
	var examples []string
	exampleAnnounced := false
	for _, p := range paragraphs {
		if p.Pointer.Type == html.TextNode && strings.TrimSpace(p.FullText()) == "" {
			continue
		}
		if p.Pointer.Type == html.ElementNode && p.Pointer.Data == "pre" && exampleAnnounced {
			if len(p.Children()) != 1 || p.Pointer.FirstChild.Type != html.ElementNode || p.Pointer.FirstChild.Data != "code" {
				return nil, fmt.Errorf("unexpected <pre> content: %v", p.Pointer.FirstChild)
			}
			ex := p.FullText()
			if ex[len(ex)-1] != '\n' {
				ex += "\n"
			}
			examples = append(examples, ex)
		}
		exampleAnnounced = IsExampleAnnounced(p)
	}
	return examples, nil
}

func IsExampleAnnounced(p soup.Root) bool {
	text := strings.ToLower(p.FullText())
	return p.Pointer.Type == html.ElementNode && p.Pointer.Data == "p" && (strings.Contains(text, "for example") ||
		strings.Contains(text, "example:") ||
		strings.Contains(text, "here is an example"))
}
