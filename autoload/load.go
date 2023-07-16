package autoload

import (
	"os"

	"github.com/lengzhao/font"
)

func init() {
	if os.Getenv("FYNE_FONT") != "" {
		return
	}
	files := font.FindFontFile("")
	if len(files) > 0 {
		os.Setenv("FYNE_FONT", files[0])
	}
}
