package network

import "terraform-provider-stackbill/http"

var (
	httpClient     http.HttpClient = http.NewHttpClient()
	networkApiObj  NetworkApiI     = NewNetworkApi()
	networkDataObj NetworkDataI    = NewNetworkData()
)
