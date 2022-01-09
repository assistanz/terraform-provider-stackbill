package sshkey

import "terraform-provider-stackbill/http"

var (
	httpClient    http.HttpClient = http.NewHttpClient()
	sshkeyApiObj  SshkeyApiI      = NewSshkeyApi()
	sshkeyDataObj SshkeyDataI     = NewSshkeyData()
)
