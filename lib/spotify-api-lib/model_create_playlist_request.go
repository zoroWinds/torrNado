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

// checks if the CreatePlaylistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePlaylistRequest{}

// CreatePlaylistRequest struct for CreatePlaylistRequest
type CreatePlaylistRequest struct {
	// The name for the new playlist, for example `\"Your Coolest Playlist\"`. This name does not need to be unique; a user may have several playlists with the same name. 
	Name string `json:"name"`
	// Defaults to `true`. If `true` the playlist will be public, if `false` it will be private. To be able to create private playlists, the user must have granted the `playlist-modify-private` [scope](/documentation/general/guides/authorization-guide/#list-of-scopes) 
	Public *bool `json:"public,omitempty"`
	// Defaults to `false`. If `true` the playlist will be collaborative. _**Note**: to create a collaborative playlist you must also set `public` to `false`. To create collaborative playlists you must have granted `playlist-modify-private` and `playlist-modify-public` [scopes](/documentation/general/guides/authorization-guide/#list-of-scopes)._ 
	Collaborative *bool `json:"collaborative,omitempty"`
	// value for playlist description as displayed in Spotify Clients and in the Web API. 
	Description *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreatePlaylistRequest CreatePlaylistRequest

// NewCreatePlaylistRequest instantiates a new CreatePlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlaylistRequest(name string) *CreatePlaylistRequest {
	this := CreatePlaylistRequest{}
	this.Name = name
	return &this
}

// NewCreatePlaylistRequestWithDefaults instantiates a new CreatePlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlaylistRequestWithDefaults() *CreatePlaylistRequest {
	this := CreatePlaylistRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreatePlaylistRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePlaylistRequest) SetName(v string) {
	o.Name = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetPublic() bool {
	if o == nil || isNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetPublicOk() (*bool, bool) {
	if o == nil || isNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CreatePlaylistRequest) SetPublic(v bool) {
	o.Public = &v
}

// GetCollaborative returns the Collaborative field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetCollaborative() bool {
	if o == nil || isNil(o.Collaborative) {
		var ret bool
		return ret
	}
	return *o.Collaborative
}

// GetCollaborativeOk returns a tuple with the Collaborative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetCollaborativeOk() (*bool, bool) {
	if o == nil || isNil(o.Collaborative) {
		return nil, false
	}
	return o.Collaborative, true
}

// HasCollaborative returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasCollaborative() bool {
	if o != nil && !isNil(o.Collaborative) {
		return true
	}

	return false
}

// SetCollaborative gets a reference to the given bool and assigns it to the Collaborative field.
func (o *CreatePlaylistRequest) SetCollaborative(v bool) {
	o.Collaborative = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePlaylistRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePlaylistRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePlaylistRequest) SetDescription(v string) {
	o.Description = &v
}

func (o CreatePlaylistRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePlaylistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !isNil(o.Collaborative) {
		toSerialize["collaborative"] = o.Collaborative
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreatePlaylistRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreatePlaylistRequest := _CreatePlaylistRequest{}

	if err = json.Unmarshal(bytes, &varCreatePlaylistRequest); err == nil {
		*o = CreatePlaylistRequest(varCreatePlaylistRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "public")
		delete(additionalProperties, "collaborative")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreatePlaylistRequest struct {
	value *CreatePlaylistRequest
	isSet bool
}

func (v NullableCreatePlaylistRequest) Get() *CreatePlaylistRequest {
	return v.value
}

func (v *NullableCreatePlaylistRequest) Set(val *CreatePlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlaylistRequest(val *CreatePlaylistRequest) *NullableCreatePlaylistRequest {
	return &NullableCreatePlaylistRequest{value: val, isSet: true}
}

func (v NullableCreatePlaylistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePlaylistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


