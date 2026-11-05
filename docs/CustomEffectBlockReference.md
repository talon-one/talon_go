# CustomEffectBlockReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the custom effect. | 
**Name** | Pointer to **string** | The name of the custom effect, as used in API requests. | [optional] [readonly] 
**Title** | Pointer to **string** | The display name of the custom effect. | [optional] [readonly] 

## Methods

### NewCustomEffectBlockReference

`func NewCustomEffectBlockReference(id int64, ) *CustomEffectBlockReference`

NewCustomEffectBlockReference instantiates a new CustomEffectBlockReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEffectBlockReferenceWithDefaults

`func NewCustomEffectBlockReferenceWithDefaults() *CustomEffectBlockReference`

NewCustomEffectBlockReferenceWithDefaults instantiates a new CustomEffectBlockReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEffectBlockReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEffectBlockReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEffectBlockReference) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CustomEffectBlockReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffectBlockReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEffectBlockReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomEffectBlockReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *CustomEffectBlockReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEffectBlockReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEffectBlockReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomEffectBlockReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


