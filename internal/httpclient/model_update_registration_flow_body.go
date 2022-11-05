/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.
 *
 * API version:
 * Contact: office@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UpdateRegistrationFlowBody - Update Registration Request Body
type UpdateRegistrationFlowBody struct {
	UpdateRegistrationFlowWithOidcMethod     *UpdateRegistrationFlowWithOidcMethod
	UpdateRegistrationFlowWithPasswordMethod *UpdateRegistrationFlowWithPasswordMethod
	UpdateRegistrationFlowWithWebAuthnMethod *UpdateRegistrationFlowWithWebAuthnMethod
}

// UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithOidcMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithOidcMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithOidcMethod: v,
	}
}

// UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithPasswordMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithPasswordMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithPasswordMethod: v,
	}
}

// UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithWebAuthnMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithWebAuthnMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithWebAuthnMethod: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRegistrationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateRegistrationFlowWithOidcMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateRegistrationFlowWithOidcMethod)
	if err == nil {
		jsonUpdateRegistrationFlowWithOidcMethod, _ := json.Marshal(dst.UpdateRegistrationFlowWithOidcMethod)
		if string(jsonUpdateRegistrationFlowWithOidcMethod) == "{}" { // empty struct
			dst.UpdateRegistrationFlowWithOidcMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRegistrationFlowWithOidcMethod = nil
	}

	// try to unmarshal data into UpdateRegistrationFlowWithPasswordMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateRegistrationFlowWithPasswordMethod)
	if err == nil {
		jsonUpdateRegistrationFlowWithPasswordMethod, _ := json.Marshal(dst.UpdateRegistrationFlowWithPasswordMethod)
		if string(jsonUpdateRegistrationFlowWithPasswordMethod) == "{}" { // empty struct
			dst.UpdateRegistrationFlowWithPasswordMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRegistrationFlowWithPasswordMethod = nil
	}

	// try to unmarshal data into UpdateRegistrationFlowWithWebAuthnMethod
	err = newStrictDecoder(data).Decode(&dst.UpdateRegistrationFlowWithWebAuthnMethod)
	if err == nil {
		jsonUpdateRegistrationFlowWithWebAuthnMethod, _ := json.Marshal(dst.UpdateRegistrationFlowWithWebAuthnMethod)
		if string(jsonUpdateRegistrationFlowWithWebAuthnMethod) == "{}" { // empty struct
			dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdateRegistrationFlowWithOidcMethod = nil
		dst.UpdateRegistrationFlowWithPasswordMethod = nil
		dst.UpdateRegistrationFlowWithWebAuthnMethod = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(UpdateRegistrationFlowBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateRegistrationFlowBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	if src.UpdateRegistrationFlowWithOidcMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithOidcMethod)
	}

	if src.UpdateRegistrationFlowWithPasswordMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithPasswordMethod)
	}

	if src.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithWebAuthnMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRegistrationFlowBody) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateRegistrationFlowWithOidcMethod != nil {
		return obj.UpdateRegistrationFlowWithOidcMethod
	}

	if obj.UpdateRegistrationFlowWithPasswordMethod != nil {
		return obj.UpdateRegistrationFlowWithPasswordMethod
	}

	if obj.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return obj.UpdateRegistrationFlowWithWebAuthnMethod
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRegistrationFlowBody struct {
	value *UpdateRegistrationFlowBody
	isSet bool
}

func (v NullableUpdateRegistrationFlowBody) Get() *UpdateRegistrationFlowBody {
	return v.value
}

func (v *NullableUpdateRegistrationFlowBody) Set(val *UpdateRegistrationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowBody(val *UpdateRegistrationFlowBody) *NullableUpdateRegistrationFlowBody {
	return &NullableUpdateRegistrationFlowBody{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
