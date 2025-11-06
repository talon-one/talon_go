# PendingActivePointsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoyaltyProgramID** | Pointer to **int64** | The ID of the loyalty program. | 
**SubledgerID** | Pointer to **string** | The ID of the subledger, when applicable. If this field is empty, the main ledger is used. | [default to ""]
**CustomerProfileID** | Pointer to **string** | The integration ID of the customer profile whose loyalty points are becoming active. | 
**Points** | Pointer to **float32** | The amount of pending loyalty points becoming active. | 
**ActiveOn** | Pointer to [**time.Time**](time.Time.md) | The date and time the loyalty points become active. | [optional] 
**ExpireOn** | Pointer to [**time.Time**](time.Time.md) | The date and time the loyalty points expire. | [optional] 
**SessionIntegrationID** | Pointer to **string** | The integration ID of the session through which the points were earned. | [optional] 

## Methods

### NewPendingActivePointsData

`func NewPendingActivePointsData(loyaltyProgramID int64, subledgerID string, customerProfileID string, points float32, ) *PendingActivePointsData`

NewPendingActivePointsData instantiates a new PendingActivePointsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingActivePointsDataWithDefaults

`func NewPendingActivePointsDataWithDefaults() *PendingActivePointsData`

NewPendingActivePointsDataWithDefaults instantiates a new PendingActivePointsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoyaltyProgramID

`func (o *PendingActivePointsData) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PendingActivePointsData) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PendingActivePointsData) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PendingActivePointsData) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PendingActivePointsData) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PendingActivePointsData) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetCustomerProfileID

`func (o *PendingActivePointsData) GetCustomerProfileID() string`

GetCustomerProfileID returns the CustomerProfileID field if non-nil, zero value otherwise.

### GetCustomerProfileIDOk

`func (o *PendingActivePointsData) GetCustomerProfileIDOk() (*string, bool)`

GetCustomerProfileIDOk returns a tuple with the CustomerProfileID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfileID

`func (o *PendingActivePointsData) SetCustomerProfileID(v string)`

SetCustomerProfileID sets CustomerProfileID field to given value.


### GetPoints

`func (o *PendingActivePointsData) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *PendingActivePointsData) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *PendingActivePointsData) SetPoints(v float32)`

SetPoints sets Points field to given value.


### GetActiveOn

`func (o *PendingActivePointsData) GetActiveOn() time.Time`

GetActiveOn returns the ActiveOn field if non-nil, zero value otherwise.

### GetActiveOnOk

`func (o *PendingActivePointsData) GetActiveOnOk() (*time.Time, bool)`

GetActiveOnOk returns a tuple with the ActiveOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOn

`func (o *PendingActivePointsData) SetActiveOn(v time.Time)`

SetActiveOn sets ActiveOn field to given value.

### HasActiveOn

`func (o *PendingActivePointsData) HasActiveOn() bool`

HasActiveOn returns a boolean if a field has been set.

### GetExpireOn

`func (o *PendingActivePointsData) GetExpireOn() time.Time`

GetExpireOn returns the ExpireOn field if non-nil, zero value otherwise.

### GetExpireOnOk

`func (o *PendingActivePointsData) GetExpireOnOk() (*time.Time, bool)`

GetExpireOnOk returns a tuple with the ExpireOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireOn

`func (o *PendingActivePointsData) SetExpireOn(v time.Time)`

SetExpireOn sets ExpireOn field to given value.

### HasExpireOn

`func (o *PendingActivePointsData) HasExpireOn() bool`

HasExpireOn returns a boolean if a field has been set.

### GetSessionIntegrationID

`func (o *PendingActivePointsData) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *PendingActivePointsData) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *PendingActivePointsData) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.

### HasSessionIntegrationID

`func (o *PendingActivePointsData) HasSessionIntegrationID() bool`

HasSessionIntegrationID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


