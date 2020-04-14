# FeatureFlag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the featureflag | 
**Value** | Pointer to **string** | The value for the featureflag | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was last created. | [optional] 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | [optional] 

## Methods

### GetName

`func (o *FeatureFlag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FeatureFlag) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *FeatureFlag) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *FeatureFlag) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetValue

`func (o *FeatureFlag) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureFlag) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *FeatureFlag) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *FeatureFlag) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetCreated

`func (o *FeatureFlag) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FeatureFlag) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *FeatureFlag) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *FeatureFlag) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *FeatureFlag) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *FeatureFlag) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *FeatureFlag) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *FeatureFlag) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


