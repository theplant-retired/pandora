package pandoratests

import (
	"github.com/kobeld/gopandoc"
	pandora "github.com/theplant/pandora/clients/go"
	"testing"
)

func BenchmarkToMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pandora.ToMarkdown(sampleHTML)
	}
	return
}

func BenchmarkToHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pandora.ToHTML(sampleMD)
	}
	return
}

func BenchmarkGopandocToMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gopandoc.ToMarkdown(sampleHTML)
	}
	return
}
