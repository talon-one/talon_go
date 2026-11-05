# OutboundLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID of the outbound message. | 
**NotificationId** | Pointer to **int64** | ID of the notification that produced the outbound request. | [optional] 
**NotificationName** | Pointer to **string** | Name of the notification that produced the outbound request. | [optional] 
**WebhookId** | Pointer to **int64** | ID of the webhook that produced the outbound request. | [optional] 
**WebhookName** | Pointer to **string** | The name of the webhook that produced the outbound request. | [optional] 
**NotificationType** | Pointer to **string** | Type of notification that produced the outbound request. | 
**ApplicationId** | Pointer to **int64** | ID of the Application associated with the outbound request. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | ID of the loyalty program associated with the outbound request. | [optional] 
**Request** | Pointer to [**OutboundLogRequest**](OutboundLogRequest.md) |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the log entry was created. | 
**ProcessingTimeMs** | Pointer to **int64** | Processing time of the outbound request in milliseconds. | 
**Response** | Pointer to [**OutboundLogResponse**](OutboundLogResponse.md) |  | [optional] 

## Methods

### NewOutboundLog

`func NewOutboundLog(uuid string, notificationType string, createdAt time.Time, processingTimeMs int64, ) *OutboundLog`

NewOutboundLog instantiates a new OutboundLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogWithDefaults

`func NewOutboundLogWithDefaults() *OutboundLog`

NewOutboundLogWithDefaults instantiates a new OutboundLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *OutboundLog) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OutboundLog) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OutboundLog) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNotificationId

`func (o *OutboundLog) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *OutboundLog) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *OutboundLog) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *OutboundLog) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetNotificationName

`func (o *OutboundLog) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *OutboundLog) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *OutboundLog) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *OutboundLog) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetWebhookId

`func (o *OutboundLog) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *OutboundLog) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *OutboundLog) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *OutboundLog) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *OutboundLog) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *OutboundLog) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *OutboundLog) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *OutboundLog) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetNotificationType

`func (o *OutboundLog) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *OutboundLog) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *OutboundLog) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetApplicationId

`func (o *OutboundLog) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OutboundLog) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OutboundLog) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OutboundLog) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *OutboundLog) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *OutboundLog) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *OutboundLog) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *OutboundLog) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetRequest

`func (o *OutboundLog) GetRequest() OutboundLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OutboundLog) GetRequestOk() (*OutboundLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OutboundLog) SetRequest(v OutboundLogRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OutboundLog) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OutboundLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutboundLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutboundLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProcessingTimeMs

`func (o *OutboundLog) GetProcessingTimeMs() int64`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *OutboundLog) GetProcessingTimeMsOk() (*int64, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *OutboundLog) SetProcessingTimeMs(v int64)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetResponse

`func (o *OutboundLog) GetResponse() OutboundLogResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OutboundLog) GetResponseOk() (*OutboundLogResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OutboundLog) SetResponse(v OutboundLogResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OutboundLog) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


