/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

// Deprecated: DeviceReport isn't utilized and needs to be removed.
type DeviceReport struct {
	Timestamps
	Id       string   `json:"id"`
	Name     string   `json:"name"`     // non-database identifier for a device report - must be unique
	Device   string   `json:"device"`   // associated device name - should be a valid and unique device name
	Action   string   `json:"action"`   // associated interval action name - should be a valid and unique interval action name
	Expected []string `json:"expected"` // array of value descriptor names describing the types of data captured in the report
}

// Custom marshaling to make empty strings null
func (dr DeviceReport) MarshalJSON() ([]byte, error) {
	test := struct {
		Timestamps
		Id       string   `json:"id"`
		Name     *string  `json:"name"`     // non-database identifier for a device report - must be unique
		Device   *string  `json:"device"`   // associated device name - should be a valid and unique device name
		Action   *string  `json:"action"`   // associated interval action name - should be a valid and unique interval action name
		Expected []string `json:"expected"` // array of value descriptor names describing the types of data captured in the report
	}{
		Timestamps: dr.Timestamps,
		Id:         dr.Id,
		Expected:   dr.Expected,
	}

	// Empty strings are null
	if dr.Name != "" {
		test.Name = &dr.Name
	}
	if dr.Device != "" {
		test.Device = &dr.Device
	}
	if dr.Action != "" {
		test.Action = &dr.Action
	}

	return json.Marshal(test)
}

/*
 * To String function for DeviceProfile
 */
func (dr DeviceReport) String() string {
	out, err := json.Marshal(dr)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
