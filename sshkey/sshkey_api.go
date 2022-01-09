package sshkey

import (
	"terraform-provider-stackbill/api"
	"terraform-provider-stackbill/auth"
)

// New Sshkey Api
func NewSshkeyApi() SshkeyApiI {
	return &networkApi{}
}

// Sshkey api interface
type SshkeyApiI interface {
	ListSshkeys(interface{}) (string, error)
}

// Sshkey api object
type networkApi struct {
}

/* List Sshkeys
@Param - Zone uuid - required
@Param - Sshkey uuid - Optional
@Param - Meta Information - Api key & Secret Key.
It will be generated autmatically By terraform
@Return Response / Error
*/
func (nt *networkApi) ListSshkeys(meta interface{}) (string, error) {
	// Meta information
	m := meta.(*auth.AuthKeys)
	apiKey := m.ApiKey
	secretKey := m.SecretKey
	endPoint := api.GetSshkeyListApi()
	response, err := httpClient.Get(endPoint, apiKey, secretKey)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
