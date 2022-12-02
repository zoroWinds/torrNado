# CursorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | The cursor to use as key to find the next page of items. | [optional] 

## Methods

### NewCursorObject

`func NewCursorObject() *CursorObject`

NewCursorObject instantiates a new CursorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCursorObjectWithDefaults

`func NewCursorObjectWithDefaults() *CursorObject`

NewCursorObjectWithDefaults instantiates a new CursorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CursorObject) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CursorObject) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CursorObject) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CursorObject) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


