package api

var (
	securityGroupApiList = "securitygroup/securityList"
)

// Security Group Api
func GetSecurityGroupListApi() string {
	return EndPoint + "/" + securityGroupApiList
}
