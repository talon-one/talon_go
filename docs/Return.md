# Return

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**ReturnedCartItems** | Pointer to [**[]ReturnedCartItem**](ReturnedCartItem.md) | List of cart items to be returned. | 
**EventId** | Pointer to **int64** | The event ID of that was generated for this return. | 
**SessionId** | Pointer to **int64** | The internal ID of the session this return was requested on. | 
**SessionIntegrationId** | Pointer to **string** | The integration ID of the session this return was requested on. | 
**ProfileId** | Pointer to **int64** | The internal ID of the profile this return was requested on. | [optional] 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the profile this return was requested on. | [optional] 
**CreatedBy** | Pointer to **int64** | ID of the user who requested this return. | [optional] 

## Methods

### NewReturn

`func NewReturn(id int64, created time.Time, applicationId int64, accountId int64, returnedCartItems []ReturnedCartItem, eventId int64, sessionId int64, sessionIntegrationId string, ) *Return`

NewReturn instantiates a new Return object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnWithDefaults

`func NewReturnWithDefaults() *Return`

NewReturnWithDefaults instantiates a new Return object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Return) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Return) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Return) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Return) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Return) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Return) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *Return) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *Return) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *Return) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *Return) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Return) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Return) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetReturnedCartItems

`func (o *Return) GetReturnedCartItems() []ReturnedCartItem`

GetReturnedCartItems returns the ReturnedCartItems field if non-nil, zero value otherwise.

### GetReturnedCartItemsOk

`func (o *Return) GetReturnedCartItemsOk() (*[]ReturnedCartItem, bool)`

GetReturnedCartItemsOk returns a tuple with the ReturnedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedCartItems

`func (o *Return) SetReturnedCartItems(v []ReturnedCartItem)`

SetReturnedCartItems sets ReturnedCartItems field to given value.


### GetEventId

`func (o *Return) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Return) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Return) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetSessionId

`func (o *Return) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Return) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Return) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.


### GetSessionIntegrationId

`func (o *Return) GetSessionIntegrationId() string`

GetSessionIntegrationId returns the SessionIntegrationId field if non-nil, zero value otherwise.

### GetSessionIntegrationIdOk

`func (o *Return) GetSessionIntegrationIdOk() (*string, bool)`

GetSessionIntegrationIdOk returns a tuple with the SessionIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationId

`func (o *Return) SetSessionIntegrationId(v string)`

SetSessionIntegrationId sets SessionIntegrationId field to given value.


### GetProfileId

`func (o *Return) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Return) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Return) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *Return) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileIntegrationId

`func (o *Return) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *Return) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *Return) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.

### HasProfileIntegrationId

`func (o *Return) HasProfileIntegrationId() bool`

HasProfileIntegrationId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Return) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Return) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Return) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Return) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


