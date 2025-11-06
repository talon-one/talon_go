# LoyaltyLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ledger** | Pointer to [**LoyaltySubLedger**](LoyaltySubLedger.md) |  | 
**SubLedgers** | Pointer to [**map[string]LoyaltySubLedger**](LoyaltySubLedger.md) | A map containing a list of all loyalty subledger balances. | [optional] 

## Methods

### NewLoyaltyLedger

`func NewLoyaltyLedger(ledger LoyaltySubLedger, ) *LoyaltyLedger`

NewLoyaltyLedger instantiates a new LoyaltyLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyLedgerWithDefaults

`func NewLoyaltyLedgerWithDefaults() *LoyaltyLedger`

NewLoyaltyLedgerWithDefaults instantiates a new LoyaltyLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedger

`func (o *LoyaltyLedger) GetLedger() LoyaltySubLedger`

GetLedger returns the Ledger field if non-nil, zero value otherwise.

### GetLedgerOk

`func (o *LoyaltyLedger) GetLedgerOk() (*LoyaltySubLedger, bool)`

GetLedgerOk returns a tuple with the Ledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedger

`func (o *LoyaltyLedger) SetLedger(v LoyaltySubLedger)`

SetLedger sets Ledger field to given value.


### GetSubLedgers

`func (o *LoyaltyLedger) GetSubLedgers() map[string]LoyaltySubLedger`

GetSubLedgers returns the SubLedgers field if non-nil, zero value otherwise.

### GetSubLedgersOk

`func (o *LoyaltyLedger) GetSubLedgersOk() (*map[string]LoyaltySubLedger, bool)`

GetSubLedgersOk returns a tuple with the SubLedgers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgers

`func (o *LoyaltyLedger) SetSubLedgers(v map[string]LoyaltySubLedger)`

SetSubLedgers sets SubLedgers field to given value.

### HasSubLedgers

`func (o *LoyaltyLedger) HasSubLedgers() bool`

HasSubLedgers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


