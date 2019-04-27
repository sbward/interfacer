package interfaces_test

import (
	"testing"

	"github.com/sbward/interfaces"
)

func TestParseQuery(t *testing.T) {
	cases := map[string]*interfaces.Query{
		`os.File`: {
			Package:  "os",
			TypeName: "File",
		},
		`github.com/sbward/interfaces.Query`: {
			Package:  "github.com/sbward/interfaces",
			TypeName: "Query",
		},
	}
	for raw, query := range cases {
		q, err := interfaces.ParseQuery(raw)
		if err != nil {
			t.Errorf("ParseQuery(%q)=%v", raw, err)
			continue
		}
		if q.Package != query.Package {
			t.Errorf("ParseQuery(%q): want package=%q; got %q", raw, query.Package, q.Package)
		}
		if q.TypeName != query.TypeName {
			t.Errorf("ParseQuery(%q): want type=%q; got %q", raw, query.TypeName, q.TypeName)
		}
	}
}
