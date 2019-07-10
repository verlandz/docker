package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func StringConcat(str ...string) string {
	var buffer bytes.Buffer
	for _, s := range str {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func GetHttpResponse(url string, header map[string]string) ([]byte, bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, false
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, false
	}

	return bodyBytes, true
}
