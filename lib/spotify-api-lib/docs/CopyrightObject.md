# CopyrightObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | The copyright text for this content.  | [optional] 
**Type** | Pointer to **string** | The type of copyright: &#x60;C&#x60; &#x3D; the copyright, &#x60;P&#x60; &#x3D; the sound recording (performance) copyright.  | [optional] 

## Methods

### NewCopyrightObject

`func NewCopyrightObject() *CopyrightObject`

NewCopyrightObject instantiates a new CopyrightObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyrightObjectWithDefaults

`func NewCopyrightObjectWithDefaults() *CopyrightObject`

NewCopyrightObjectWithDefaults instantiates a new CopyrightObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CopyrightObject) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CopyrightObject) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CopyrightObject) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CopyrightObject) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *CopyrightObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CopyrightObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CopyrightObject) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CopyrightObject) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


