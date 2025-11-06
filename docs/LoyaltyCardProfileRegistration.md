# LoyaltyCardProfileRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | Integration ID of the customer profile linked to the card. | 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | Timestamp the customer profile was linked to the card. | 

## Methods

### NewLoyaltyCardProfileRegistration

`func NewLoyaltyCardProfileRegistration(integrationId string, timestamp time.Time, ) *LoyaltyCardProfileRegistration`

NewLoyaltyCardProfileRegistration instantiates a new LoyaltyCardProfileRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardProfileRegistrationWithDefaults

`func NewLoyaltyCardProfileRegistrationWithDefaults() *LoyaltyCardProfileRegistration`

NewLoyaltyCardProfileRegistrationWithDefaults instantiates a new LoyaltyCardProfileRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *LoyaltyCardProfileRegistration) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *LoyaltyCardProfileRegistration) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *LoyaltyCardProfileRegistration) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetTimestamp

`func (o *LoyaltyCardProfileRegistration) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoyaltyCardProfileRegistration) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LoyaltyCardProfileRegistration) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


