# OutboundMessage

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
**Request** | Pointer to [**OutboundLogRequest**](OutboundLogRequest.md) |  | 
**FirstLogAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the first log entry for this message. | 
**LastLogAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp of the last log entry for this message. | 
**LastResponseCode** | Pointer to **int64** | HTTP status code from the latest response. | [optional] 
**Status** | Pointer to **string** |  | 
**RetryCount** | Pointer to **int64** | Number of retries. | [optional] 
**Responses** | Pointer to [**[]OutboundMessageResponse**](OutboundMessageResponse.md) | Log entries for this message. Omitted when &#x60;includeLogs&#x3D;false&#x60;. | [optional] 

## Methods

### NewOutboundMessage

`func NewOutboundMessage(uuid string, notificationType string, request OutboundLogRequest, firstLogAt time.Time, lastLogAt time.Time, status string, ) *OutboundMessage`

NewOutboundMessage instantiates a new OutboundMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessageWithDefaults

`func NewOutboundMessageWithDefaults() *OutboundMessage`

NewOutboundMessageWithDefaults instantiates a new OutboundMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *OutboundMessage) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OutboundMessage) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OutboundMessage) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNotificationId

`func (o *OutboundMessage) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *OutboundMessage) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *OutboundMessage) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *OutboundMessage) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetNotificationName

`func (o *OutboundMessage) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *OutboundMessage) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *OutboundMessage) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *OutboundMessage) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetWebhookId

`func (o *OutboundMessage) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *OutboundMessage) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *OutboundMessage) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *OutboundMessage) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *OutboundMessage) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *OutboundMessage) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *OutboundMessage) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *OutboundMessage) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetNotificationType

`func (o *OutboundMessage) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *OutboundMessage) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *OutboundMessage) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetApplicationId

`func (o *OutboundMessage) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OutboundMessage) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OutboundMessage) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OutboundMessage) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *OutboundMessage) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *OutboundMessage) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *OutboundMessage) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *OutboundMessage) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetRequest

`func (o *OutboundMessage) GetRequest() OutboundLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OutboundMessage) GetRequestOk() (*OutboundLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OutboundMessage) SetRequest(v OutboundLogRequest)`

SetRequest sets Request field to given value.


### GetFirstLogAt

`func (o *OutboundMessage) GetFirstLogAt() time.Time`

GetFirstLogAt returns the FirstLogAt field if non-nil, zero value otherwise.

### GetFirstLogAtOk

`func (o *OutboundMessage) GetFirstLogAtOk() (*time.Time, bool)`

GetFirstLogAtOk returns a tuple with the FirstLogAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstLogAt

`func (o *OutboundMessage) SetFirstLogAt(v time.Time)`

SetFirstLogAt sets FirstLogAt field to given value.


### GetLastLogAt

`func (o *OutboundMessage) GetLastLogAt() time.Time`

GetLastLogAt returns the LastLogAt field if non-nil, zero value otherwise.

### GetLastLogAtOk

`func (o *OutboundMessage) GetLastLogAtOk() (*time.Time, bool)`

GetLastLogAtOk returns a tuple with the LastLogAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogAt

`func (o *OutboundMessage) SetLastLogAt(v time.Time)`

SetLastLogAt sets LastLogAt field to given value.


### GetLastResponseCode

`func (o *OutboundMessage) GetLastResponseCode() int64`

GetLastResponseCode returns the LastResponseCode field if non-nil, zero value otherwise.

### GetLastResponseCodeOk

`func (o *OutboundMessage) GetLastResponseCodeOk() (*int64, bool)`

GetLastResponseCodeOk returns a tuple with the LastResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseCode

`func (o *OutboundMessage) SetLastResponseCode(v int64)`

SetLastResponseCode sets LastResponseCode field to given value.

### HasLastResponseCode

`func (o *OutboundMessage) HasLastResponseCode() bool`

HasLastResponseCode returns a boolean if a field has been set.

### GetStatus

`func (o *OutboundMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OutboundMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OutboundMessage) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRetryCount

`func (o *OutboundMessage) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *OutboundMessage) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *OutboundMessage) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *OutboundMessage) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetResponses

`func (o *OutboundMessage) GetResponses() []OutboundMessageResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *OutboundMessage) GetResponsesOk() (*[]OutboundMessageResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *OutboundMessage) SetResponses(v []OutboundMessageResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *OutboundMessage) HasResponses() bool`

HasResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


