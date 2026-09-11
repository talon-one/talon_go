# WebhookAuthenticationBaseCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the webhook authentication. | 
**Type** | Pointer to **string** | A webhook authentication discriminator of type &#x60;custom&#x60;. | 
**Data** | Pointer to [**WebhookAuthenticationDataCustom**](WebhookAuthenticationDataCustom.md) |  | 

## Methods

### NewWebhookAuthenticationBaseCustom

`func NewWebhookAuthenticationBaseCustom(name string, type_ string, data WebhookAuthenticationDataCustom, ) *WebhookAuthenticationBaseCustom`

NewWebhookAuthenticationBaseCustom instantiates a new WebhookAuthenticationBaseCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBaseCustomWithDefaults

`func NewWebhookAuthenticationBaseCustomWithDefaults() *WebhookAuthenticationBaseCustom`

NewWebhookAuthenticationBaseCustomWithDefaults instantiates a new WebhookAuthenticationBaseCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBaseCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBaseCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBaseCustom) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthenticationBaseCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBaseCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBaseCustom) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookAuthenticationBaseCustom) GetData() WebhookAuthenticationDataCustom`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBaseCustom) GetDataOk() (*WebhookAuthenticationDataCustom, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBaseCustom) SetData(v WebhookAuthenticationDataCustom)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


