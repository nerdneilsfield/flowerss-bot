package translator

import (
	"net/http"
)

func r(e string, t string) {
	for i := 0; i < len(t); i=i+3 {
		rr := t[i + 2];
		if rr =
	}
}

func googleTranslate(data string) string {
	// new a http client
	client := &http.Client{}
	// new a post request
	req, err := http.NewRequest("POST", "https://translate.google.cn/translate_a/single", nil)
	if err != nil {
		return ""
	}
}
