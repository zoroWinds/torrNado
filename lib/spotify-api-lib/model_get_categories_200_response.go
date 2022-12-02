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

// checks if the GetCategories200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCategories200Response{}

// GetCategories200Response struct for GetCategories200Response
type GetCategories200Response struct {
	Categories PagingObject `json:"categories"`
}

// NewGetCategories200Response instantiates a new GetCategories200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCategories200Response(categories PagingObject) *GetCategories200Response {
	this := GetCategories200Response{}
	this.Categories = categories
	return &this
}

// NewGetCategories200ResponseWithDefaults instantiates a new GetCategories200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCategories200ResponseWithDefaults() *GetCategories200Response {
	this := GetCategories200Response{}
	return &this
}

// GetCategories returns the Categories field value
func (o *GetCategories200Response) GetCategories() PagingObject {
	if o == nil {
		var ret PagingObject
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *GetCategories200Response) GetCategoriesOk() (*PagingObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Categories, true
}

// SetCategories sets field value
func (o *GetCategories200Response) SetCategories(v PagingObject) {
	o.Categories = v
}

func (o GetCategories200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCategories200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["categories"] = o.Categories
	return toSerialize, nil
}

type NullableGetCategories200Response struct {
	value *GetCategories200Response
	isSet bool
}

func (v NullableGetCategories200Response) Get() *GetCategories200Response {
	return v.value
}

func (v *NullableGetCategories200Response) Set(val *GetCategories200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCategories200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCategories200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCategories200Response(val *GetCategories200Response) *NullableGetCategories200Response {
	return &NullableGetCategories200Response{value: val, isSet: true}
}

func (v NullableGetCategories200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCategories200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


