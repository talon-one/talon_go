# CustomerSessionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstSession** | Pointer to **bool** | Indicates whether this is the first session for the customer&#39;s profile. Will always be true for anonymous sessions. | 
**Discounts** | Pointer to **map[string]float32** | A map of labelled discount values, values will be in the same currency as the application associated with the session. | 
**IntegrationId** | Pointer to **string** |  | 
**Updated** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the most recent event received on this session. | 

## Methods

### GetFirstSession

`func (o *CustomerSessionAllOf) GetFirstSession() bool`

GetFirstSession returns the FirstSession field if non-nil, zero value otherwise.

### GetFirstSessionOk

`func (o *CustomerSessionAllOf) GetFirstSessionOk() (bool, bool)`

GetFirstSessionOk returns a tuple with the FirstSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFirstSession

`func (o *CustomerSessionAllOf) HasFirstSession() bool`

HasFirstSession returns a boolean if a field has been set.

### SetFirstSession

`func (o *CustomerSessionAllOf) SetFirstSession(v bool)`

SetFirstSession gets a reference to the given bool and assigns it to the FirstSession field.

### GetDiscounts

`func (o *CustomerSessionAllOf) GetDiscounts() map[string]float32`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *CustomerSessionAllOf) GetDiscountsOk() (map[string]float32, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscounts

`func (o *CustomerSessionAllOf) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### SetDiscounts

`func (o *CustomerSessionAllOf) SetDiscounts(v map[string]float32)`

SetDiscounts gets a reference to the given map[string]float32 and assigns it to the Discounts field.

### GetIntegrationId

`func (o *CustomerSessionAllOf) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerSessionAllOf) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *CustomerSessionAllOf) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *CustomerSessionAllOf) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetUpdated

`func (o *CustomerSessionAllOf) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CustomerSessionAllOf) GetUpdatedOk() (time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUpdated

`func (o *CustomerSessionAllOf) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdated

`func (o *CustomerSessionAllOf) SetUpdated(v time.Time)`

SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


