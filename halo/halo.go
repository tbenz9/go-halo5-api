package halo

type Halo struct {
	baseurl   string
	title     string
	apikey    string
	callLimit int
}

func NewHalo(baseurl, title, apikey string, callLimit int) *Halo {
	h := &Halo{
		baseurl:   baseurl,
		title:     title,
		apikey:    apikey,
		callLimit: callLimit,
	}
	return h
}
