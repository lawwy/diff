package diff

import (
	_strings "strings"
)

func HtmlTokenize(html string) []string {
	tokens := []string{}
	stack := []string{}
	for _, c := range _strings.Split(html, "") {
		if len(stack) > 0 && c != ">" {
			stack = append(stack, c)
			continue
		}
		if c == "<" {
			stack = append(stack, c)
			continue
		}
		if c == ">" {
			stack = append(stack, c)
			tokens = append(tokens, _strings.Join(stack, ""))
			stack = []string{}
			continue
		}
		tokens = append(tokens, c)
	}
	return tokens
}

type htmls struct{ a, b []string }

func (d *htmls) Equal(i, j int) bool { return d.a[i] == d.b[j] }

func (d *htmls) Translate(changes []Change) []Part {
	// var x,y int
	parts := []Part{}
	x, y := 0, 0
	for _, c := range changes {
		if x < c.A && y < c.B {
			v := _strings.Join(d.a[x:c.A], "")
			parts = append(parts, Part{v, false, false})
		}
		if c.Del > 0 {
			v := _strings.Join(d.a[c.A:c.A+c.Del], "")
			parts = append(parts, Part{v, true, false})
		}
		if c.Ins > 0 {
			v := _strings.Join(d.b[c.B:c.B+c.Ins], "")
			parts = append(parts, Part{v, false, true})
		}
		x = c.A + c.Del
		y = c.B + c.Ins
	}
	if x < len(d.a) {
		v := _strings.Join(d.a[x:], "")
		parts = append(parts, Part{v, false, false})
	}
	return parts
}

func DiffHtmls(h1, h2 string) []Part {
	a := HtmlTokenize(h1)
	b := HtmlTokenize(h2)
	htmls := &htmls{a, b}
	diff := Diff(len(a), len(b), htmls)
	return htmls.Translate(diff)
}
