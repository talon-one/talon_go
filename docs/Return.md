# Return

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the Application that owns this entity. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**ReturnedCartItems** | Pointer to [**[]ReturnedCartItem**](ReturnedCartItem.md) | List of cart items to be returned. | 
**EventId** | Pointer to **int32** | The event ID of that was generated for this return. | 
**SessionId** | Pointer to **int32** | The internal ID of the session this return was requested on. | 
**SessionIntegrationId** | Pointer to **string** | The integration ID of the session this return was requested on. | 
**ProfileId** | Pointer to **int32** | The internal ID of the profile this return was requested on. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the profile this return was requested on. | [optional] 
**CreatedBy** | Pointer to **int32** | ID of the user who requested this return. | [optional] 

## Methods

### GetId

`func (o *Return) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Return) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Return) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Return) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Return) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Return) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Return) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Return) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *Return) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Return) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *Return) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *Return) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetAccountId

`func (o *Return) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Return) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Return) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Return) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetReturnedCartItems

`func (o *Return) GetReturnedCartItems() []ReturnedCartItem`

GetReturnedCartItems returns the ReturnedCartItems field if non-nil, zero value otherwise.

### GetReturnedCartItemsOk

`func (o *Return) GetReturnedCartItemsOk() ([]ReturnedCartItem, bool)`

GetReturnedCartItemsOk returns a tuple with the ReturnedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReturnedCartItems

`func (o *Return) HasReturnedCartItems() bool`

HasReturnedCartItems returns a boolean if a field has been set.

### SetReturnedCartItems

`func (o *Return) SetReturnedCartItems(v []ReturnedCartItem)`

SetReturnedCartItems gets a reference to the given []ReturnedCartItem and assigns it to the ReturnedCartItems field.

### GetEventId

`func (o *Return) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Return) GetEventIdOk() (int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventId

`func (o *Return) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### SetEventId

`func (o *Return) SetEventId(v int32)`

SetEventId gets a reference to the given int32 and assigns it to the EventId field.

### GetSessionId

`func (o *Return) GetSessionId() int32`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Return) GetSessionIdOk() (int32, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionId

`func (o *Return) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionId

`func (o *Return) SetSessionId(v int32)`

SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.

### GetSessionIntegrationId

`func (o *Return) GetSessionIntegrationId() string`

GetSessionIntegrationId returns the SessionIntegrationId field if non-nil, zero value otherwise.

### GetSessionIntegrationIdOk

`func (o *Return) GetSessionIntegrationIdOk() (string, bool)`

GetSessionIntegrationIdOk returns a tuple with the SessionIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSessionIntegrationId

`func (o *Return) HasSessionIntegrationId() bool`

HasSessionIntegrationId returns a boolean if a field has been set.

### SetSessionIntegrationId

`func (o *Return) SetSessionIntegrationId(v string)`

SetSessionIntegrationId gets a reference to the given string and assigns it to the SessionIntegrationId field.

### GetProfileId

`func (o *Return) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Return) GetProfileIdOk() (int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *Return) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *Return) SetProfileId(v int32)`

SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.

### GetProfileIntegrationId

`func (o *Return) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *Return) GetProfileIntegrationIdOk() (string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileIntegrationId

`func (o *Return) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### SetProfileIntegrationId

`func (o *Return) SetProfileIntegrationId(v string)`

SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.

### GetCreatedBy

`func (o *Return) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Return) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Return) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Return) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


