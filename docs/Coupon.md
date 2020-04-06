# Coupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**Value** | Pointer to **string** | The actual coupon code. | 
**UsageLimit** | Pointer to **int32** | The number of times a coupon code can be redeemed. This can be set to 0 for no limit, but any campaign usage limits will still apply.  | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative. | [optional] 
**UsageCounter** | Pointer to **int32** | The number of times this coupon has been successfully used. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item | [optional] 
**ReferralId** | Pointer to **int32** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The integration ID of a referred customer profile. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this coupon. | [optional] 
**Reservation** | Pointer to **bool** | This value controls what reservations mean to a coupon. If set to true the coupon reservation is used to mark it as a favourite, if set to false the coupon reservation is used as a requirement of usage. This value defaults to true if not specified. | [optional] 
**BatchId** | Pointer to **string** | The id of the batch the coupon belongs to. | [optional] 

## Methods

### GetId

`func (o *Coupon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Coupon) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Coupon) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Coupon) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Coupon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Coupon) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Coupon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Coupon) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *Coupon) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Coupon) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *Coupon) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *Coupon) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetValue

`func (o *Coupon) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Coupon) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *Coupon) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *Coupon) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetUsageLimit

`func (o *Coupon) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Coupon) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *Coupon) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *Coupon) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetStartDate

`func (o *Coupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Coupon) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *Coupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *Coupon) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *Coupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Coupon) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *Coupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *Coupon) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetUsageCounter

`func (o *Coupon) GetUsageCounter() int32`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *Coupon) GetUsageCounterOk() (int32, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageCounter

`func (o *Coupon) HasUsageCounter() bool`

HasUsageCounter returns a boolean if a field has been set.

### SetUsageCounter

`func (o *Coupon) SetUsageCounter(v int32)`

SetUsageCounter gets a reference to the given int32 and assigns it to the UsageCounter field.

### GetAttributes

`func (o *Coupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Coupon) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Coupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Coupon) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetReferralId

`func (o *Coupon) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *Coupon) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *Coupon) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *Coupon) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetRecipientIntegrationId

`func (o *Coupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *Coupon) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *Coupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *Coupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetImportId

`func (o *Coupon) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Coupon) GetImportIdOk() (int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportId

`func (o *Coupon) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportId

`func (o *Coupon) SetImportId(v int32)`

SetImportId gets a reference to the given int32 and assigns it to the ImportId field.

### GetReservation

`func (o *Coupon) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Coupon) GetReservationOk() (bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservation

`func (o *Coupon) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservation

`func (o *Coupon) SetReservation(v bool)`

SetReservation gets a reference to the given bool and assigns it to the Reservation field.

### GetBatchId

`func (o *Coupon) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Coupon) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *Coupon) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *Coupon) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


