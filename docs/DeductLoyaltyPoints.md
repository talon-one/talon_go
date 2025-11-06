# DeductLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points. | 
**Name** | Pointer to **string** | Name / reason for the point deduction. | [optional] 
**SubledgerId** | Pointer to **string** | ID of the subledger the points are deducted from. | [optional] 
**ApplicationId** | Pointer to **int64** | ID of the Application that is connected to the loyalty program. | [optional] 

## Methods

### NewDeductLoyaltyPoints

`func NewDeductLoyaltyPoints(points float32, ) *DeductLoyaltyPoints`

NewDeductLoyaltyPoints instantiates a new DeductLoyaltyPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductLoyaltyPointsWithDefaults

`func NewDeductLoyaltyPointsWithDefaults() *DeductLoyaltyPoints`

NewDeductLoyaltyPointsWithDefaults instantiates a new DeductLoyaltyPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoints

`func (o *DeductLoyaltyPoints) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *DeductLoyaltyPoints) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *DeductLoyaltyPoints) SetPoints(v float32)`

SetPoints sets Points field to given value.


### GetName

`func (o *DeductLoyaltyPoints) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeductLoyaltyPoints) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeductLoyaltyPoints) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeductLoyaltyPoints) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubledgerId

`func (o *DeductLoyaltyPoints) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *DeductLoyaltyPoints) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *DeductLoyaltyPoints) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.

### HasSubledgerId

`func (o *DeductLoyaltyPoints) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### GetApplicationId

`func (o *DeductLoyaltyPoints) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DeductLoyaltyPoints) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DeductLoyaltyPoints) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *DeductLoyaltyPoints) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


