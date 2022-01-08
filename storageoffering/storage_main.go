package storageoffering

import "terraform-provider-stackbill/http"

var (
	httpClient             http.HttpClient      = http.NewHttpClient()
	storageOfferingApiObj  StorageOfferingApiI  = NewStorageOfferingApi()
	storageOfferingDataObj StorageOfferingDataI = NewStorageOfferingData()
)
