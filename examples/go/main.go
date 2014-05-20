package main

import (
	"fmt"
	"github.com/kobeld/gopandoc"
	pandora "github.com/theplant/pandora/clients/go"
)

func main() {
	md2Html()
	Html2Md()
	runParallelExample()
}

func md2Html() {
	html := `
       <ul>
        <li>
                你好
        </li>
        <li>
                你妹
        </li>
       </ul>
<p>好不好看不重要。有钱有权紧要点。这样就可以威胁他们拿钱来换了。<br />。。不知不觉间又发现了一条发家致富的道路了。也和还可以和习大大合作，有选择性地反腐反贪。</p><p>当然这一切都建立在群策成功打入政府市场。。。</p><p>意淫无罪，意淫万岁。</p>
       `

	md, _ := pandora.ToMarkdown(html)
	fmt.Println("------ HTML convert to markdown------")
	fmt.Println("HTML:")
	fmt.Println(html)
	fmt.Println("")
	fmt.Println("Markdown:")
	fmt.Println(md)
}

func Html2Md() {
	md := `- foo
- bar `
	html, _ := pandora.ToHTML(md)
	fmt.Println("------ markdown convert to HTML------")
	fmt.Println("Markdown:")
	fmt.Println(md)
	fmt.Println("")
	fmt.Println("HTML:")
	fmt.Println(html)
}

func runParallelExample() {
	parallel(pandora.ToMarkdown)
	parallel(gopandoc.ToMarkdown)
	return
}

func parallel(f func(string) (string, error)) {
	mdChan := make(chan string, 100)
	parallelCount := 2
	sampleHTML := `
	<ul>
	<li>
	你好
	</li>
	<li>
	你妹
	</li>
	</ul>
	`

	go func() {
		for i := 0; i < parallelCount; i++ {
			mdText, _ := f(sampleHTML)
			mdChan <- mdText
		}
		close(mdChan)
	}()

	for mdText := range mdChan {
		println(mdText)
	}

	return
}
