# WebhookAuthenticationBaseBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the webhook authentication. | 
**Type** | Pointer to **string** | A webhook authentication discriminator of type &#x60;basic&#x60;. | 
**Data** | Pointer to [**WebhookAuthenticationDataBasic**](WebhookAuthenticationDataBasic.md) |  | 

## Methods

### NewWebhookAuthenticationBaseBasic

`func NewWebhookAuthenticationBaseBasic(name string, type_ string, data WebhookAuthenticationDataBasic, ) *WebhookAuthenticationBaseBasic`

NewWebhookAuthenticationBaseBasic instantiates a new WebhookAuthenticationBaseBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBaseBasicWithDefaults

`func NewWebhookAuthenticationBaseBasicWithDefaults() *WebhookAuthenticationBaseBasic`

NewWebhookAuthenticationBaseBasicWithDefaults instantiates a new WebhookAuthenticationBaseBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBaseBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBaseBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBaseBasic) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthenticationBaseBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBaseBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBaseBasic) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookAuthenticationBaseBasic) GetData() WebhookAuthenticationDataBasic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBaseBasic) GetDataOk() (*WebhookAuthenticationDataBasic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBaseBasic) SetData(v WebhookAuthenticationDataBasic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


