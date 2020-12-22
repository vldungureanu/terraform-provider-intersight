/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-12-22T00:49:18Z.
 *
 * API version: 1.0.9-3127
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// TelemetryDruidDataSource - A data source is the Apache Druid equivalent of a database table. However, a query can also masquerade as a data source, providing subquery-like functionality. Query data sources are currently supported only by GroupBy queries.
type TelemetryDruidDataSource struct {
	TelemetryDruidInlineDataSource *TelemetryDruidInlineDataSource
	TelemetryDruidJoinDataSource   *TelemetryDruidJoinDataSource
	TelemetryDruidLookupDataSource *TelemetryDruidLookupDataSource
	TelemetryDruidQueryDataSource  *TelemetryDruidQueryDataSource
	TelemetryDruidTableDataSource  *TelemetryDruidTableDataSource
	TelemetryDruidUnionDataSource  *TelemetryDruidUnionDataSource
}

// TelemetryDruidInlineDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidInlineDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidInlineDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidInlineDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidInlineDataSource: v}
}

// TelemetryDruidJoinDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidJoinDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidJoinDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidJoinDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidJoinDataSource: v}
}

// TelemetryDruidLookupDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidLookupDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidLookupDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidLookupDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidLookupDataSource: v}
}

// TelemetryDruidQueryDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidQueryDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidQueryDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidQueryDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidQueryDataSource: v}
}

// TelemetryDruidTableDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidTableDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidTableDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidTableDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidTableDataSource: v}
}

// TelemetryDruidUnionDataSourceAsTelemetryDruidDataSource is a convenience function that returns TelemetryDruidUnionDataSource wrapped in TelemetryDruidDataSource
func TelemetryDruidUnionDataSourceAsTelemetryDruidDataSource(v *TelemetryDruidUnionDataSource) TelemetryDruidDataSource {
	return TelemetryDruidDataSource{TelemetryDruidUnionDataSource: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidDataSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'join'
	if jsonDict["type"] == "join" {
		// try to unmarshal JSON data into TelemetryDruidJoinDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidJoinDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidJoinDataSource, return on the first match
		} else {
			dst.TelemetryDruidJoinDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidJoinDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'lookup'
	if jsonDict["type"] == "lookup" {
		// try to unmarshal JSON data into TelemetryDruidLookupDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidLookupDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidLookupDataSource, return on the first match
		} else {
			dst.TelemetryDruidLookupDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidLookupDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'query'
	if jsonDict["type"] == "query" {
		// try to unmarshal JSON data into TelemetryDruidQueryDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidQueryDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidQueryDataSource, return on the first match
		} else {
			dst.TelemetryDruidQueryDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidQueryDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'scan'
	if jsonDict["type"] == "scan" {
		// try to unmarshal JSON data into TelemetryDruidInlineDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidInlineDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInlineDataSource, return on the first match
		} else {
			dst.TelemetryDruidInlineDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidInlineDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'table'
	if jsonDict["type"] == "table" {
		// try to unmarshal JSON data into TelemetryDruidTableDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidTableDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidTableDataSource, return on the first match
		} else {
			dst.TelemetryDruidTableDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidTableDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidInlineDataSource'
	if jsonDict["type"] == "telemetry.DruidInlineDataSource" {
		// try to unmarshal JSON data into TelemetryDruidInlineDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidInlineDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidInlineDataSource, return on the first match
		} else {
			dst.TelemetryDruidInlineDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidInlineDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidJoinDataSource'
	if jsonDict["type"] == "telemetry.DruidJoinDataSource" {
		// try to unmarshal JSON data into TelemetryDruidJoinDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidJoinDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidJoinDataSource, return on the first match
		} else {
			dst.TelemetryDruidJoinDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidJoinDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidLookupDataSource'
	if jsonDict["type"] == "telemetry.DruidLookupDataSource" {
		// try to unmarshal JSON data into TelemetryDruidLookupDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidLookupDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidLookupDataSource, return on the first match
		} else {
			dst.TelemetryDruidLookupDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidLookupDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidQueryDataSource'
	if jsonDict["type"] == "telemetry.DruidQueryDataSource" {
		// try to unmarshal JSON data into TelemetryDruidQueryDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidQueryDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidQueryDataSource, return on the first match
		} else {
			dst.TelemetryDruidQueryDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidQueryDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidTableDataSource'
	if jsonDict["type"] == "telemetry.DruidTableDataSource" {
		// try to unmarshal JSON data into TelemetryDruidTableDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidTableDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidTableDataSource, return on the first match
		} else {
			dst.TelemetryDruidTableDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidTableDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidUnionDataSource'
	if jsonDict["type"] == "telemetry.DruidUnionDataSource" {
		// try to unmarshal JSON data into TelemetryDruidUnionDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidUnionDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidUnionDataSource, return on the first match
		} else {
			dst.TelemetryDruidUnionDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidUnionDataSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'union'
	if jsonDict["type"] == "union" {
		// try to unmarshal JSON data into TelemetryDruidUnionDataSource
		err = json.Unmarshal(data, &dst.TelemetryDruidUnionDataSource)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidUnionDataSource, return on the first match
		} else {
			dst.TelemetryDruidUnionDataSource = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDataSource as TelemetryDruidUnionDataSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidDataSource) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidInlineDataSource != nil {
		return json.Marshal(&src.TelemetryDruidInlineDataSource)
	}

	if src.TelemetryDruidJoinDataSource != nil {
		return json.Marshal(&src.TelemetryDruidJoinDataSource)
	}

	if src.TelemetryDruidLookupDataSource != nil {
		return json.Marshal(&src.TelemetryDruidLookupDataSource)
	}

	if src.TelemetryDruidQueryDataSource != nil {
		return json.Marshal(&src.TelemetryDruidQueryDataSource)
	}

	if src.TelemetryDruidTableDataSource != nil {
		return json.Marshal(&src.TelemetryDruidTableDataSource)
	}

	if src.TelemetryDruidUnionDataSource != nil {
		return json.Marshal(&src.TelemetryDruidUnionDataSource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidDataSource) GetActualInstance() interface{} {
	if obj.TelemetryDruidInlineDataSource != nil {
		return obj.TelemetryDruidInlineDataSource
	}

	if obj.TelemetryDruidJoinDataSource != nil {
		return obj.TelemetryDruidJoinDataSource
	}

	if obj.TelemetryDruidLookupDataSource != nil {
		return obj.TelemetryDruidLookupDataSource
	}

	if obj.TelemetryDruidQueryDataSource != nil {
		return obj.TelemetryDruidQueryDataSource
	}

	if obj.TelemetryDruidTableDataSource != nil {
		return obj.TelemetryDruidTableDataSource
	}

	if obj.TelemetryDruidUnionDataSource != nil {
		return obj.TelemetryDruidUnionDataSource
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidDataSource struct {
	value *TelemetryDruidDataSource
	isSet bool
}

func (v NullableTelemetryDruidDataSource) Get() *TelemetryDruidDataSource {
	return v.value
}

func (v *NullableTelemetryDruidDataSource) Set(val *TelemetryDruidDataSource) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDataSource(val *TelemetryDruidDataSource) *NullableTelemetryDruidDataSource {
	return &NullableTelemetryDruidDataSource{value: val, isSet: true}
}

func (v NullableTelemetryDruidDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
