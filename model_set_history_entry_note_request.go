/*
sipgate API

This is the sipgate REST API documentation. We build our applications on this API and we invite you to use it too.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SetHistoryEntryNoteRequest struct for SetHistoryEntryNoteRequest
type SetHistoryEntryNoteRequest struct {
	Note string `json:"note"`
}

// NewSetHistoryEntryNoteRequest instantiates a new SetHistoryEntryNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetHistoryEntryNoteRequest(note string) *SetHistoryEntryNoteRequest {
	this := SetHistoryEntryNoteRequest{}
	this.Note = note
	return &this
}

// NewSetHistoryEntryNoteRequestWithDefaults instantiates a new SetHistoryEntryNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetHistoryEntryNoteRequestWithDefaults() *SetHistoryEntryNoteRequest {
	this := SetHistoryEntryNoteRequest{}
	return &this
}

// GetNote returns the Note field value
func (o *SetHistoryEntryNoteRequest) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *SetHistoryEntryNoteRequest) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *SetHistoryEntryNoteRequest) SetNote(v string) {
	o.Note = v
}

func (o SetHistoryEntryNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["note"] = o.Note
	}
	return json.Marshal(toSerialize)
}

type NullableSetHistoryEntryNoteRequest struct {
	value *SetHistoryEntryNoteRequest
	isSet bool
}

func (v NullableSetHistoryEntryNoteRequest) Get() *SetHistoryEntryNoteRequest {
	return v.value
}

func (v *NullableSetHistoryEntryNoteRequest) Set(val *SetHistoryEntryNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetHistoryEntryNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetHistoryEntryNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetHistoryEntryNoteRequest(val *SetHistoryEntryNoteRequest) *NullableSetHistoryEntryNoteRequest {
	return &NullableSetHistoryEntryNoteRequest{value: val, isSet: true}
}

func (v NullableSetHistoryEntryNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetHistoryEntryNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


