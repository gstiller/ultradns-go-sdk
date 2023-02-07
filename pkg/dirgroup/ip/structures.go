package ip

type AccGroupGeoIP struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Codes       []string `json:"codes"`
}

type Response struct {
	AccGroupGeoIP
}
