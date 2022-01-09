package iso

import "terraform-provider-stackbill/http"

var (
	httpClient http.HttpClient = http.NewHttpClient()
	isoApiObj  IsoApiI         = NewIsoApi()
	isoDataObj IsoDataI        = NewIsoData()
)
