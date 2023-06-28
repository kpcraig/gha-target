package ghtarget

import "testing"

func TestSayHello(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "Alice",
			out: "Hello, Alice!",
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			out := SayHello(c.in)
			if out != c.out {
				t.Fail()
			}
		})
	}
}
