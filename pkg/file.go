package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FileInfoGet(url string) string{
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return "400"
	}
	reader := response.Body
	defer reader.Close()
	body, _ := ioutil.ReadAll(response.Body)
	AllMap := make(map[string]interface{})
    err = json.Unmarshal(body, &AllMap)
    if err != nil {
        fmt.Println("err")
    }
	return AllMap["data"].(map[string]interface{})["url"].(string)
}
