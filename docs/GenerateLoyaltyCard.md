# GenerateLoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the loyalty card. | [optional] [default to "active"]
**CustomerProfileIds** | Pointer to **[]string** | Integration IDs of the customer profiles linked to the card. | [optional] 
**CardIdentifier** | Pointer to **string** | The alphanumeric identifier of the loyalty card.  | [optional] 

## Methods

### NewGenerateLoyaltyCard

`func NewGenerateLoyaltyCard() *GenerateLoyaltyCard`

NewGenerateLoyaltyCard instantiates a new GenerateLoyaltyCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateLoyaltyCardWithDefaults

`func NewGenerateLoyaltyCardWithDefaults() *GenerateLoyaltyCard`

NewGenerateLoyaltyCardWithDefaults instantiates a new GenerateLoyaltyCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GenerateLoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateLoyaltyCard) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerateLoyaltyCard) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerateLoyaltyCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomerProfileIds

`func (o *GenerateLoyaltyCard) GetCustomerProfileIds() []string`

GetCustomerProfileIds returns the CustomerProfileIds field if non-nil, zero value otherwise.

### GetCustomerProfileIdsOk

`func (o *GenerateLoyaltyCard) GetCustomerProfileIdsOk() (*[]string, bool)`

GetCustomerProfileIdsOk returns a tuple with the CustomerProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileIds

`func (o *GenerateLoyaltyCard) SetCustomerProfileIds(v []string)`

SetCustomerProfileIds sets CustomerProfileIds field to given value.

### HasCustomerProfileIds

`func (o *GenerateLoyaltyCard) HasCustomerProfileIds() bool`

HasCustomerProfileIds returns a boolean if a field has been set.

### GetCardIdentifier

`func (o *GenerateLoyaltyCard) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *GenerateLoyaltyCard) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *GenerateLoyaltyCard) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *GenerateLoyaltyCard) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


