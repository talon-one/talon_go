# NewPriceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The API name of the price type. This is an immutable value. | 
**Title** | Pointer to **string** | The name displayed in the Campaign Manager for the price type. | 
**Description** | Pointer to **string** | A description of the price type. | [optional] 
**TargetedAudiencesIds** | Pointer to **[]int64** | A list of the IDs of the audiences targeted by this price type. | [optional] 

## Methods

### NewNewPriceType

`func NewNewPriceType(name string, title string, ) *NewPriceType`

NewNewPriceType instantiates a new NewPriceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewPriceTypeWithDefaults

`func NewNewPriceTypeWithDefaults() *NewPriceType`

NewNewPriceTypeWithDefaults instantiates a new NewPriceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewPriceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewPriceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewPriceType) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewPriceType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewPriceType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewPriceType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewPriceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewPriceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewPriceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewPriceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetedAudiencesIds

`func (o *NewPriceType) GetTargetedAudiencesIds() []int64`

GetTargetedAudiencesIds returns the TargetedAudiencesIds field if non-nil, zero value otherwise.

### GetTargetedAudiencesIdsOk

`func (o *NewPriceType) GetTargetedAudiencesIdsOk() (*[]int64, bool)`

GetTargetedAudiencesIdsOk returns a tuple with the TargetedAudiencesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedAudiencesIds

`func (o *NewPriceType) SetTargetedAudiencesIds(v []int64)`

SetTargetedAudiencesIds sets TargetedAudiencesIds field to given value.

### HasTargetedAudiencesIds

`func (o *NewPriceType) HasTargetedAudiencesIds() bool`

HasTargetedAudiencesIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


