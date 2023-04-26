# LoyaltyCardProfileRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | Integration ID of the customer profile linked to the card. | 
**Timestamp** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the registration to the card. | 

## Methods

### GetIntegrationId

`func (o *LoyaltyCardProfileRegistration) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *LoyaltyCardProfileRegistration) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *LoyaltyCardProfileRegistration) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *LoyaltyCardProfileRegistration) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetTimestamp

`func (o *LoyaltyCardProfileRegistration) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LoyaltyCardProfileRegistration) GetTimestampOk() (time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimestamp

`func (o *LoyaltyCardProfileRegistration) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestamp

`func (o *LoyaltyCardProfileRegistration) SetTimestamp(v time.Time)`

SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


