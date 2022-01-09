package api

var (
	sshKeyList  = "sshkey/sshkeyList"
	resetSshKey = "instance/resetSSHkey"
)

// SshKey List api
func GetSshkeyListApi() string {
	return EndPoint + "/" + sshKeyList
}

// SshKey List api
func GetResetSshkeyApi(uuid string, sshKeyId string) string {
	return EndPoint + "/" + resetSshKey + "?uuid=" + uuid + "&sshKeyId=" + sshKeyId
}
