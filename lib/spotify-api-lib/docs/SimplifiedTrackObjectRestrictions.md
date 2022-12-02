# SimplifiedTrackObjectRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason for the restriction. Supported values:&lt;br&gt; - &#x60;market&#x60; - The content item is not available in the given market.&lt;br&gt; - &#x60;product&#x60; - The content item is not available for the user&#39;s subscription type.&lt;br&gt; - &#x60;explicit&#x60; - The content item is explicit and the user&#39;s account is set to not play explicit content.&lt;br&gt; Additional reasons may be added in the future. **Note**: If you use this field, make sure that your application safely handles unknown values.  | [optional] 

## Methods

### NewSimplifiedTrackObjectRestrictions

`func NewSimplifiedTrackObjectRestrictions() *SimplifiedTrackObjectRestrictions`

NewSimplifiedTrackObjectRestrictions instantiates a new SimplifiedTrackObjectRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedTrackObjectRestrictionsWithDefaults

`func NewSimplifiedTrackObjectRestrictionsWithDefaults() *SimplifiedTrackObjectRestrictions`

NewSimplifiedTrackObjectRestrictionsWithDefaults instantiates a new SimplifiedTrackObjectRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *SimplifiedTrackObjectRestrictions) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SimplifiedTrackObjectRestrictions) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SimplifiedTrackObjectRestrictions) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SimplifiedTrackObjectRestrictions) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


