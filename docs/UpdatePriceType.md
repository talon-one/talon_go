# UpdatePriceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The name displayed in the Campaign Manager for the price type. | 
**Description** | Pointer to **string** | A description of the price type. | [optional] 
**TargetedAudiencesIds** | Pointer to **[]int64** | A list of the IDs of the audiences targeted by this price type. | 

## Methods

### NewUpdatePriceType

`func NewUpdatePriceType(title string, targetedAudiencesIds []int64, ) *UpdatePriceType`

NewUpdatePriceType instantiates a new UpdatePriceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceTypeWithDefaults

`func NewUpdatePriceTypeWithDefaults() *UpdatePriceType`

NewUpdatePriceTypeWithDefaults instantiates a new UpdatePriceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdatePriceType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePriceType) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePriceType) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *UpdatePriceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePriceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePriceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePriceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetedAudiencesIds

`func (o *UpdatePriceType) GetTargetedAudiencesIds() []int64`

GetTargetedAudiencesIds returns the TargetedAudiencesIds field if non-nil, zero value otherwise.

### GetTargetedAudiencesIdsOk

`func (o *UpdatePriceType) GetTargetedAudiencesIdsOk() (*[]int64, bool)`

GetTargetedAudiencesIdsOk returns a tuple with the TargetedAudiencesIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedAudiencesIds

`func (o *UpdatePriceType) SetTargetedAudiencesIds(v []int64)`

SetTargetedAudiencesIds sets TargetedAudiencesIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


