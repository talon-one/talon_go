# CustomerProfileIntegrationRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**EvaluableCampaignIds** | Pointer to **[]int64** | When using the &#x60;dry&#x60; query parameter, use this property to list the campaign to be evaluated by the Rule Engine.  These campaigns will be evaluated, even if they are disabled, allowing you to test specific campaigns before activating them.  | [optional] 
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data back as needed from one request instead of sending extra requests to other endpoints.  | [optional] 
**AudiencesChanges** | Pointer to [**ProfileAudiencesChanges**](ProfileAudiencesChanges.md) |  | [optional] 

## Methods

### NewCustomerProfileIntegrationRequestV2

`func NewCustomerProfileIntegrationRequestV2() *CustomerProfileIntegrationRequestV2`

NewCustomerProfileIntegrationRequestV2 instantiates a new CustomerProfileIntegrationRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProfileIntegrationRequestV2WithDefaults

`func NewCustomerProfileIntegrationRequestV2WithDefaults() *CustomerProfileIntegrationRequestV2`

NewCustomerProfileIntegrationRequestV2WithDefaults instantiates a new CustomerProfileIntegrationRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *CustomerProfileIntegrationRequestV2) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerProfileIntegrationRequestV2) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerProfileIntegrationRequestV2) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CustomerProfileIntegrationRequestV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIds() []int64`

GetEvaluableCampaignIds returns the EvaluableCampaignIds field if non-nil, zero value otherwise.

### GetEvaluableCampaignIdsOk

`func (o *CustomerProfileIntegrationRequestV2) GetEvaluableCampaignIdsOk() (*[]int64, bool)`

GetEvaluableCampaignIdsOk returns a tuple with the EvaluableCampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) SetEvaluableCampaignIds(v []int64)`

SetEvaluableCampaignIds sets EvaluableCampaignIds field to given value.

### HasEvaluableCampaignIds

`func (o *CustomerProfileIntegrationRequestV2) HasEvaluableCampaignIds() bool`

HasEvaluableCampaignIds returns a boolean if a field has been set.

### GetResponseContent

`func (o *CustomerProfileIntegrationRequestV2) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *CustomerProfileIntegrationRequestV2) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *CustomerProfileIntegrationRequestV2) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *CustomerProfileIntegrationRequestV2) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### GetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChanges() ProfileAudiencesChanges`

GetAudiencesChanges returns the AudiencesChanges field if non-nil, zero value otherwise.

### GetAudiencesChangesOk

`func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChangesOk() (*ProfileAudiencesChanges, bool)`

GetAudiencesChangesOk returns a tuple with the AudiencesChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) SetAudiencesChanges(v ProfileAudiencesChanges)`

SetAudiencesChanges sets AudiencesChanges field to given value.

### HasAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2) HasAudiencesChanges() bool`

HasAudiencesChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


