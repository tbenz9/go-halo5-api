package halo

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func (h *Halo) EmblemImage(player string, size int) (string, error) {
	url, err := url.Parse(
		fmt.Sprintf("%s/profile/%s/profiles/%s/emblem",
			h.baseurl,
			h.title,
			player,
		))
	if err != nil {
		return "", err
	}
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", strconv.Itoa(size))
	}
	url.RawQuery = q.Encode()
	response, err := h.sendRequest(url.String())
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (h *Halo) SpartanImage(player string, size int, crop string) (string, error) {
	url, err := url.Parse(fmt.Sprintf("%s/profile/%s/profiles/%s/spartan", h.baseurl, h.title, player))
	if err != nil {
		return "", err
	}
	q := url.Query()
	if (size == 95) || (size == 128) || (size == 190) || (size == 256) || (size == 512) {
		q.Set("size", strconv.Itoa(size))
	}
	if (strings.ToLower(crop) == "full") || (strings.ToLower(crop) == "portrait") {
		q.Set("crop", crop)
	}
	url.RawQuery = q.Encode()
	response, err := h.sendRequest(url.String())
	if err != nil {
		return "", err
	}
	return string(response), nil
}
