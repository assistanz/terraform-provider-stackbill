package instance

import "terraform-provider-stackbill/http"

var (
	httpClient        http.HttpClient  = http.NewHttpClient()
	instanceApiObj    InstanceApiI     = NewInstanceApi()
	instanceUtilsObj  InstanceUtilsI   = NewInstanceUtils()
	actionInstanceObj InstanceActionsI = NewInstanceActions()
	instanceDataObj   InstanceDataI    = NewInstanceData()
	instanceObj       InstanceI        = NewInstance()
)
