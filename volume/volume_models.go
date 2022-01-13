package volume

// Volume Action Request
type VolumeActionRequest struct {
	Uuid         string `json:"uuid" validate:"required"`
	InstanceUuid string `json:"instanceUuid" validate:"required"`
	Action       string `json:"action" validate:"required"`
}

// Volume Create Request
type VolumeCreateRequest struct {
	DiskSize            int    `json:"diskSize" validate:"required"`
	Name                string `json:"name" validate:"required"`
	StorageOfferingUuid string `json:"storageOfferingUuid" validate:"required"`
	ZoneUuid            string `json:"zoneUuid" validate:"required"`
}
