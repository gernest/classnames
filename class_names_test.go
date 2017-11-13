package classnames

import "testing"

func TestJoin(t *testing.T) {

	sample := []struct {
		names  []interface{}
		expect string
	}{
		{
			names: []interface{}{
				Name{Class: "a"}, Name{Class: "b", Skip: true}, Name{Class: "f"},
			},
			expect: "a f",
		},
		{
			names: []interface{}{
				Name{Class: ""}, Name{Class: "b"}, "",
			},
			expect: "b",
		},
		{
			names:  []interface{}{},
			expect: "",
		},
	}

	for _, s := range sample {
		n := Join(s.names...)
		if n != s.expect {
			t.Errorf("expected %s got %s", s.expect, n)
		}
	}
}
