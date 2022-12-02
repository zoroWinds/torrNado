# StartAUsersPlaybackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextUri** | Pointer to **map[string]interface{}** | Optional. Spotify URI of the context to play. Valid contexts are albums, artists &amp; playlists. {context_uri:\&quot;spotify:album:1Je1IMUlBXcx1Fz0WE7oPT\&quot;}  | [optional] 
**Uris** | Pointer to **[]string** | Optional. A JSON array of the Spotify track URIs to play. For example: {\&quot;uris\&quot;: [\&quot;spotify:track:4iV5W9uYEdYUVa79Axb7Rh\&quot;, \&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;]}  | [optional] 
**Offset** | Pointer to **map[string]interface{}** | Optional. Indicates from where in the context playback should start. Only available when context_uri corresponds to an album or playlist object \&quot;position\&quot; is zero based and can’t be negative. Example: \&quot;offset\&quot;: {\&quot;position\&quot;: 5} \&quot;uri\&quot; is a string representing the uri of the item to start at. Example: \&quot;offset\&quot;: {\&quot;uri\&quot;: \&quot;spotify:track:1301WleyT98MSxVHPZCA6M\&quot;}  | [optional] 
**PositionMs** | Pointer to **map[string]interface{}** | integer | [optional] 

## Methods

### NewStartAUsersPlaybackRequest

`func NewStartAUsersPlaybackRequest() *StartAUsersPlaybackRequest`

NewStartAUsersPlaybackRequest instantiates a new StartAUsersPlaybackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartAUsersPlaybackRequestWithDefaults

`func NewStartAUsersPlaybackRequestWithDefaults() *StartAUsersPlaybackRequest`

NewStartAUsersPlaybackRequestWithDefaults instantiates a new StartAUsersPlaybackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextUri

`func (o *StartAUsersPlaybackRequest) GetContextUri() map[string]interface{}`

GetContextUri returns the ContextUri field if non-nil, zero value otherwise.

### GetContextUriOk

`func (o *StartAUsersPlaybackRequest) GetContextUriOk() (*map[string]interface{}, bool)`

GetContextUriOk returns a tuple with the ContextUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextUri

`func (o *StartAUsersPlaybackRequest) SetContextUri(v map[string]interface{})`

SetContextUri sets ContextUri field to given value.

### HasContextUri

`func (o *StartAUsersPlaybackRequest) HasContextUri() bool`

HasContextUri returns a boolean if a field has been set.

### GetUris

`func (o *StartAUsersPlaybackRequest) GetUris() []string`

GetUris returns the Uris field if non-nil, zero value otherwise.

### GetUrisOk

`func (o *StartAUsersPlaybackRequest) GetUrisOk() (*[]string, bool)`

GetUrisOk returns a tuple with the Uris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUris

`func (o *StartAUsersPlaybackRequest) SetUris(v []string)`

SetUris sets Uris field to given value.

### HasUris

`func (o *StartAUsersPlaybackRequest) HasUris() bool`

HasUris returns a boolean if a field has been set.

### GetOffset

`func (o *StartAUsersPlaybackRequest) GetOffset() map[string]interface{}`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *StartAUsersPlaybackRequest) GetOffsetOk() (*map[string]interface{}, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *StartAUsersPlaybackRequest) SetOffset(v map[string]interface{})`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *StartAUsersPlaybackRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPositionMs

`func (o *StartAUsersPlaybackRequest) GetPositionMs() map[string]interface{}`

GetPositionMs returns the PositionMs field if non-nil, zero value otherwise.

### GetPositionMsOk

`func (o *StartAUsersPlaybackRequest) GetPositionMsOk() (*map[string]interface{}, bool)`

GetPositionMsOk returns a tuple with the PositionMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionMs

`func (o *StartAUsersPlaybackRequest) SetPositionMs(v map[string]interface{})`

SetPositionMs sets PositionMs field to given value.

### HasPositionMs

`func (o *StartAUsersPlaybackRequest) HasPositionMs() bool`

HasPositionMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


