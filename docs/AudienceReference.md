# AudienceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the audience. | 
**Integration** | Pointer to **string** | The third-party integration of the audience. | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration. | [optional] 

## Methods

### NewAudienceReference

`func NewAudienceReference(id int64, ) *AudienceReference`

NewAudienceReference instantiates a new AudienceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceReferenceWithDefaults

`func NewAudienceReferenceWithDefaults() *AudienceReference`

NewAudienceReferenceWithDefaults instantiates a new AudienceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceReference) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegration

`func (o *AudienceReference) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AudienceReference) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AudienceReference) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AudienceReference) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationId

`func (o *AudienceReference) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *AudienceReference) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *AudienceReference) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *AudienceReference) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


