# IntegrationEventV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int32** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**Type** | Pointer to **string** | A string representing the event name. Must not be a reserved event name. You create this value when you [create an attribute](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) of type &#x60;event&#x60; in the Campaign Manager.  | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute). | [optional] 
**ResponseContent** | Pointer to **[]string** | Optional list of requested information to be present on the response related to the tracking custom event.  | [optional] 

## Methods

### GetProfileId

`func (o *IntegrationEventV2Request) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *IntegrationEventV2Request) GetProfileIdOk() (string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileId

`func (o *IntegrationEventV2Request) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileId

`func (o *IntegrationEventV2Request) SetProfileId(v string)`

SetProfileId gets a reference to the given string and assigns it to the ProfileId field.

### GetEvaluableCampaignIds

`func (o *IntegrationEventV2Request) GetEvaluableCampaignIds() []int32`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *IntegrationEventV2Request) GetEvaluableCampaignIdsOk() ([]int32, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluableCampaignIds

`func (o *IntegrationEventV2Request) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### SetEvaluableCampaignIds

`func (o *IntegrationEventV2Request) SetEvaluableCampaignIds(v []int32)`

SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.

### GetType

`func (o *IntegrationEventV2Request) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationEventV2Request) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *IntegrationEventV2Request) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *IntegrationEventV2Request) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetAttributes

`func (o *IntegrationEventV2Request) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationEventV2Request) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *IntegrationEventV2Request) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *IntegrationEventV2Request) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetResponseContent

`func (o *IntegrationEventV2Request) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationEventV2Request) GetResponseContentOk() ([]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContent

`func (o *IntegrationEventV2Request) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### SetResponseContent

`func (o *IntegrationEventV2Request) SetResponseContent(v []string)`

SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


