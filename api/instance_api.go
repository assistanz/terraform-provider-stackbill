package api

var (
	instancePostApi = "instance/createInstance"
)

func GetInstanceCreateApi() string {
	return EndPoint + "/" + instancePostApi
}
