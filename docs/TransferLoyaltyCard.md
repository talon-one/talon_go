# TransferLoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewCardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | 
**BlockReason** | Pointer to **string** | Reason for transferring and blocking the loyalty card.  | [optional] 

## Methods

### NewTransferLoyaltyCard

`func NewTransferLoyaltyCard(newCardIdentifier string, ) *TransferLoyaltyCard`

NewTransferLoyaltyCard instantiates a new TransferLoyaltyCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferLoyaltyCardWithDefaults

`func NewTransferLoyaltyCardWithDefaults() *TransferLoyaltyCard`

NewTransferLoyaltyCardWithDefaults instantiates a new TransferLoyaltyCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewCardIdentifier

`func (o *TransferLoyaltyCard) GetNewCardIdentifier() string`

GetNewCardIdentifier returns the NewCardIdentifier field if non-nil, zero value otherwise.

### GetNewCardIdentifierOk

`func (o *TransferLoyaltyCard) GetNewCardIdentifierOk() (*string, bool)`

GetNewCardIdentifierOk returns a tuple with the NewCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCardIdentifier

`func (o *TransferLoyaltyCard) SetNewCardIdentifier(v string)`

SetNewCardIdentifier sets NewCardIdentifier field to given value.


### GetBlockReason

`func (o *TransferLoyaltyCard) GetBlockReason() string`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *TransferLoyaltyCard) GetBlockReasonOk() (*string, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReason

`func (o *TransferLoyaltyCard) SetBlockReason(v string)`

SetBlockReason sets BlockReason field to given value.

### HasBlockReason

`func (o *TransferLoyaltyCard) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


