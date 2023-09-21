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

// checks if the PaymentChannelAllOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentChannelAllOf{}

// PaymentChannelAllOf struct for PaymentChannelAllOf
type PaymentChannelAllOf struct {
	// The specific Xendit code used to identify the partner channel
	ChannelCode *string `json:"channel_code,omitempty"`
	// The payment method type
	Type *string `json:"type,omitempty"`
	// The country where the channel operates  in ISO 3166-2 country code
	Country *string `json:"country,omitempty"`
	// Official parter name of the payment channel
	ChannelName *string `json:"channel_name,omitempty"`
	ChannelProperties []ChannelProperty `json:"channel_properties,omitempty"`
	// If available, this contains a URL to the logo of the partner channel
	LogoUrl *string `json:"logo_url,omitempty"`
	AmountLimits []ChannelAmountLimits `json:"amount_limits,omitempty"`
}

// NewPaymentChannelAllOf instantiates a new PaymentChannelAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelAllOf() *PaymentChannelAllOf {
	this := PaymentChannelAllOf{}
	return &this
}

// NewPaymentChannelAllOfWithDefaults instantiates a new PaymentChannelAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelAllOfWithDefaults() *PaymentChannelAllOf {
	this := PaymentChannelAllOf{}
	return &this
}

// GetChannelCode returns the ChannelCode field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetChannelCode() string {
	if o == nil || utils.IsNil(o.ChannelCode) {
		var ret string
		return ret
	}
	return *o.ChannelCode
}

// GetChannelCodeOk returns a tuple with the ChannelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetChannelCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelCode) {
		return nil, false
	}
	return o.ChannelCode, true
}

// HasChannelCode returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasChannelCode() bool {
	if o != nil && !utils.IsNil(o.ChannelCode) {
		return true
	}

	return false
}

// SetChannelCode gets a reference to the given string and assigns it to the ChannelCode field.
func (o *PaymentChannelAllOf) SetChannelCode(v string) {
	o.ChannelCode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentChannelAllOf) SetType(v string) {
	o.Type = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetCountry() string {
	if o == nil || utils.IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasCountry() bool {
	if o != nil && !utils.IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PaymentChannelAllOf) SetCountry(v string) {
	o.Country = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetChannelName() string {
	if o == nil || utils.IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetChannelNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasChannelName() bool {
	if o != nil && !utils.IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *PaymentChannelAllOf) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetChannelProperties returns the ChannelProperties field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetChannelProperties() []ChannelProperty {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		var ret []ChannelProperty
		return ret
	}
	return o.ChannelProperties
}

// GetChannelPropertiesOk returns a tuple with the ChannelProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetChannelPropertiesOk() ([]ChannelProperty, bool) {
	if o == nil || utils.IsNil(o.ChannelProperties) {
		return nil, false
	}
	return o.ChannelProperties, true
}

// HasChannelProperties returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasChannelProperties() bool {
	if o != nil && !utils.IsNil(o.ChannelProperties) {
		return true
	}

	return false
}

// SetChannelProperties gets a reference to the given []ChannelProperty and assigns it to the ChannelProperties field.
func (o *PaymentChannelAllOf) SetChannelProperties(v []ChannelProperty) {
	o.ChannelProperties = v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetLogoUrl() string {
	if o == nil || utils.IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetLogoUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasLogoUrl() bool {
	if o != nil && !utils.IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *PaymentChannelAllOf) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetAmountLimits returns the AmountLimits field value if set, zero value otherwise.
func (o *PaymentChannelAllOf) GetAmountLimits() []ChannelAmountLimits {
	if o == nil || utils.IsNil(o.AmountLimits) {
		var ret []ChannelAmountLimits
		return ret
	}
	return o.AmountLimits
}

// GetAmountLimitsOk returns a tuple with the AmountLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelAllOf) GetAmountLimitsOk() ([]ChannelAmountLimits, bool) {
	if o == nil || utils.IsNil(o.AmountLimits) {
		return nil, false
	}
	return o.AmountLimits, true
}

// HasAmountLimits returns a boolean if a field has been set.
func (o *PaymentChannelAllOf) HasAmountLimits() bool {
	if o != nil && !utils.IsNil(o.AmountLimits) {
		return true
	}

	return false
}

// SetAmountLimits gets a reference to the given []ChannelAmountLimits and assigns it to the AmountLimits field.
func (o *PaymentChannelAllOf) SetAmountLimits(v []ChannelAmountLimits) {
	o.AmountLimits = v
}

func (o PaymentChannelAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ChannelCode) {
		toSerialize["channel_code"] = o.ChannelCode
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !utils.IsNil(o.ChannelName) {
		toSerialize["channel_name"] = o.ChannelName
	}
	if !utils.IsNil(o.ChannelProperties) {
		toSerialize["channel_properties"] = o.ChannelProperties
	}
	if !utils.IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !utils.IsNil(o.AmountLimits) {
		toSerialize["amount_limits"] = o.AmountLimits
	}
	return toSerialize, nil
}

type NullablePaymentChannelAllOf struct {
	value *PaymentChannelAllOf
	isSet bool
}

func (v NullablePaymentChannelAllOf) Get() *PaymentChannelAllOf {
	return v.value
}

func (v *NullablePaymentChannelAllOf) Set(val *PaymentChannelAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelAllOf(val *PaymentChannelAllOf) *NullablePaymentChannelAllOf {
	return &NullablePaymentChannelAllOf{value: val, isSet: true}
}

func (v NullablePaymentChannelAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

