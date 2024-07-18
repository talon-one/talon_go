# CustomerProfileIntegrationRequestV2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudiencesChanges** | Pointer to [**ProfileAudiencesChanges**](.md) |  | [optional] 
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data as you need in one _Update customer profile_ request instead of sending extra requests to other endpoints.  | [optional] 

## Methods

### GetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2AllOf) GetAudiencesChanges() ProfileAudiencesChanges`

GetAudiencesChanges returns the AudiencesChanges field if non-nil, zero value otherwise.

### GetAudiencesChangesOk

`func (o *CustomerProfileIntegrationRequestV2AllOf) GetAudiencesChangesOk() (ProfileAudiencesChanges, bool)`

GetAudiencesChangesOk returns a tuple with the AudiencesChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2AllOf) HasAudiencesChanges() bool`

HasAudiencesChanges returns a boolean if a field has been set.

### SetAudiencesChanges

`func (o *CustomerProfileIntegrationRequestV2AllOf) SetAudiencesChanges(v ProfileAudiencesChanges)`

SetAudiencesChanges gets a reference to the given ProfileAudiencesChanges and assigns it to the AudiencesChanges field.

### GetResponseContent

`func (o *CustomerProfileIntegrationRequestV2AllOf) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *CustomerProfileIntegrationRequestV2AllOf) GetResponseContentOk() ([]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContent

`func (o *CustomerProfileIntegrationRequestV2AllOf) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### SetResponseContent

`func (o *CustomerProfileIntegrationRequestV2AllOf) SetResponseContent(v []string)`

SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


