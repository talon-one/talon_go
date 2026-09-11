# SupportCustomerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the customer profile. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time the customer profile was created. | 
**IntegrationId** | Pointer to **string** | The integration ID set by your integration layer. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | 
**ApplicationMemberships** | Pointer to [**[]ApplicationMembership**](ApplicationMembership.md) | The applications the customer belongs to. | 

## Methods

### NewSupportCustomerProfile

`func NewSupportCustomerProfile(id int64, created time.Time, integrationId string, attributes map[string]interface{}, applicationMemberships []ApplicationMembership, ) *SupportCustomerProfile`

NewSupportCustomerProfile instantiates a new SupportCustomerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCustomerProfileWithDefaults

`func NewSupportCustomerProfileWithDefaults() *SupportCustomerProfile`

NewSupportCustomerProfileWithDefaults instantiates a new SupportCustomerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SupportCustomerProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupportCustomerProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupportCustomerProfile) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *SupportCustomerProfile) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SupportCustomerProfile) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SupportCustomerProfile) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetIntegrationId

`func (o *SupportCustomerProfile) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *SupportCustomerProfile) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *SupportCustomerProfile) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetAttributes

`func (o *SupportCustomerProfile) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SupportCustomerProfile) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SupportCustomerProfile) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetApplicationMemberships

`func (o *SupportCustomerProfile) GetApplicationMemberships() []ApplicationMembership`

GetApplicationMemberships returns the ApplicationMemberships field if non-nil, zero value otherwise.

### GetApplicationMembershipsOk

`func (o *SupportCustomerProfile) GetApplicationMembershipsOk() (*[]ApplicationMembership, bool)`

GetApplicationMembershipsOk returns a tuple with the ApplicationMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMemberships

`func (o *SupportCustomerProfile) SetApplicationMemberships(v []ApplicationMembership)`

SetApplicationMemberships sets ApplicationMemberships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


