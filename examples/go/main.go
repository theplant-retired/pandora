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
