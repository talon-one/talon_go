# LoyaltyMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Joined** | Pointer to [**time.Time**](time.Time.md) | The moment in which the loyalty program was joined. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | The ID of the loyalty program belonging to this entity. | 

## Methods

### NewLoyaltyMembership

`func NewLoyaltyMembership(loyaltyProgramId int64, ) *LoyaltyMembership`

NewLoyaltyMembership instantiates a new LoyaltyMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyMembershipWithDefaults

`func NewLoyaltyMembershipWithDefaults() *LoyaltyMembership`

NewLoyaltyMembershipWithDefaults instantiates a new LoyaltyMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJoined

`func (o *LoyaltyMembership) GetJoined() time.Time`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *LoyaltyMembership) GetJoinedOk() (*time.Time, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *LoyaltyMembership) SetJoined(v time.Time)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *LoyaltyMembership) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *LoyaltyMembership) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *LoyaltyMembership) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *LoyaltyMembership) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


