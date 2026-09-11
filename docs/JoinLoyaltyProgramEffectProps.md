# JoinLoyaltyProgramEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | The ID of the loyalty program the customer profile is joined to. | 
**JoinDate** | Pointer to [**time.Time**](time.Time.md) | The date and time when the customer profile joined the loyalty program. | 

## Methods

### NewJoinLoyaltyProgramEffectProps

`func NewJoinLoyaltyProgramEffectProps(programId int64, joinDate time.Time, ) *JoinLoyaltyProgramEffectProps`

NewJoinLoyaltyProgramEffectProps instantiates a new JoinLoyaltyProgramEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinLoyaltyProgramEffectPropsWithDefaults

`func NewJoinLoyaltyProgramEffectPropsWithDefaults() *JoinLoyaltyProgramEffectProps`

NewJoinLoyaltyProgramEffectPropsWithDefaults instantiates a new JoinLoyaltyProgramEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *JoinLoyaltyProgramEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *JoinLoyaltyProgramEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *JoinLoyaltyProgramEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetJoinDate

`func (o *JoinLoyaltyProgramEffectProps) GetJoinDate() time.Time`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *JoinLoyaltyProgramEffectProps) GetJoinDateOk() (*time.Time, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *JoinLoyaltyProgramEffectProps) SetJoinDate(v time.Time)`

SetJoinDate sets JoinDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


