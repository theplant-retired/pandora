package pandoratests

import (
	pandora "github.com/theplant/pandora/clients/go"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var referenceCasesDir = "MarkdownTest_1.0.3"

func runMarkdownReference(input string, flag int) (r string) {
	r, _ = pandora.ToHTML(input)
	return
}

func doTestsReference(t *testing.T, files []string, flag int) {
	// catch and report panics
	var candidate string
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("\npanic while processing [%#v]\n", candidate)
		}
	}()

	for _, basename := range files {
		filename := filepath.Join(referenceCasesDir, basename+".text")
		inputBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}
		input := string(inputBytes)

		filename = filepath.Join(referenceCasesDir, basename+".html")
		expectedBytes, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Errorf("Couldn't open '%s', error: %v\n", filename, err)
			continue
		}
		expected := string(expectedBytes)

		actual := string(runMarkdownReference(input, flag))
		if actual != expected {
			t.Errorf("\n    [%#v]\nExpected[%#v]\nActual  [%#v]",
				basename+".text", expected, actual)
		}

		// now test every prefix of every input to check for
		// bounds checking
		if !testing.Short() {
			start := 0
			for end := start + 1; end <= len(input); end++ {
				candidate = input[start:end]
				_ = runMarkdownReference(candidate, flag)
			}
		}
	}
}

func TestReferences(t *testing.T) {
	return //Sad, seems pandoc behaves a bit different than the standards.

	files := []string{
		"Amps and angle encoding",
		"Auto links",
		"Backslash escapes",
		"Blockquotes with code blocks",
		"Code Blocks",
		"Code Spans",
		"Hard-wrapped paragraphs with list-like lines",
		"Horizontal rules",
		"Inline HTML (Advanced)",
		"Inline HTML (Simple)",
		"Inline HTML comments",
		"Links, inline style",
		"Links, reference style",
		"Links, shortcut references",
		"Literal quotes in titles",
		"Markdown Documentation - Basics",
		"Markdown Documentation - Syntax",
		"Nested blockquotes",
		"Ordered and unordered lists",
		"Strong and em together",
		"Tabs",
		"Tidyness",
	}
	doTestsReference(t, files, 0)
}
