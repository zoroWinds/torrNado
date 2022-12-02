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

// checks if the SavedTrackObjectTrack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SavedTrackObjectTrack{}

// SavedTrackObjectTrack Information about the track.
type SavedTrackObjectTrack struct {
	Album *TrackObjectAlbum `json:"album,omitempty"`
	// The artists who performed the track. Each artist object includes a link in `href` to more detailed information about the artist. 
	Artists []ArtistObject `json:"artists,omitempty"`
	// A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. 
	AvailableMarkets []string `json:"available_markets,omitempty"`
	// The disc number (usually `1` unless the album consists of more than one disc). 
	DiscNumber *int32 `json:"disc_number,omitempty"`
	// The track length in milliseconds. 
	DurationMs *int32 `json:"duration_ms,omitempty"`
	// Whether or not the track has explicit lyrics ( `true` = yes it does; `false` = no it does not OR unknown). 
	Explicit *bool `json:"explicit,omitempty"`
	ExternalIds *TrackObjectExternalIds `json:"external_ids,omitempty"`
	ExternalUrls *LinkedTrackObjectExternalUrls `json:"external_urls,omitempty"`
	// A link to the Web API endpoint providing full details of the track. 
	Href *string `json:"href,omitempty"`
	// The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the track. 
	Id *string `json:"id,omitempty"`
	// Part of the response when [Track Relinking](/documentation/general/guides/track-relinking-guide/) is applied. If `true`, the track is playable in the given market. Otherwise `false`. 
	IsPlayable *bool `json:"is_playable,omitempty"`
	LinkedFrom *TrackObjectLinkedFrom `json:"linked_from,omitempty"`
	Restrictions *SimplifiedTrackObjectRestrictions `json:"restrictions,omitempty"`
	// The name of the track. 
	Name *string `json:"name,omitempty"`
	// The popularity of the track. The value will be between 0 and 100, with 100 being the most popular.<br>The popularity of a track is a value between 0 and 100, with 100 being the most popular. The popularity is calculated by algorithm and is based, in the most part, on the total number of plays the track has had and how recent those plays are.<br>Generally speaking, songs that are being played a lot now will have a higher popularity than songs that were played a lot in the past. Duplicate tracks (e.g. the same track from a single and an album) are rated independently. Artist and album popularity is derived mathematically from track popularity. _**Note**: the popularity value may lag actual popularity by a few days: the value is not updated in real time._ 
	Popularity *int32 `json:"popularity,omitempty"`
	// A link to a 30 second preview (MP3 format) of the track. Can be `null` 
	PreviewUrl *string `json:"preview_url,omitempty"`
	// The number of the track. If an album has several discs, the track number is the number on the specified disc. 
	TrackNumber *int32 `json:"track_number,omitempty"`
	// The object type: \"track\". 
	Type *string `json:"type,omitempty"`
	// The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the track. 
	Uri *string `json:"uri,omitempty"`
	// Whether or not the track is from a local file. 
	IsLocal *bool `json:"is_local,omitempty"`
}

// NewSavedTrackObjectTrack instantiates a new SavedTrackObjectTrack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedTrackObjectTrack() *SavedTrackObjectTrack {
	this := SavedTrackObjectTrack{}
	return &this
}

// NewSavedTrackObjectTrackWithDefaults instantiates a new SavedTrackObjectTrack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedTrackObjectTrackWithDefaults() *SavedTrackObjectTrack {
	this := SavedTrackObjectTrack{}
	return &this
}

// GetAlbum returns the Album field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetAlbum() TrackObjectAlbum {
	if o == nil || isNil(o.Album) {
		var ret TrackObjectAlbum
		return ret
	}
	return *o.Album
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetAlbumOk() (*TrackObjectAlbum, bool) {
	if o == nil || isNil(o.Album) {
		return nil, false
	}
	return o.Album, true
}

// HasAlbum returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasAlbum() bool {
	if o != nil && !isNil(o.Album) {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given TrackObjectAlbum and assigns it to the Album field.
func (o *SavedTrackObjectTrack) SetAlbum(v TrackObjectAlbum) {
	o.Album = &v
}

// GetArtists returns the Artists field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetArtists() []ArtistObject {
	if o == nil || isNil(o.Artists) {
		var ret []ArtistObject
		return ret
	}
	return o.Artists
}

// GetArtistsOk returns a tuple with the Artists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetArtistsOk() ([]ArtistObject, bool) {
	if o == nil || isNil(o.Artists) {
		return nil, false
	}
	return o.Artists, true
}

// HasArtists returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasArtists() bool {
	if o != nil && !isNil(o.Artists) {
		return true
	}

	return false
}

// SetArtists gets a reference to the given []ArtistObject and assigns it to the Artists field.
func (o *SavedTrackObjectTrack) SetArtists(v []ArtistObject) {
	o.Artists = v
}

// GetAvailableMarkets returns the AvailableMarkets field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetAvailableMarkets() []string {
	if o == nil || isNil(o.AvailableMarkets) {
		var ret []string
		return ret
	}
	return o.AvailableMarkets
}

// GetAvailableMarketsOk returns a tuple with the AvailableMarkets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetAvailableMarketsOk() ([]string, bool) {
	if o == nil || isNil(o.AvailableMarkets) {
		return nil, false
	}
	return o.AvailableMarkets, true
}

// HasAvailableMarkets returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasAvailableMarkets() bool {
	if o != nil && !isNil(o.AvailableMarkets) {
		return true
	}

	return false
}

// SetAvailableMarkets gets a reference to the given []string and assigns it to the AvailableMarkets field.
func (o *SavedTrackObjectTrack) SetAvailableMarkets(v []string) {
	o.AvailableMarkets = v
}

// GetDiscNumber returns the DiscNumber field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetDiscNumber() int32 {
	if o == nil || isNil(o.DiscNumber) {
		var ret int32
		return ret
	}
	return *o.DiscNumber
}

// GetDiscNumberOk returns a tuple with the DiscNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetDiscNumberOk() (*int32, bool) {
	if o == nil || isNil(o.DiscNumber) {
		return nil, false
	}
	return o.DiscNumber, true
}

// HasDiscNumber returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasDiscNumber() bool {
	if o != nil && !isNil(o.DiscNumber) {
		return true
	}

	return false
}

// SetDiscNumber gets a reference to the given int32 and assigns it to the DiscNumber field.
func (o *SavedTrackObjectTrack) SetDiscNumber(v int32) {
	o.DiscNumber = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetDurationMs() int32 {
	if o == nil || isNil(o.DurationMs) {
		var ret int32
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetDurationMsOk() (*int32, bool) {
	if o == nil || isNil(o.DurationMs) {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasDurationMs() bool {
	if o != nil && !isNil(o.DurationMs) {
		return true
	}

	return false
}

// SetDurationMs gets a reference to the given int32 and assigns it to the DurationMs field.
func (o *SavedTrackObjectTrack) SetDurationMs(v int32) {
	o.DurationMs = &v
}

// GetExplicit returns the Explicit field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetExplicit() bool {
	if o == nil || isNil(o.Explicit) {
		var ret bool
		return ret
	}
	return *o.Explicit
}

// GetExplicitOk returns a tuple with the Explicit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetExplicitOk() (*bool, bool) {
	if o == nil || isNil(o.Explicit) {
		return nil, false
	}
	return o.Explicit, true
}

// HasExplicit returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasExplicit() bool {
	if o != nil && !isNil(o.Explicit) {
		return true
	}

	return false
}

// SetExplicit gets a reference to the given bool and assigns it to the Explicit field.
func (o *SavedTrackObjectTrack) SetExplicit(v bool) {
	o.Explicit = &v
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetExternalIds() TrackObjectExternalIds {
	if o == nil || isNil(o.ExternalIds) {
		var ret TrackObjectExternalIds
		return ret
	}
	return *o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetExternalIdsOk() (*TrackObjectExternalIds, bool) {
	if o == nil || isNil(o.ExternalIds) {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasExternalIds() bool {
	if o != nil && !isNil(o.ExternalIds) {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given TrackObjectExternalIds and assigns it to the ExternalIds field.
func (o *SavedTrackObjectTrack) SetExternalIds(v TrackObjectExternalIds) {
	o.ExternalIds = &v
}

// GetExternalUrls returns the ExternalUrls field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetExternalUrls() LinkedTrackObjectExternalUrls {
	if o == nil || isNil(o.ExternalUrls) {
		var ret LinkedTrackObjectExternalUrls
		return ret
	}
	return *o.ExternalUrls
}

// GetExternalUrlsOk returns a tuple with the ExternalUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetExternalUrlsOk() (*LinkedTrackObjectExternalUrls, bool) {
	if o == nil || isNil(o.ExternalUrls) {
		return nil, false
	}
	return o.ExternalUrls, true
}

// HasExternalUrls returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasExternalUrls() bool {
	if o != nil && !isNil(o.ExternalUrls) {
		return true
	}

	return false
}

// SetExternalUrls gets a reference to the given LinkedTrackObjectExternalUrls and assigns it to the ExternalUrls field.
func (o *SavedTrackObjectTrack) SetExternalUrls(v LinkedTrackObjectExternalUrls) {
	o.ExternalUrls = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SavedTrackObjectTrack) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SavedTrackObjectTrack) SetId(v string) {
	o.Id = &v
}

// GetIsPlayable returns the IsPlayable field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetIsPlayable() bool {
	if o == nil || isNil(o.IsPlayable) {
		var ret bool
		return ret
	}
	return *o.IsPlayable
}

// GetIsPlayableOk returns a tuple with the IsPlayable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetIsPlayableOk() (*bool, bool) {
	if o == nil || isNil(o.IsPlayable) {
		return nil, false
	}
	return o.IsPlayable, true
}

// HasIsPlayable returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasIsPlayable() bool {
	if o != nil && !isNil(o.IsPlayable) {
		return true
	}

	return false
}

// SetIsPlayable gets a reference to the given bool and assigns it to the IsPlayable field.
func (o *SavedTrackObjectTrack) SetIsPlayable(v bool) {
	o.IsPlayable = &v
}

// GetLinkedFrom returns the LinkedFrom field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetLinkedFrom() TrackObjectLinkedFrom {
	if o == nil || isNil(o.LinkedFrom) {
		var ret TrackObjectLinkedFrom
		return ret
	}
	return *o.LinkedFrom
}

// GetLinkedFromOk returns a tuple with the LinkedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetLinkedFromOk() (*TrackObjectLinkedFrom, bool) {
	if o == nil || isNil(o.LinkedFrom) {
		return nil, false
	}
	return o.LinkedFrom, true
}

// HasLinkedFrom returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasLinkedFrom() bool {
	if o != nil && !isNil(o.LinkedFrom) {
		return true
	}

	return false
}

// SetLinkedFrom gets a reference to the given TrackObjectLinkedFrom and assigns it to the LinkedFrom field.
func (o *SavedTrackObjectTrack) SetLinkedFrom(v TrackObjectLinkedFrom) {
	o.LinkedFrom = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetRestrictions() SimplifiedTrackObjectRestrictions {
	if o == nil || isNil(o.Restrictions) {
		var ret SimplifiedTrackObjectRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetRestrictionsOk() (*SimplifiedTrackObjectRestrictions, bool) {
	if o == nil || isNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasRestrictions() bool {
	if o != nil && !isNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given SimplifiedTrackObjectRestrictions and assigns it to the Restrictions field.
func (o *SavedTrackObjectTrack) SetRestrictions(v SimplifiedTrackObjectRestrictions) {
	o.Restrictions = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SavedTrackObjectTrack) SetName(v string) {
	o.Name = &v
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetPopularity() int32 {
	if o == nil || isNil(o.Popularity) {
		var ret int32
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetPopularityOk() (*int32, bool) {
	if o == nil || isNil(o.Popularity) {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasPopularity() bool {
	if o != nil && !isNil(o.Popularity) {
		return true
	}

	return false
}

// SetPopularity gets a reference to the given int32 and assigns it to the Popularity field.
func (o *SavedTrackObjectTrack) SetPopularity(v int32) {
	o.Popularity = &v
}

// GetPreviewUrl returns the PreviewUrl field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetPreviewUrl() string {
	if o == nil || isNil(o.PreviewUrl) {
		var ret string
		return ret
	}
	return *o.PreviewUrl
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetPreviewUrlOk() (*string, bool) {
	if o == nil || isNil(o.PreviewUrl) {
		return nil, false
	}
	return o.PreviewUrl, true
}

// HasPreviewUrl returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasPreviewUrl() bool {
	if o != nil && !isNil(o.PreviewUrl) {
		return true
	}

	return false
}

// SetPreviewUrl gets a reference to the given string and assigns it to the PreviewUrl field.
func (o *SavedTrackObjectTrack) SetPreviewUrl(v string) {
	o.PreviewUrl = &v
}

// GetTrackNumber returns the TrackNumber field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetTrackNumber() int32 {
	if o == nil || isNil(o.TrackNumber) {
		var ret int32
		return ret
	}
	return *o.TrackNumber
}

// GetTrackNumberOk returns a tuple with the TrackNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetTrackNumberOk() (*int32, bool) {
	if o == nil || isNil(o.TrackNumber) {
		return nil, false
	}
	return o.TrackNumber, true
}

// HasTrackNumber returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasTrackNumber() bool {
	if o != nil && !isNil(o.TrackNumber) {
		return true
	}

	return false
}

// SetTrackNumber gets a reference to the given int32 and assigns it to the TrackNumber field.
func (o *SavedTrackObjectTrack) SetTrackNumber(v int32) {
	o.TrackNumber = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SavedTrackObjectTrack) SetType(v string) {
	o.Type = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *SavedTrackObjectTrack) SetUri(v string) {
	o.Uri = &v
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *SavedTrackObjectTrack) GetIsLocal() bool {
	if o == nil || isNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedTrackObjectTrack) GetIsLocalOk() (*bool, bool) {
	if o == nil || isNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *SavedTrackObjectTrack) HasIsLocal() bool {
	if o != nil && !isNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *SavedTrackObjectTrack) SetIsLocal(v bool) {
	o.IsLocal = &v
}

func (o SavedTrackObjectTrack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SavedTrackObjectTrack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Album) {
		toSerialize["album"] = o.Album
	}
	if !isNil(o.Artists) {
		toSerialize["artists"] = o.Artists
	}
	if !isNil(o.AvailableMarkets) {
		toSerialize["available_markets"] = o.AvailableMarkets
	}
	if !isNil(o.DiscNumber) {
		toSerialize["disc_number"] = o.DiscNumber
	}
	if !isNil(o.DurationMs) {
		toSerialize["duration_ms"] = o.DurationMs
	}
	if !isNil(o.Explicit) {
		toSerialize["explicit"] = o.Explicit
	}
	if !isNil(o.ExternalIds) {
		toSerialize["external_ids"] = o.ExternalIds
	}
	if !isNil(o.ExternalUrls) {
		toSerialize["external_urls"] = o.ExternalUrls
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsPlayable) {
		toSerialize["is_playable"] = o.IsPlayable
	}
	if !isNil(o.LinkedFrom) {
		toSerialize["linked_from"] = o.LinkedFrom
	}
	if !isNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Popularity) {
		toSerialize["popularity"] = o.Popularity
	}
	if !isNil(o.PreviewUrl) {
		toSerialize["preview_url"] = o.PreviewUrl
	}
	if !isNil(o.TrackNumber) {
		toSerialize["track_number"] = o.TrackNumber
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !isNil(o.IsLocal) {
		toSerialize["is_local"] = o.IsLocal
	}
	return toSerialize, nil
}

type NullableSavedTrackObjectTrack struct {
	value *SavedTrackObjectTrack
	isSet bool
}

func (v NullableSavedTrackObjectTrack) Get() *SavedTrackObjectTrack {
	return v.value
}

func (v *NullableSavedTrackObjectTrack) Set(val *SavedTrackObjectTrack) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedTrackObjectTrack) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedTrackObjectTrack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedTrackObjectTrack(val *SavedTrackObjectTrack) *NullableSavedTrackObjectTrack {
	return &NullableSavedTrackObjectTrack{value: val, isSet: true}
}

func (v NullableSavedTrackObjectTrack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedTrackObjectTrack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


