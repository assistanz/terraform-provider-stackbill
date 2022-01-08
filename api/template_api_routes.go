package api

var (
	templateList = "template/templateList"
)

// Template List api
func GetNewTemplateListApi(zoneId string) string {
	return EndPoint + "/" + templateList + "?zoneUuid=" + zoneId
}
