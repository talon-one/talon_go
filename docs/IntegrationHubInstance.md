# IntegrationHubInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** | The ID of the Prismatic integration instance. | 
**InstanceName** | Pointer to **string** | The name of the Prismatic integration instance. | 

## Methods

### NewIntegrationHubInstance

`func NewIntegrationHubInstance(instanceId string, instanceName string, ) *IntegrationHubInstance`

NewIntegrationHubInstance instantiates a new IntegrationHubInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubInstanceWithDefaults

`func NewIntegrationHubInstanceWithDefaults() *IntegrationHubInstance`

NewIntegrationHubInstanceWithDefaults instantiates a new IntegrationHubInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *IntegrationHubInstance) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *IntegrationHubInstance) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *IntegrationHubInstance) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceName

`func (o *IntegrationHubInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *IntegrationHubInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *IntegrationHubInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


