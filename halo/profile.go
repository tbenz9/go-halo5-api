package halo

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func (h *Halo) EmblemImage(player string, size int) []byte {
	url, err := url.Parse(
		fmt.Sprintf("%s/profile/%s/profiles/%s",
			h.baseurl,
			h.title,
			player,
		))
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	url.RawQuery = q.Encode()
	response, err := h.sendRequest(url.String())
	return response
}

func (h *Halo) SpartanImage(player string, size int, crop string) []byte {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/spartan", h.baseurl, h.title, player))
	if err != nil {
		log.Fatal("Error parsing URL: ", err)
	}
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", string(size))
	}
	if (strings.ToLower(crop) == "full") || (strings.ToLower(crop) == "portrait") {
		q.Set("crop", crop)
	}
	url.RawQuery = q.Encode()
	response, err := h.sendRequest(url.String())
	return response
}
