# LoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Points** | Pointer to **float32** | Amount of loyalty points. | 
**Name** | Pointer to **string** | Allows to specify a name/reason for the point addition or deduction. | [optional] 
**ValidityDuration** | Pointer to **string** | Indicates the duration after which the added loyalty points should expire. The format is a number followed by one letter indicating the time unit, like &#39;1h&#39; or &#39;40m&#39; (defined by Go time package). | [optional] 
**PendingDuration** | Pointer to **string** | Indicates the amount of time before the points are considered valid. The format is a number followed by one letter indicating the time unit, like &#39;1h&#39; or &#39;40m&#39; (defined by Go time package). | [optional] 
**SubLedgerID** | Pointer to **string** | This specifies if we are adding loyalty points to the main ledger or a subledger. | [optional] 

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

### GetValidityDuration

`func (o *LoyaltyPoints) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LoyaltyPoints) GetValidityDurationOk() (string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidityDuration

`func (o *LoyaltyPoints) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### SetValidityDuration

`func (o *LoyaltyPoints) SetValidityDuration(v string)`

SetValidityDuration gets a reference to the given string and assigns it to the ValidityDuration field.

### GetPendingDuration

`func (o *LoyaltyPoints) GetPendingDuration() string`

GetPendingDuration returns the PendingDuration field if non-nil, zero value otherwise.

### GetPendingDurationOk

`func (o *LoyaltyPoints) GetPendingDurationOk() (string, bool)`

GetPendingDurationOk returns a tuple with the PendingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPendingDuration

`func (o *LoyaltyPoints) HasPendingDuration() bool`

HasPendingDuration returns a boolean if a field has been set.

### SetPendingDuration

`func (o *LoyaltyPoints) SetPendingDuration(v string)`

SetPendingDuration gets a reference to the given string and assigns it to the PendingDuration field.

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


