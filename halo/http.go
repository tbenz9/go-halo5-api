package halo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var callList = make([]int64, 200)

//var callCount int = 0

func (h *Halo) metadataRequest(datatype, id string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s/metadata/%s/metadata/%s/%s", h.baseurl, h.title, datatype, id))
	if err != nil {
		return nil, err
	}
	q := url.Query()
	url.RawQuery = q.Encode()
	response, err := h.sendRequest(url.String())
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (h *Halo) sendRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Ocp-Apim-Subscription-Key", h.apikey)
	currentTime := time.Now().Unix()
	callList = append(callList[:0], callList[1:]...)
	callList = append(callList, currentTime)
	if callList[0]+10 > currentTime {
		sleepTime := (callList[0] + 10) - currentTime
		//		fmt.Println(sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
	//	callCount++
	//	fmt.Println(callCount)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//check for response code
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(response.Status)
	}

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
	re, _ := regexp.Compile("^\\w{8}-\\w{4}-\\w{4}-\\w{4}-\\w{12}$")
	if !re.MatchString(id) {
		return fmt.Errorf("Not a Valid id: ", id)
	}
	return nil
}
