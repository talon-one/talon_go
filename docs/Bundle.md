# Bundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An identifier derived from the bundle content. | 
**Name** | Pointer to **string** | The name of the bundle. | 
**Type** | Pointer to **string** | A binding of type &#x60;bundle&#x60;. | 
**Sources** | Pointer to **[]string** | The selector sources of bundle items. Each source is expressed as a &#x60;{{$selectorName}}&#x60; reference. | 
**Counts** | Pointer to **[]int64** | The number of items to retrieve from each corresponding source in &#x60;sources&#x60;. | 
**Matchers** | Pointer to **[]string** | Attribute names that the bundled items must share. | [optional] 

## Methods

### NewBundle

`func NewBundle(id string, name string, type_ string, sources []string, counts []int64, ) *Bundle`

NewBundle instantiates a new Bundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleWithDefaults

`func NewBundleWithDefaults() *Bundle`

NewBundleWithDefaults instantiates a new Bundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bundle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bundle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bundle) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Bundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bundle) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Bundle) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bundle) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bundle) SetType(v string)`

SetType sets Type field to given value.


### GetSources

`func (o *Bundle) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Bundle) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Bundle) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetCounts

`func (o *Bundle) GetCounts() []int64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *Bundle) GetCountsOk() (*[]int64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *Bundle) SetCounts(v []int64)`

SetCounts sets Counts field to given value.


### GetMatchers

`func (o *Bundle) GetMatchers() []string`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *Bundle) GetMatchersOk() (*[]string, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *Bundle) SetMatchers(v []string)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *Bundle) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


