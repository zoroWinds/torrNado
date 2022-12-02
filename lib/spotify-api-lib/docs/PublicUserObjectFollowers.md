# PublicUserObjectFollowers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **NullableString** | This will always be set to null, as the Web API does not support it at the moment.  | [optional] 
**Total** | Pointer to **int32** | The total number of followers.  | [optional] 

## Methods

### NewPublicUserObjectFollowers

`func NewPublicUserObjectFollowers() *PublicUserObjectFollowers`

NewPublicUserObjectFollowers instantiates a new PublicUserObjectFollowers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUserObjectFollowersWithDefaults

`func NewPublicUserObjectFollowersWithDefaults() *PublicUserObjectFollowers`

NewPublicUserObjectFollowersWithDefaults instantiates a new PublicUserObjectFollowers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PublicUserObjectFollowers) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PublicUserObjectFollowers) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PublicUserObjectFollowers) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PublicUserObjectFollowers) HasHref() bool`

HasHref returns a boolean if a field has been set.

### SetHrefNil

`func (o *PublicUserObjectFollowers) SetHrefNil(b bool)`

 SetHrefNil sets the value for Href to be an explicit nil

### UnsetHref
`func (o *PublicUserObjectFollowers) UnsetHref()`

UnsetHref ensures that no value is present for Href, not even an explicit nil
### GetTotal

`func (o *PublicUserObjectFollowers) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PublicUserObjectFollowers) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PublicUserObjectFollowers) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PublicUserObjectFollowers) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


