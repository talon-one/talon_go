# AddedDeductedPointsBalancesAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The amount of added or deducted loyalty points. | 
**Reason** | Pointer to **string** | The reason for the points addition or deduction. | 
**Operation** | Pointer to **string** | The action (addition or subtraction) made with loyalty points. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | The start date for loyalty points. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The expiration date for loyalty points. | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of the transaction in the loyalty ledger. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The timestamp when the loyalty transaction occurred. | [optional] 
**AwaitsActivation** | Pointer to **bool** | Indicates whether the points have an action-based start date. This property is returned only for point transactions with an action-based start date. | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, calculated relative to their start date.   **Note**: This only applies to points for which &#x60;AwaitsActivation&#x60; is &#x60;true&#x60; and &#x60;ExpiryDate&#x60; is not set.  | [optional] 

## Methods

### NewAddedDeductedPointsBalancesAction

`func NewAddedDeductedPointsBalancesAction(amount float32, reason string, operation string, transactionUUID string, ) *AddedDeductedPointsBalancesAction`

NewAddedDeductedPointsBalancesAction instantiates a new AddedDeductedPointsBalancesAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedDeductedPointsBalancesActionWithDefaults

`func NewAddedDeductedPointsBalancesActionWithDefaults() *AddedDeductedPointsBalancesAction`

NewAddedDeductedPointsBalancesActionWithDefaults instantiates a new AddedDeductedPointsBalancesAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddedDeductedPointsBalancesAction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddedDeductedPointsBalancesAction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddedDeductedPointsBalancesAction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *AddedDeductedPointsBalancesAction) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *AddedDeductedPointsBalancesAction) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *AddedDeductedPointsBalancesAction) SetReason(v string)`

SetReason sets Reason field to given value.


### GetOperation

`func (o *AddedDeductedPointsBalancesAction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AddedDeductedPointsBalancesAction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AddedDeductedPointsBalancesAction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStartDate

`func (o *AddedDeductedPointsBalancesAction) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddedDeductedPointsBalancesAction) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddedDeductedPointsBalancesAction) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddedDeductedPointsBalancesAction) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *AddedDeductedPointsBalancesAction) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AddedDeductedPointsBalancesAction) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *AddedDeductedPointsBalancesAction) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *AddedDeductedPointsBalancesAction) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *AddedDeductedPointsBalancesAction) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *AddedDeductedPointsBalancesAction) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *AddedDeductedPointsBalancesAction) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *AddedDeductedPointsBalancesAction) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AddedDeductedPointsBalancesAction) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AddedDeductedPointsBalancesAction) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AddedDeductedPointsBalancesAction) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetAwaitsActivation

`func (o *AddedDeductedPointsBalancesAction) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *AddedDeductedPointsBalancesAction) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *AddedDeductedPointsBalancesAction) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *AddedDeductedPointsBalancesAction) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetValidityDuration

`func (o *AddedDeductedPointsBalancesAction) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *AddedDeductedPointsBalancesAction) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *AddedDeductedPointsBalancesAction) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *AddedDeductedPointsBalancesAction) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


