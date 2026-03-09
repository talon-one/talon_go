# ActivateLoyaltyPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionUUIDs** | Pointer to **[]string** | An array of transaction UUIDs used to activate specific pending point transactions.   If provided, do not include the &#x60;sessionId&#x60; parameter.  | [optional] 
**SessionId** | Pointer to **string** | The ID of the session containing the pending point transactions to activate.  If provided, do not include the &#x60;transactionUUIDs&#x60; parameter.  | [optional] 

## Methods

### NewActivateLoyaltyPoints

`func NewActivateLoyaltyPoints() *ActivateLoyaltyPoints`

NewActivateLoyaltyPoints instantiates a new ActivateLoyaltyPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateLoyaltyPointsWithDefaults

`func NewActivateLoyaltyPointsWithDefaults() *ActivateLoyaltyPoints`

NewActivateLoyaltyPointsWithDefaults instantiates a new ActivateLoyaltyPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionUUIDs

`func (o *ActivateLoyaltyPoints) GetTransactionUUIDs() []string`

GetTransactionUUIDs returns the TransactionUUIDs field if non-nil, zero value otherwise.

### GetTransactionUUIDsOk

`func (o *ActivateLoyaltyPoints) GetTransactionUUIDsOk() (*[]string, bool)`

GetTransactionUUIDsOk returns a tuple with the TransactionUUIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUIDs

`func (o *ActivateLoyaltyPoints) SetTransactionUUIDs(v []string)`

SetTransactionUUIDs sets TransactionUUIDs field to given value.

### HasTransactionUUIDs

`func (o *ActivateLoyaltyPoints) HasTransactionUUIDs() bool`

HasTransactionUUIDs returns a boolean if a field has been set.

### GetSessionId

`func (o *ActivateLoyaltyPoints) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ActivateLoyaltyPoints) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ActivateLoyaltyPoints) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ActivateLoyaltyPoints) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


