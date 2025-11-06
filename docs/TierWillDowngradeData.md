# TierWillDowngradeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerProfileID** | Pointer to **string** | The integration ID of the customer profile whose tier was downgraded. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger, when applicable. If this field is empty, the main ledger is used. | [default to ""]
**CurrentTier** | Pointer to **string** | The name of the customer&#39;s current tier. | 
**CurrentPoints** | Pointer to **float32** | The number of points the customer will have after the tier downgrade. | 
**PointsRequiredToRemain** | Pointer to **float32** | The number of points needed for a customer to remain on the same tier. | 
**NextTier** | Pointer to **string** | The name of the customer&#39;s next tier. | [optional] 
**TierExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The date and time the tier expires. | [optional] 

## Methods

### NewTierWillDowngradeData

`func NewTierWillDowngradeData(customerProfileID string, loyaltyProgramID int64, subledgerID string, currentTier string, currentPoints float32, pointsRequiredToRemain float32, ) *TierWillDowngradeData`

NewTierWillDowngradeData instantiates a new TierWillDowngradeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierWillDowngradeDataWithDefaults

`func NewTierWillDowngradeDataWithDefaults() *TierWillDowngradeData`

NewTierWillDowngradeDataWithDefaults instantiates a new TierWillDowngradeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerProfileID

`func (o *TierWillDowngradeData) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *TierWillDowngradeData) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *TierWillDowngradeData) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.


### GetLoyaltyProgramID

`func (o *TierWillDowngradeData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *TierWillDowngradeData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *TierWillDowngradeData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *TierWillDowngradeData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *TierWillDowngradeData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *TierWillDowngradeData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetCurrentTier

`func (o *TierWillDowngradeData) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *TierWillDowngradeData) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *TierWillDowngradeData) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.


### GetCurrentPoints

`func (o *TierWillDowngradeData) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *TierWillDowngradeData) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *TierWillDowngradeData) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetPointsRequiredToRemain

`func (o *TierWillDowngradeData) GetPointsRequiredToRemain() float32`

GetPointsRequiredToRemain returns the PointsRequiredToRemain field if non-nil, zero value otherwise.

### GetPointsRequiredToRemainOk

`func (o *TierWillDowngradeData) GetPointsRequiredToRemainOk() (*float32, bool)`

GetPointsRequiredToRemainOk returns a tuple with the PointsRequiredToRemain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToRemain

`func (o *TierWillDowngradeData) SetPointsRequiredToRemain(v float32)`

SetPointsRequiredToRemain sets PointsRequiredToRemain field to given value.


### GetNextTier

`func (o *TierWillDowngradeData) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *TierWillDowngradeData) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *TierWillDowngradeData) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *TierWillDowngradeData) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *TierWillDowngradeData) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *TierWillDowngradeData) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *TierWillDowngradeData) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *TierWillDowngradeData) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


