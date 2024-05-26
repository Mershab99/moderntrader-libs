/*
ModernTrader-Go

ModernTrader-Go

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BrokerAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrokerAccount{}

// BrokerAccount struct for BrokerAccount
type BrokerAccount struct {
	// Unique identifier for the article
	BrokerAccountId *string `json:"brokerAccountId,omitempty"`
	// brokerage type (enum)
	BrokerageType *string `json:"brokerageType,omitempty"`
	// broker account username
	BrokerUsername *string `json:"brokerUsername,omitempty"`
	// broker account password
	BrokerPassword *string `json:"brokerPassword,omitempty"`
}

// NewBrokerAccount instantiates a new BrokerAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerAccount() *BrokerAccount {
	this := BrokerAccount{}
	return &this
}

// NewBrokerAccountWithDefaults instantiates a new BrokerAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerAccountWithDefaults() *BrokerAccount {
	this := BrokerAccount{}
	return &this
}

// GetBrokerAccountId returns the BrokerAccountId field value if set, zero value otherwise.
func (o *BrokerAccount) GetBrokerAccountId() string {
	if o == nil || IsNil(o.BrokerAccountId) {
		var ret string
		return ret
	}
	return *o.BrokerAccountId
}

// GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerAccount) GetBrokerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerAccountId) {
		return nil, false
	}
	return o.BrokerAccountId, true
}

// HasBrokerAccountId returns a boolean if a field has been set.
func (o *BrokerAccount) HasBrokerAccountId() bool {
	if o != nil && !IsNil(o.BrokerAccountId) {
		return true
	}

	return false
}

// SetBrokerAccountId gets a reference to the given string and assigns it to the BrokerAccountId field.
func (o *BrokerAccount) SetBrokerAccountId(v string) {
	o.BrokerAccountId = &v
}

// GetBrokerageType returns the BrokerageType field value if set, zero value otherwise.
func (o *BrokerAccount) GetBrokerageType() string {
	if o == nil || IsNil(o.BrokerageType) {
		var ret string
		return ret
	}
	return *o.BrokerageType
}

// GetBrokerageTypeOk returns a tuple with the BrokerageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerAccount) GetBrokerageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerageType) {
		return nil, false
	}
	return o.BrokerageType, true
}

// HasBrokerageType returns a boolean if a field has been set.
func (o *BrokerAccount) HasBrokerageType() bool {
	if o != nil && !IsNil(o.BrokerageType) {
		return true
	}

	return false
}

// SetBrokerageType gets a reference to the given string and assigns it to the BrokerageType field.
func (o *BrokerAccount) SetBrokerageType(v string) {
	o.BrokerageType = &v
}

// GetBrokerUsername returns the BrokerUsername field value if set, zero value otherwise.
func (o *BrokerAccount) GetBrokerUsername() string {
	if o == nil || IsNil(o.BrokerUsername) {
		var ret string
		return ret
	}
	return *o.BrokerUsername
}

// GetBrokerUsernameOk returns a tuple with the BrokerUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerAccount) GetBrokerUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerUsername) {
		return nil, false
	}
	return o.BrokerUsername, true
}

// HasBrokerUsername returns a boolean if a field has been set.
func (o *BrokerAccount) HasBrokerUsername() bool {
	if o != nil && !IsNil(o.BrokerUsername) {
		return true
	}

	return false
}

// SetBrokerUsername gets a reference to the given string and assigns it to the BrokerUsername field.
func (o *BrokerAccount) SetBrokerUsername(v string) {
	o.BrokerUsername = &v
}

// GetBrokerPassword returns the BrokerPassword field value if set, zero value otherwise.
func (o *BrokerAccount) GetBrokerPassword() string {
	if o == nil || IsNil(o.BrokerPassword) {
		var ret string
		return ret
	}
	return *o.BrokerPassword
}

// GetBrokerPasswordOk returns a tuple with the BrokerPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrokerAccount) GetBrokerPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerPassword) {
		return nil, false
	}
	return o.BrokerPassword, true
}

// HasBrokerPassword returns a boolean if a field has been set.
func (o *BrokerAccount) HasBrokerPassword() bool {
	if o != nil && !IsNil(o.BrokerPassword) {
		return true
	}

	return false
}

// SetBrokerPassword gets a reference to the given string and assigns it to the BrokerPassword field.
func (o *BrokerAccount) SetBrokerPassword(v string) {
	o.BrokerPassword = &v
}

func (o BrokerAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrokerAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrokerAccountId) {
		toSerialize["brokerAccountId"] = o.BrokerAccountId
	}
	if !IsNil(o.BrokerageType) {
		toSerialize["brokerageType"] = o.BrokerageType
	}
	if !IsNil(o.BrokerUsername) {
		toSerialize["brokerUsername"] = o.BrokerUsername
	}
	if !IsNil(o.BrokerPassword) {
		toSerialize["brokerPassword"] = o.BrokerPassword
	}
	return toSerialize, nil
}

type NullableBrokerAccount struct {
	value *BrokerAccount
	isSet bool
}

func (v NullableBrokerAccount) Get() *BrokerAccount {
	return v.value
}

func (v *NullableBrokerAccount) Set(val *BrokerAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerAccount(val *BrokerAccount) *NullableBrokerAccount {
	return &NullableBrokerAccount{value: val, isSet: true}
}

func (v NullableBrokerAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrokerAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


