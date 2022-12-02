# AlbumObjectAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artists** | Pointer to [**[]ArtistObject**](ArtistObject.md) | The artists of the album. Each artist object includes a link in &#x60;href&#x60; to more detailed information about the artist.  | [optional] 
**Tracks** | Pointer to [**PagingObject**](PagingObject.md) | The tracks of the album.  | [optional] 

## Methods

### NewAlbumObjectAllOf

`func NewAlbumObjectAllOf() *AlbumObjectAllOf`

NewAlbumObjectAllOf instantiates a new AlbumObjectAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumObjectAllOfWithDefaults

`func NewAlbumObjectAllOfWithDefaults() *AlbumObjectAllOf`

NewAlbumObjectAllOfWithDefaults instantiates a new AlbumObjectAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtists

`func (o *AlbumObjectAllOf) GetArtists() []ArtistObject`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *AlbumObjectAllOf) GetArtistsOk() (*[]ArtistObject, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *AlbumObjectAllOf) SetArtists(v []ArtistObject)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *AlbumObjectAllOf) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetTracks

`func (o *AlbumObjectAllOf) GetTracks() PagingObject`

GetTracks returns the Tracks field if non-nil, zero value otherwise.

### GetTracksOk

`func (o *AlbumObjectAllOf) GetTracksOk() (*PagingObject, bool)`

GetTracksOk returns a tuple with the Tracks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracks

`func (o *AlbumObjectAllOf) SetTracks(v PagingObject)`

SetTracks sets Tracks field to given value.

### HasTracks

`func (o *AlbumObjectAllOf) HasTracks() bool`

HasTracks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


