# \IntegrationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudienceV2**](IntegrationApi.md#CreateAudienceV2) | **Post** /v2/audiences | Create audience
[**CreateCouponReservation**](IntegrationApi.md#CreateCouponReservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
[**CreateReferral**](IntegrationApi.md#CreateReferral) | **Post** /v1/referrals | Create referral code for an advocate
[**CreateReferralsForMultipleAdvocates**](IntegrationApi.md#CreateReferralsForMultipleAdvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates
[**DeleteAudienceMembershipsV2**](IntegrationApi.md#DeleteAudienceMembershipsV2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
[**DeleteAudienceV2**](IntegrationApi.md#DeleteAudienceV2) | **Delete** /v2/audiences/{audienceId} | Delete audience
[**DeleteCouponReservation**](IntegrationApi.md#DeleteCouponReservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
[**DeleteCustomerData**](IntegrationApi.md#DeleteCustomerData) | **Delete** /v1/customer_data/{integrationId} | Delete the personal data of a customer
[**GetCustomerInventory**](IntegrationApi.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List data associated with a specific customer profile
[**GetReservedCustomers**](IntegrationApi.md#GetReservedCustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List users that have this coupon reserved
[**TrackEvent**](IntegrationApi.md#TrackEvent) | **Post** /v1/events | Track an Event
[**UpdateAudienceCustomersAttributes**](IntegrationApi.md#UpdateAudienceCustomersAttributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
[**UpdateAudienceV2**](IntegrationApi.md#UpdateAudienceV2) | **Put** /v2/audiences/{audienceId} | Update audience
[**UpdateCustomerProfileAudiences**](IntegrationApi.md#UpdateCustomerProfileAudiences) | **Post** /v2/customer_audiences | Update multiple customer profiles&#39; audiences
[**UpdateCustomerProfileV2**](IntegrationApi.md#UpdateCustomerProfileV2) | **Put** /v2/customer_profiles/{integrationId} | Update customer profile
[**UpdateCustomerProfilesV2**](IntegrationApi.md#UpdateCustomerProfilesV2) | **Put** /v2/customer_profiles | Update multiple customer profiles
[**UpdateCustomerSessionV2**](IntegrationApi.md#UpdateCustomerSessionV2) | **Put** /v2/customer_sessions/{customerSessionId} | Update customer session



## CreateAudienceV2

> Audience CreateAudienceV2(ctx).Body(body).Execute()

Create audience



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAudience**](NewAudience.md) |  | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponReservation

> Coupon CreateCouponReservation(ctx, couponValue).Body(body).Execute()

Create coupon reservation



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The value of a coupon | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) |  | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferral

> Referral CreateReferral(ctx).Body(body).Execute()

Create referral code for an advocate



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewReferral**](NewReferral.md) |  | 

### Return type

[**Referral**](Referral.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferralsForMultipleAdvocates

> InlineResponse201 CreateReferralsForMultipleAdvocates(ctx).Body(body).Silent(silent).Execute()

Create referral codes for multiple advocates



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralsForMultipleAdvocatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewReferralsForMultipleAdvocates**](NewReferralsForMultipleAdvocates.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to yes]

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceMembershipsV2

> DeleteAudienceMembershipsV2(ctx, audienceId).Execute()

Delete audience memberships



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceMembershipsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceV2

> DeleteAudienceV2(ctx, audienceId).Execute()

Delete audience



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCouponReservation

> DeleteCouponReservation(ctx, couponValue).Body(body).Execute()

Delete coupon reservations



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The value of a coupon | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerData

> DeleteCustomerData(ctx, integrationId).Execute()

Delete the personal data of a customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).Execute()

List data associated with a specific customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile, must be unique within the account.  To get the &#x60;integrationId&#x60; of the profile from a &#x60;sessionId&#x60;, use the [Update customer session](/integration-api/#operation/updateCustomerSessionV2).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **bool** | Set to &#x60;true&#x60; to include customer profile information in the response. | 
 **referrals** | **bool** | Set to &#x60;true&#x60; to include referral information in the response. | 
 **coupons** | **bool** | Set to &#x60;true&#x60; to include coupon information in the response. | 
 **loyalty** | **bool** | Set to &#x60;true&#x60; to include loyalty information in the response. | 
 **giveaways** | **bool** | Set to &#x60;true&#x60; to include giveaways information in the response. | 

### Return type

[**CustomerInventory**](CustomerInventory.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCustomers

> InlineResponse200 GetReservedCustomers(ctx, couponValue).Execute()

List users that have this coupon reserved



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The value of a coupon | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservedCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEvent

> IntegrationState TrackEvent(ctx).Body(body).Dry(dry).Execute()

Track an Event



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewEvent**](NewEvent.md) |  | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. | 

### Return type

[**IntegrationState**](IntegrationState.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceCustomersAttributes

> UpdateAudienceCustomersAttributes(ctx, audienceId).Body(body).Execute()

Update profile attributes for all customers in audience



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceCustomersAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceV2

> Audience UpdateAudienceV2(ctx, audienceId).Body(body).Execute()

Update audience



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. You get it via the &#x60;id&#x60; property when [creating an audience](#operation/createAudienceV2). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAudience**](UpdateAudience.md) |  | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileAudiences

> UpdateCustomerProfileAudiences(ctx).Body(body).Execute()

Update multiple customer profiles' audiences



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerProfileAudienceRequest**](CustomerProfileAudienceRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileV2

> IntegrationStateV2 UpdateCustomerProfileV2(ctx, integrationId).Body(body).RunRuleEngine(runRuleEngine).Dry(dry).Execute()

Update customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile. Must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomerProfileIntegrationRequestV2**](CustomerProfileIntegrationRequestV2.md) |  | 
 **runRuleEngine** | **bool** | Indicates whether to run the rule engine. Setting this property to &#x60;false&#x60; improves response times. | [default to false]
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. Only used when &#x60;runRuleEngine&#x60; is set to &#x60;true&#x60;.  | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfilesV2

> MultipleCustomerProfileIntegrationResponseV2 UpdateCustomerProfilesV2(ctx).Body(body).Silent(silent).Execute()

Update multiple customer profiles



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfilesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MultipleCustomerProfileIntegrationRequest**](MultipleCustomerProfileIntegrationRequest.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to yes]

### Return type

[**MultipleCustomerProfileIntegrationResponseV2**](MultipleCustomerProfileIntegrationResponseV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSessionV2

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId).Body(body).Dry(dry).Execute()

Update customer session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The custom identifier for this session, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSessionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IntegrationRequest**](IntegrationRequest.md) |  | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;. | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

