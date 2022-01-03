package api

var (
	instancePostApi       = "instance/createInstance"
	instanceNameUpdateApi = "instance/updateName"
	instanceStart         = "instance/startInstance"
	instanceStop          = "instance/stopInstance"
)

// Instance create api
func GetInstanceCreateApi() string {
	return EndPoint + "/" + instancePostApi
}

// Instance name update api
func GetInstanceNameUpdateApi() string {
	return EndPoint + "/" + instanceNameUpdateApi
}

// Start instance api
func GetInstanceStartApi() string {
	return EndPoint + "/" + instanceStart
}

// Stop instance api
func GetInstanceStopApi() string {
	return EndPoint + "/" + instanceStop
}
