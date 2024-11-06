package writer

import (
	"runtime"
)

var NewLine string

func init() {
	if runtime.GOOS == "windows" {
		NewLine = "\r\n"
	} else {
		NewLine = "\n"
	}
}
