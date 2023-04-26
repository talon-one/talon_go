# DeductLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points. | 
**Name** | Pointer to **string** | Name / reason for the point deduction. | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger the points are deducted from. | [optional] 
**ApplicationId** | Pointer to **int32** | ID of the Application that is connected to the loyalty program. | [optional] 

## Methods

### GetPoints

`func (o *DeductLoyaltyPoints) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DeductLoyaltyPoints) GetPointsOk() (float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoints

`func (o *DeductLoyaltyPoints) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPoints

`func (o *DeductLoyaltyPoints) SetPoints(v float32)`

SetPoints gets a reference to the given float32 and assigns it to the Points field.

### GetName

`func (o *DeductLoyaltyPoints) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeductLoyaltyPoints) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *DeductLoyaltyPoints) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *DeductLoyaltyPoints) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSubledgerId

`func (o *DeductLoyaltyPoints) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *DeductLoyaltyPoints) GetSubledgerIdOk() (string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubledgerId

`func (o *DeductLoyaltyPoints) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### SetSubledgerId

`func (o *DeductLoyaltyPoints) SetSubledgerId(v string)`

SetSubledgerId gets a reference to the given string and assigns it to the SubledgerId field.

### GetApplicationId

`func (o *DeductLoyaltyPoints) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DeductLoyaltyPoints) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *DeductLoyaltyPoints) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *DeductLoyaltyPoints) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


