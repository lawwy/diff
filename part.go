package diff

type Part struct {
	Value   string `json:"value"`
	Removed bool   `json:"removed"`
	Added   bool   `json:"added"`
}

func (d *strings) Translate(changes []Change) []Part {
	// var x,y int
	parts := []Part{}
	x, y := 0, 0
	for _, c := range changes {
		if x < c.A && y < c.B {
			v := d.a[x:c.A]
			parts = append(parts, Part{v, false, false})
		}
		if c.Del > 0 {
			v := d.a[c.A : c.A+c.Del]
			parts = append(parts, Part{v, true, false})
		}
		if c.Ins > 0 {
			v := d.b[c.B : c.B+c.Ins]
			parts = append(parts, Part{v, false, true})
		}
		x = c.A + c.Del
		y = c.B + c.Ins
	}
	if x < len(d.a) {
		v := d.a[x:]
		parts = append(parts, Part{v, false, false})
	}
	return parts
}

func DiffStrings(a, b string) []Part {
	ss := &strings{a, b}
	diff := Diff(len(a), len(b), ss)
	return ss.Translate(diff)
}
