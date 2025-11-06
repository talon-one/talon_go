# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the feature flag. | 
**Value** | Pointer to **string** | The value of the feature flag. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last created. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | [optional] 

## Methods

### NewFeatureFlag

`func NewFeatureFlag(name string, value string, ) *FeatureFlag`

NewFeatureFlag instantiates a new FeatureFlag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureFlagWithDefaults

`func NewFeatureFlagWithDefaults() *FeatureFlag`

NewFeatureFlagWithDefaults instantiates a new FeatureFlag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *FeatureFlag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureFlag) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeatureFlag) SetValue(v string)`

SetValue sets Value field to given value.


### GetCreated

`func (o *FeatureFlag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FeatureFlag) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FeatureFlag) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FeatureFlag) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *FeatureFlag) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *FeatureFlag) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *FeatureFlag) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *FeatureFlag) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


