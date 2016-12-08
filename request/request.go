package request

import (
	"io/ioutil"
	"net/http"
)

func GetBodyString(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
