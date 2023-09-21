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

// checks if the ChannelProperty type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChannelProperty{}

// ChannelProperty struct for ChannelProperty
type ChannelProperty struct {
	// The corresponding parameter name
	Name *string `json:"name,omitempty"`
	// Data type of the parameter
	Type *string `json:"type,omitempty"`
	// Indicates whether or not the parameter is required
	IsRequired *bool `json:"is_required,omitempty"`
}

// NewChannelProperty instantiates a new ChannelProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelProperty() *ChannelProperty {
	this := ChannelProperty{}
	return &this
}

// NewChannelPropertyWithDefaults instantiates a new ChannelProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPropertyWithDefaults() *ChannelProperty {
	this := ChannelProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelProperty) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelProperty) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelProperty) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelProperty) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ChannelProperty) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelProperty) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ChannelProperty) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ChannelProperty) SetType(v string) {
	o.Type = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ChannelProperty) GetIsRequired() bool {
	if o == nil || utils.IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelProperty) GetIsRequiredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ChannelProperty) HasIsRequired() bool {
	if o != nil && !utils.IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ChannelProperty) SetIsRequired(v bool) {
	o.IsRequired = &v
}

func (o ChannelProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.IsRequired) {
		toSerialize["is_required"] = o.IsRequired
	}
	return toSerialize, nil
}

type NullableChannelProperty struct {
	value *ChannelProperty
	isSet bool
}

func (v NullableChannelProperty) Get() *ChannelProperty {
	return v.value
}

func (v *NullableChannelProperty) Set(val *ChannelProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelProperty(val *ChannelProperty) *NullableChannelProperty {
	return &NullableChannelProperty{value: val, isSet: true}
}

func (v NullableChannelProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


