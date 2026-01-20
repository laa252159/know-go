package printme

import (
	"fmt"
	"io"
)

// Your PrintAnythingTo function goes here!

func PrintAnythingTo[T any](w io.Writer, p T) {
	bytes, err := fmt.Fprint(w, p)
	fmt.Println(bytes, err)
	if err != nil {
		return
	}
}
