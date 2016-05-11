package halo

type Halo struct {
	baseurl string
	title   string
	apikey  string
}

func NewHalo(baseurl, title, apikey string) *Halo {
	h := &Halo{
		baseurl: baseurl,
		title:   title,
		apikey:  apikey,
	}
	return h
}
