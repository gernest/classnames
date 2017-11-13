package classnames

// Name represent a css class. If Skip is set to true this class will be
// ignored.
type Name struct {
	Class string
	Skip  bool
}

// Join joins the names to form css classes.
func Join(names ...interface{}) string {
	c := &class{}
	for _, name := range names {
		switch t := name.(type) {
		case string:
			if t == "" {
				continue
			}
			c.add(t)
		case Name:
			if t.Skip || t.Class == "" {
				continue
			}
			c.add(t.Class)
		}
	}
	return c.String()
}

type class struct {
	n []string
}

func (c *class) add(cl string) {
	c.n = append(c.n, cl)
}

func (c *class) String() string {
	buf := ""
	for k, v := range c.n {
		if k == 0 {
			buf += v
		} else {
			buf += " " + v
		}
	}
	return buf
}
