package zone

import "terraform-provider-stackbill/http"

var (
	httpClient  http.HttpClient = http.NewHttpClient()
	zoneApiObj  ZoneApiI        = NewZoneApi()
	zoneDataObj ZoneDataI       = NewZoneData()
)
