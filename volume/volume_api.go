package volume

import (
	"errors"
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Volume Api
func NewVolumeApi() VolumeApi {
	return &volumeApi{}
}

// Volume api interface
type VolumeApi interface {
	ListVolumes(string, string, interface{}) (string, error)
	VolumeActions(VolumeActionRequest, interface{}) (string, error)
	CreateVolume(VolumeCreateRequest, interface{}) (string, error)
}

// Volume api object
type volumeApi struct {
}

// List Zones
// TODO - Documentation
func (v *volumeApi) ListVolumes(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetVolumeListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Actions -  start / Stop / Restart
// TODO - Documentation
func (v *volumeApi) VolumeActions(vr VolumeActionRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := ""
	switch vr.Action {
	case ATTACH:
		endPoint = api.GetVolumeAttachApi(vr.Uuid) + "&instanceUuid=" + vr.InstanceUuid
	case DETACH:
		endPoint = api.GetVolumeDetachApi(vr.Uuid)
	default:
		return "", errors.New("Invalid action provided...!")
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

// List Zones
// TODO - Documentation
func (v *volumeApi) CreateVolume(vc VolumeCreateRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetVolumeCreateApi()
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, vc)
	if err != nil {
		return "", err
	}
	return response, nil
}
