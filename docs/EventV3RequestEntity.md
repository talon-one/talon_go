# EventV3RequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int64** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**Type** | Pointer to **string** | The name of the event. Must be a [custom event](https://docs.talon.one/docs/dev/concepts/entities/events#custom-events), not a built-in event. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute). | [optional] 
**IntegrationId** | Pointer to **string** | The unique ID of the event. Only one event with this ID can be registered.  | 
**ConnectedSessionId** | Pointer to **string** | The ID of the session to reference. The session must be in &#x60;closed&#x60; state. Otherwise, the API call will fail. | [optional] 
**ReferralCode** | Pointer to **string** | The referral code submitted with the event. The endpoint does not validate the code, and submitting a code does not redeem it. Use the \&quot;Referral code is valid\&quot; condition in the Rule Builder to validate and redeem the code, or \&quot;Referral code is valid (without redemption)\&quot; to validate without redeeming.  | [optional] 

## Methods

### NewEventV3RequestEntity

`func NewEventV3RequestEntity(profileId string, type_ string, integrationId string, ) *EventV3RequestEntity`

NewEventV3RequestEntity instantiates a new EventV3RequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventV3RequestEntityWithDefaults

`func NewEventV3RequestEntityWithDefaults() *EventV3RequestEntity`

NewEventV3RequestEntityWithDefaults instantiates a new EventV3RequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *EventV3RequestEntity) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EventV3RequestEntity) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *EventV3RequestEntity) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetStoreIntegrationId

`func (o *EventV3RequestEntity) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *EventV3RequestEntity) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *EventV3RequestEntity) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *EventV3RequestEntity) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetEvaluableCampaignIds

`func (o *EventV3RequestEntity) GetEvaluableCampaignIds() []int64`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *EventV3RequestEntity) GetEvaluableCampaignIdsOk() (*[]int64, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluableCampaignIds

`func (o *EventV3RequestEntity) SetEvaluableCampaignIds(v []int64)`

SetEvaluableCampaignIds sets EvaluableCampaignIds field to given value.

### HasEvaluableCampaignIds

`func (o *EventV3RequestEntity) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### GetType

`func (o *EventV3RequestEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventV3RequestEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventV3RequestEntity) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EventV3RequestEntity) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventV3RequestEntity) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventV3RequestEntity) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EventV3RequestEntity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIntegrationId

`func (o *EventV3RequestEntity) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *EventV3RequestEntity) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *EventV3RequestEntity) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetConnectedSessionId

`func (o *EventV3RequestEntity) GetConnectedSessionId() string`

GetConnectedSessionId returns the ConnectedSessionId field if non-nil, zero value otherwise.

### GetConnectedSessionIdOk

`func (o *EventV3RequestEntity) GetConnectedSessionIdOk() (*string, bool)`

GetConnectedSessionIdOk returns a tuple with the ConnectedSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSessionId

`func (o *EventV3RequestEntity) SetConnectedSessionId(v string)`

SetConnectedSessionId sets ConnectedSessionId field to given value.

### HasConnectedSessionId

`func (o *EventV3RequestEntity) HasConnectedSessionId() bool`

HasConnectedSessionId returns a boolean if a field has been set.

### GetReferralCode

`func (o *EventV3RequestEntity) GetReferralCode() string`

GetReferralCode returns the ReferralCode field if non-nil, zero value otherwise.

### GetReferralCodeOk

`func (o *EventV3RequestEntity) GetReferralCodeOk() (*string, bool)`

GetReferralCodeOk returns a tuple with the ReferralCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralCode

`func (o *EventV3RequestEntity) SetReferralCode(v string)`

SetReferralCode sets ReferralCode field to given value.

### HasReferralCode

`func (o *EventV3RequestEntity) HasReferralCode() bool`

HasReferralCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


