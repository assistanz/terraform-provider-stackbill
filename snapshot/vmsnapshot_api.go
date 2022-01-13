package snapshot

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// Snapshot Object
func NewVmSnapshotApi() VmSnapshotApi {
	return &vmSnapshotApi{}
}

// Vm Snapshot Interfae
type VmSnapshotApi interface {
	SnapshotVm(VmSnapshotRequest, interface{}) (string, error)
	ListVmSnapshots(string, string, interface{}) (string, error)
	DeleteVmSnapshot(string, interface{}) (string, error)
}

// Vm Snapshot utils Object
type vmSnapshotApi struct {
}

// Update Instance Name
// TODO - Documentation
func (vs *vmSnapshotApi) SnapshotVm(vr VmSnapshotRequest, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetVmSnapshotCreateApi()
	response, err := httpClient.PostJson(endPoint, apiKey, secretKey, vr)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Delete Instance
// TODO - Documentation
func (vs *vmSnapshotApi) DeleteVmSnapshot(uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetVmSnapshotDeleteApi(uuid)
	response, err := httpClient.Delete(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}

// Delete Instance
// TODO - Documentation
func (vs *vmSnapshotApi) ListVmSnapshots(zoneUuid string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetVmSnapshotListApi(zoneUuid)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return response, nil
}