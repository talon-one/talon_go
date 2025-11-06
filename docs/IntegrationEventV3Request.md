# IntegrationEventV3Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int64** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**IntegrationId** | Pointer to **string** | The unique ID of the current event. Only one event with this ID could be activated, duplicated events are forbidden.  | 
**Type** | Pointer to **string** | A string representing the event name. Must not be a reserved event name. You create this value when you [create an attribute](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) of type &#x60;event&#x60; in the Campaign Manager.  | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute). | [optional] 
**ConnectedSessionID** | Pointer to **string** | The ID of the session that happened in the past. | [optional] 
**PreviousEventID** | Pointer to **string** | The unique identifier of the event that happened in the past. | [optional] 
**LoyaltyCards** | Pointer to **[]string** | Identifiers of the loyalty cards used during this event. | [optional] 
**ResponseContent** | Pointer to **[]string** | Optional list of requested information to be present on the response related to the tracking custom event.  | [optional] 

## Methods

### NewIntegrationEventV3Request

`func NewIntegrationEventV3Request(profileId string, integrationId string, type_ string, ) *IntegrationEventV3Request`

NewIntegrationEventV3Request instantiates a new IntegrationEventV3Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEventV3RequestWithDefaults

`func NewIntegrationEventV3RequestWithDefaults() *IntegrationEventV3Request`

NewIntegrationEventV3RequestWithDefaults instantiates a new IntegrationEventV3Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *IntegrationEventV3Request) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationEventV3Request) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *IntegrationEventV3Request) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetStoreIntegrationId

`func (o *IntegrationEventV3Request) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *IntegrationEventV3Request) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *IntegrationEventV3Request) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *IntegrationEventV3Request) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetEvaluableCampaignIds

`func (o *IntegrationEventV3Request) GetEvaluableCampaignIds() []int64`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *IntegrationEventV3Request) GetEvaluableCampaignIdsOk() (*[]int64, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluableCampaignIds

`func (o *IntegrationEventV3Request) SetEvaluableCampaignIds(v []int64)`

SetEvaluableCampaignIds sets EvaluableCampaignIds field to given value.

### HasEvaluableCampaignIds

`func (o *IntegrationEventV3Request) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### GetIntegrationId

`func (o *IntegrationEventV3Request) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationEventV3Request) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationEventV3Request) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetType

`func (o *IntegrationEventV3Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationEventV3Request) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationEventV3Request) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *IntegrationEventV3Request) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationEventV3Request) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationEventV3Request) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationEventV3Request) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetConnectedSessionID

`func (o *IntegrationEventV3Request) GetConnectedSessionID() string`

GetConnectedSessionID returns the ConnectedSessionID field if non-nil, zero value otherwise.

### GetConnectedSessionIDOk

`func (o *IntegrationEventV3Request) GetConnectedSessionIDOk() (*string, bool)`

GetConnectedSessionIDOk returns a tuple with the ConnectedSessionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSessionID

`func (o *IntegrationEventV3Request) SetConnectedSessionID(v string)`

SetConnectedSessionID sets ConnectedSessionID field to given value.

### HasConnectedSessionID

`func (o *IntegrationEventV3Request) HasConnectedSessionID() bool`

HasConnectedSessionID returns a boolean if a field has been set.

### GetPreviousEventID

`func (o *IntegrationEventV3Request) GetPreviousEventID() string`

GetPreviousEventID returns the PreviousEventID field if non-nil, zero value otherwise.

### GetPreviousEventIDOk

`func (o *IntegrationEventV3Request) GetPreviousEventIDOk() (*string, bool)`

GetPreviousEventIDOk returns a tuple with the PreviousEventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousEventID

`func (o *IntegrationEventV3Request) SetPreviousEventID(v string)`

SetPreviousEventID sets PreviousEventID field to given value.

### HasPreviousEventID

`func (o *IntegrationEventV3Request) HasPreviousEventID() bool`

HasPreviousEventID returns a boolean if a field has been set.

### GetLoyaltyCards

`func (o *IntegrationEventV3Request) GetLoyaltyCards() []string`

GetLoyaltyCards returns the LoyaltyCards field if non-nil, zero value otherwise.

### GetLoyaltyCardsOk

`func (o *IntegrationEventV3Request) GetLoyaltyCardsOk() (*[]string, bool)`

GetLoyaltyCardsOk returns a tuple with the LoyaltyCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCards

`func (o *IntegrationEventV3Request) SetLoyaltyCards(v []string)`

SetLoyaltyCards sets LoyaltyCards field to given value.

### HasLoyaltyCards

`func (o *IntegrationEventV3Request) HasLoyaltyCards() bool`

HasLoyaltyCards returns a boolean if a field has been set.

### GetResponseContent

`func (o *IntegrationEventV3Request) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationEventV3Request) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntegrationEventV3Request) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntegrationEventV3Request) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


