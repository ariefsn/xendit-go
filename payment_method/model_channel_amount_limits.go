/*
Payment Method Service v2

This API is used for Payment Method Service v2

API version: 2.87.2
*/


package payment_method

import (
	"encoding/json"
	
	utils "github.com/xendit/xendit-go/v3/utils"
)

// checks if the ChannelAmountLimits type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelAmountLimits{}

// ChannelAmountLimits struct for ChannelAmountLimits
type ChannelAmountLimits struct {
	// Currency supported by the payment channel
	Currency *string `json:"currency,omitempty"`
	// The minimum allowed transaction amount for the payment channel
	MinLimit *float32 `json:"min_limit,omitempty"`
	// The minimum allowed transaction amount for the payment channel
	MaxLimit *float32 `json:"max_limit,omitempty"`
}

// NewChannelAmountLimits instantiates a new ChannelAmountLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelAmountLimits() *ChannelAmountLimits {
	this := ChannelAmountLimits{}
	return &this
}

// NewChannelAmountLimitsWithDefaults instantiates a new ChannelAmountLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelAmountLimitsWithDefaults() *ChannelAmountLimits {
	this := ChannelAmountLimits{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ChannelAmountLimits) GetCurrency() string {
	if o == nil || utils.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetCurrencyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ChannelAmountLimits) HasCurrency() bool {
	if o != nil && !utils.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ChannelAmountLimits) SetCurrency(v string) {
	o.Currency = &v
}

// GetMinLimit returns the MinLimit field value if set, zero value otherwise.
func (o *ChannelAmountLimits) GetMinLimit() float32 {
	if o == nil || utils.IsNil(o.MinLimit) {
		var ret float32
		return ret
	}
	return *o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetMinLimitOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MinLimit) {
		return nil, false
	}
	return o.MinLimit, true
}

// HasMinLimit returns a boolean if a field has been set.
func (o *ChannelAmountLimits) HasMinLimit() bool {
	if o != nil && !utils.IsNil(o.MinLimit) {
		return true
	}

	return false
}

// SetMinLimit gets a reference to the given float32 and assigns it to the MinLimit field.
func (o *ChannelAmountLimits) SetMinLimit(v float32) {
	o.MinLimit = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *ChannelAmountLimits) GetMaxLimit() float32 {
	if o == nil || utils.IsNil(o.MaxLimit) {
		var ret float32
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAmountLimits) GetMaxLimitOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MaxLimit) {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *ChannelAmountLimits) HasMaxLimit() bool {
	if o != nil && !utils.IsNil(o.MaxLimit) {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given float32 and assigns it to the MaxLimit field.
func (o *ChannelAmountLimits) SetMaxLimit(v float32) {
	o.MaxLimit = &v
}

func (o ChannelAmountLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelAmountLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !utils.IsNil(o.MinLimit) {
		toSerialize["min_limit"] = o.MinLimit
	}
	if !utils.IsNil(o.MaxLimit) {
		toSerialize["max_limit"] = o.MaxLimit
	}
	return toSerialize, nil
}

type NullableChannelAmountLimits struct {
	value *ChannelAmountLimits
	isSet bool
}

func (v NullableChannelAmountLimits) Get() *ChannelAmountLimits {
	return v.value
}

func (v *NullableChannelAmountLimits) Set(val *ChannelAmountLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAmountLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAmountLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAmountLimits(val *ChannelAmountLimits) *NullableChannelAmountLimits {
	return &NullableChannelAmountLimits{value: val, isSet: true}
}

func (v NullableChannelAmountLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAmountLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


