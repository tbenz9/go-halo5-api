package halo

import (
	"fmt"
	"net/url"
	"strings"
)

func EmblemImage(baseurl, title, player string, size int) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s", baseurl, title, player))
	checkErr(err)
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}

func SpartanImage(baseurl, title, player string, size int, crop string) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/spartan", baseurl, title, player))
	checkErr(err)
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	if (strings.ToLower(crop) == "full") || (strings.ToLower(crop) == "portrait") {
		q.Set("crop", crop)
	}
	url.RawQuery = q.Encode()
	response := sendRequest(url.String())
	return response
}
