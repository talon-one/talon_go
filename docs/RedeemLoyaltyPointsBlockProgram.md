# RedeemLoyaltyPointsBlockProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the loyalty program. | 
**Name** | Pointer to **string** | The internal name of the loyalty program. | [optional] [readonly] 
**Title** | Pointer to **string** | The display name of the loyalty program. | [optional] [readonly] 

## Methods

### NewRedeemLoyaltyPointsBlockProgram

`func NewRedeemLoyaltyPointsBlockProgram(id int64, ) *RedeemLoyaltyPointsBlockProgram`

NewRedeemLoyaltyPointsBlockProgram instantiates a new RedeemLoyaltyPointsBlockProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemLoyaltyPointsBlockProgramWithDefaults

`func NewRedeemLoyaltyPointsBlockProgramWithDefaults() *RedeemLoyaltyPointsBlockProgram`

NewRedeemLoyaltyPointsBlockProgramWithDefaults instantiates a new RedeemLoyaltyPointsBlockProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedeemLoyaltyPointsBlockProgram) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedeemLoyaltyPointsBlockProgram) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedeemLoyaltyPointsBlockProgram) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *RedeemLoyaltyPointsBlockProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedeemLoyaltyPointsBlockProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedeemLoyaltyPointsBlockProgram) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RedeemLoyaltyPointsBlockProgram) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *RedeemLoyaltyPointsBlockProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RedeemLoyaltyPointsBlockProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RedeemLoyaltyPointsBlockProgram) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RedeemLoyaltyPointsBlockProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


