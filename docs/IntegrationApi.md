# \IntegrationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCouponReservation**](IntegrationApi.md#CreateCouponReservation) | **Post** /v1/coupon_reservations/{couponValue} | Create a new coupon reservation
[**CreateReferral**](IntegrationApi.md#CreateReferral) | **Post** /v1/referrals | Create a referral code for an advocate
[**DeleteCouponReservation**](IntegrationApi.md#DeleteCouponReservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
[**DeleteCustomerData**](IntegrationApi.md#DeleteCustomerData) | **Delete** /v1/customer_data/{integrationId} | Delete the personal data of a customer.
[**GetCustomerInventory**](IntegrationApi.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | Get an inventory of all data associated with a specific customer profile.
[**GetReservedCoupons**](IntegrationApi.md#GetReservedCoupons) | **Get** /v1/coupon_reservations/coupons/{integrationId} | Get all valid reserved coupons
[**GetReservedCustomers**](IntegrationApi.md#GetReservedCustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | Get the users that have this coupon reserved
[**TrackEvent**](IntegrationApi.md#TrackEvent) | **Post** /v1/events | Track an Event
[**UpdateCustomerProfile**](IntegrationApi.md#UpdateCustomerProfile) | **Put** /v1/customer_profiles/{integrationId} | Update a Customer Profile
[**UpdateCustomerSession**](IntegrationApi.md#UpdateCustomerSession) | **Put** /v1/customer_sessions/{customerSessionId} | Update a Customer Session
[**UpdateCustomerSessionV2**](IntegrationApi.md#UpdateCustomerSessionV2) | **Put** /v2/customer_sessions/{customerSessionId} | Update a Customer Session



## CreateCouponReservation

> Coupon CreateCouponReservation(ctx, couponValue).Body(body).Execute()

Create a new coupon reservation



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

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferral

> Referral CreateReferral(ctx).Body(body).Execute()

Create a referral code for an advocate



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewReferral**](NewReferral.md) |  | 

### Return type

[**Referral**](Referral.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerData

> DeleteCustomerData(ctx, integrationId).Execute()

Delete the personal data of a customer.



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

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId).Profile(profile).Referrals(referrals).Execute()

Get an inventory of all data associated with a specific customer profile.



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **bool** | optional flag to decide if you would like customer profile information in the response | 
 **referrals** | **bool** | optional flag to decide if you would like referral information in the response | 

### Return type

[**CustomerInventory**](CustomerInventory.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCoupons

> InlineResponse2001 GetReservedCoupons(ctx, integrationId).Execute()

Get all valid reserved coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservedCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCustomers

> InlineResponse200 GetReservedCustomers(ctx, couponValue).Execute()

Get the users that have this coupon reserved



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

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEvent

> IntegrationState TrackEvent(ctx).Body(body).Execute()

Track an Event



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewEvent**](NewEvent.md) |  | 

### Return type

[**IntegrationState**](IntegrationState.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfile

> IntegrationState UpdateCustomerProfile(ctx, integrationId).Body(body).Execute()

Update a Customer Profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The custom identifier for this profile, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewCustomerProfile**](NewCustomerProfile.md) |  | 

### Return type

[**IntegrationState**](IntegrationState.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSession

> IntegrationState UpdateCustomerSession(ctx, customerSessionId).Body(body).Execute()

Update a Customer Session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The custom identifier for this session, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewCustomerSession**](NewCustomerSession.md) |  | 

### Return type

[**IntegrationState**](IntegrationState.md)

### Authorization

[api_key_v1](../README.md#api_key_v1), [integration_auth](../README.md#integration_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSessionV2

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId).Body(body).Execute()

Update a Customer Session



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

