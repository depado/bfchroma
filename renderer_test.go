package bfchroma

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"
	bf "github.com/russross/blackfriday/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtend(t *testing.T) {
	var b bf.Renderer
	var r *Renderer
	b = bf.NewHTMLRenderer(bf.HTMLRendererParameters{
		Flags: bf.CommonHTMLFlags,
	})
	r = NewRenderer(Extend(b))
	assert.Equal(t, r.Base, b, "should be the same renderer")
}

func ExampleExtend() {
	md := "```go\npackage main\n\nfunc main() {\n}\n```"

	b := bf.NewHTMLRenderer(bf.HTMLRendererParameters{
		Flags: bf.CommonHTMLFlags,
	})
	r := NewRenderer(Extend(b))

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}

func TestStyle(t *testing.T) {
	var r *Renderer
	for k, v := range styles.Registry {
		r = NewRenderer(Style(k))
		assert.Equal(t, r.Style, v, "Style should match")
	}
	for _, v := range []string{"♥", "inexistent", "fallback!"} {
		r = NewRenderer(Style(v))
		assert.Equal(t, r.Style, styles.Fallback)
	}
}

func ExampleStyle() {
	md := "```go\npackage main\n\nfunc main() {\n}\n```"

	r := NewRenderer(Style("github"))

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}

func TestChromaStyle(t *testing.T) {
	var r *Renderer
	for _, v := range styles.Registry {
		r = NewRenderer(ChromaStyle(v))
		assert.Equal(t, r.Style, v, "Style should match")
	}
}

func ExampleChromaStyle() {
	md := "```go\npackage main\n\nfunc main() {\n}\n```"

	r := NewRenderer(ChromaStyle(styles.Get("github")))

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}

func TestWithoutAutodetect(t *testing.T) {
	r := NewRenderer(WithoutAutodetect())
	assert.False(t, r.Autodetect, "Should set Autodetect to false")
	r = NewRenderer()
	assert.True(t, r.Autodetect, "Not using option should leave Autodetect to true")
}

func ExampleWithoutAutodetect() {
	md := "```\npackage main\n\nfunc main() {\n}\n```"

	r := NewRenderer(WithoutAutodetect())

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}

func TestChromaOptions(t *testing.T) {
	NewRenderer(ChromaOptions(html.WithClasses(true)))
}

func ExampleChromaOptions() {
	md := "```go\npackage main\n\nfunc main() {\n}\n```"

	r := NewRenderer(ChromaOptions(html.WithLineNumbers(true)))

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}

func TestRenderWithChroma(t *testing.T) {
	var err error
	var b *bytes.Buffer
	r := NewRenderer()
	tests := []struct {
		in  []byte
		cbd bf.CodeBlockData
		out string
	}{
		{[]byte{0}, bf.CodeBlockData{}, "<pre style=\"color:#f8f8f2;background-color:#272822;\"><code><span style=\"display:flex;\"><span>\x00</span></span></code></pre>"},
		{[]byte{0, 1, 2}, bf.CodeBlockData{}, "<pre style=\"color:#f8f8f2;background-color:#272822;\"><code><span style=\"display:flex;\"><span>\x00\x01\x02</span></span></code></pre>"},
		{[]byte("Hello World"), bf.CodeBlockData{}, "<pre style=\"color:#f8f8f2;background-color:#272822;\"><code><span style=\"display:flex;\"><span>Hello World</span></span></code></pre>"},
	}
	for _, test := range tests {
		b = new(bytes.Buffer)
		err = r.RenderWithChroma(b, test.in, test.cbd)
		assert.NoError(t, err, "Should not fail")
		assert.Equal(t, test.out, b.String())
	}
}

func TestRender(t *testing.T) {
	md := "```go\npackage main\n\nfunc main() {\n}\n```"
	r := NewRenderer()
	bg := r.Style.Get(chroma.Background).Background.String()

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	assert.Contains(t, string(h), r.Style.Get(chroma.NameFunction).Colour.String())
	assert.Contains(t, string(h), bg)
	assert.Contains(t, string(h), "<pre")

	// Check if auto-detection works on Go example
	md = "```\npackage main\n\nfunc main() {\n\tfmt.Println()\n}\n```"
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	assert.Contains(t, string(h), r.Style.Get(chroma.NameFunction).Colour.String())
	assert.Contains(t, string(h), bg)
	assert.Contains(t, string(h), "<pre")

	// Check if disabling the Autodetect feature works
	r = NewRenderer(WithoutAutodetect())
	h = bf.Run([]byte(md), bf.WithRenderer(r))
	assert.NotContains(t, string(h), r.Style.Get(chroma.NameFunction).Colour.String())
	assert.Contains(t, string(h), bg)
	assert.Contains(t, string(h), "<pre")
}

func TestRender_EmbedCSS(t *testing.T) {
	r := NewRenderer(EmbedCSS())
	h := bf.Run([]byte(""), bf.WithRenderer(r))
	assert.Contains(t, string(h), "<style>")
	assert.Contains(t, string(h), ".chroma")
	assert.Contains(t, string(h), "</style>")
}

func TestRenderer_ChromaCSS(t *testing.T) {
	r := NewRenderer()
	var w bytes.Buffer
	err := r.ChromaCSS(&w)
	require.NoError(t, err)
	assert.Contains(t, w.String(), ".chroma")

}

func ExampleNewRenderer() {
	// Complex example on how to initialize the renderer
	md := "```go\npackage main\n\nfunc main() {\n}\n```"

	r := NewRenderer(
		Extend(bf.NewHTMLRenderer(bf.HTMLRendererParameters{
			Flags: bf.CommonHTMLFlags,
		})),
		WithoutAutodetect(),
		ChromaStyle(styles.Get("github")),
		ChromaOptions(html.WithLineNumbers(true)),
	)

	h := bf.Run([]byte(md), bf.WithRenderer(r))
	fmt.Println(string(h))
}
