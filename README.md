# blackfriday-chroma
Integrating Chroma syntax highlighter as a blackfriday renderer

## Usage

```go
package main

import (
	"fmt"

	"github.com/Depado/blackfriday-chroma"

	bf "gopkg.in/russross/blackfriday.v2"
)

var md = "This is some sample code.\n\n```go\n" +
	`func main() {
fmt.Println("Hi")
}
` + "```"

func main() {
	r := bfchroma.ChromaRenderer{
		Base: bf.NewHTMLRenderer(bf.HTMLRendererParameters{
			Flags: bf.CommonHTMLFlags,
		}),
		Style: "monokai",
	}
	html := bf.Run([]byte(md), bf.WithRenderer(&r))
	fmt.Println(string(html))
}
```
