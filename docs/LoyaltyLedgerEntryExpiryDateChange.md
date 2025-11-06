# LoyaltyLedgerEntryExpiryDateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionUUID** | Pointer to **string** | The identifier of the transaction affected by the extension or update. | 
**PreviousExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the transactions before applying the extension or update. | [optional] 
**NewExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the transaction after applying the extension or update. | 

## Methods

### NewLoyaltyLedgerEntryExpiryDateChange

`func NewLoyaltyLedgerEntryExpiryDateChange(transactionUUID string, newExpiryDate time.Time, ) *LoyaltyLedgerEntryExpiryDateChange`

NewLoyaltyLedgerEntryExpiryDateChange instantiates a new LoyaltyLedgerEntryExpiryDateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyLedgerEntryExpiryDateChangeWithDefaults

`func NewLoyaltyLedgerEntryExpiryDateChangeWithDefaults() *LoyaltyLedgerEntryExpiryDateChange`

NewLoyaltyLedgerEntryExpiryDateChangeWithDefaults instantiates a new LoyaltyLedgerEntryExpiryDateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionUUID

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LoyaltyLedgerEntryExpiryDateChange) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetPreviousExpiryDate

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetPreviousExpiryDate() time.Time`

GetPreviousExpiryDate returns the PreviousExpiryDate field if non-nil, zero value otherwise.

### GetPreviousExpiryDateOk

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetPreviousExpiryDateOk() (*time.Time, bool)`

GetPreviousExpiryDateOk returns a tuple with the PreviousExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousExpiryDate

`func (o *LoyaltyLedgerEntryExpiryDateChange) SetPreviousExpiryDate(v time.Time)`

SetPreviousExpiryDate sets PreviousExpiryDate field to given value.

### HasPreviousExpiryDate

`func (o *LoyaltyLedgerEntryExpiryDateChange) HasPreviousExpiryDate() bool`

HasPreviousExpiryDate returns a boolean if a field has been set.

### GetNewExpiryDate

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetNewExpiryDate() time.Time`

GetNewExpiryDate returns the NewExpiryDate field if non-nil, zero value otherwise.

### GetNewExpiryDateOk

`func (o *LoyaltyLedgerEntryExpiryDateChange) GetNewExpiryDateOk() (*time.Time, bool)`

GetNewExpiryDateOk returns a tuple with the NewExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExpiryDate

`func (o *LoyaltyLedgerEntryExpiryDateChange) SetNewExpiryDate(v time.Time)`

SetNewExpiryDate sets NewExpiryDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


