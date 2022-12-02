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

// checks if the ShowObjectAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowObjectAllOf{}

// ShowObjectAllOf struct for ShowObjectAllOf
type ShowObjectAllOf struct {
	// The episodes of the show. 
	Episodes PagingObject `json:"episodes"`
}

// NewShowObjectAllOf instantiates a new ShowObjectAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowObjectAllOf(episodes PagingObject) *ShowObjectAllOf {
	this := ShowObjectAllOf{}
	this.Episodes = episodes
	return &this
}

// NewShowObjectAllOfWithDefaults instantiates a new ShowObjectAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowObjectAllOfWithDefaults() *ShowObjectAllOf {
	this := ShowObjectAllOf{}
	return &this
}

// GetEpisodes returns the Episodes field value
func (o *ShowObjectAllOf) GetEpisodes() PagingObject {
	if o == nil {
		var ret PagingObject
		return ret
	}

	return o.Episodes
}

// GetEpisodesOk returns a tuple with the Episodes field value
// and a boolean to check if the value has been set.
func (o *ShowObjectAllOf) GetEpisodesOk() (*PagingObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Episodes, true
}

// SetEpisodes sets field value
func (o *ShowObjectAllOf) SetEpisodes(v PagingObject) {
	o.Episodes = v
}

func (o ShowObjectAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowObjectAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["episodes"] = o.Episodes
	return toSerialize, nil
}

type NullableShowObjectAllOf struct {
	value *ShowObjectAllOf
	isSet bool
}

func (v NullableShowObjectAllOf) Get() *ShowObjectAllOf {
	return v.value
}

func (v *NullableShowObjectAllOf) Set(val *ShowObjectAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableShowObjectAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableShowObjectAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowObjectAllOf(val *ShowObjectAllOf) *NullableShowObjectAllOf {
	return &NullableShowObjectAllOf{value: val, isSet: true}
}

func (v NullableShowObjectAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowObjectAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


