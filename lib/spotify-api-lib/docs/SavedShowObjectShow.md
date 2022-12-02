# SavedShowObjectShow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableMarkets** | **[]string** | A list of the countries in which the show can be played, identified by their [ISO 3166-1 alpha-2](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code.  | 
**Copyrights** | [**[]CopyrightObject**](CopyrightObject.md) | The copyright statements of the show.  | 
**Description** | **string** | A description of the show. HTML tags are stripped away from this field, use &#x60;html_description&#x60; field in case HTML tags are needed.  | 
**HtmlDescription** | **string** | A description of the show. This field may contain HTML tags.  | 
**Explicit** | **bool** | Whether or not the show has explicit content (true &#x3D; yes it does; false &#x3D; no it does not OR unknown).  | 
**ExternalUrls** | [**ShowBaseExternalUrls**](ShowBaseExternalUrls.md) |  | 
**Href** | **string** | A link to the Web API endpoint providing full details of the show.  | 
**Id** | **string** | The [Spotify ID](/documentation/web-api/#spotify-uris-and-ids) for the show.  | 
**Images** | [**[]ImageObject**](ImageObject.md) | The cover art for the show in various sizes, widest first.  | 
**IsExternallyHosted** | **bool** | True if all of the shows episodes are hosted outside of Spotify&#39;s CDN. This field might be &#x60;null&#x60; in some cases.  | 
**Languages** | **[]string** | A list of the languages used in the show, identified by their [ISO 639](https://en.wikipedia.org/wiki/ISO_639) code.  | 
**MediaType** | **string** | The media type of the show.  | 
**Name** | **string** | The name of the episode.  | 
**Publisher** | **string** | The publisher of the show.  | 
**Type** | **string** | The object type.  | 
**Uri** | **string** | The [Spotify URI](/documentation/web-api/#spotify-uris-and-ids) for the show.  | 

## Methods

### NewSavedShowObjectShow

`func NewSavedShowObjectShow(availableMarkets []string, copyrights []CopyrightObject, description string, htmlDescription string, explicit bool, externalUrls ShowBaseExternalUrls, href string, id string, images []ImageObject, isExternallyHosted bool, languages []string, mediaType string, name string, publisher string, type_ string, uri string, ) *SavedShowObjectShow`

NewSavedShowObjectShow instantiates a new SavedShowObjectShow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedShowObjectShowWithDefaults

`func NewSavedShowObjectShowWithDefaults() *SavedShowObjectShow`

NewSavedShowObjectShowWithDefaults instantiates a new SavedShowObjectShow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableMarkets

`func (o *SavedShowObjectShow) GetAvailableMarkets() []string`

GetAvailableMarkets returns the AvailableMarkets field if non-nil, zero value otherwise.

### GetAvailableMarketsOk

`func (o *SavedShowObjectShow) GetAvailableMarketsOk() (*[]string, bool)`

GetAvailableMarketsOk returns a tuple with the AvailableMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMarkets

`func (o *SavedShowObjectShow) SetAvailableMarkets(v []string)`

SetAvailableMarkets sets AvailableMarkets field to given value.


### GetCopyrights

`func (o *SavedShowObjectShow) GetCopyrights() []CopyrightObject`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *SavedShowObjectShow) GetCopyrightsOk() (*[]CopyrightObject, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *SavedShowObjectShow) SetCopyrights(v []CopyrightObject)`

SetCopyrights sets Copyrights field to given value.


### GetDescription

`func (o *SavedShowObjectShow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedShowObjectShow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedShowObjectShow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHtmlDescription

`func (o *SavedShowObjectShow) GetHtmlDescription() string`

GetHtmlDescription returns the HtmlDescription field if non-nil, zero value otherwise.

### GetHtmlDescriptionOk

`func (o *SavedShowObjectShow) GetHtmlDescriptionOk() (*string, bool)`

GetHtmlDescriptionOk returns a tuple with the HtmlDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlDescription

`func (o *SavedShowObjectShow) SetHtmlDescription(v string)`

SetHtmlDescription sets HtmlDescription field to given value.


### GetExplicit

`func (o *SavedShowObjectShow) GetExplicit() bool`

GetExplicit returns the Explicit field if non-nil, zero value otherwise.

### GetExplicitOk

`func (o *SavedShowObjectShow) GetExplicitOk() (*bool, bool)`

GetExplicitOk returns a tuple with the Explicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicit

`func (o *SavedShowObjectShow) SetExplicit(v bool)`

SetExplicit sets Explicit field to given value.


### GetExternalUrls

`func (o *SavedShowObjectShow) GetExternalUrls() ShowBaseExternalUrls`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *SavedShowObjectShow) GetExternalUrlsOk() (*ShowBaseExternalUrls, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *SavedShowObjectShow) SetExternalUrls(v ShowBaseExternalUrls)`

SetExternalUrls sets ExternalUrls field to given value.


### GetHref

`func (o *SavedShowObjectShow) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SavedShowObjectShow) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SavedShowObjectShow) SetHref(v string)`

SetHref sets Href field to given value.


### GetId

`func (o *SavedShowObjectShow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedShowObjectShow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedShowObjectShow) SetId(v string)`

SetId sets Id field to given value.


### GetImages

`func (o *SavedShowObjectShow) GetImages() []ImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *SavedShowObjectShow) GetImagesOk() (*[]ImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *SavedShowObjectShow) SetImages(v []ImageObject)`

SetImages sets Images field to given value.


### GetIsExternallyHosted

`func (o *SavedShowObjectShow) GetIsExternallyHosted() bool`

GetIsExternallyHosted returns the IsExternallyHosted field if non-nil, zero value otherwise.

### GetIsExternallyHostedOk

`func (o *SavedShowObjectShow) GetIsExternallyHostedOk() (*bool, bool)`

GetIsExternallyHostedOk returns a tuple with the IsExternallyHosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternallyHosted

`func (o *SavedShowObjectShow) SetIsExternallyHosted(v bool)`

SetIsExternallyHosted sets IsExternallyHosted field to given value.


### GetLanguages

`func (o *SavedShowObjectShow) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *SavedShowObjectShow) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *SavedShowObjectShow) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.


### GetMediaType

`func (o *SavedShowObjectShow) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SavedShowObjectShow) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SavedShowObjectShow) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetName

`func (o *SavedShowObjectShow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedShowObjectShow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedShowObjectShow) SetName(v string)`

SetName sets Name field to given value.


### GetPublisher

`func (o *SavedShowObjectShow) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *SavedShowObjectShow) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *SavedShowObjectShow) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.


### GetType

`func (o *SavedShowObjectShow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedShowObjectShow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedShowObjectShow) SetType(v string)`

SetType sets Type field to given value.


### GetUri

`func (o *SavedShowObjectShow) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *SavedShowObjectShow) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *SavedShowObjectShow) SetUri(v string)`

SetUri sets Uri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


