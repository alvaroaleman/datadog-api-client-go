// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAttributionSort The field to sort by.
type UsageAttributionSort string

// List of UsageAttributionSort.
const (
	USAGEATTRIBUTIONSORT_API_PERCENTAGE                               UsageAttributionSort = "api_percentage"
	USAGEATTRIBUTIONSORT_SNMP_USAGE                                   UsageAttributionSort = "snmp_usage"
	USAGEATTRIBUTIONSORT_APM_HOST_USAGE                               UsageAttributionSort = "apm_host_usage"
	USAGEATTRIBUTIONSORT_API_USAGE                                    UsageAttributionSort = "api_usage"
	USAGEATTRIBUTIONSORT_APPSEC_USAGE                                 UsageAttributionSort = "appsec_usage"
	USAGEATTRIBUTIONSORT_APPSEC_PERCENTAGE                            UsageAttributionSort = "appsec_percentage"
	USAGEATTRIBUTIONSORT_CONTAINER_USAGE                              UsageAttributionSort = "container_usage"
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_PERCENTAGE                 UsageAttributionSort = "custom_timeseries_percentage"
	USAGEATTRIBUTIONSORT_CONTAINER_PERCENTAGE                         UsageAttributionSort = "container_percentage"
	USAGEATTRIBUTIONSORT_APM_HOST_PERCENTAGE                          UsageAttributionSort = "apm_host_percentage"
	USAGEATTRIBUTIONSORT_NPM_HOST_PERCENTAGE                          UsageAttributionSort = "npm_host_percentage"
	USAGEATTRIBUTIONSORT_BROWSER_PERCENTAGE                           UsageAttributionSort = "browser_percentage"
	USAGEATTRIBUTIONSORT_BROWSER_USAGE                                UsageAttributionSort = "browser_usage"
	USAGEATTRIBUTIONSORT_INFRA_HOST_PERCENTAGE                        UsageAttributionSort = "infra_host_percentage"
	USAGEATTRIBUTIONSORT_SNMP_PERCENTAGE                              UsageAttributionSort = "snmp_percentage"
	USAGEATTRIBUTIONSORT_NPM_HOST_USAGE                               UsageAttributionSort = "npm_host_usage"
	USAGEATTRIBUTIONSORT_INFRA_HOST_USAGE                             UsageAttributionSort = "infra_host_usage"
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_USAGE                      UsageAttributionSort = "custom_timeseries_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_USAGE                       UsageAttributionSort = "lambda_functions_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_PERCENTAGE                  UsageAttributionSort = "lambda_functions_percentage"
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_USAGE                     UsageAttributionSort = "lambda_invocations_usage"
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_PERCENTAGE                UsageAttributionSort = "lambda_invocations_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_USAGE                 UsageAttributionSort = "estimated_indexed_logs_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_PERCENTAGE            UsageAttributionSort = "estimated_indexed_logs_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_LOGS_USAGE                UsageAttributionSort = "estimated_ingested_logs_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_LOGS_PERCENTAGE           UsageAttributionSort = "estimated_ingested_logs_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_SPANS_USAGE                UsageAttributionSort = "estimated_indexed_spans_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_SPANS_PERCENTAGE           UsageAttributionSort = "estimated_indexed_spans_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_SPANS_USAGE               UsageAttributionSort = "estimated_ingested_spans_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_SPANS_PERCENTAGE          UsageAttributionSort = "estimated_ingested_spans_percentage"
	USAGEATTRIBUTIONSORT_APM_FARGATE_USAGE                            UsageAttributionSort = "apm_fargate_usage"
	USAGEATTRIBUTIONSORT_APM_FARGATE_PERCENTAGE                       UsageAttributionSort = "apm_fargate_percentage"
	USAGEATTRIBUTIONSORT_APPSEC_FARGATE_USAGE                         UsageAttributionSort = "appsec_fargate_usage"
	USAGEATTRIBUTIONSORT_APPSEC_FARGATE_PERCENTAGE                    UsageAttributionSort = "appsec_fargate_percentage"
	USAGEATTRIBUTIONSORT_ESTIMATED_RUM_USAGE_ATTRIBUTION_USAGE        UsageAttributionSort = "estimated_rum_usage_attribution_usage"
	USAGEATTRIBUTIONSORT_ESTIMATED_RUM_USAGE_ATTRIBUTION_PERCENTAGE   UsageAttributionSort = "estimated_rum_usage_attribution_percentage"
	USAGEATTRIBUTIONSORT_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE      UsageAttributionSort = "asm_serverless_traced_invocations_usage"
	USAGEATTRIBUTIONSORT_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE UsageAttributionSort = "asm_serverless_traced_invocations_percentage"
)

var allowedUsageAttributionSortEnumValues = []UsageAttributionSort{
	USAGEATTRIBUTIONSORT_API_PERCENTAGE,
	USAGEATTRIBUTIONSORT_SNMP_USAGE,
	USAGEATTRIBUTIONSORT_APM_HOST_USAGE,
	USAGEATTRIBUTIONSORT_API_USAGE,
	USAGEATTRIBUTIONSORT_APPSEC_USAGE,
	USAGEATTRIBUTIONSORT_APPSEC_PERCENTAGE,
	USAGEATTRIBUTIONSORT_CONTAINER_USAGE,
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_PERCENTAGE,
	USAGEATTRIBUTIONSORT_CONTAINER_PERCENTAGE,
	USAGEATTRIBUTIONSORT_APM_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_NPM_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_BROWSER_PERCENTAGE,
	USAGEATTRIBUTIONSORT_BROWSER_USAGE,
	USAGEATTRIBUTIONSORT_INFRA_HOST_PERCENTAGE,
	USAGEATTRIBUTIONSORT_SNMP_PERCENTAGE,
	USAGEATTRIBUTIONSORT_NPM_HOST_USAGE,
	USAGEATTRIBUTIONSORT_INFRA_HOST_USAGE,
	USAGEATTRIBUTIONSORT_CUSTOM_TIMESERIES_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_FUNCTIONS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_USAGE,
	USAGEATTRIBUTIONSORT_LAMBDA_INVOCATIONS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_LOGS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_LOGS_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_LOGS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_SPANS_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INDEXED_SPANS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_SPANS_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_INGESTED_SPANS_PERCENTAGE,
	USAGEATTRIBUTIONSORT_APM_FARGATE_USAGE,
	USAGEATTRIBUTIONSORT_APM_FARGATE_PERCENTAGE,
	USAGEATTRIBUTIONSORT_APPSEC_FARGATE_USAGE,
	USAGEATTRIBUTIONSORT_APPSEC_FARGATE_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_RUM_USAGE_ATTRIBUTION_USAGE,
	USAGEATTRIBUTIONSORT_ESTIMATED_RUM_USAGE_ATTRIBUTION_PERCENTAGE,
	USAGEATTRIBUTIONSORT_ASM_SERVERLESS_TRACED_INVOCATIONS_USAGE,
	USAGEATTRIBUTIONSORT_ASM_SERVERLESS_TRACED_INVOCATIONS_PERCENTAGE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *UsageAttributionSort) GetAllowedValues() []UsageAttributionSort {
	return allowedUsageAttributionSortEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *UsageAttributionSort) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = UsageAttributionSort(value)
	return nil
}

// NewUsageAttributionSortFromValue returns a pointer to a valid UsageAttributionSort
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewUsageAttributionSortFromValue(v string) (*UsageAttributionSort, error) {
	ev := UsageAttributionSort(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for UsageAttributionSort: valid values are %v", v, allowedUsageAttributionSortEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v UsageAttributionSort) IsValid() bool {
	for _, existing := range allowedUsageAttributionSortEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageAttributionSort value.
func (v UsageAttributionSort) Ptr() *UsageAttributionSort {
	return &v
}
