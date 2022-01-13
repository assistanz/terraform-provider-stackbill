package network

// Create network request
type CreateNetworkRequest struct {
	Name                string `json:"name" validate:"required"`
	NetworkOfferingUuid string `json:"networkOfferingUuid"`
	VirtualmachineUuid  string `json:"virtualmachineUuid"`
	ZoneUuid            string `json:"zoneUuid"`
}

type CreateNetworkRequestOptional struct {
	IsPublic        bool   `json:"isPublic" validate:"required"`
	SecurityGroupId string `json:"networkOfferingUuid"`
}

type NetworkActionRequest struct {
	NetworkUuid string `json:"networkUuid" validate:"required"`
	Uuid        string `json:"uuid"`
}
