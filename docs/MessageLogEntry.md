# MessageLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the message. | 
**Service** | Pointer to **string** | Name of the service that generated the log entry. | 
**ChangeType** | Pointer to **string** | Type of change that triggered the notification. | [optional] 
**NotificationId** | Pointer to **int64** | ID of the notification. | [optional] 
**NotificationName** | Pointer to **string** | The name of the notification. | [optional] 
**WebhookId** | Pointer to **int64** | ID of the webhook. | [optional] 
**WebhookName** | Pointer to **string** | The name of the webhook. | [optional] 
**Request** | Pointer to [**MessageLogRequest**](MessageLogRequest.md) |  | [optional] 
**Response** | Pointer to [**MessageLogResponse**](MessageLogResponse.md) |  | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the log entry was created. | 
**EntityType** | Pointer to **string** | The entity type the log is related to.  | 
**Url** | Pointer to **string** | The target URL of the request. | [optional] 
**ApplicationId** | Pointer to **int64** | Identifier of the Application. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | Identifier of the loyalty program. | [optional] 
**CampaignId** | Pointer to **int64** | Identifier of the campaign. | [optional] 

## Methods

### NewMessageLogEntry

`func NewMessageLogEntry(id string, service string, createdAt time.Time, entityType string, ) *MessageLogEntry`

NewMessageLogEntry instantiates a new MessageLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageLogEntryWithDefaults

`func NewMessageLogEntryWithDefaults() *MessageLogEntry`

NewMessageLogEntryWithDefaults instantiates a new MessageLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageLogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetService

`func (o *MessageLogEntry) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MessageLogEntry) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MessageLogEntry) SetService(v string)`

SetService sets Service field to given value.


### GetChangeType

`func (o *MessageLogEntry) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *MessageLogEntry) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *MessageLogEntry) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *MessageLogEntry) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetNotificationId

`func (o *MessageLogEntry) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *MessageLogEntry) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *MessageLogEntry) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *MessageLogEntry) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetNotificationName

`func (o *MessageLogEntry) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *MessageLogEntry) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *MessageLogEntry) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *MessageLogEntry) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetWebhookId

`func (o *MessageLogEntry) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *MessageLogEntry) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *MessageLogEntry) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *MessageLogEntry) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *MessageLogEntry) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *MessageLogEntry) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *MessageLogEntry) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *MessageLogEntry) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetRequest

`func (o *MessageLogEntry) GetRequest() MessageLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *MessageLogEntry) GetRequestOk() (*MessageLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *MessageLogEntry) SetRequest(v MessageLogRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *MessageLogEntry) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *MessageLogEntry) GetResponse() MessageLogResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MessageLogEntry) GetResponseOk() (*MessageLogResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MessageLogEntry) SetResponse(v MessageLogResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MessageLogEntry) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MessageLogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageLogEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageLogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEntityType

`func (o *MessageLogEntry) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MessageLogEntry) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MessageLogEntry) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetUrl

`func (o *MessageLogEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageLogEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageLogEntry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageLogEntry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetApplicationId

`func (o *MessageLogEntry) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *MessageLogEntry) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *MessageLogEntry) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *MessageLogEntry) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *MessageLogEntry) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *MessageLogEntry) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *MessageLogEntry) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *MessageLogEntry) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetCampaignId

`func (o *MessageLogEntry) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *MessageLogEntry) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *MessageLogEntry) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *MessageLogEntry) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


