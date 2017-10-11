package bfchroma

import (
	"testing"

	"github.com/alecthomas/chroma/styles"

	"github.com/stretchr/testify/assert"
	bf "gopkg.in/russross/blackfriday.v2"
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

func TestChromaStyle(t *testing.T) {
	var r *Renderer
	for _, v := range styles.Registry {
		r = NewRenderer(ChromaStyle(v))
		assert.Equal(t, r.Style, v, "Style should match")
	}
}

func TestWithoutAutodetect(t *testing.T) {
	r := NewRenderer(WithoutAutodetect())
	assert.False(t, r.Autodetect, "Should set Autodetect to false")
	r = NewRenderer()
	assert.True(t, r.Autodetect, "Not using option should leave Autodetect to true")
}
