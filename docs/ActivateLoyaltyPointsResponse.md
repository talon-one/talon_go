# ActivateLoyaltyPointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedgerEntries** | Pointer to [**[]LoyaltyLedgerEntry**](LoyaltyLedgerEntry.md) | Updated ledger entries after activation. | [optional] 

## Methods

### NewActivateLoyaltyPointsResponse

`func NewActivateLoyaltyPointsResponse() *ActivateLoyaltyPointsResponse`

NewActivateLoyaltyPointsResponse instantiates a new ActivateLoyaltyPointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateLoyaltyPointsResponseWithDefaults

`func NewActivateLoyaltyPointsResponseWithDefaults() *ActivateLoyaltyPointsResponse`

NewActivateLoyaltyPointsResponseWithDefaults instantiates a new ActivateLoyaltyPointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedgerEntries

`func (o *ActivateLoyaltyPointsResponse) GetLedgerEntries() []LoyaltyLedgerEntry`

GetLedgerEntries returns the LedgerEntries field if non-nil, zero value otherwise.

### GetLedgerEntriesOk

`func (o *ActivateLoyaltyPointsResponse) GetLedgerEntriesOk() (*[]LoyaltyLedgerEntry, bool)`

GetLedgerEntriesOk returns a tuple with the LedgerEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerEntries

`func (o *ActivateLoyaltyPointsResponse) SetLedgerEntries(v []LoyaltyLedgerEntry)`

SetLedgerEntries sets LedgerEntries field to given value.

### HasLedgerEntries

`func (o *ActivateLoyaltyPointsResponse) HasLedgerEntries() bool`

HasLedgerEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


