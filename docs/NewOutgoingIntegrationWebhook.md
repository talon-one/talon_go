# NewOutgoingIntegrationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Webhook title. | 
**Description** | Pointer to **string** | A description of the webhook. | [optional] 
**ApplicationIds** | Pointer to **[]int64** | IDs of the Applications to which a webhook must be linked. | 

## Methods

### NewNewOutgoingIntegrationWebhook

`func NewNewOutgoingIntegrationWebhook(title string, applicationIds []int64, ) *NewOutgoingIntegrationWebhook`

NewNewOutgoingIntegrationWebhook instantiates a new NewOutgoingIntegrationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewOutgoingIntegrationWebhookWithDefaults

`func NewNewOutgoingIntegrationWebhookWithDefaults() *NewOutgoingIntegrationWebhook`

NewNewOutgoingIntegrationWebhookWithDefaults instantiates a new NewOutgoingIntegrationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewOutgoingIntegrationWebhook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewOutgoingIntegrationWebhook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewOutgoingIntegrationWebhook) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewOutgoingIntegrationWebhook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewOutgoingIntegrationWebhook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewOutgoingIntegrationWebhook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewOutgoingIntegrationWebhook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *NewOutgoingIntegrationWebhook) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *NewOutgoingIntegrationWebhook) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *NewOutgoingIntegrationWebhook) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


