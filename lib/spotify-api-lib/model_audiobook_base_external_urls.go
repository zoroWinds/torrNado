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

// checks if the AudiobookBaseExternalUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudiobookBaseExternalUrls{}

// AudiobookBaseExternalUrls External URLs for this audiobook. 
type AudiobookBaseExternalUrls struct {
	// The [Spotify URL](/documentation/web-api/#spotify-uris-and-ids) for the object. 
	Spotify *string `json:"spotify,omitempty"`
}

// NewAudiobookBaseExternalUrls instantiates a new AudiobookBaseExternalUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudiobookBaseExternalUrls() *AudiobookBaseExternalUrls {
	this := AudiobookBaseExternalUrls{}
	return &this
}

// NewAudiobookBaseExternalUrlsWithDefaults instantiates a new AudiobookBaseExternalUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudiobookBaseExternalUrlsWithDefaults() *AudiobookBaseExternalUrls {
	this := AudiobookBaseExternalUrls{}
	return &this
}

// GetSpotify returns the Spotify field value if set, zero value otherwise.
func (o *AudiobookBaseExternalUrls) GetSpotify() string {
	if o == nil || isNil(o.Spotify) {
		var ret string
		return ret
	}
	return *o.Spotify
}

// GetSpotifyOk returns a tuple with the Spotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudiobookBaseExternalUrls) GetSpotifyOk() (*string, bool) {
	if o == nil || isNil(o.Spotify) {
		return nil, false
	}
	return o.Spotify, true
}

// HasSpotify returns a boolean if a field has been set.
func (o *AudiobookBaseExternalUrls) HasSpotify() bool {
	if o != nil && !isNil(o.Spotify) {
		return true
	}

	return false
}

// SetSpotify gets a reference to the given string and assigns it to the Spotify field.
func (o *AudiobookBaseExternalUrls) SetSpotify(v string) {
	o.Spotify = &v
}

func (o AudiobookBaseExternalUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AudiobookBaseExternalUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Spotify) {
		toSerialize["spotify"] = o.Spotify
	}
	return toSerialize, nil
}

type NullableAudiobookBaseExternalUrls struct {
	value *AudiobookBaseExternalUrls
	isSet bool
}

func (v NullableAudiobookBaseExternalUrls) Get() *AudiobookBaseExternalUrls {
	return v.value
}

func (v *NullableAudiobookBaseExternalUrls) Set(val *AudiobookBaseExternalUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableAudiobookBaseExternalUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableAudiobookBaseExternalUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudiobookBaseExternalUrls(val *AudiobookBaseExternalUrls) *NullableAudiobookBaseExternalUrls {
	return &NullableAudiobookBaseExternalUrls{value: val, isSet: true}
}

func (v NullableAudiobookBaseExternalUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudiobookBaseExternalUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


