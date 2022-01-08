package computeoffering

import "terraform-provider-stackbill/http"

var (
	httpClient             http.HttpClient      = http.NewHttpClient()
	computeOfferingApiObj  ComputeOfferingApiI  = NewComputeOfferingApi()
	computeOfferingDataObj ComputeOfferingDataI = NewComputeOfferingData()
)
