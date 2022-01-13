package network

import (
	"errors"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Network Api
func NewNetworkApi() NetworkApi {
	return &networkApi{}
}

// Network api interface
type NetworkApi interface {
	ListNetworks(string, string, interface{}) (string, error)
	ListNetworkOfferings(string, string, interface{}) (string, error)
	ListVpcNetworkOfferings(string, interface{}) (string, error)
	CreateNetwork(map[string]interface{}, CreateNetworkRequestOptional, interface{}) (string, error)
	DeleteNetwork(string, interface{}) (string, error)
	NetworkActions(string, NetworkActionRequest, interface{}) (string, error)
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
func (nt *networkApi) ListNetworks(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNewNetworkOfferingListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListNetworkOfferings(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNewNetworkOfferingListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* List Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListVpcNetworkOfferings(uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNewVpcNetworkOfferingListApi()
	if uuid != "" {
		endPoint += "?uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* Create Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) CreateNetwork(cn map[string]interface{}, co CreateNetworkRequestOptional, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNetworkCreateApi(co.IsPublic, co.SecurityGroupId)
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, cn)
	if err != nil {
		return "", err
	}
	return response, nil
}

/* Delete Networks
@Param - Zone uuid - required
@Param - Network uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) DeleteNetwork(uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNetworkDeleteApi(uuid)
	response, err := httpClient.Delete(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (nt *networkApi) NetworkActions(action string, nr NetworkActionRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := ""
	switch action {
	case ADD:
		endPoint = api.GetVmNetworkAddApi()
	case DELETE:
		endPoint = api.GetVmNetworkDeleteApi()
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, nr)
	if err != nil {
		return "", err
	}
	return response, nil
}
