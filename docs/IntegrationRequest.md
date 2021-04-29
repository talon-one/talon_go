# IntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | Pointer to [**NewCustomerSessionV2**](NewCustomerSessionV2.md) |  | 
**ResponseContent** | Pointer to **[]string** | Optional list of requested information to be present on the response related to the customer session update. Currently supported: \&quot;customerSession\&quot;, \&quot;customerProfile\&quot;, \&quot;coupons\&quot;, \&quot;triggeredCampaigns\&quot;, \&quot;referral\&quot;, \&quot;loyalty\&quot;, \&quot;event\&quot;, \&quot;awardedGiveaways\&quot; and \&quot;ruleFailureReasons\&quot;.  | [optional] 

## Methods

### GetCustomerSession

`func (o *IntegrationRequest) GetCustomerSession() NewCustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationRequest) GetCustomerSessionOk() (NewCustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSession

`func (o *IntegrationRequest) HasCustomerSession() bool`

HasCustomerSession returns a boolean if a field has been set.

### SetCustomerSession

`func (o *IntegrationRequest) SetCustomerSession(v NewCustomerSessionV2)`

SetCustomerSession gets a reference to the given NewCustomerSessionV2 and assigns it to the CustomerSession field.

### GetResponseContent

`func (o *IntegrationRequest) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationRequest) GetResponseContentOk() ([]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContent

`func (o *IntegrationRequest) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### SetResponseContent

`func (o *IntegrationRequest) SetResponseContent(v []string)`

SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


