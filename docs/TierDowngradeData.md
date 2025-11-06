# TierDowngradeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerProfileID** | Pointer to **string** | The integration ID of the customer profile whose tier was downgraded. | 
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger, when applicable. If this field is empty, the main ledger is used. | [default to ""]
**CurrentTier** | Pointer to **string** | The name of the customer&#39;s current tier. | [optional] 
**CurrentPoints** | Pointer to **float32** | The number of points the customer had at the time of tier downgrade. | 
**OldTier** | Pointer to **string** | The name of the customer&#39;s previous tier. | 
**TierExpirationDate** | Pointer to [**time.Time**](time.Time.md) | The exact date and time the tier expires. | [optional] 
**TimestampOfTierChange** | Pointer to [**time.Time**](time.Time.md) | The exact date and time the tier was changed. | 

## Methods

### NewTierDowngradeData

`func NewTierDowngradeData(customerProfileID string, loyaltyProgramID int64, subledgerID string, currentPoints float32, oldTier string, timestampOfTierChange time.Time, ) *TierDowngradeData`

NewTierDowngradeData instantiates a new TierDowngradeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTierDowngradeDataWithDefaults

`func NewTierDowngradeDataWithDefaults() *TierDowngradeData`

NewTierDowngradeDataWithDefaults instantiates a new TierDowngradeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerProfileID

`func (o *TierDowngradeData) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *TierDowngradeData) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *TierDowngradeData) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.


### GetLoyaltyProgramID

`func (o *TierDowngradeData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *TierDowngradeData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *TierDowngradeData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *TierDowngradeData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *TierDowngradeData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *TierDowngradeData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetCurrentTier

`func (o *TierDowngradeData) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *TierDowngradeData) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *TierDowngradeData) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *TierDowngradeData) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *TierDowngradeData) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *TierDowngradeData) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *TierDowngradeData) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *TierDowngradeData) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *TierDowngradeData) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *TierDowngradeData) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.


### GetTierExpirationDate

`func (o *TierDowngradeData) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *TierDowngradeData) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *TierDowngradeData) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *TierDowngradeData) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *TierDowngradeData) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *TierDowngradeData) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *TierDowngradeData) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


