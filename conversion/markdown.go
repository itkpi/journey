package conversion

import (
	// "github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
)

// Using blackfriday library for markdown to html conversion. At least for now.

// func GenerateHtmlFromMarkdown(input []byte) []byte {
// 	// return blackfriday.MarkdownCommon(input)
// 	return bluemonday.UGCPolicy().SanitizeBytes(input)
// }

func SanitizeHtml(input []byte) []byte {
	// return blackfriday.MarkdownCommon(input)
	policy := bluemonday.UGCPolicy()
	policy.AllowAttrs("width", "height", "src", "allowfullscreen", "frameborder").
		OnElements("iframe")
	return policy.SanitizeBytes(input)
}
