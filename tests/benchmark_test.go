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

func pushHTML(htmlText string, out chan string) {
	msg, _ := pandora.ToMarkdown(sampleHTML)
	out <- msg
	return
}

func parallel(f func(string) (string, error)) {
	mdChan := make(chan string)
	parallelCount := 100

	go func() {
		for i := 0; i < parallelCount; i++ {
			msg, _ := f(sampleHTML)
			mdChan <- msg
		}
		close(mdChan)
	}()

	for mdText := range mdChan {
		dummy(mdText)
	}
	return
}

func dummy(text string) {
	return
}

func BenchmarkToMarkdownParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallel(pandora.ToMarkdown)
	}
	return
}

func BenchmarkGopandocToMarkdownParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallel(gopandoc.ToMarkdown)
	}
	return
}
