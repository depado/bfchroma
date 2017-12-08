# bfchroma

![Go Version](https://img.shields.io/badge/go-1.8-brightgreen.svg)
![Go Version](https://img.shields.io/badge/go-1.9-brightgreen.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Depado/bfchroma)](https://goreportcard.com/report/github.com/Depado/bfchroma)
[![Build Status](https://drone.depado.eu/api/badges/Depado/bfchroma/status.svg)](https://drone.depado.eu/Depado/bfchroma)
[![codecov](https://codecov.io/gh/Depado/bfchroma/branch/master/graph/badge.svg)](https://codecov.io/gh/Depado/bfchroma)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/bfchroma/blob/master/LICENSE)
[![Godoc](https://godoc.org/github.com/Depado/bfchroma?status.svg)](https://godoc.org/github.com/Depado/bfchroma)


Integrating [Chroma](https://github.com/alecthomas/chroma) syntax highlighter as a [Blackfriday](https://github.com/russross/blackfriday) renderer

## Install

`go get -u github.com/Depado/bfchroma`

## Features

This renderer integrates chroma to highlight code with triple backtick notation.
It will try to use the given language when available otherwise it will try to
detect the language. If none of these two method works it will fallback to sane
defaults.

## Usage

bfchroma uses the functional options approach so you can customize the behavior
of the renderer. It uses sane defaults when no option is passed so you can use
the renderer simply by doing so :

```go
html := bf.Run([]byte(md), bf.WithRenderer(bfchroma.NewRenderer()))
```

### Options

- `Style(s string)`  
Define the style used by chroma for the rendering. The full list can be found [here](https://github.com/alecthomas/chroma/tree/master/styles)
- `ChromaStyle(*chroma.Style)`  
This option can be used to passe directly a `*chroma.Style` instead of the 
string representing the style as with the `Style(string)` option. 
- `WithoutAutodetect()`  
By default when no language information is written in the code block, this 
renderer will try to auto-detect the used language. This option disables
this behavior and will fallback to a sane default when no language
information is avaiable.
- `Extend(bf.Renderer)`  
This option allows to define the base blackfriday that will be extended.
- `ChromaOptions(...html.Option)`  
This option allows you to pass Chroma's html options in the renderer. Such
options can be found [here](https://github.com/alecthomas/chroma#the-html-formatter).
There is currently an issue with the `html.WithClasses()` option as it expects
the CSS classes to be written separately. I'll come up with a fix later.

### Option examples

Disabling language auto-detection and displaying line numbers

```go
r := bfchroma.NewRenderer(
	bfchroma.WithoutAutodetect(),
	bfchroma.ChromaOptions(html.WithLineNumbers()),
)
```

Extend a blackfriday renderer

```go
b := bf.NewHTMLRenderer(bf.HTMLRendererParameters{
	Flags: bf.CommonHTMLFlags,
})

r := bfchroma.NewRenderer(bfchroma.Extend(b))
```

Use a different style

```go
r := bfchroma.NewRenderer(bfchroma.Style("dracula"))
// Or
r = bfchroma.NewRenderer(bfchroma.ChromaStyle(styles.Dracula))
```

## Examples

```go
package main

import (
	"fmt"

	"github.com/Depado/bfchroma"

	bf "gopkg.in/russross/blackfriday.v2"
)

var md = "This is some sample code.\n\n```go\n" +
	`func main() {
	fmt.Println("Hi")
}
` + "```"

func main() {
	html := bf.Run([]byte(md), bf.WithRenderer(bfchroma.NewRenderer()))
	fmt.Println(string(html))
}
```


Will output :

```html
<p>This is some sample code.</p>
<pre style="color:#f8f8f2;background-color:#272822"><span style="color:#66d9ef">func</span> <span style="color:#a6e22e">main</span>() {
<span style="color:#a6e22e">fmt</span>.<span style="color:#a6e22e">Println</span>(<span style="color:#e6db74">&#34;Hi&#34;</span>)
}
</pre>
```

## Real-life example

In [smallblog](https://github.com/Depado/smallblog) I'm using bfchroma to render
my articles. It's using a combination of both bfchroma's options and blackfriday
extensions and flags.

```go
package main

import (
	"github.com/Depado/bfchroma"
	"github.com/alecthomas/chroma/formatters/html"
	bf "github.com/russross/blackfriday"
)

// Defines the extensions that are used
var exts = bf.NoIntraEmphasis | bf.Tables | bf.FencedCode | bf.Autolink |
	bf.Strikethrough | bf.SpaceHeadings | bf.BackslashLineBreak |
	bf.DefinitionLists | bf.Footnotes

// Defines the HTML rendering flags that are used
var flags = bf.UseXHTML | bf.Smartypants | bf.SmartypantsFractions |
	bf.SmartypantsDashes | bf.SmartypantsLatexDashes | bf.TOC

// render will take a []byte input and will render it using a new renderer each
// time because reusing the same can mess with TOC and header IDs
func render(input []byte) []byte {
	return bf.Run(
		input,
		bf.WithRenderer(
			bfchroma.NewRenderer(
				bfchroma.WithoutAutodetect(),
				bfchroma.ChromaOptions(
					html.WithLineNumbers(),
				),
				bfchroma.Extend(
					bf.NewHTMLRenderer(bf.HTMLRendererParameters{
						Flags: flags,
					}),
				),
			),
		),
		bf.WithExtensions(exts),
	)
}
```


## ToDo

- [ ] Allow the use of `html.WithClasses()` 
