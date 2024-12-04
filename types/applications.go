package types

import v1 "github.com/sacloud/apprun-api-go/apis/v1"

var ApplicationMaxCPUs = []string{
	(string)(v1.PostApplicationBodyComponentMaxCpuN01),
	(string)(v1.PostApplicationBodyComponentMaxCpuN02),
	(string)(v1.PostApplicationBodyComponentMaxCpuN03),
	(string)(v1.PostApplicationBodyComponentMaxCpuN04),
	(string)(v1.PostApplicationBodyComponentMaxCpuN05),
	(string)(v1.PostApplicationBodyComponentMaxCpuN06),
	(string)(v1.PostApplicationBodyComponentMaxCpuN07),
	(string)(v1.PostApplicationBodyComponentMaxCpuN08),
	(string)(v1.PostApplicationBodyComponentMaxCpuN09),
	(string)(v1.PostApplicationBodyComponentMaxCpuN1),
}

var ApplicationMaxMemories = []string{
	(string)(v1.PostApplicationBodyComponentMaxMemoryN256Mi),
	(string)(v1.PostApplicationBodyComponentMaxMemoryN512Mi),
	(string)(v1.PostApplicationBodyComponentMaxMemoryN1Gi),
	(string)(v1.PostApplicationBodyComponentMaxMemoryN2Gi),
}

var ApplicationSortOrders = []string{
	(string)(v1.ListApplicationsParamsSortOrderAsc),
	(string)(v1.ListApplicationsParamsSortOrderDesc),
}

var ApplicationStatuses = []string{
	(string)(v1.ApplicationStatusFail),
	(string)(v1.ApplicationStatusSuccess),
	(string)(v1.ApplicationStatusUnknown),
}
