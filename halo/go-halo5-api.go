package halo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (h *Halo) metadataRequest(datatype, id string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s/metadata/%s/metadata/%s/%s", h.baseurl, h.title, datatype, id))
	if err != nil {
		return nil, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	response := h.sendRequest(url.String())
	return response, nil
}

func (h *Halo) sendRequest(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	request.Header.Set("Ocp-Apim-Subscription-Key", h.apikey)

	response, err := http.DefaultClient.Do(request)
	checkErr(err)
	defer response.Body.Close()

	// Return the URL of the image for SpartanImage and EmblemImage
	if url != response.Request.URL.String() {
		return []byte(response.Request.URL.String())
	}

	// If its not SpartanImage or EmblemImage return the body
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	return contents
}

func sendRequest(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	request.Header.Set("Ocp-Apim-Subscription-Key", getAPIKey())

	response, err := http.DefaultClient.Do(request)
	checkErr(err)
	defer response.Body.Close()

	// Return the URL of the image for SpartanImage and EmblemImage
	if url != response.Request.URL.String() {
		return []byte(response.Request.URL.String())
	}

	// If its not SpartanImage or EmblemImage return the body
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)

	return contents
}

func getAPIKey() string {
	apikey := os.Getenv("HALO_API_KEY")
	if len(apikey) != 32 {
		fmt.Println("Invalid API Key")
	}
	return apikey
}

func verifyValidID(ID, name string) {
	re, _ := regexp.Compile("^[a-z0-9]{8}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{4}-[a-z0-9]{12}$")
	valid := re.MatchString(ID)
	if valid == false {
		fmt.Printf("%s is not a valid %s\n", ID, name)
	}
}
