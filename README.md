# font

autoload font file for fyne app

zh: 自动加载系统中的字体文件

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    _ "github.com/lengzhao/font/autoload"
)

func main() {
    a := app.New()
    w := a.NewWindow("Hello")

    hello := widget.NewLabel("Hello Fyne!你好")
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("Hi!", func() {
            hello.SetText("Welcome :)")
        }),
    ))

    w.ShowAndRun()
}
```

## description

FindFontFile: finds font files based on the specified language.

It takes a language code as a parameter and returns a slice of strings representing the file paths of the font files.
If no font files are found, an error is returned.

params: language string, e.g. "en", "fr", "de", etc. if language is "", the system language is used.

See <https://www.microsoft.com/typography/otspec/languagetags.htm>

```go
package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/lengzhao/font"
)

func main() {
    files := font.FindFontFile("ZHS")
    if len(files) > 0 {
        os.Setenv("FYNE_FONT", files[0])
    }

    a := app.New()
    w := a.NewWindow("Hello")

    hello := widget.NewLabel("Hello Fyne!你好")
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("Hi!", func() {
            hello.SetText("Welcome :)")
        }),
    ))

    w.ShowAndRun()
}
```
