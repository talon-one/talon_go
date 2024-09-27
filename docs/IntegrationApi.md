# \IntegrationApi

All URIs are relative to *https://yourbaseurl.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudienceV2**](IntegrationApi.md#CreateAudienceV2) | **Post** /v2/audiences | Create audience
[**CreateCouponReservation**](IntegrationApi.md#CreateCouponReservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
[**CreateReferral**](IntegrationApi.md#CreateReferral) | **Post** /v1/referrals | Create referral code for an advocate
[**CreateReferralsForMultipleAdvocates**](IntegrationApi.md#CreateReferralsForMultipleAdvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates
[**DeleteAudienceMembershipsV2**](IntegrationApi.md#DeleteAudienceMembershipsV2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
[**DeleteAudienceV2**](IntegrationApi.md#DeleteAudienceV2) | **Delete** /v2/audiences/{audienceId} | Delete audience
[**DeleteCouponReservation**](IntegrationApi.md#DeleteCouponReservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
[**DeleteCustomerData**](IntegrationApi.md#DeleteCustomerData) | **Delete** /v1/customer_data/{integrationId} | Delete customer&#39;s personal data
[**GenerateLoyaltyCard**](IntegrationApi.md#GenerateLoyaltyCard) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards | Generate loyalty card
[**GetCustomerInventory**](IntegrationApi.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
[**GetCustomerSession**](IntegrationApi.md#GetCustomerSession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
[**GetLoyaltyBalances**](IntegrationApi.md#GetLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/balances | Get customer&#39;s loyalty points
[**GetLoyaltyCardBalances**](IntegrationApi.md#GetLoyaltyCardBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/balances | Get card&#39;s point balances
[**GetLoyaltyCardPoints**](IntegrationApi.md#GetLoyaltyCardPoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/points | List card&#39;s unused loyalty points
[**GetLoyaltyCardTransactions**](IntegrationApi.md#GetLoyaltyCardTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transactions | List card&#39;s transactions
[**GetLoyaltyProgramProfilePoints**](IntegrationApi.md#GetLoyaltyProgramProfilePoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/points | List customer&#39;s unused loyalty points
[**GetLoyaltyProgramProfileTransactions**](IntegrationApi.md#GetLoyaltyProgramProfileTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/transactions | List customer&#39;s loyalty transactions
[**GetReservedCustomers**](IntegrationApi.md#GetReservedCustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List customers that have this coupon reserved
[**LinkLoyaltyCardToProfile**](IntegrationApi.md#LinkLoyaltyCardToProfile) | **Post** /v2/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/link_profile | Link customer profile to card
[**ReopenCustomerSession**](IntegrationApi.md#ReopenCustomerSession) | **Put** /v2/customer_sessions/{customerSessionId}/reopen | Reopen customer session
[**ReturnCartItems**](IntegrationApi.md#ReturnCartItems) | **Post** /v2/customer_sessions/{customerSessionId}/returns | Return cart items
[**SyncCatalog**](IntegrationApi.md#SyncCatalog) | **Put** /v1/catalogs/{catalogId}/sync | Sync cart item catalog
[**TrackEventV2**](IntegrationApi.md#TrackEventV2) | **Post** /v2/events | Track event
[**UpdateAudienceCustomersAttributes**](IntegrationApi.md#UpdateAudienceCustomersAttributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
[**UpdateAudienceV2**](IntegrationApi.md#UpdateAudienceV2) | **Put** /v2/audiences/{audienceId} | Update audience name
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
 **body** | [**NewAudience**](NewAudience.md) | body | 

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
**couponValue** | **string** | The code of the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) | body | 

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
 **body** | [**NewReferral**](NewReferral.md) | body | 

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
 **body** | [**NewReferralsForMultipleAdvocates**](NewReferralsForMultipleAdvocates.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

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
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceMembershipsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
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
**couponValue** | **string** | The code of the coupon. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CouponReservations**](CouponReservations.md) | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerData

> DeleteCustomerData(ctx, integrationId).Execute()

Delete customer's personal data



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateLoyaltyCard

> LoyaltyCard GenerateLoyaltyCard(ctx, loyaltyProgramId).Body(body).Execute()

Generate loyalty card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GenerateLoyaltyCard**](GenerateLoyaltyCard.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).Achievements(achievements).Execute()

List customer data



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **bool** | Set to &#x60;true&#x60; to include customer profile information in the response. | 
 **referrals** | **bool** | Set to &#x60;true&#x60; to include referral information in the response. | 
 **coupons** | **bool** | Set to &#x60;true&#x60; to include coupon information in the response. | 
 **loyalty** | **bool** | Set to &#x60;true&#x60; to include loyalty information in the response. | 
 **giveaways** | **bool** | Set to &#x60;true&#x60; to include giveaways information in the response. | 
 **achievements** | **bool** | Set to &#x60;true&#x60; to include achievement information in the response. | 

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


## GetCustomerSession

> IntegrationCustomerSessionResponse GetCustomerSession(ctx, customerSessionId).Execute()

Get customer session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationCustomerSessionResponse**](IntegrationCustomerSessionResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyBalances

> LoyaltyBalancesWithTiers GetLoyaltyBalances(ctx, loyaltyProgramId, integrationId).EndDate(endDate).SubledgerId(subledgerId).IncludeTiers(includeTiers).IncludeProjectedTier(includeProjectedTier).Execute()

Get customer's loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **includeTiers** | **bool** | Indicates whether tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the current tier and the number of points required to move to next tier.  | [default to false]
 **includeProjectedTier** | **bool** | Indicates whether the customer&#39;s projected tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the customer’s active points and the name of the projected tier.  **Note** We recommend filtering by &#x60;subledgerId&#x60; for better performance.  | [default to false]

### Return type

[**LoyaltyBalancesWithTiers**](LoyaltyBalancesWithTiers.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardBalances

> LoyaltyCardBalances GetLoyaltyCardBalances(ctx, loyaltyProgramId, loyaltyCardId).EndDate(endDate).SubledgerId(subledgerId).Execute()

Get card's point balances



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | [**[]string**](string.md) | Filter results by one or more subledger IDs. Must be exact match. | 

### Return type

[**LoyaltyCardBalances**](LoyaltyCardBalances.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardPoints

> InlineResponse2003 GetLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).Status(status).SubledgerId(subledgerId).PageSize(pageSize).Skip(skip).Execute()

List card's unused loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter points based on their status. | [default to active]
 **subledgerId** | [**[]string**](string.md) | Filter results by one or more subledger IDs. Must be exact match. | 
 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactions

> InlineResponse2001 GetLoyaltyCardTransactions(ctx, loyaltyProgramId, loyaltyCardId).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).Execute()

List card's transactions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subledgerId** | [**[]string**](string.md) | Filter results by one or more subledger IDs. Must be exact match. | 
 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfilePoints

> InlineResponse2004 GetLoyaltyProgramProfilePoints(ctx, loyaltyProgramId, integrationId).Status(status).SubledgerId(subledgerId).PageSize(pageSize).Skip(skip).Execute()

List customer's unused loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramProfilePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter points based on their status. | [default to active]
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfileTransactions

> InlineResponse2002 GetLoyaltyProgramProfileTransactions(ctx, loyaltyProgramId, integrationId).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).Execute()

List customer's loyalty transactions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramProfileTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

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

List customers that have this coupon reserved



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon. | 

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


## LinkLoyaltyCardToProfile

> LoyaltyCard LinkLoyaltyCardToProfile(ctx, loyaltyProgramId, loyaltyCardId).Body(body).Execute()

Link customer profile to card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkLoyaltyCardToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**LoyaltyCardRegistration**](LoyaltyCardRegistration.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReopenCustomerSession

> ReopenSessionResponse ReopenCustomerSession(ctx, customerSessionId).Execute()

Reopen customer session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReopenCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReopenSessionResponse**](ReopenSessionResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCartItems

> IntegrationStateV2 ReturnCartItems(ctx, customerSessionId).Body(body).Dry(dry).Execute()

Return cart items



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnCartItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReturnIntegrationRequest**](ReturnIntegrationRequest.md) | body | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 

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


## SyncCatalog

> Catalog SyncCatalog(ctx, catalogId).Body(body).Execute()

Sync cart item catalog



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int32** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CatalogSyncRequest**](CatalogSyncRequest.md) | body | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEventV2

> TrackEventV2Response TrackEventV2(ctx).Body(body).Silent(silent).Dry(dry).Execute()

Track event



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IntegrationEventV2Request**](IntegrationEventV2Request.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 

### Return type

[**TrackEventV2Response**](TrackEventV2Response.md)

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
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceCustomersAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceV2

> Audience UpdateAudienceV2(ctx, audienceId).Body(body).Execute()

Update audience name



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAudience**](UpdateAudience.md) | body | 

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
 **body** | [**CustomerProfileAudienceRequest**](CustomerProfileAudienceRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileV2

> CustomerProfileIntegrationResponseV2 UpdateCustomerProfileV2(ctx, integrationId).Body(body).RunRuleEngine(runRuleEngine).Dry(dry).Execute()

Update customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomerProfileIntegrationRequestV2**](CustomerProfileIntegrationRequestV2.md) | body | 
 **runRuleEngine** | **bool** | Indicates whether to run the Rule Engine.  If &#x60;true&#x60;, the response includes: - The effects generated by the triggered campaigns are returned in the &#x60;effects&#x60; property. - The created coupons and referral objects.  If &#x60;false&#x60;: - The rules are not executed and the &#x60;effects&#x60; property is always empty. - The response time improves. - You cannot use &#x60;responseContent&#x60; in the body.  | [default to false]
 **dry** | **bool** | (Only works when &#x60;runRuleEngine&#x3D;true&#x60;) Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;, you can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  | 

### Return type

[**CustomerProfileIntegrationResponseV2**](CustomerProfileIntegrationResponseV2.md)

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
 **body** | [**MultipleCustomerProfileIntegrationRequest**](MultipleCustomerProfileIntegrationRequest.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

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

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId).Body(body).Dry(dry).Now(now).Execute()

Update customer session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSessionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IntegrationRequest**](IntegrationRequest.md) | body | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;: - The endpoint considers **only** the payload that you pass when **closing** the session.   When you do not use the &#x60;dry&#x60; parameter, the endpoint behaves as a typical PUT endpoint. Each update builds upon the previous ones. - You can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  [See the docs](https://docs.talon.one/docs/dev/integration-api/dry-requests).  | 
 **now** | **time.Time** | A timestamp value of a future date that acts as a current date when included in the query.  Use this parameter, for example, to test campaigns that would be evaluated for this customer session in the future (say, [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule)).  **Note:**  - It must be an RFC3339 timestamp string. - It can **only** be a date in the future. - It can **only** be used if the &#x60;dry&#x60; parameter in the query is set to &#x60;true&#x60;.  | 

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

