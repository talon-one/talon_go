# ProjectedTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectedActivePoints** | Pointer to **float32** | The active points of the customer when their current tier expires. | 
**StayInTierPoints** | Pointer to **float32** | The number of points the customer needs to stay in the current tier.  **Note**: This is included in the response when the customer is projected to be downgraded.  | [optional] 
**ProjectedTierName** | Pointer to **string** | The name of the tier the user will enter once their current tier expires. | [optional] 

## Methods

### GetProjectedActivePoints

`func (o *ProjectedTier) GetProjectedActivePoints() float32`

GetProjectedActivePoints returns the ProjectedActivePoints field if non-nil, zero value otherwise.

### GetProjectedActivePointsOk

`func (o *ProjectedTier) GetProjectedActivePointsOk() (float32, bool)`

GetProjectedActivePointsOk returns a tuple with the ProjectedActivePoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectedActivePoints

`func (o *ProjectedTier) HasProjectedActivePoints() bool`

HasProjectedActivePoints returns a boolean if a field has been set.

### SetProjectedActivePoints

`func (o *ProjectedTier) SetProjectedActivePoints(v float32)`

SetProjectedActivePoints gets a reference to the given float32 and assigns it to the ProjectedActivePoints field.

### GetStayInTierPoints

`func (o *ProjectedTier) GetStayInTierPoints() float32`

GetStayInTierPoints returns the StayInTierPoints field if non-nil, zero value otherwise.

### GetStayInTierPointsOk

`func (o *ProjectedTier) GetStayInTierPointsOk() (float32, bool)`

GetStayInTierPointsOk returns a tuple with the StayInTierPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStayInTierPoints

`func (o *ProjectedTier) HasStayInTierPoints() bool`

HasStayInTierPoints returns a boolean if a field has been set.

### SetStayInTierPoints

`func (o *ProjectedTier) SetStayInTierPoints(v float32)`

SetStayInTierPoints gets a reference to the given float32 and assigns it to the StayInTierPoints field.

### GetProjectedTierName

`func (o *ProjectedTier) GetProjectedTierName() string`

GetProjectedTierName returns the ProjectedTierName field if non-nil, zero value otherwise.

### GetProjectedTierNameOk

`func (o *ProjectedTier) GetProjectedTierNameOk() (string, bool)`

GetProjectedTierNameOk returns a tuple with the ProjectedTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectedTierName

`func (o *ProjectedTier) HasProjectedTierName() bool`

HasProjectedTierName returns a boolean if a field has been set.

### SetProjectedTierName

`func (o *ProjectedTier) SetProjectedTierName(v string)`

SetProjectedTierName gets a reference to the given string and assigns it to the ProjectedTierName field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


