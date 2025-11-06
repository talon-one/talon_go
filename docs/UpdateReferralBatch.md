# UpdateReferralBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**BatchID** | Pointer to **string** | The id of the batch the referral belongs to. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the referral code. Referral never expires if this is omitted. | [optional] 
**UsageLimit** | Pointer to **int64** | The number of times a referral code can be used. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | [optional] 

## Methods

### NewUpdateReferralBatch

`func NewUpdateReferralBatch(batchID string, ) *UpdateReferralBatch`

NewUpdateReferralBatch instantiates a new UpdateReferralBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReferralBatchWithDefaults

`func NewUpdateReferralBatchWithDefaults() *UpdateReferralBatch`

NewUpdateReferralBatchWithDefaults instantiates a new UpdateReferralBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *UpdateReferralBatch) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateReferralBatch) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateReferralBatch) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateReferralBatch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBatchID

`func (o *UpdateReferralBatch) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *UpdateReferralBatch) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *UpdateReferralBatch) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.


### GetStartDate

`func (o *UpdateReferralBatch) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateReferralBatch) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateReferralBatch) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateReferralBatch) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *UpdateReferralBatch) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *UpdateReferralBatch) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *UpdateReferralBatch) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *UpdateReferralBatch) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *UpdateReferralBatch) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UpdateReferralBatch) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *UpdateReferralBatch) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *UpdateReferralBatch) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


