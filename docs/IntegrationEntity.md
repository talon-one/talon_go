# IntegrationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 

## Methods

### NewIntegrationEntity

`func NewIntegrationEntity(integrationId string, created time.Time, ) *IntegrationEntity`

NewIntegrationEntity instantiates a new IntegrationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationEntityWithDefaults

`func NewIntegrationEntityWithDefaults() *IntegrationEntity`

NewIntegrationEntityWithDefaults instantiates a new IntegrationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *IntegrationEntity) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationEntity) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationEntity) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetCreated

`func (o *IntegrationEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IntegrationEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IntegrationEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


