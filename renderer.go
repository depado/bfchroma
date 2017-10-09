package bfchroma

import (
	"io"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"

	bf "gopkg.in/russross/blackfriday.v2"
)

func (r *ChromaRenderer) renderWithChroma(w io.Writer, text []byte, data bf.CodeBlockData) error {
	var lexer chroma.Lexer
	if len(data.Info) > 0 {
		lexer = lexers.Get(string(data.Info))
	} else {
		lexer = lexers.Analyse(string(text))
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}
	cstyle := styles.Get(r.Style)
	if cstyle == nil {
		cstyle = styles.Fallback
	}
	formatter := html.New()
	iterator, err := lexer.Tokenise(nil, string(text))
	if err != nil {
		return err
	}
	return formatter.Format(w, cstyle, iterator)
}

// ChromaRenderer is a custom Blackfriday renderer that uses the capabilities of
// chroma to highlight code with triple backtick notation
type ChromaRenderer struct {
	Base  *bf.HTMLRenderer
	Style string
}

// RenderNode satisfies the Renderer interface
func (r *ChromaRenderer) RenderNode(w io.Writer, node *bf.Node, entering bool) bf.WalkStatus {
	switch node.Type {
	case bf.CodeBlock:
		if err := r.renderWithChroma(w, node.Literal, node.CodeBlockData); err != nil {
			return r.Base.RenderNode(w, node, entering)
		}
		return bf.SkipChildren
	default:
		return r.Base.RenderNode(w, node, entering)
	}
}

// RenderHeader satisfies the Renderer interface
func (r *ChromaRenderer) RenderHeader(w io.Writer, ast *bf.Node) {
	r.Base.RenderHeader(w, ast)
}

// RenderFooter satisfies the Renderer interface
func (r *ChromaRenderer) RenderFooter(w io.Writer, ast *bf.Node) {
	r.Base.RenderFooter(w, ast)
}
