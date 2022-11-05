// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

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

// IdentityCredentialsType and so on.
type IdentityCredentialsType string

// List of identityCredentialsType
const (
	IDENTITYCREDENTIALSTYPE_PASSWORD      IdentityCredentialsType = "password"
	IDENTITYCREDENTIALSTYPE_TOTP          IdentityCredentialsType = "totp"
	IDENTITYCREDENTIALSTYPE_OIDC          IdentityCredentialsType = "oidc"
	IDENTITYCREDENTIALSTYPE_WEBAUTHN      IdentityCredentialsType = "webauthn"
	IDENTITYCREDENTIALSTYPE_LOOKUP_SECRET IdentityCredentialsType = "lookup_secret"
)

func (v *IdentityCredentialsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IdentityCredentialsType(value)
	for _, existing := range []IdentityCredentialsType{"password", "totp", "oidc", "webauthn", "lookup_secret"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IdentityCredentialsType", value)
}

// Ptr returns reference to identityCredentialsType value
func (v IdentityCredentialsType) Ptr() *IdentityCredentialsType {
	return &v
}

type NullableIdentityCredentialsType struct {
	value *IdentityCredentialsType
	isSet bool
}

func (v NullableIdentityCredentialsType) Get() *IdentityCredentialsType {
	return v.value
}

func (v *NullableIdentityCredentialsType) Set(val *IdentityCredentialsType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCredentialsType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCredentialsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCredentialsType(val *IdentityCredentialsType) *NullableIdentityCredentialsType {
	return &NullableIdentityCredentialsType{value: val, isSet: true}
}

func (v NullableIdentityCredentialsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCredentialsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
