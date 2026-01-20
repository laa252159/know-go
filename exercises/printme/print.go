package printme

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!

func PrintAnythingTo[T any](w io.Writer, p T) {
	_, _ = fmt.Fprint(w, p)
}
