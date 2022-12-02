/*
Spotify Web API

You can use Spotify's Web API to discover music and podcasts, manage your Spotify library, control audio playback, and much more. Browse our available Web API endpoints using the sidebar at left, or via the navigation bar on top of this page on smaller screens.  In order to make successful Web API requests your app will need a valid access token. One can be obtained through <a href=\"https://developer.spotify.com/documentation/general/guides/authorization-guide/\">OAuth 2.0</a>.  The base URI for all Web API requests is `https://api.spotify.com/v1`.  Need help? See our <a href=\"https://developer.spotify.com/documentation/web-api/guides/\">Web API guides</a> for more information, or visit the <a href=\"https://community.spotify.com/t5/Spotify-for-Developers/bd-p/Spotify_Developer\">Spotify for Developers community forum</a> to ask questions and connect with other developers. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the SavedTrackObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrackObject{}

// SavedTrackObject struct for SavedTrackObject
type SavedTrackObject struct {
	// The date and time the track was saved. Timestamps are returned in ISO 8601 format as Coordinated Universal Time (UTC) with a zero offset: YYYY-MM-DDTHH:MM:SSZ. If the time is imprecise (for example, the date/time of an album release), an additional field indicates the precision; see for example, release_date in an album object. 
	AddedAt *time.Time `json:"added_at,omitempty"`
	Track *SavedTrackObjectTrack `json:"track,omitempty"`
}

// NewSavedTrackObject instantiates a new SavedTrackObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrackObject() *SavedTrackObject {
	this := SavedTrackObject{}
	return &this
}

// NewSavedTrackObjectWithDefaults instantiates a new SavedTrackObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrackObjectWithDefaults() *SavedTrackObject {
	this := SavedTrackObject{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *SavedTrackObject) GetAddedAt() time.Time {
	if o == nil || isNil(o.AddedAt) {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObject) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.AddedAt) {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *SavedTrackObject) HasAddedAt() bool {
	if o != nil && !isNil(o.AddedAt) {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *SavedTrackObject) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *SavedTrackObject) GetTrack() SavedTrackObjectTrack {
	if o == nil || isNil(o.Track) {
		var ret SavedTrackObjectTrack
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObject) GetTrackOk() (*SavedTrackObjectTrack, bool) {
	if o == nil || isNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *SavedTrackObject) HasTrack() bool {
	if o != nil && !isNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given SavedTrackObjectTrack and assigns it to the Track field.
func (o *SavedTrackObject) SetTrack(v SavedTrackObjectTrack) {
	o.Track = &v
}

func (o SavedTrackObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrackObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AddedAt) {
		toSerialize["added_at"] = o.AddedAt
	}
	if !isNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	return toSerialize, nil
}

type NullableSavedTrackObject struct {
	value *SavedTrackObject
	isSet bool
}

func (v NullableSavedTrackObject) Get() *SavedTrackObject {
	return v.value
}

func (v *NullableSavedTrackObject) Set(val *SavedTrackObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrackObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrackObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrackObject(val *SavedTrackObject) *NullableSavedTrackObject {
	return &NullableSavedTrackObject{value: val, isSet: true}
}

func (v NullableSavedTrackObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrackObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

