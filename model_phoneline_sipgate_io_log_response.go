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

// PhonelineSipgateIoLogResponse struct for PhonelineSipgateIoLogResponse
type PhonelineSipgateIoLogResponse struct {
	Created *string `json:"created,omitempty"`
	Url *string `json:"url,omitempty"`
	Request *string `json:"request,omitempty"`
	Response *string `json:"response,omitempty"`
	XmlError *string `json:"xmlError,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewPhonelineSipgateIoLogResponse instantiates a new PhonelineSipgateIoLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonelineSipgateIoLogResponse() *PhonelineSipgateIoLogResponse {
	this := PhonelineSipgateIoLogResponse{}
	return &this
}

// NewPhonelineSipgateIoLogResponseWithDefaults instantiates a new PhonelineSipgateIoLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonelineSipgateIoLogResponseWithDefaults() *PhonelineSipgateIoLogResponse {
	this := PhonelineSipgateIoLogResponse{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *PhonelineSipgateIoLogResponse) SetCreated(v string) {
	o.Created = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PhonelineSipgateIoLogResponse) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *PhonelineSipgateIoLogResponse) SetRequest(v string) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetResponse() string {
	if o == nil || o.Response == nil {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetResponseOk() (*string, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasResponse() bool {
	if o != nil && o.Response != nil {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *PhonelineSipgateIoLogResponse) SetResponse(v string) {
	o.Response = &v
}

// GetXmlError returns the XmlError field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetXmlError() string {
	if o == nil || o.XmlError == nil {
		var ret string
		return ret
	}
	return *o.XmlError
}

// GetXmlErrorOk returns a tuple with the XmlError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetXmlErrorOk() (*string, bool) {
	if o == nil || o.XmlError == nil {
		return nil, false
	}
	return o.XmlError, true
}

// HasXmlError returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasXmlError() bool {
	if o != nil && o.XmlError != nil {
		return true
	}

	return false
}

// SetXmlError gets a reference to the given string and assigns it to the XmlError field.
func (o *PhonelineSipgateIoLogResponse) SetXmlError(v string) {
	o.XmlError = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PhonelineSipgateIoLogResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhonelineSipgateIoLogResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PhonelineSipgateIoLogResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PhonelineSipgateIoLogResponse) SetStatus(v string) {
	o.Status = &v
}

func (o PhonelineSipgateIoLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.XmlError != nil {
		toSerialize["xmlError"] = o.XmlError
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePhonelineSipgateIoLogResponse struct {
	value *PhonelineSipgateIoLogResponse
	isSet bool
}

func (v NullablePhonelineSipgateIoLogResponse) Get() *PhonelineSipgateIoLogResponse {
	return v.value
}

func (v *NullablePhonelineSipgateIoLogResponse) Set(val *PhonelineSipgateIoLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonelineSipgateIoLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonelineSipgateIoLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonelineSipgateIoLogResponse(val *PhonelineSipgateIoLogResponse) *NullablePhonelineSipgateIoLogResponse {
	return &NullablePhonelineSipgateIoLogResponse{value: val, isSet: true}
}

func (v NullablePhonelineSipgateIoLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonelineSipgateIoLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


