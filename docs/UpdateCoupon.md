# UpdateCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**RecipientIntegrationId** | Pointer to **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile. | [optional] 

## Methods

### GetAttributes

`func (o *UpdateCoupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCoupon) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *UpdateCoupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *UpdateCoupon) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetImplicitlyReserved

`func (o *UpdateCoupon) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *UpdateCoupon) GetImplicitlyReservedOk() (bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImplicitlyReserved

`func (o *UpdateCoupon) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.

### SetImplicitlyReserved

`func (o *UpdateCoupon) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.

### GetIsReservationMandatory

`func (o *UpdateCoupon) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *UpdateCoupon) GetIsReservationMandatoryOk() (bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReservationMandatory

`func (o *UpdateCoupon) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### SetIsReservationMandatory

`func (o *UpdateCoupon) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.

### GetRecipientIntegrationId

`func (o *UpdateCoupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *UpdateCoupon) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *UpdateCoupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *UpdateCoupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


