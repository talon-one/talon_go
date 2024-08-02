# MessageLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | Identifier of the Application. | [optional] 
**ChangeType** | Pointer to **string** | Type of change that triggered the notification. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the log entry was created. | 
**EntityType** | Pointer to **string** | The entity type the notification is related to.  | [optional] 
**Id** | Pointer to **string** | Unique identifier of the message. | 
**LoyaltyProgramId** | Pointer to **int32** | Identifier of the loyalty program. | [optional] 
**NotificationId** | Pointer to **int32** | ID of the notification. | [optional] 
**NotificationName** | Pointer to **string** | The name of the notification. | [optional] 
**Request** | Pointer to [**MessageLogRequest**](MessageLogRequest.md) |  | [optional] 
**Response** | Pointer to [**MessageLogResponse**](MessageLogResponse.md) |  | [optional] 
**Service** | Pointer to **string** | Name of the service that generated the log entry. | 

## Methods

### GetApplicationId

`func (o *MessageLogEntry) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageLogEntry) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *MessageLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *MessageLogEntry) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetChangeType

`func (o *MessageLogEntry) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *MessageLogEntry) GetChangeTypeOk() (string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChangeType

`func (o *MessageLogEntry) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeType

`func (o *MessageLogEntry) SetChangeType(v string)`

SetChangeType gets a reference to the given string and assigns it to the ChangeType field.

### GetCreatedAt

`func (o *MessageLogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageLogEntry) GetCreatedAtOk() (time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *MessageLogEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *MessageLogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.

### GetEntityType

`func (o *MessageLogEntry) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MessageLogEntry) GetEntityTypeOk() (string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntityType

`func (o *MessageLogEntry) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### SetEntityType

`func (o *MessageLogEntry) SetEntityType(v string)`

SetEntityType gets a reference to the given string and assigns it to the EntityType field.

### GetId

`func (o *MessageLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageLogEntry) GetIdOk() (string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MessageLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MessageLogEntry) SetId(v string)`

SetId gets a reference to the given string and assigns it to the Id field.

### GetLoyaltyProgramId

`func (o *MessageLogEntry) GetLoyaltyProgramId() int32`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *MessageLogEntry) GetLoyaltyProgramIdOk() (int32, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyProgramId

`func (o *MessageLogEntry) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### SetLoyaltyProgramId

`func (o *MessageLogEntry) SetLoyaltyProgramId(v int32)`

SetLoyaltyProgramId gets a reference to the given int32 and assigns it to the LoyaltyProgramId field.

### GetNotificationId

`func (o *MessageLogEntry) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *MessageLogEntry) GetNotificationIdOk() (int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationId

`func (o *MessageLogEntry) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### SetNotificationId

`func (o *MessageLogEntry) SetNotificationId(v int32)`

SetNotificationId gets a reference to the given int32 and assigns it to the NotificationId field.

### GetNotificationName

`func (o *MessageLogEntry) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *MessageLogEntry) GetNotificationNameOk() (string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNotificationName

`func (o *MessageLogEntry) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### SetNotificationName

`func (o *MessageLogEntry) SetNotificationName(v string)`

SetNotificationName gets a reference to the given string and assigns it to the NotificationName field.

### GetRequest

`func (o *MessageLogEntry) GetRequest() MessageLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *MessageLogEntry) GetRequestOk() (MessageLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequest

`func (o *MessageLogEntry) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### SetRequest

`func (o *MessageLogEntry) SetRequest(v MessageLogRequest)`

SetRequest gets a reference to the given MessageLogRequest and assigns it to the Request field.

### GetResponse

`func (o *MessageLogEntry) GetResponse() MessageLogResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MessageLogEntry) GetResponseOk() (MessageLogResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponse

`func (o *MessageLogEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponse

`func (o *MessageLogEntry) SetResponse(v MessageLogResponse)`

SetResponse gets a reference to the given MessageLogResponse and assigns it to the Response field.

### GetService

`func (o *MessageLogEntry) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MessageLogEntry) GetServiceOk() (string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasService

`func (o *MessageLogEntry) HasService() bool`

HasService returns a boolean if a field has been set.

### SetService

`func (o *MessageLogEntry) SetService(v string)`

SetService gets a reference to the given string and assigns it to the Service field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


