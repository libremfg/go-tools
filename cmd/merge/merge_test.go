package main_test

import (
	"os"
	"testing"

	merge "github.com/libremfg/go-tools/cmd/merge"
)

func TestMerge(t *testing.T) {
	d := "./test/"
	e := []string{".txt"}
	o := "./output"
	v := false

	t.Cleanup(func() {
		_, err := os.Stat(o)
		if err == nil {
			os.Remove(o)
		}
	})

	merge.Merge(d, e, o, v)

	expect := []byte("1\r\n2\r\n3\r\n")

	actual, err := os.ReadFile(o)
	if err != nil {
		t.Error(err)
	}

	if len(expect) != len(actual) {
		t.Errorf("lengths don't match expect %d got %d", len(expect), len(actual))
		return
	}

	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			t.Errorf("index %d don't match expect %b got %b", i, expect[i], actual[i])
		}
	}
}
