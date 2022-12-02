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

// checks if the ContextObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextObject{}

// ContextObject struct for ContextObject
type ContextObject struct {
	// The object type, e.g. \"artist\", \"playlist\", \"album\", \"show\". 
	Type *string `json:"type,omitempty"`
	// A link to the Web API endpoint providing full details of the track.
	Href *string `json:"href,omitempty"`
	ExternalUrls *ContextObjectExternalUrls `json:"external_urls,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the context. 
	Uri *string `json:"uri,omitempty"`
}

// NewContextObject instantiates a new ContextObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextObject() *ContextObject {
	this := ContextObject{}
	return &this
}

// NewContextObjectWithDefaults instantiates a new ContextObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextObjectWithDefaults() *ContextObject {
	this := ContextObject{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContextObject) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextObject) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContextObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContextObject) SetType(v string) {
	o.Type = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ContextObject) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextObject) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ContextObject) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ContextObject) SetHref(v string) {
	o.Href = &v
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *ContextObject) GetExternalUrls() ContextObjectExternalUrls {
	if o == nil || isNil(o.ExternalUrls) {
		var ret ContextObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextObject) GetExternalUrlsOk() (*ContextObjectExternalUrls, bool) {
	if o == nil || isNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *ContextObject) HasExternalUrls() bool {
	if o != nil && !isNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given ContextObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *ContextObject) SetExternalUrls(v ContextObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ContextObject) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextObject) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ContextObject) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ContextObject) SetUri(v string) {
	o.Uri = &v
}

func (o ContextObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableContextObject struct {
	value *ContextObject
	isSet bool
}

func (v NullableContextObject) Get() *ContextObject {
	return v.value
}

func (v *NullableContextObject) Set(val *ContextObject) {
	v.value = val
	v.isSet = true
}

func (v NullableContextObject) IsSet() bool {
	return v.isSet
}

func (v *NullableContextObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextObject(val *ContextObject) *NullableContextObject {
	return &NullableContextObject{value: val, isSet: true}
}

func (v NullableContextObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

