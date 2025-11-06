# NewCustomEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications that are related to this entity. | 
**IsPerItem** | Pointer to **bool** | Indicates if this effect is per item or not. | [optional] 
**Name** | Pointer to **string** | The name of this effect. | 
**Title** | Pointer to **string** | The title of this effect. | 
**Payload** | Pointer to **string** | The JSON payload of this effect. | 
**Description** | Pointer to **string** | The description of this effect. | [optional] 
**Enabled** | Pointer to **bool** | Determines if this effect is active. | 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | [optional] 

## Methods

### NewNewCustomEffect

`func NewNewCustomEffect(applicationIds []int64, name string, title string, payload string, enabled bool, ) *NewCustomEffect`

NewNewCustomEffect instantiates a new NewCustomEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCustomEffectWithDefaults

`func NewNewCustomEffectWithDefaults() *NewCustomEffect`

NewNewCustomEffectWithDefaults instantiates a new NewCustomEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *NewCustomEffect) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewCustomEffect) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewCustomEffect) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetIsPerItem

`func (o *NewCustomEffect) GetIsPerItem() bool`

GetIsPerItem returns the IsPerItem field if non-nil, zero value otherwise.

### GetIsPerItemOk

`func (o *NewCustomEffect) GetIsPerItemOk() (*bool, bool)`

GetIsPerItemOk returns a tuple with the IsPerItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerItem

`func (o *NewCustomEffect) SetIsPerItem(v bool)`

SetIsPerItem sets IsPerItem field to given value.

### HasIsPerItem

`func (o *NewCustomEffect) HasIsPerItem() bool`

HasIsPerItem returns a boolean if a field has been set.

### GetName

`func (o *NewCustomEffect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewCustomEffect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewCustomEffect) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *NewCustomEffect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewCustomEffect) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewCustomEffect) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPayload

`func (o *NewCustomEffect) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NewCustomEffect) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NewCustomEffect) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetDescription

`func (o *NewCustomEffect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCustomEffect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCustomEffect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewCustomEffect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NewCustomEffect) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewCustomEffect) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewCustomEffect) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParams

`func (o *NewCustomEffect) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NewCustomEffect) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NewCustomEffect) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.

### HasParams

`func (o *NewCustomEffect) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


