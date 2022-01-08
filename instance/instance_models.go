package instance

// Create instance request
type CreateRequest struct {
	ComputeOfferingUuid string `json:"computeOfferingUuid" validate:"required"`
	CpuCore             string `json:"cpuCore"`
	DiskSize            int    `json:"diskSize"`
	HypervisorName      string `json:"hypervisorName"`
	Memory              string `json:"memory"`
	Name                string `json:"name" validate:"required"`
	NetworkUuid         string `json:"networkUuid" validate:"required"`
	SecuritygroupName   string `json:"securitygroupName"`
	SshKeyName          string `json:"sshKeyName"`
	StorageOfferingUuid string `json:"storageOfferingUuid"`
	TemplateUuid        string `json:"templateUuid" validate:"required"`
	ZoneUuid            string `json:"zoneUuid" validate:"required"`
}

// Update Request
type UpdateNameRequest struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"displayName" validate:"required"`
}

// Resize VM request
type ResizeRequest struct {
	Uuid         string `json:"uuid" validate:"required"`
	OfferingUuid string `json:"offeringUuid" validate:"required"`
	CpuCore      string `json:"cpuCore" validate:"required"`
	Memory       string `json:"memory" validate:"required"`
}

// Action Request
type InstanceActionRequest struct {
	Uuid   string `json:"uuid" validate:"required"`
	Action string `json:"action" validate:"required"`
}
