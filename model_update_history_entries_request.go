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

// UpdateHistoryEntriesRequest struct for UpdateHistoryEntriesRequest
type UpdateHistoryEntriesRequest struct {
	Id string `json:"id"`
	Archived *bool `json:"archived,omitempty"`
	Read *bool `json:"read,omitempty"`
	Starred *bool `json:"starred,omitempty"`
}

// NewUpdateHistoryEntriesRequest instantiates a new UpdateHistoryEntriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateHistoryEntriesRequest(id string) *UpdateHistoryEntriesRequest {
	this := UpdateHistoryEntriesRequest{}
	this.Id = id
	var archived bool = false
	this.Archived = &archived
	var read bool = false
	this.Read = &read
	var starred bool = false
	this.Starred = &starred
	return &this
}

// NewUpdateHistoryEntriesRequestWithDefaults instantiates a new UpdateHistoryEntriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateHistoryEntriesRequestWithDefaults() *UpdateHistoryEntriesRequest {
	this := UpdateHistoryEntriesRequest{}
	var archived bool = false
	this.Archived = &archived
	var read bool = false
	this.Read = &read
	var starred bool = false
	this.Starred = &starred
	return &this
}

// GetId returns the Id field value
func (o *UpdateHistoryEntriesRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateHistoryEntriesRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateHistoryEntriesRequest) SetId(v string) {
	o.Id = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *UpdateHistoryEntriesRequest) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHistoryEntriesRequest) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *UpdateHistoryEntriesRequest) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *UpdateHistoryEntriesRequest) SetArchived(v bool) {
	o.Archived = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *UpdateHistoryEntriesRequest) GetRead() bool {
	if o == nil || o.Read == nil {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHistoryEntriesRequest) GetReadOk() (*bool, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *UpdateHistoryEntriesRequest) HasRead() bool {
	if o != nil && o.Read != nil {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *UpdateHistoryEntriesRequest) SetRead(v bool) {
	o.Read = &v
}

// GetStarred returns the Starred field value if set, zero value otherwise.
func (o *UpdateHistoryEntriesRequest) GetStarred() bool {
	if o == nil || o.Starred == nil {
		var ret bool
		return ret
	}
	return *o.Starred
}

// GetStarredOk returns a tuple with the Starred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateHistoryEntriesRequest) GetStarredOk() (*bool, bool) {
	if o == nil || o.Starred == nil {
		return nil, false
	}
	return o.Starred, true
}

// HasStarred returns a boolean if a field has been set.
func (o *UpdateHistoryEntriesRequest) HasStarred() bool {
	if o != nil && o.Starred != nil {
		return true
	}

	return false
}

// SetStarred gets a reference to the given bool and assigns it to the Starred field.
func (o *UpdateHistoryEntriesRequest) SetStarred(v bool) {
	o.Starred = &v
}

func (o UpdateHistoryEntriesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	if o.Starred != nil {
		toSerialize["starred"] = o.Starred
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateHistoryEntriesRequest struct {
	value *UpdateHistoryEntriesRequest
	isSet bool
}

func (v NullableUpdateHistoryEntriesRequest) Get() *UpdateHistoryEntriesRequest {
	return v.value
}

func (v *NullableUpdateHistoryEntriesRequest) Set(val *UpdateHistoryEntriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateHistoryEntriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateHistoryEntriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateHistoryEntriesRequest(val *UpdateHistoryEntriesRequest) *NullableUpdateHistoryEntriesRequest {
	return &NullableUpdateHistoryEntriesRequest{value: val, isSet: true}
}

func (v NullableUpdateHistoryEntriesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateHistoryEntriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


