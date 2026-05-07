# DeleteLoyaltyTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** | &#x60;AllSubledgers&#x60; deletes all transactions for the specified customer profile from all ledgers in the loyalty program.  &#x60;SelectedSubledgers&#x60; deletes all transactions for the specified customer profile only from the given ledgers in the loyalty program.  | 
**SubledgerIds** | Pointer to **[]string** | The IDs of the ledgers from which to delete the customer&#39;s transactions. This parameter is required if the &#x60;scope&#x60; is set to &#x60;SelectedSubledgers&#x60;.  To specify the main ledger, provide an empty string (\&quot;\&quot;).  | [optional] 

## Methods

### NewDeleteLoyaltyTransactionsRequest

`func NewDeleteLoyaltyTransactionsRequest(scope string, ) *DeleteLoyaltyTransactionsRequest`

NewDeleteLoyaltyTransactionsRequest instantiates a new DeleteLoyaltyTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteLoyaltyTransactionsRequestWithDefaults

`func NewDeleteLoyaltyTransactionsRequestWithDefaults() *DeleteLoyaltyTransactionsRequest`

NewDeleteLoyaltyTransactionsRequestWithDefaults instantiates a new DeleteLoyaltyTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *DeleteLoyaltyTransactionsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DeleteLoyaltyTransactionsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DeleteLoyaltyTransactionsRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSubledgerIds

`func (o *DeleteLoyaltyTransactionsRequest) GetSubledgerIds() []string`

GetSubledgerIds returns the SubledgerIds field if non-nil, zero value otherwise.

### GetSubledgerIdsOk

`func (o *DeleteLoyaltyTransactionsRequest) GetSubledgerIdsOk() (*[]string, bool)`

GetSubledgerIdsOk returns a tuple with the SubledgerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerIds

`func (o *DeleteLoyaltyTransactionsRequest) SetSubledgerIds(v []string)`

SetSubledgerIds sets SubledgerIds field to given value.

### HasSubledgerIds

`func (o *DeleteLoyaltyTransactionsRequest) HasSubledgerIds() bool`

HasSubledgerIds returns a boolean if a field has been set.

### SetSubledgerIdsNil

`func (o *DeleteLoyaltyTransactionsRequest) SetSubledgerIdsNil(b bool)`

 SetSubledgerIdsNil sets the value for SubledgerIds to be an explicit nil

### UnsetSubledgerIds
`func (o *DeleteLoyaltyTransactionsRequest) UnsetSubledgerIds()`

UnsetSubledgerIds ensures that no value is present for SubledgerIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


