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

// checks if the PublicUserObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicUserObject{}

// PublicUserObject struct for PublicUserObject
type PublicUserObject struct {
	// The name displayed on the user's profile. `null` if not available. 
	DisplayName NullableString `json:"display_name,omitempty"`
	ExternalUrls *PublicUserObjectExternalUrls `json:"external_urls,omitempty"`
	Followers *PublicUserObjectFollowers `json:"followers,omitempty"`
	// A link to the Web API endpoint for this user. 
	Href *string `json:"href,omitempty"`
	// The [Spotify user ID](/documentation/web-api/#spotify-uris-and-ids) for this user. 
	Id *string `json:"id,omitempty"`
	// The user's profile image. 
	Images []ImageObject `json:"images,omitempty"`
	// The object type. 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for this user. 
	Uri *string `json:"uri,omitempty"`
}

// NewPublicUserObject instantiates a new PublicUserObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicUserObject() *PublicUserObject {
	this := PublicUserObject{}
	return &this
}

// NewPublicUserObjectWithDefaults instantiates a new PublicUserObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicUserObjectWithDefaults() *PublicUserObject {
	this := PublicUserObject{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicUserObject) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicUserObject) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PublicUserObject) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *PublicUserObject) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *PublicUserObject) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *PublicUserObject) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *PublicUserObject) GetExternalUrls() PublicUserObjectExternalUrls {
	if o == nil || isNil(o.ExternalUrls) {
		var ret PublicUserObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetExternalUrlsOk() (*PublicUserObjectExternalUrls, bool) {
	if o == nil || isNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *PublicUserObject) HasExternalUrls() bool {
	if o != nil && !isNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given PublicUserObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *PublicUserObject) SetExternalUrls(v PublicUserObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *PublicUserObject) GetFollowers() PublicUserObjectFollowers {
	if o == nil || isNil(o.Followers) {
		var ret PublicUserObjectFollowers
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetFollowersOk() (*PublicUserObjectFollowers, bool) {
	if o == nil || isNil(o.Followers) {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *PublicUserObject) HasFollowers() bool {
	if o != nil && !isNil(o.Followers) {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given PublicUserObjectFollowers and assigns it to the Followers field.
func (o *PublicUserObject) SetFollowers(v PublicUserObjectFollowers) {
	o.Followers = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *PublicUserObject) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *PublicUserObject) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *PublicUserObject) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicUserObject) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicUserObject) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicUserObject) SetId(v string) {
	o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *PublicUserObject) GetImages() []ImageObject {
	if o == nil || isNil(o.Images) {
		var ret []ImageObject
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetImagesOk() ([]ImageObject, bool) {
	if o == nil || isNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *PublicUserObject) HasImages() bool {
	if o != nil && !isNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ImageObject and assigns it to the Images field.
func (o *PublicUserObject) SetImages(v []ImageObject) {
	o.Images = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PublicUserObject) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PublicUserObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PublicUserObject) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PublicUserObject) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicUserObject) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PublicUserObject) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PublicUserObject) SetUri(v string) {
	o.Uri = &v
}

func (o PublicUserObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicUserObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if !isNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !isNil(o.Followers) {
		toSerialize["followers"] = o.Followers
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullablePublicUserObject struct {
	value *PublicUserObject
	isSet bool
}

func (v NullablePublicUserObject) Get() *PublicUserObject {
	return v.value
}

func (v *NullablePublicUserObject) Set(val *PublicUserObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicUserObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicUserObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicUserObject(val *PublicUserObject) *NullablePublicUserObject {
	return &NullablePublicUserObject{value: val, isSet: true}
}

func (v NullablePublicUserObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicUserObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


