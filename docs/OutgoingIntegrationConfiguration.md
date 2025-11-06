# OutgoingIntegrationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**AccountId** | Pointer to **int64** | The ID of the account to which this configuration belongs. | 
**TypeId** | Pointer to **int64** | The outgoing integration type ID. | 
**Policy** | Pointer to [**map[string]interface{}**](.md) | The outgoing integration policy specific to each integration type. | 

## Methods

### NewOutgoingIntegrationConfiguration

`func NewOutgoingIntegrationConfiguration(id int64, accountId int64, typeId int64, policy map[string]interface{}, ) *OutgoingIntegrationConfiguration`

NewOutgoingIntegrationConfiguration instantiates a new OutgoingIntegrationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutgoingIntegrationConfigurationWithDefaults

`func NewOutgoingIntegrationConfigurationWithDefaults() *OutgoingIntegrationConfiguration`

NewOutgoingIntegrationConfigurationWithDefaults instantiates a new OutgoingIntegrationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutgoingIntegrationConfiguration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationConfiguration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutgoingIntegrationConfiguration) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *OutgoingIntegrationConfiguration) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OutgoingIntegrationConfiguration) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OutgoingIntegrationConfiguration) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetTypeId

`func (o *OutgoingIntegrationConfiguration) GetTypeId() int64`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *OutgoingIntegrationConfiguration) GetTypeIdOk() (*int64, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *OutgoingIntegrationConfiguration) SetTypeId(v int64)`

SetTypeId sets TypeId field to given value.


### GetPolicy

`func (o *OutgoingIntegrationConfiguration) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OutgoingIntegrationConfiguration) GetPolicyOk() (*map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *OutgoingIntegrationConfiguration) SetPolicy(v map[string]interface{})`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


