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

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription struct for Subscription
type Subscription struct {
	// Unique identifier for the article
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// Unique identifier for the article
	StrategyId *string `json:"strategyId,omitempty"`
	// Unique identifier for the article
	UserId *string `json:"userId,omitempty"`
	// Unique identifier for the article
	BrokerAccountId *string `json:"brokerAccountId,omitempty"`
	// Percentage of your portfolio
	PortfolioPercentage *int32 `json:"portfolioPercentage,omitempty"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription() *Subscription {
	this := Subscription{}
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Subscription) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Subscription) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Subscription) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *Subscription) GetStrategyId() string {
	if o == nil || IsNil(o.StrategyId) {
		var ret string
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetStrategyIdOk() (*string, bool) {
	if o == nil || IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *Subscription) HasStrategyId() bool {
	if o != nil && !IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given string and assigns it to the StrategyId field.
func (o *Subscription) SetStrategyId(v string) {
	o.StrategyId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Subscription) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Subscription) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Subscription) SetUserId(v string) {
	o.UserId = &v
}

// GetBrokerAccountId returns the BrokerAccountId field value if set, zero value otherwise.
func (o *Subscription) GetBrokerAccountId() string {
	if o == nil || IsNil(o.BrokerAccountId) {
		var ret string
		return ret
	}
	return *o.BrokerAccountId
}

// GetBrokerAccountIdOk returns a tuple with the BrokerAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetBrokerAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrokerAccountId) {
		return nil, false
	}
	return o.BrokerAccountId, true
}

// HasBrokerAccountId returns a boolean if a field has been set.
func (o *Subscription) HasBrokerAccountId() bool {
	if o != nil && !IsNil(o.BrokerAccountId) {
		return true
	}

	return false
}

// SetBrokerAccountId gets a reference to the given string and assigns it to the BrokerAccountId field.
func (o *Subscription) SetBrokerAccountId(v string) {
	o.BrokerAccountId = &v
}

// GetPortfolioPercentage returns the PortfolioPercentage field value if set, zero value otherwise.
func (o *Subscription) GetPortfolioPercentage() int32 {
	if o == nil || IsNil(o.PortfolioPercentage) {
		var ret int32
		return ret
	}
	return *o.PortfolioPercentage
}

// GetPortfolioPercentageOk returns a tuple with the PortfolioPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetPortfolioPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.PortfolioPercentage) {
		return nil, false
	}
	return o.PortfolioPercentage, true
}

// HasPortfolioPercentage returns a boolean if a field has been set.
func (o *Subscription) HasPortfolioPercentage() bool {
	if o != nil && !IsNil(o.PortfolioPercentage) {
		return true
	}

	return false
}

// SetPortfolioPercentage gets a reference to the given int32 and assigns it to the PortfolioPercentage field.
func (o *Subscription) SetPortfolioPercentage(v int32) {
	o.PortfolioPercentage = &v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.StrategyId) {
		toSerialize["strategyId"] = o.StrategyId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.BrokerAccountId) {
		toSerialize["brokerAccountId"] = o.BrokerAccountId
	}
	if !IsNil(o.PortfolioPercentage) {
		toSerialize["portfolioPercentage"] = o.PortfolioPercentage
	}
	return toSerialize, nil
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


