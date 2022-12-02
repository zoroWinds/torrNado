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

// checks if the LinkedTrackObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedTrackObject{}

// LinkedTrackObject struct for LinkedTrackObject
type LinkedTrackObject struct {
	ExternalUrls *LinkedTrackObjectExternalUrls `json:"external_urls,omitempty"`
	// A link to the Web API endpoint providing full details of the track. 
	Href *string `json:"href,omitempty"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the track. 
	Id *string `json:"id,omitempty"`
	// The object type: \"track\". 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the track. 
	Uri *string `json:"uri,omitempty"`
}

// NewLinkedTrackObject instantiates a new LinkedTrackObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedTrackObject() *LinkedTrackObject {
	this := LinkedTrackObject{}
	return &this
}

// NewLinkedTrackObjectWithDefaults instantiates a new LinkedTrackObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedTrackObjectWithDefaults() *LinkedTrackObject {
	this := LinkedTrackObject{}
	return &this
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *LinkedTrackObject) GetExternalUrls() LinkedTrackObjectExternalUrls {
	if o == nil || isNil(o.ExternalUrls) {
		var ret LinkedTrackObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObject) GetExternalUrlsOk() (*LinkedTrackObjectExternalUrls, bool) {
	if o == nil || isNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *LinkedTrackObject) HasExternalUrls() bool {
	if o != nil && !isNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given LinkedTrackObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *LinkedTrackObject) SetExternalUrls(v LinkedTrackObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LinkedTrackObject) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObject) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LinkedTrackObject) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LinkedTrackObject) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LinkedTrackObject) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObject) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LinkedTrackObject) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LinkedTrackObject) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinkedTrackObject) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObject) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinkedTrackObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinkedTrackObject) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *LinkedTrackObject) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkedTrackObject) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *LinkedTrackObject) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *LinkedTrackObject) SetUri(v string) {
	o.Uri = &v
}

func (o LinkedTrackObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedTrackObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableLinkedTrackObject struct {
	value *LinkedTrackObject
	isSet bool
}

func (v NullableLinkedTrackObject) Get() *LinkedTrackObject {
	return v.value
}

func (v *NullableLinkedTrackObject) Set(val *LinkedTrackObject) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedTrackObject) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedTrackObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedTrackObject(val *LinkedTrackObject) *NullableLinkedTrackObject {
	return &NullableLinkedTrackObject{value: val, isSet: true}
}

func (v NullableLinkedTrackObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedTrackObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

