package utils

import "terraform-provider-stackbill/config"

func GetProviderName() string {
	return config.Provider
}
