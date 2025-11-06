# NewLoyaltyTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the tier. | 
**MinPoints** | Pointer to **float32** | The minimum amount of points required to enter the tier. | 

## Methods

### NewNewLoyaltyTier

`func NewNewLoyaltyTier(name string, minPoints float32, ) *NewLoyaltyTier`

NewNewLoyaltyTier instantiates a new NewLoyaltyTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewLoyaltyTierWithDefaults

`func NewNewLoyaltyTierWithDefaults() *NewLoyaltyTier`

NewNewLoyaltyTierWithDefaults instantiates a new NewLoyaltyTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewLoyaltyTier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewLoyaltyTier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewLoyaltyTier) SetName(v string)`

SetName sets Name field to given value.


### GetMinPoints

`func (o *NewLoyaltyTier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *NewLoyaltyTier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *NewLoyaltyTier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


