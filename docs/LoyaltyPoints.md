# LoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points | 
**Name** | Pointer to **string** | Allows to specify a name for the addition or deduction | [optional] 
**ExpiryDuration** | Pointer to **string** | Indicates the duration after which the added loyalty points should expire. The format is a number followed by one letter indicating the unit, like &#39;1h&#39; or &#39;40m&#39; or &#39;30d&#39;. | [optional] 
**SubLedgerID** | Pointer to **string** | This specifies if we are adding loyalty points to the main ledger or a subledger | [optional] 

## Methods

### GetPoints

`func (o *LoyaltyPoints) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *LoyaltyPoints) GetPointsOk() (float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPoints

`func (o *LoyaltyPoints) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPoints

`func (o *LoyaltyPoints) SetPoints(v float32)`

SetPoints gets a reference to the given float32 and assigns it to the Points field.

### GetName

`func (o *LoyaltyPoints) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyPoints) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyPoints) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyPoints) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetExpiryDuration

`func (o *LoyaltyPoints) GetExpiryDuration() string`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *LoyaltyPoints) GetExpiryDurationOk() (string, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDuration

`func (o *LoyaltyPoints) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDuration

`func (o *LoyaltyPoints) SetExpiryDuration(v string)`

SetExpiryDuration gets a reference to the given string and assigns it to the ExpiryDuration field.

### GetSubLedgerID

`func (o *LoyaltyPoints) GetSubLedgerID() string`

GetSubLedgerID returns the SubLedgerID field if non-nil, zero value otherwise.

### GetSubLedgerIDOk

`func (o *LoyaltyPoints) GetSubLedgerIDOk() (string, bool)`

GetSubLedgerIDOk returns a tuple with the SubLedgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubLedgerID

`func (o *LoyaltyPoints) HasSubLedgerID() bool`

HasSubLedgerID returns a boolean if a field has been set.

### SetSubLedgerID

`func (o *LoyaltyPoints) SetSubLedgerID(v string)`

SetSubLedgerID gets a reference to the given string and assigns it to the SubLedgerID field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


