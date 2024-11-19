# GenerateLoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the loyalty card. | [optional] [default to STATUS_ACTIVE]
**CustomerProfileIds** | Pointer to **[]string** | Integration IDs of the customer profiles linked to the card. | [optional] 

## Methods

### GetStatus

`func (o *GenerateLoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateLoyaltyCard) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *GenerateLoyaltyCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *GenerateLoyaltyCard) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetCustomerProfileIds

`func (o *GenerateLoyaltyCard) GetCustomerProfileIds() []string`

GetCustomerProfileIds returns the CustomerProfileIds field if non-nil, zero value otherwise.

### GetCustomerProfileIdsOk

`func (o *GenerateLoyaltyCard) GetCustomerProfileIdsOk() ([]string, bool)`

GetCustomerProfileIdsOk returns a tuple with the CustomerProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerProfileIds

`func (o *GenerateLoyaltyCard) HasCustomerProfileIds() bool`

HasCustomerProfileIds returns a boolean if a field has been set.

### SetCustomerProfileIds

`func (o *GenerateLoyaltyCard) SetCustomerProfileIds(v []string)`

SetCustomerProfileIds gets a reference to the given []string and assigns it to the CustomerProfileIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


