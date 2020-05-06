// +build !windows

package colors

import (
	"io"
	"os"
)

// NewColorable -
func NewColorable(file *os.File) io.Writer {
	return file
}
