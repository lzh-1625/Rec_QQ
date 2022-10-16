package pkg

import (
	"io"
	"net/http"
	"os"
)


func Download(url,name string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./Downloads/"+name)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}
