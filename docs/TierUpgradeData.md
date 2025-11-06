# TierUpgradeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerProfileID** | Pointer to **string** | The integration ID of the customer profile whose tier was upgraded. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger, when applicable. If this field is empty, the main ledger is used. | [default to ""]
**CurrentTier** | Pointer to **string** | The name of the customer&#39;s current tier. | 
**CurrentPoints** | Pointer to **float32** | The number of points the customer had at the time of tier upgrade. | 
**OldTier** | Pointer to **string** | The name of the customer&#39;s previous tier. | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** | The number of points needed for a customer to reach the next tier. | [optional] 
**NextTier** | Pointer to **string** | The name of the customer&#39;s next tier. | [optional] 
**TierExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The exact date and time the tier expires. | 
**TimestampOfTierChange** | Pointer to [**time.Time**](time.Time.md) | The exact date and time the tier was changed. | 

## Methods

### NewTierUpgradeData

`func NewTierUpgradeData(customerProfileID string, loyaltyProgramID int64, subledgerID string, currentTier string, currentPoints float32, tierExpirationDate time.Time, timestampOfTierChange time.Time, ) *TierUpgradeData`

NewTierUpgradeData instantiates a new TierUpgradeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierUpgradeDataWithDefaults

`func NewTierUpgradeDataWithDefaults() *TierUpgradeData`

NewTierUpgradeDataWithDefaults instantiates a new TierUpgradeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerProfileID

`func (o *TierUpgradeData) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *TierUpgradeData) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *TierUpgradeData) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.


### GetLoyaltyProgramID

`func (o *TierUpgradeData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *TierUpgradeData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *TierUpgradeData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *TierUpgradeData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *TierUpgradeData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *TierUpgradeData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetCurrentTier

`func (o *TierUpgradeData) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *TierUpgradeData) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *TierUpgradeData) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.


### GetCurrentPoints

`func (o *TierUpgradeData) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *TierUpgradeData) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *TierUpgradeData) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *TierUpgradeData) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *TierUpgradeData) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *TierUpgradeData) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *TierUpgradeData) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *TierUpgradeData) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *TierUpgradeData) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *TierUpgradeData) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *TierUpgradeData) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *TierUpgradeData) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *TierUpgradeData) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *TierUpgradeData) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *TierUpgradeData) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *TierUpgradeData) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *TierUpgradeData) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *TierUpgradeData) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.


### GetTimestampOfTierChange

`func (o *TierUpgradeData) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *TierUpgradeData) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *TierUpgradeData) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


