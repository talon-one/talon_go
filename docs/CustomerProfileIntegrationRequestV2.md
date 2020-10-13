# CustomerProfileIntegrationRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 
**ResponseContent** | Pointer to **[]string** | Optional list of requested information to be present on the response related to the customer profile update. Currently supported: \&quot;customerProfile\&quot;, \&quot;triggeredCampaigns\&quot;, \&quot;loyalty\&quot; and \&quot;event\&quot;.  | [optional] 

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


