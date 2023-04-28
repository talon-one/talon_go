# CustomEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**ApplicationIds** | Pointer to **[]int32** | The IDs of the Applications that are related to this entity. | 
**IsPerItem** | Pointer to **bool** | Indicates if this effect is per item or not. | [optional] 
**Name** | Pointer to **string** | The name of this effect. | 
**Title** | Pointer to **string** | The title of this effect. | 
**Payload** | Pointer to **string** | The JSON payload of this effect. | 
**Description** | Pointer to **string** | The description of this effect. | [optional] 
**Enabled** | Pointer to **bool** | Determines if this effect is active. | 
**Params** | Pointer to [**[]TemplateArgDef**](TemplateArgDef.md) | Array of template argument definitions. | [optional] 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 

## Methods

### GetId

`func (o *CustomEffect) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomEffect) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CustomEffect) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CustomEffect) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CustomEffect) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomEffect) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CustomEffect) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CustomEffect) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetAccountId

`func (o *CustomEffect) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CustomEffect) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CustomEffect) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CustomEffect) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetModified

`func (o *CustomEffect) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CustomEffect) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *CustomEffect) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *CustomEffect) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetApplicationIds

`func (o *CustomEffect) GetApplicationIds() []int32`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *CustomEffect) GetApplicationIdsOk() ([]int32, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationIds

`func (o *CustomEffect) HasApplicationIds() bool`

HasApplicationIds returns a boolean if a field has been set.

### SetApplicationIds

`func (o *CustomEffect) SetApplicationIds(v []int32)`

SetApplicationIds gets a reference to the given []int32 and assigns it to the ApplicationIds field.

### GetIsPerItem

`func (o *CustomEffect) GetIsPerItem() bool`

GetIsPerItem returns the IsPerItem field if non-nil, zero value otherwise.

### GetIsPerItemOk

`func (o *CustomEffect) GetIsPerItemOk() (bool, bool)`

GetIsPerItemOk returns a tuple with the IsPerItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsPerItem

`func (o *CustomEffect) HasIsPerItem() bool`

HasIsPerItem returns a boolean if a field has been set.

### SetIsPerItem

`func (o *CustomEffect) SetIsPerItem(v bool)`

SetIsPerItem gets a reference to the given bool and assigns it to the IsPerItem field.

### GetName

`func (o *CustomEffect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEffect) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CustomEffect) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CustomEffect) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *CustomEffect) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomEffect) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *CustomEffect) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *CustomEffect) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetPayload

`func (o *CustomEffect) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEffect) GetPayloadOk() (string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *CustomEffect) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *CustomEffect) SetPayload(v string)`

SetPayload gets a reference to the given string and assigns it to the Payload field.

### GetDescription

`func (o *CustomEffect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomEffect) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *CustomEffect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *CustomEffect) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetEnabled

`func (o *CustomEffect) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomEffect) GetEnabledOk() (bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnabled

`func (o *CustomEffect) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabled

`func (o *CustomEffect) SetEnabled(v bool)`

SetEnabled gets a reference to the given bool and assigns it to the Enabled field.

### GetParams

`func (o *CustomEffect) GetParams() []TemplateArgDef`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CustomEffect) GetParamsOk() ([]TemplateArgDef, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParams

`func (o *CustomEffect) HasParams() bool`

HasParams returns a boolean if a field has been set.

### SetParams

`func (o *CustomEffect) SetParams(v []TemplateArgDef)`

SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.

### GetModifiedBy

`func (o *CustomEffect) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CustomEffect) GetModifiedByOk() (int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *CustomEffect) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *CustomEffect) SetModifiedBy(v int32)`

SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.

### GetCreatedBy

`func (o *CustomEffect) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CustomEffect) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CustomEffect) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CustomEffect) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


