package topwords

import (
	"testing"
)

func TestTopWords(t *testing.T) {

	testTable := []struct {
		text    string
		n       int
		expects []Word
	}{
		{"a a. a!b b!c", 2, []Word{
			{"a", 3},
			{"b", 2},
		}},
		{"h he he !hel hel. hel hello hello? hello     hello hello", 3, []Word{
			{"hello", 5},
			{"hel", 3},
			{"he", 2},
		}},
	}

	for _, test := range testTable {
		res := TopWords(test.text, test.n)
		exps := test.expects
		for i := range res {

			if res[i].N != exps[i].N {
				t.Errorf(`Incorrect result. For "%s" expect: %v, got: %v`, res[i].W, exps[i].N, res[i].N)
			}
		}
	}
}
