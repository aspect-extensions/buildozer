package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/bazelbuild/buildtools/edit"
)

//go:wasmexport edit
func Edit(command string, label string) {
	opts := edit.NewOpts()
	edit.Buildozer(opts, []string{
		command,
		label,
	})
}

//go:wasmexport alloc
func Alloc(len uint32) unsafe.Pointer {
	sl := make([]uint8, len)
	return unsafe.Pointer(&sl[0])
}

func main() {
	fmt.Println(os.Getwd())
}
