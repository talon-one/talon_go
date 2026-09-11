# ApplicationMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application the customer belongs to. | 
**ApplicationName** | Pointer to **string** | The name of the Application the customer belongs to. | 

## Methods

### NewApplicationMembership

`func NewApplicationMembership(applicationId int64, applicationName string, ) *ApplicationMembership`

NewApplicationMembership instantiates a new ApplicationMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMembershipWithDefaults

`func NewApplicationMembershipWithDefaults() *ApplicationMembership`

NewApplicationMembershipWithDefaults instantiates a new ApplicationMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationMembership) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationMembership) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationMembership) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationName

`func (o *ApplicationMembership) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApplicationMembership) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApplicationMembership) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


