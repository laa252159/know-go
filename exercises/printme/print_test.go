package printme_test

import (
	"bytes"
	"strconv"
	"testing"

	printme "github.com/bitfield/print"
)

func TestPrintAnythingTo_PrintsToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	printme.PrintAnythingTo(buf, "Hello, world\n")
	got := buf.String()
	want := "Hello, world\n"
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	want2 := 1

	buf2 := new(bytes.Buffer)
	printme.PrintAnythingTo(buf2, want2)
	got2, err := strconv.Atoi(buf2.String())
	if err == nil && want2 != got2 {
		t.Errorf("want2 %q, got %q", want2, got2)
	}
}
