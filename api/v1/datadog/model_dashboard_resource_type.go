/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// DashboardResourceType Dashboard resource type.
type DashboardResourceType string

// List of DashboardResourceType
const (
	DASHBOARDRESOURCETYPE_DASHBOARD DashboardResourceType = "dashboard"
)

var allowedDashboardResourceTypeEnumValues = []DashboardResourceType{
	"dashboard",
}

func (w *DashboardResourceType) GetAllowedValues() []DashboardResourceType {
	return allowedDashboardResourceTypeEnumValues
}

func (v *DashboardResourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DashboardResourceType(value)
	for _, existing := range allowedDashboardResourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DashboardResourceType", value)
}

// NewDashboardResourceTypeFromValue returns a pointer to a valid DashboardResourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardResourceTypeFromValue(v string) (*DashboardResourceType, error) {
	ev := DashboardResourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardResourceType: valid values are %v", v, allowedDashboardResourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardResourceType) IsValid() bool {
	for _, existing := range allowedDashboardResourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardResourceType value
func (v DashboardResourceType) Ptr() *DashboardResourceType {
	return &v
}

type NullableDashboardResourceType struct {
	value *DashboardResourceType
	isSet bool
}

func (v NullableDashboardResourceType) Get() *DashboardResourceType {
	return v.value
}

func (v *NullableDashboardResourceType) Set(val *DashboardResourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardResourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardResourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardResourceType(val *DashboardResourceType) *NullableDashboardResourceType {
	return &NullableDashboardResourceType{value: val, isSet: true}
}

func (v NullableDashboardResourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardResourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
