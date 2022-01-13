package api

var (
	templateList = "template/templateList"
)

// Template List api
func GetNewTemplateListApi(zoneUuid string) string {
	return EndPoint + "/" + templateList + "?zoneUuid=" + zoneUuid
}
