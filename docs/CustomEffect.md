# CustomEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications that are related to this entity. | 
**IsPerItem** | Pointer to **bool** | Indicates if this effect is per item or not. | [optional] 
**Name** | Pointer to **string** | The name of this effect. | 
**Title** | Pointer to **string** | The title of this effect. | 
**Payload** | Pointer to **string** | The JSON payload of this effect. | 
**Description** | Pointer to **string** | The description of this effect. | [optional] 
**Enabled** | Pointer to **bool** | Determines if this effect is active. | 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | [optional] 
**ModifiedBy** | Pointer to **int64** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | Pointer to **int64** | ID of the user who created this effect. | 

## Methods

### NewCustomEffect

`func NewCustomEffect(id int64, created time.Time, accountId int64, modified time.Time, applicationIds []int64, name string, title string, payload string, enabled bool, createdBy int64, ) *CustomEffect`

NewCustomEffect instantiates a new CustomEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEffectWithDefaults

`func NewCustomEffectWithDefaults() *CustomEffect`

NewCustomEffectWithDefaults instantiates a new CustomEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomEffect) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEffect) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomEffect) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CustomEffect) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomEffect) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomEffect) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *CustomEffect) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomEffect) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CustomEffect) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetModified

`func (o *CustomEffect) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomEffect) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CustomEffect) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetApplicationIds

`func (o *CustomEffect) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *CustomEffect) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *CustomEffect) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetIsPerItem

`func (o *CustomEffect) GetIsPerItem() bool`

GetIsPerItem returns the IsPerItem field if non-nil, zero value otherwise.

### GetIsPerItemOk

`func (o *CustomEffect) GetIsPerItemOk() (*bool, bool)`

GetIsPerItemOk returns a tuple with the IsPerItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerItem

`func (o *CustomEffect) SetIsPerItem(v bool)`

SetIsPerItem sets IsPerItem field to given value.

### HasIsPerItem

`func (o *CustomEffect) HasIsPerItem() bool`

HasIsPerItem returns a boolean if a field has been set.

### GetName

`func (o *CustomEffect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEffect) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CustomEffect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEffect) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomEffect) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPayload

`func (o *CustomEffect) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffect) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomEffect) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetDescription

`func (o *CustomEffect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEffect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomEffect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomEffect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomEffect) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEffect) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomEffect) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetParams

`func (o *CustomEffect) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CustomEffect) GetParamsOk() (*[]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CustomEffect) SetParams(v []TemplateArgDef)`

SetParams sets Params field to given value.

### HasParams

`func (o *CustomEffect) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetModifiedBy

`func (o *CustomEffect) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CustomEffect) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CustomEffect) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CustomEffect) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CustomEffect) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CustomEffect) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CustomEffect) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


