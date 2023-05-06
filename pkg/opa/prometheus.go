package opa

import "github.com/prometheus/client_golang/prometheus"

var (
	namespace = "default"
	Up        = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"was te last scorecard was sucessfull",
		nil, nil,
	)
	ConstraintViolation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "constraint_violations"),
		"Opa Violations Needed For Constraints",
		[]string{"kind", "name", "violating_kind", "violating_name", "violating_namespace", "violating_msg", "violating_enforcement"}, nil,
	)
	ConstraintInformation = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "constraint_information"),
		"some information on all constraints",
		[]string{"kind", "name", "enforcementAction", "totalViolations"}, nil,
	)
)
