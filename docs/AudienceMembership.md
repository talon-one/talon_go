# AudienceMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the audience belonging to this entity. | 
**Name** | Pointer to **string** | The Name of the audience belonging to this entity. | 

## Methods

### NewAudienceMembership

`func NewAudienceMembership(id int64, name string, ) *AudienceMembership`

NewAudienceMembership instantiates a new AudienceMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceMembershipWithDefaults

`func NewAudienceMembershipWithDefaults() *AudienceMembership`

NewAudienceMembershipWithDefaults instantiates a new AudienceMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceMembership) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceMembership) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceMembership) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *AudienceMembership) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceMembership) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceMembership) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


