# OutboundLogBase

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

## Methods

### NewOutboundLogBase

`func NewOutboundLogBase(uuid string, notificationType string, ) *OutboundLogBase`

NewOutboundLogBase instantiates a new OutboundLogBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogBaseWithDefaults

`func NewOutboundLogBaseWithDefaults() *OutboundLogBase`

NewOutboundLogBaseWithDefaults instantiates a new OutboundLogBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *OutboundLogBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OutboundLogBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OutboundLogBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNotificationId

`func (o *OutboundLogBase) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *OutboundLogBase) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *OutboundLogBase) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *OutboundLogBase) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetNotificationName

`func (o *OutboundLogBase) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *OutboundLogBase) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *OutboundLogBase) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *OutboundLogBase) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetWebhookId

`func (o *OutboundLogBase) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *OutboundLogBase) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *OutboundLogBase) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *OutboundLogBase) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *OutboundLogBase) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *OutboundLogBase) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *OutboundLogBase) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *OutboundLogBase) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetNotificationType

`func (o *OutboundLogBase) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *OutboundLogBase) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *OutboundLogBase) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetApplicationId

`func (o *OutboundLogBase) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OutboundLogBase) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OutboundLogBase) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OutboundLogBase) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *OutboundLogBase) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *OutboundLogBase) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *OutboundLogBase) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *OutboundLogBase) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetRequest

`func (o *OutboundLogBase) GetRequest() OutboundLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OutboundLogBase) GetRequestOk() (*OutboundLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OutboundLogBase) SetRequest(v OutboundLogRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OutboundLogBase) HasRequest() bool`

HasRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


