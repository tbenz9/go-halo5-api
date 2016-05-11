package halo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

func (h *Halo) metadataRequest(datatype, id string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s/metadata/%s/metadata/%s/%s", h.baseurl, h.title, datatype, id))
	if err != nil {
		return nil, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	response, _ := h.sendRequest(url.String())
	return response, nil
}

func (h *Halo) sendRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Ocp-Apim-Subscription-Key", h.apikey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Return the URL of the image for SpartanImage and EmblemImage
	if url != response.Request.URL.String() {
		return []byte(response.Request.URL.String()), nil
	}

	// If its not SpartanImage or EmblemImage return the body
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func verifyValidID(id string) error {
	re, _ := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}$")
	if !re.MatchString(id) {
		return fmt.Errorf("Not a Valid id: ", id)
	}
	return nil
}
