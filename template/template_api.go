package template

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Template Api
func NewTemplateApi() TemplateApiI {
	return &templateApi{}
}

// Template api interface
type TemplateApiI interface {
	ListTemplates(string, string, interface{}) (string, error)
}

// Template api object
type templateApi struct {
}

// List Templates
// TODO - Documentation
func (nt *templateApi) ListTemplates(zoneId string, uuid string, meta interface{}) (string, error) {
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