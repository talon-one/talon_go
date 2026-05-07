# IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TransactionUUID** | Pointer to **string** | The identifier of the transaction in the loyalty ledger. | 

## Methods

### NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction(amount float32, operation string, transactionUUID string, ) *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationActionWithDefaults

`func NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationActionWithDefaults() *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

NewIntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationActionWithDefaults instantiates a new IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetOperation

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStartDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


