/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the RemoveEpisodesUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveEpisodesUserRequest{}

// RemoveEpisodesUserRequest struct for RemoveEpisodesUserRequest
type RemoveEpisodesUserRequest struct {
	// A JSON array of the [Spotify IDs](/documentation/web-api/#spotify-uris-and-ids). <br>A maximum of 50 items can be specified in one request. _**Note**: if the `ids` parameter is present in the query string, any IDs listed here in the body will be ignored._ 
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RemoveEpisodesUserRequest RemoveEpisodesUserRequest

// NewRemoveEpisodesUserRequest instantiates a new RemoveEpisodesUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveEpisodesUserRequest() *RemoveEpisodesUserRequest {
	this := RemoveEpisodesUserRequest{}
	return &this
}

// NewRemoveEpisodesUserRequestWithDefaults instantiates a new RemoveEpisodesUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveEpisodesUserRequestWithDefaults() *RemoveEpisodesUserRequest {
	this := RemoveEpisodesUserRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *RemoveEpisodesUserRequest) GetIds() []string {
	if o == nil || isNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveEpisodesUserRequest) GetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *RemoveEpisodesUserRequest) HasIds() bool {
	if o != nil && !isNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *RemoveEpisodesUserRequest) SetIds(v []string) {
	o.Ids = v
}

func (o RemoveEpisodesUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveEpisodesUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveEpisodesUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRemoveEpisodesUserRequest := _RemoveEpisodesUserRequest{}

	if err = json.Unmarshal(bytes, &varRemoveEpisodesUserRequest); err == nil {
		*o = RemoveEpisodesUserRequest(varRemoveEpisodesUserRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveEpisodesUserRequest struct {
	value *RemoveEpisodesUserRequest
	isSet bool
}

func (v NullableRemoveEpisodesUserRequest) Get() *RemoveEpisodesUserRequest {
	return v.value
}

func (v *NullableRemoveEpisodesUserRequest) Set(val *RemoveEpisodesUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveEpisodesUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveEpisodesUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveEpisodesUserRequest(val *RemoveEpisodesUserRequest) *NullableRemoveEpisodesUserRequest {
	return &NullableRemoveEpisodesUserRequest{value: val, isSet: true}
}

func (v NullableRemoveEpisodesUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveEpisodesUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


