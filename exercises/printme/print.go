package printme

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!

func PrintAnythingTo[T any](w io.Writer, p T) {
	fprint, err := fmt.Fprint(w, p)
	fmt.Println(fprint, err)
	if err != nil {
		return
	}
}
