# LoyaltyLedgerTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | If &#x60;true&#x60;, it means that there is more data to request in the source collection. | [optional] 
**Data** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | List of transaction entries from a loyalty ledger. | [optional] 

## Methods

### NewLoyaltyLedgerTransactions

`func NewLoyaltyLedgerTransactions() *LoyaltyLedgerTransactions`

NewLoyaltyLedgerTransactions instantiates a new LoyaltyLedgerTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyLedgerTransactionsWithDefaults

`func NewLoyaltyLedgerTransactionsWithDefaults() *LoyaltyLedgerTransactions`

NewLoyaltyLedgerTransactionsWithDefaults instantiates a new LoyaltyLedgerTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *LoyaltyLedgerTransactions) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *LoyaltyLedgerTransactions) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *LoyaltyLedgerTransactions) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *LoyaltyLedgerTransactions) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *LoyaltyLedgerTransactions) GetData() []LoyaltyLedgerEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoyaltyLedgerTransactions) GetDataOk() (*[]LoyaltyLedgerEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoyaltyLedgerTransactions) SetData(v []LoyaltyLedgerEntry)`

SetData sets Data field to given value.

### HasData

`func (o *LoyaltyLedgerTransactions) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


