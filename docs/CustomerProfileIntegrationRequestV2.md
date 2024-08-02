# CustomerProfileIntegrationRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**AudiencesChanges** | Pointer to [**ProfileAudiencesChanges**](ProfileAudiencesChanges.md) |  | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int32** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data as you need in one _Update customer profile_ request instead of sending extra requests to other endpoints.  | [optional] 

## Methods

### GetAttributes

`func (o *CustomerProfileIntegrationRequestV2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerProfileIntegrationRequestV2) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CustomerProfileIntegrationRequestV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CustomerProfileIntegrationRequestV2) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChanges() ProfileAudiencesChanges`

GetAudiencesChanges returns the AudiencesChanges field if non-nil, zero value otherwise.

### GetAudiencesChangesOk

`func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChangesOk() (ProfileAudiencesChanges, bool)`

GetAudiencesChangesOk returns a tuple with the AudiencesChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) HasAudiencesChanges() bool`

HasAudiencesChanges returns a boolean if a field has been set.

### SetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) SetAudiencesChanges(v ProfileAudiencesChanges)`

SetAudiencesChanges gets a reference to the given ProfileAudiencesChanges and assigns it to the AudiencesChanges field.

### GetEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIds() []int32`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIdsOk() ([]int32, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### SetEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) SetEvaluableCampaignIds(v []int32)`

SetEvaluableCampaignIds gets a reference to the given []int32 and assigns it to the EvaluableCampaignIds field.

### GetResponseContent

`func (o *CustomerProfileIntegrationRequestV2) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *CustomerProfileIntegrationRequestV2) GetResponseContentOk() ([]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContent

`func (o *CustomerProfileIntegrationRequestV2) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### SetResponseContent

`func (o *CustomerProfileIntegrationRequestV2) SetResponseContent(v []string)`

SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


