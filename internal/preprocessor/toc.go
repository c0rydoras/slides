package preprocessor

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func collectH1s(slideContent []byte) []string {
	parser := goldmark.DefaultParser()
	reader := text.NewReader(slideContent)
	doc := parser.Parse(reader)

	var h1s []string

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if h, ok := n.(*ast.Heading); ok {
			if h.Level == 1 {
				h1s = append(h1s, string(h.Lines().Value(slideContent)))
			}
		}
		return ast.WalkContinue, nil
	})

	return h1s
}

func GenerateTOC(slides []string, title string, description string) string {
	var toc strings.Builder
	toc.WriteString(fmt.Sprintf("# %s\n\n", title))
	if description != "" {
		toc.WriteString(fmt.Sprintf("%s\n\n", description))
	}

	var headings = collectH1s([]byte(strings.Join(slides, "\n")))

	for _, text := range headings {
		toc.WriteString(fmt.Sprintf("- %s\n", text))
	}

	// make toc first slide
	if len(slides) == 0 {
		return ""
	}
	return toc.String()
}
