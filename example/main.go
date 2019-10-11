package main

import (
	"fmt"

	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"

	bf "github.com/russross/blackfriday/v2"
)

var md = "This is some sample code.\n\n```go\n" +
	`func main() {
	fmt.Println("Hi")
}
` + "```"

func main() {
	var r *bfchroma.Renderer
	var h []byte

	// Basic usage
	r = bfchroma.NewRenderer()
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))

	// Option examples and extending a specific blackfriday renderer
	b := bf.NewHTMLRenderer(bf.HTMLRendererParameters{
		Flags: bf.CommonHTMLFlags,
	})
	r = bfchroma.NewRenderer(
		bfchroma.WithoutAutodetect(),
		bfchroma.Extend(b),
		bfchroma.ChromaOptions(html.WithLineNumbers()),
	)
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))

	md := "```\npackage main\n\nfunc main() {\n}\n```"
	r = bfchroma.NewRenderer(bfchroma.WithoutAutodetect())
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))

	md = "```go\npackage main\n\nfunc main() {\n}\n```"
	r = bfchroma.NewRenderer(bfchroma.ChromaOptions(html.WithLineNumbers()))
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}
