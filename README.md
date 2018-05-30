wkhtmltoimage
-------------
Go binding of [wkhtmltoimage](http://wkhtmltopdf.org). Generate image from html documents. 

If you need PDF generation binding, please visit https://github.com/adrg/go-wkhtmltopdf

## Sample
```
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	img "github.com/bbirec/wkhtmltoimage"
)

func conv(url string, out string) {
	converter, err := img.NewConverter(map[string]string{
		"in":  url,
		"fmt": "png",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

	output, err := converter.Convert()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	f.Write(output)
}

func main() {
	img.Init()
	defer img.Destroy()

	conv("https://beep-notice.firebaseapp.com/ko/billing_popup.html", "ko.png")
	conv("https://beep-notice.firebaseapp.com/en/billing_popup.html", "en.png")
}

```