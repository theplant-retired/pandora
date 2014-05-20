package pandoratests

import (
	pandora "github.com/theplant/pandora/clients/go"
	"testing"
)

func TestToMarkdownWithCustomOptions(t *testing.T) {
	html := `<p>好不好看不重要。有钱有权紧要点。这样就可以威胁他们拿钱来换了。<br />。。不知不觉间又发现了一条发家致富的道路了。 。 。。`

	expectedMD := "好不好看不重要。有钱有权紧要点。这样就可以威胁他们拿钱来换了。  \n。。不知不觉间又发现了一条发家致富的道路了。 。 。。"

	md, err := pandora.ToMarkdown(html)
	if err != nil {
		t.Error("ToMarkdown get error: ", err)
	}

	if md != expectedMD {
		t.Errorf("ToMarkdown expect:\n %v,\n but actually got:\n %v \n", expectedMD, md)
	}

	return
}

func TestToMarkdown(t *testing.T) {
	html := `
         <ul>
                <li>foo</li>
                <li>bar</li>
         </ul>
        `

	expectedMD := `-   foo
-   bar
`

	md, err := pandora.ToMarkdown(html)
	if err != nil {
		t.Error("ToMarkdown get error: ", err)
	}

	if md != expectedMD {
		t.Errorf("ToMarkdown expect:\n %v,\n but actually got:\n %v \n", expectedMD, md)
	}

	return
}

func TestToHTML(t *testing.T) {
	md := `
- aang
- roku
- kyoshi
- kuruk
        `

	expectedHTML := `<ul>
<li>aang</li>
<li>roku</li>
<li>kyoshi</li>
<li>kuruk</li>
</ul>`

	html, err := pandora.ToHTML(md)
	if err != nil {
		t.Error("ToMarkdown get error: ", err)
	}

	if html != expectedHTML {
		t.Errorf("ToMarkdown expect:\n |%v|,\n but actually got:\n |%v| \n", expectedHTML, html)
	}

	return
}
