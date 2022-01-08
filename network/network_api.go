package network

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Network Api
func NewNetworkApi() NetworkApiI {
	return &networkApi{}
}

// Network api interface
type NetworkApiI interface {
	ListNetworks(string, string, interface{}) (string, error)
}

// Network api object
type networkApi struct {
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListNetworks(zoneId string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNewNetworkListApi(zoneId)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
