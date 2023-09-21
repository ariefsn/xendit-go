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

// checks if the PaymentChannelListLinksInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaymentChannelListLinksInner{}

// PaymentChannelListLinksInner struct for PaymentChannelListLinksInner
type PaymentChannelListLinksInner struct {
	// Target URI that should contain a target to Internationalized Resource Identifiers (IRI)
	Href *string `json:"href,omitempty"`
	// The link relation type described how the current context (source) is related to target resource
	Rel *string `json:"rel,omitempty"`
	// The HTTP method to be used to call `href`
	Method *string `json:"method,omitempty"`
}

// NewPaymentChannelListLinksInner instantiates a new PaymentChannelListLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentChannelListLinksInner() *PaymentChannelListLinksInner {
	this := PaymentChannelListLinksInner{}
	return &this
}

// NewPaymentChannelListLinksInnerWithDefaults instantiates a new PaymentChannelListLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentChannelListLinksInnerWithDefaults() *PaymentChannelListLinksInner {
	this := PaymentChannelListLinksInner{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInner) GetHref() string {
	if o == nil || utils.IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInner) HasHref() bool {
	if o != nil && !utils.IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PaymentChannelListLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetRel returns the Rel field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInner) GetRel() string {
	if o == nil || utils.IsNil(o.Rel) {
		var ret string
		return ret
	}
	return *o.Rel
}

// GetRelOk returns a tuple with the Rel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInner) GetRelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Rel) {
		return nil, false
	}
	return o.Rel, true
}

// HasRel returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInner) HasRel() bool {
	if o != nil && !utils.IsNil(o.Rel) {
		return true
	}

	return false
}

// SetRel gets a reference to the given string and assigns it to the Rel field.
func (o *PaymentChannelListLinksInner) SetRel(v string) {
	o.Rel = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentChannelListLinksInner) GetMethod() string {
	if o == nil || utils.IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentChannelListLinksInner) GetMethodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentChannelListLinksInner) HasMethod() bool {
	if o != nil && !utils.IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentChannelListLinksInner) SetMethod(v string) {
	o.Method = &v
}

func (o PaymentChannelListLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentChannelListLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !utils.IsNil(o.Rel) {
		toSerialize["rel"] = o.Rel
	}
	if !utils.IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	return toSerialize, nil
}

type NullablePaymentChannelListLinksInner struct {
	value *PaymentChannelListLinksInner
	isSet bool
}

func (v NullablePaymentChannelListLinksInner) Get() *PaymentChannelListLinksInner {
	return v.value
}

func (v *NullablePaymentChannelListLinksInner) Set(val *PaymentChannelListLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentChannelListLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentChannelListLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentChannelListLinksInner(val *PaymentChannelListLinksInner) *NullablePaymentChannelListLinksInner {
	return &NullablePaymentChannelListLinksInner{value: val, isSet: true}
}

func (v NullablePaymentChannelListLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentChannelListLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


