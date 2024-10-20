/*
usermng

usermng api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usermngclientgo

import (
	"encoding/json"
)

// checks if the Signup201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Signup201Response{}

// Signup201Response struct for Signup201Response
type Signup201Response struct {
	Id *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewSignup201Response instantiates a new Signup201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignup201Response() *Signup201Response {
	this := Signup201Response{}
	return &this
}

// NewSignup201ResponseWithDefaults instantiates a new Signup201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignup201ResponseWithDefaults() *Signup201Response {
	this := Signup201Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Signup201Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signup201Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Signup201Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Signup201Response) SetId(v string) {
	o.Id = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Signup201Response) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Signup201Response) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Signup201Response) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Signup201Response) SetUsername(v string) {
	o.Username = &v
}

func (o Signup201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Signup201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableSignup201Response struct {
	value *Signup201Response
	isSet bool
}

func (v NullableSignup201Response) Get() *Signup201Response {
	return v.value
}

func (v *NullableSignup201Response) Set(val *Signup201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSignup201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSignup201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignup201Response(val *Signup201Response) *NullableSignup201Response {
	return &NullableSignup201Response{value: val, isSet: true}
}

func (v NullableSignup201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignup201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


