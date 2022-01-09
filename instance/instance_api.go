package instance

import (
	"errors"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

const (
	START = "Start"
	STOP  = "Stop"
)

// Instance util Object
func NewInstanceApi() InstanceApiI {
	return &instanceApi{}
}

//Instance Interfae
type InstanceApiI interface {
	CreateInstance(CreateRequest, interface{}) (string, error)
	UpdateInstanceName(UpdateNameRequest, interface{}) (string, error)
	ResetSshKey(ResetSshkeyRequest, interface{}) (string, error)
	ResizeInstance(ResizeRequest, interface{}) (string, error)
	InstanceActions(InstanceActionRequest, interface{}) (string, error)
	ListInstances(string, string, interface{}) (string, error)
	DeleteInstance(string, interface{}) (string, error)
}

//Instance utils Object
type instanceApi struct {
}

// Update Instance Name
// TODO - Documentation
func (ia *instanceApi) CreateInstance(cr CreateRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetInstanceCreateApi()
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, cr)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Update Instance Name
// TODO - Documentation
func (ia *instanceApi) ResetSshKey(rsr ResetSshkeyRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetResetSshkeyApi(rsr.Uuid, rsr.SshkeyId)
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Reset ssh key
// TODO - Documentation
func (ia *instanceApi) UpdateInstanceName(ur UpdateNameRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetInstanceNameUpdateApi()
	response, err := httpClient.PutJson(endPoint, apiKey, secretKey, ur)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Resize VM
// TODO - Documentation
func (vs *instanceApi) ResizeInstance(r ResizeRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	qryString := "?uuid=" + r.Uuid + "&offeringUuid=" + r.OfferingUuid
	qryString += "&memory=" + r.Memory + "&cpuCore=" + r.CpuCore
	endPoint := api.GetInstanceResizeApi() + qryString
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (vs *instanceApi) InstanceActions(ar InstanceActionRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := ""
	switch ar.Action {
	case START:
		endPoint = api.GetInstanceStartApi() + "?uuid=" + ar.Uuid
	case STOP:
		endPoint = api.GetInstanceStopApi() + "?uuid=" + ar.Uuid + "&forceStop=true"
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Delete Instance
// TODO - Documentation
func (vs *instanceApi) DeleteInstance(uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	qryString := "?uuid=" + uuid + "&expunge=true"
	endPoint := api.GetInstanceDeleteApi() + qryString
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// Delete Instance
// TODO - Documentation
func (vs *instanceApi) ListInstances(zoneId string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetInstanceListApi(zoneId)
	if uuid != "" {
		endPoint += "&vmUuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
