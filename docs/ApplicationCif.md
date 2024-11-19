# ApplicationCif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The name of the Application cart item filter used in API requests. | 
**Description** | Pointer to **string** | A short description of the Application cart item filter. | [optional] 
**ActiveExpressionId** | Pointer to **int32** | The ID of the expression that the Application cart item filter uses. | [optional] 
**ModifiedBy** | Pointer to **int32** | The ID of the user who last updated the Application cart item filter. | [optional] 
**CreatedBy** | Pointer to **int32** | The ID of the user who created the Application cart item filter. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the Application cart item filter. | [optional] 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 

## Methods

### GetId

`func (o *ApplicationCif) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCif) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *ApplicationCif) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *ApplicationCif) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *ApplicationCif) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApplicationCif) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *ApplicationCif) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *ApplicationCif) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetName

`func (o *ApplicationCif) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCif) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ApplicationCif) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ApplicationCif) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *ApplicationCif) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCif) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *ApplicationCif) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *ApplicationCif) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetActiveExpressionId

`func (o *ApplicationCif) GetActiveExpressionId() int32`

GetActiveExpressionId returns the ActiveExpressionId field if non-nil, zero value otherwise.

### GetActiveExpressionIdOk

`func (o *ApplicationCif) GetActiveExpressionIdOk() (int32, bool)`

GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveExpressionId

`func (o *ApplicationCif) HasActiveExpressionId() bool`

HasActiveExpressionId returns a boolean if a field has been set.

### SetActiveExpressionId

`func (o *ApplicationCif) SetActiveExpressionId(v int32)`

SetActiveExpressionId gets a reference to the given int32 and assigns it to the ActiveExpressionId field.

### GetModifiedBy

`func (o *ApplicationCif) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ApplicationCif) GetModifiedByOk() (int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *ApplicationCif) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *ApplicationCif) SetModifiedBy(v int32)`

SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.

### GetCreatedBy

`func (o *ApplicationCif) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApplicationCif) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *ApplicationCif) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *ApplicationCif) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetModified

`func (o *ApplicationCif) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ApplicationCif) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *ApplicationCif) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *ApplicationCif) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetApplicationId

`func (o *ApplicationCif) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationCif) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *ApplicationCif) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *ApplicationCif) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


