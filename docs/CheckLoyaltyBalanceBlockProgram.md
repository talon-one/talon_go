# CheckLoyaltyBalanceBlockProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the loyalty program. | 
**Name** | Pointer to **string** | The internal name of the loyalty program. | 
**Title** | Pointer to **string** | The display name of the loyalty program. | 

## Methods

### NewCheckLoyaltyBalanceBlockProgram

`func NewCheckLoyaltyBalanceBlockProgram(id int64, name string, title string, ) *CheckLoyaltyBalanceBlockProgram`

NewCheckLoyaltyBalanceBlockProgram instantiates a new CheckLoyaltyBalanceBlockProgram object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckLoyaltyBalanceBlockProgramWithDefaults

`func NewCheckLoyaltyBalanceBlockProgramWithDefaults() *CheckLoyaltyBalanceBlockProgram`

NewCheckLoyaltyBalanceBlockProgramWithDefaults instantiates a new CheckLoyaltyBalanceBlockProgram object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckLoyaltyBalanceBlockProgram) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckLoyaltyBalanceBlockProgram) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckLoyaltyBalanceBlockProgram) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CheckLoyaltyBalanceBlockProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckLoyaltyBalanceBlockProgram) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckLoyaltyBalanceBlockProgram) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CheckLoyaltyBalanceBlockProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CheckLoyaltyBalanceBlockProgram) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CheckLoyaltyBalanceBlockProgram) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


