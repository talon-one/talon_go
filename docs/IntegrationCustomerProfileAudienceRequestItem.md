# IntegrationCustomerProfileAudienceRequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Defines the action to perform: - &#x60;add&#x60;: Adds the customer profile to the audience.    **Note**: If the customer profile does not exist, it will be created. The profile will not be visible in any Application   until a session or profile update is received for that profile. - &#x60;delete&#x60;: Removes the customer profile from the audience.  | 
**ProfileIntegrationId** | Pointer to **string** | The ID of this customer profile in the third-party integration. | 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration. | 

## Methods

### NewIntegrationCustomerProfileAudienceRequestItem

`func NewIntegrationCustomerProfileAudienceRequestItem(action string, profileIntegrationId string, integrationId string, ) *IntegrationCustomerProfileAudienceRequestItem`

NewIntegrationCustomerProfileAudienceRequestItem instantiates a new IntegrationCustomerProfileAudienceRequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationCustomerProfileAudienceRequestItemWithDefaults

`func NewIntegrationCustomerProfileAudienceRequestItemWithDefaults() *IntegrationCustomerProfileAudienceRequestItem`

NewIntegrationCustomerProfileAudienceRequestItemWithDefaults instantiates a new IntegrationCustomerProfileAudienceRequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetProfileIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationCustomerProfileAudienceRequestItem) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationCustomerProfileAudienceRequestItem) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


