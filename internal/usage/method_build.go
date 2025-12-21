package usage

import "strings"

type methodBuilder struct {
	b *strings.Builder
}

func NewMethod(root string) *methodBuilder {
	b := &strings.Builder{}
	b.Grow(64)
	b.WriteString("minecraft:")
	b.WriteString(root)

	return &methodBuilder{b: b}
}

func (mb *methodBuilder) Add(path string) *methodBuilder {
	if path != "" {
		mb.b.WriteString("/")
		mb.b.WriteString(path)
	}
	return mb
}

func (mb *methodBuilder) String() string {
	return mb.b.String()
}
