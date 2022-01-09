package iso

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Iso Api
func NewIsoApi() IsoApiI {
	return &isoApi{}
}

// Iso api interface
type IsoApiI interface {
	ListIso(string, string, interface{}) (string, error)
}

// Iso api object
type isoApi struct {
}

// List Isos
// TODO - Documentation
func (nt *isoApi) ListIso(zoneId string, uuid string, meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetNewTemplateListApi(zoneId)
	if uuid != "" {
		endPoint += "&uuid=" + uuid
	}
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
