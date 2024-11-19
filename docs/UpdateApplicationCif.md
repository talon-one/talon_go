# UpdateApplicationCif

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the Application cart item filter. | [optional] 
**ActiveExpressionId** | Pointer to **int32** | The ID of the expression that the Application cart item filter uses. | [optional] 
**ModifiedBy** | Pointer to **int32** | The ID of the user who last updated the Application cart item filter. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent update to the Application cart item filter. | [optional] 

## Methods

### GetDescription

`func (o *UpdateApplicationCif) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplicationCif) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateApplicationCif) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateApplicationCif) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetActiveExpressionId

`func (o *UpdateApplicationCif) GetActiveExpressionId() int32`

GetActiveExpressionId returns the ActiveExpressionId field if non-nil, zero value otherwise.

### GetActiveExpressionIdOk

`func (o *UpdateApplicationCif) GetActiveExpressionIdOk() (int32, bool)`

GetActiveExpressionIdOk returns a tuple with the ActiveExpressionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActiveExpressionId

`func (o *UpdateApplicationCif) HasActiveExpressionId() bool`

HasActiveExpressionId returns a boolean if a field has been set.

### SetActiveExpressionId

`func (o *UpdateApplicationCif) SetActiveExpressionId(v int32)`

SetActiveExpressionId gets a reference to the given int32 and assigns it to the ActiveExpressionId field.

### GetModifiedBy

`func (o *UpdateApplicationCif) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *UpdateApplicationCif) GetModifiedByOk() (int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *UpdateApplicationCif) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *UpdateApplicationCif) SetModifiedBy(v int32)`

SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.

### GetModified

`func (o *UpdateApplicationCif) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *UpdateApplicationCif) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *UpdateApplicationCif) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *UpdateApplicationCif) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


