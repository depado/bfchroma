# bfchroma

![Go Version](https://img.shields.io/badge/go-1.8-brightgreen.svg)
![Go Version](https://img.shields.io/badge/go-1.9-brightgreen.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/Depado/bfchroma)](https://goreportcard.com/report/github.com/Depado/bfchroma)
[![Build Status](https://drone.depado.eu/api/badges/Depado/bfchroma/status.svg)](https://drone.depado.eu/Depado/bfchroma)
[![codecov](https://codecov.io/gh/Depado/bfchroma/branch/master/graph/badge.svg)](https://codecov.io/gh/Depado/bfchroma)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/bfchroma/blob/master/LICENSE)


Integrating Chroma syntax highlighter as a blackfriday renderer

## Install

`go get -u github.com/Depado/bfchroma`

## Features

This renderer integrates chroma to highlight code with triple backtick notation.
It will try to use the given language when available otherwise it will try to
detect the language. If none of these two method works it will fallback.

````markdown
```go
fmt.Println("Chroma will use the given language : go")
```

```
fmt.Println("Chroma will auto-detect the go language")
```

```
chroma will fallback to a sane default as it can't determine the used language
```
````

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

```

## Example

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

## ToDo

- [ ] Add tests
- [ ] Use directly `chroma.Style` in the structure ?
- [ ] Allow the use of `html.WithClasses()` 
