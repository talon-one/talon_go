# CampaignLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the loyalty program. | 
**Name** | Pointer to **string** | The name of the loyalty program. | 
**Tiers** | Pointer to **[]string** | The names of the tiers in the loyalty program. | 
**CardBased** | Pointer to **bool** | Whether the loyalty program is card-based. | 

## Methods

### NewCampaignLoyaltyProgram

`func NewCampaignLoyaltyProgram(id int64, name string, tiers []string, cardBased bool, ) *CampaignLoyaltyProgram`

NewCampaignLoyaltyProgram instantiates a new CampaignLoyaltyProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignLoyaltyProgramWithDefaults

`func NewCampaignLoyaltyProgramWithDefaults() *CampaignLoyaltyProgram`

NewCampaignLoyaltyProgramWithDefaults instantiates a new CampaignLoyaltyProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignLoyaltyProgram) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignLoyaltyProgram) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignLoyaltyProgram) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CampaignLoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignLoyaltyProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignLoyaltyProgram) SetName(v string)`

SetName sets Name field to given value.


### GetTiers

`func (o *CampaignLoyaltyProgram) GetTiers() []string`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *CampaignLoyaltyProgram) GetTiersOk() (*[]string, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *CampaignLoyaltyProgram) SetTiers(v []string)`

SetTiers sets Tiers field to given value.


### GetCardBased

`func (o *CampaignLoyaltyProgram) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *CampaignLoyaltyProgram) GetCardBasedOk() (*bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBased

`func (o *CampaignLoyaltyProgram) SetCardBased(v bool)`

SetCardBased sets CardBased field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


