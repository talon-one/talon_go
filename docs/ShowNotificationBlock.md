# ShowNotificationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**NotificationType** | Pointer to **string** | The type of notification to display. | 
**Title** | Pointer to **string** | The notification heading shown to the customer. | 
**Body** | Pointer to **string** | The notification body text. Supports template placeholders (e.g. \&quot;{{$Session.Total}}\&quot;) evaluated at rule execution time. | [optional] 
**OnFailure** | Pointer to **[]map[string]interface{}** | Blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewShowNotificationBlock

`func NewShowNotificationBlock(type_ string, notificationType string, title string, ) *ShowNotificationBlock`

NewShowNotificationBlock instantiates a new ShowNotificationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShowNotificationBlockWithDefaults

`func NewShowNotificationBlockWithDefaults() *ShowNotificationBlock`

NewShowNotificationBlockWithDefaults instantiates a new ShowNotificationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShowNotificationBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShowNotificationBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShowNotificationBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShowNotificationBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ShowNotificationBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShowNotificationBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShowNotificationBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *ShowNotificationBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ShowNotificationBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ShowNotificationBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ShowNotificationBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotificationType

`func (o *ShowNotificationBlock) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *ShowNotificationBlock) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *ShowNotificationBlock) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTitle

`func (o *ShowNotificationBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShowNotificationBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShowNotificationBlock) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *ShowNotificationBlock) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ShowNotificationBlock) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ShowNotificationBlock) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ShowNotificationBlock) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetOnFailure

`func (o *ShowNotificationBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *ShowNotificationBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *ShowNotificationBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *ShowNotificationBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *ShowNotificationBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *ShowNotificationBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *ShowNotificationBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *ShowNotificationBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


