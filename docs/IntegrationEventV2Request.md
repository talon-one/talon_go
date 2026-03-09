# IntegrationEventV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int64** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**Type** | Pointer to **string** | A string representing the event name. Must not be a reserved event name. You create this value when you [create an attribute](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) of type &#x60;event&#x60; in the Campaign Manager.  | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute). | [optional] 
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data back as needed from one request instead of sending extra requests to other endpoints.  | [optional] 
**LoyaltyCards** | Pointer to **[]string** | Identifiers of the loyalty cards used during this event. | [optional] 

## Methods

### NewIntegrationEventV2Request

`func NewIntegrationEventV2Request(type_ string, ) *IntegrationEventV2Request`

NewIntegrationEventV2Request instantiates a new IntegrationEventV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEventV2RequestWithDefaults

`func NewIntegrationEventV2RequestWithDefaults() *IntegrationEventV2Request`

NewIntegrationEventV2RequestWithDefaults instantiates a new IntegrationEventV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *IntegrationEventV2Request) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationEventV2Request) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *IntegrationEventV2Request) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *IntegrationEventV2Request) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetStoreIntegrationId

`func (o *IntegrationEventV2Request) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *IntegrationEventV2Request) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *IntegrationEventV2Request) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *IntegrationEventV2Request) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetEvaluableCampaignIds

`func (o *IntegrationEventV2Request) GetEvaluableCampaignIds() []int64`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *IntegrationEventV2Request) GetEvaluableCampaignIdsOk() (*[]int64, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluableCampaignIds

`func (o *IntegrationEventV2Request) SetEvaluableCampaignIds(v []int64)`

SetEvaluableCampaignIds sets EvaluableCampaignIds field to given value.

### HasEvaluableCampaignIds

`func (o *IntegrationEventV2Request) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### GetType

`func (o *IntegrationEventV2Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationEventV2Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationEventV2Request) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *IntegrationEventV2Request) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationEventV2Request) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationEventV2Request) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationEventV2Request) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetResponseContent

`func (o *IntegrationEventV2Request) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationEventV2Request) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntegrationEventV2Request) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntegrationEventV2Request) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetLoyaltyCards

`func (o *IntegrationEventV2Request) GetLoyaltyCards() []string`

GetLoyaltyCards returns the LoyaltyCards field if non-nil, zero value otherwise.

### GetLoyaltyCardsOk

`func (o *IntegrationEventV2Request) GetLoyaltyCardsOk() (*[]string, bool)`

GetLoyaltyCardsOk returns a tuple with the LoyaltyCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCards

`func (o *IntegrationEventV2Request) SetLoyaltyCards(v []string)`

SetLoyaltyCards sets LoyaltyCards field to given value.

### HasLoyaltyCards

`func (o *IntegrationEventV2Request) HasLoyaltyCards() bool`

HasLoyaltyCards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


