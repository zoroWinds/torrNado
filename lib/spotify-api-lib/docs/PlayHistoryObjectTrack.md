# PlayHistoryObjectTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artists** | Pointer to [**[]SimplifiedArtistObject**](SimplifiedArtistObject.md) | The artists who performed the track. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist. | [optional] 
**AvailableMarkets** | Pointer to **[]string** | A list of the countries in which the track can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | [optional] 
**DiscNumber** | Pointer to **int32** | The disc number (usually &#x60;1&#x60; unless the album consists of more than one disc). | [optional] 
**DurationMs** | Pointer to **int32** | The track length in milliseconds. | [optional] 
**Explicit** | Pointer to **bool** | Whether or not the track has explicit lyrics ( &#x60;true&#x60; &#x3D; yes it does; &#x60;false&#x60; &#x3D; no it does not OR unknown). | [optional] 
**ExternalUrls** | Pointer to [**SimplifiedTrackObjectExternalUrls**](SimplifiedTrackObjectExternalUrls.md) |  | [optional] 
**Href** | Pointer to **string** | A link to the Web API endpoint providing full details of the track. | [optional] 
**Id** | Pointer to **string** | The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the track.  | [optional] 
**IsPlayable** | Pointer to **bool** | Part of the response when [Track Relinking](/documentation/general/guides/track-relinking-guide/) is applied. If &#x60;true&#x60;, the track is playable in the given market. Otherwise &#x60;false&#x60;.  | [optional] 
**LinkedFrom** | Pointer to [**SimplifiedTrackObjectLinkedFrom**](SimplifiedTrackObjectLinkedFrom.md) |  | [optional] 
**Restrictions** | Pointer to [**SimplifiedTrackObjectRestrictions**](SimplifiedTrackObjectRestrictions.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the track. | [optional] 
**PreviewUrl** | Pointer to **string** | A URL to a 30 second preview (MP3 format) of the track.  | [optional] 
**TrackNumber** | Pointer to **int32** | The number of the track. If an album has several discs, the track number is the number on the specified disc.  | [optional] 
**Type** | Pointer to **string** | The object type: \&quot;track\&quot;.  | [optional] 
**Uri** | Pointer to **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the track.  | [optional] 
**IsLocal** | Pointer to **bool** | Whether or not the track is from a local file.  | [optional] 

## Methods

### NewPlayHistoryObjectTrack

`func NewPlayHistoryObjectTrack() *PlayHistoryObjectTrack`

NewPlayHistoryObjectTrack instantiates a new PlayHistoryObjectTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayHistoryObjectTrackWithDefaults

`func NewPlayHistoryObjectTrackWithDefaults() *PlayHistoryObjectTrack`

NewPlayHistoryObjectTrackWithDefaults instantiates a new PlayHistoryObjectTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtists

`func (o *PlayHistoryObjectTrack) GetArtists() []SimplifiedArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *PlayHistoryObjectTrack) GetArtistsOk() (*[]SimplifiedArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *PlayHistoryObjectTrack) SetArtists(v []SimplifiedArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *PlayHistoryObjectTrack) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetAvailableMarkets

`func (o *PlayHistoryObjectTrack) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *PlayHistoryObjectTrack) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *PlayHistoryObjectTrack) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.

### HasAvailableMarkets

`func (o *PlayHistoryObjectTrack) HasAvailableMarkets() bool`

HasAvailableMarkets returns a boolean if a field has been set.

### GetDiscNumber

`func (o *PlayHistoryObjectTrack) GetDiscNumber() int32`

GetDiscNumber returns the DiscNumber field if non-nil, zero value otherwise.

### GetDiscNumberOk

`func (o *PlayHistoryObjectTrack) GetDiscNumberOk() (*int32, bool)`

GetDiscNumberOk returns a tuple with the DiscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscNumber

`func (o *PlayHistoryObjectTrack) SetDiscNumber(v int32)`

SetDiscNumber sets DiscNumber field to given value.

### HasDiscNumber

`func (o *PlayHistoryObjectTrack) HasDiscNumber() bool`

HasDiscNumber returns a boolean if a field has been set.

### GetDurationMs

`func (o *PlayHistoryObjectTrack) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *PlayHistoryObjectTrack) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *PlayHistoryObjectTrack) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *PlayHistoryObjectTrack) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetExplicit

`func (o *PlayHistoryObjectTrack) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *PlayHistoryObjectTrack) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *PlayHistoryObjectTrack) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.

### HasExplicit

`func (o *PlayHistoryObjectTrack) HasExplicit() bool`

HasExplicit returns a boolean if a field has been set.

### GetExternalUrls

`func (o *PlayHistoryObjectTrack) GetExternalUrls() SimplifiedTrackObjectExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *PlayHistoryObjectTrack) GetExternalUrlsOk() (*SimplifiedTrackObjectExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *PlayHistoryObjectTrack) SetExternalUrls(v SimplifiedTrackObjectExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *PlayHistoryObjectTrack) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetHref

`func (o *PlayHistoryObjectTrack) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PlayHistoryObjectTrack) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PlayHistoryObjectTrack) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PlayHistoryObjectTrack) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PlayHistoryObjectTrack) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlayHistoryObjectTrack) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlayHistoryObjectTrack) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlayHistoryObjectTrack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPlayable

`func (o *PlayHistoryObjectTrack) GetIsPlayable() bool`

GetIsPlayable returns the IsPlayable field if non-nil, zero value otherwise.

### GetIsPlayableOk

`func (o *PlayHistoryObjectTrack) GetIsPlayableOk() (*bool, bool)`

GetIsPlayableOk returns a tuple with the IsPlayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayable

`func (o *PlayHistoryObjectTrack) SetIsPlayable(v bool)`

SetIsPlayable sets IsPlayable field to given value.

### HasIsPlayable

`func (o *PlayHistoryObjectTrack) HasIsPlayable() bool`

HasIsPlayable returns a boolean if a field has been set.

### GetLinkedFrom

`func (o *PlayHistoryObjectTrack) GetLinkedFrom() SimplifiedTrackObjectLinkedFrom`

GetLinkedFrom returns the LinkedFrom field if non-nil, zero value otherwise.

### GetLinkedFromOk

`func (o *PlayHistoryObjectTrack) GetLinkedFromOk() (*SimplifiedTrackObjectLinkedFrom, bool)`

GetLinkedFromOk returns a tuple with the LinkedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedFrom

`func (o *PlayHistoryObjectTrack) SetLinkedFrom(v SimplifiedTrackObjectLinkedFrom)`

SetLinkedFrom sets LinkedFrom field to given value.

### HasLinkedFrom

`func (o *PlayHistoryObjectTrack) HasLinkedFrom() bool`

HasLinkedFrom returns a boolean if a field has been set.

### GetRestrictions

`func (o *PlayHistoryObjectTrack) GetRestrictions() SimplifiedTrackObjectRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *PlayHistoryObjectTrack) GetRestrictionsOk() (*SimplifiedTrackObjectRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *PlayHistoryObjectTrack) SetRestrictions(v SimplifiedTrackObjectRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *PlayHistoryObjectTrack) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetName

`func (o *PlayHistoryObjectTrack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlayHistoryObjectTrack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlayHistoryObjectTrack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlayHistoryObjectTrack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *PlayHistoryObjectTrack) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *PlayHistoryObjectTrack) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *PlayHistoryObjectTrack) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *PlayHistoryObjectTrack) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetTrackNumber

`func (o *PlayHistoryObjectTrack) GetTrackNumber() int32`

GetTrackNumber returns the TrackNumber field if non-nil, zero value otherwise.

### GetTrackNumberOk

`func (o *PlayHistoryObjectTrack) GetTrackNumberOk() (*int32, bool)`

GetTrackNumberOk returns a tuple with the TrackNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackNumber

`func (o *PlayHistoryObjectTrack) SetTrackNumber(v int32)`

SetTrackNumber sets TrackNumber field to given value.

### HasTrackNumber

`func (o *PlayHistoryObjectTrack) HasTrackNumber() bool`

HasTrackNumber returns a boolean if a field has been set.

### GetType

`func (o *PlayHistoryObjectTrack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayHistoryObjectTrack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayHistoryObjectTrack) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlayHistoryObjectTrack) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *PlayHistoryObjectTrack) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PlayHistoryObjectTrack) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PlayHistoryObjectTrack) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PlayHistoryObjectTrack) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetIsLocal

`func (o *PlayHistoryObjectTrack) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *PlayHistoryObjectTrack) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *PlayHistoryObjectTrack) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *PlayHistoryObjectTrack) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


