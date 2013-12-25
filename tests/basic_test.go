package pandoratests

import (
	pandora "github.com/theplant/pandora/clients/go"
	"testing"
)

func TestToMD(t *testing.T) {
	html := `
         <ul>       
                <li>foo</li>
                <li>bar</li>
         </ul>       
        `

	expectedMD := `-   foo
-   bar
`

	md, err := pandora.ToMD(html)
	if err != nil {
		t.Error("ToMD get error: ", err)
	}

	if md != expectedMD {
		t.Errorf("ToMD expect:\n %v,\n but actually got:\n %v \n", expectedMD, md)
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
		t.Error("ToMD get error: ", err)
	}

	if html != expectedHTML {
		t.Errorf("ToMD expect:\n |%v|,\n but actually got:\n |%v| \n", expectedHTML, html)
	}

	return
}

