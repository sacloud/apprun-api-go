package types

import v1 "github.com/sacloud/apprun-api-go/apis/v1"

var VersionSortOrders = []string{
	(string)(v1.ListApplicationVersionsParamsSortOrderAsc),
	(string)(v1.ListApplicationVersionsParamsSortOrderDesc),
}

var VersionStatuses = []string{
	(string)(v1.VersionStatusFail),
	(string)(v1.VersionStatusSuccess),
	(string)(v1.VersionStatusUnknown),
}
