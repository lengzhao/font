package font

import (
	"fmt"
	"os"
	"strings"

	"github.com/ConradIrwin/font/sfnt"
	"github.com/flopp/go-findfont"
	"github.com/jeandeaual/go-locale"
)

// FindFontFile finds font files based on the specified language.
//
// It takes a language code as a parameter and returns a slice of strings representing the file paths of the font files.
// If no font files are found, an error is returned.
//
// params: language string, e.g. "en", "fr", "de", etc.
//
//	if language is "", the system language is used.
//
// See https://www.microsoft.com/typography/otspec/languagetags.htm
func FindFontFile(lg string) []string {
	var out []string
	language := lg
	if len(language) == 0 {
		var err error
		language, err = locale.GetLanguage()
		if err != nil {
			return nil
		}
	}
	language = strings.ToUpper(language)
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if check(language, path) == nil {
			out = append(out, path)
		}
	}
	return out
}

func check(language, fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	ft, err := sfnt.Parse(f)
	if err != nil {
		return err
	}
	// fmt.Println(path, ft.String())
	tl, err := ft.GposTable()
	if err == nil {
		for _, it := range tl.Scripts {
			for _, sit := range it.Languages {
				if strings.HasPrefix(sit.Tag.String(), language) {
					return nil
				}
			}
		}
	}
	tl, err = ft.GsubTable()
	if err == nil {
		for _, it := range tl.Scripts {
			for _, sit := range it.Languages {
				if strings.HasPrefix(sit.Tag.String(), language) {
					return nil
				}
			}
		}
	}
	return fmt.Errorf("not founc")
}
