# \IntegrationApi

All URLs are relative to `https://yourbaseurl.talon.one`.

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
[**GetCustomerAchievementHistory**](IntegrationApi.md#GetCustomerAchievementHistory) | **Get** /v1/customer_profiles/{integrationId}/achievements/{achievementId} | List customer&#39;s achievement history
[**GetCustomerAchievements**](IntegrationApi.md#GetCustomerAchievements) | **Get** /v1/customer_profiles/{integrationId}/achievements | List customer&#39;s available achievements
[**GetCustomerInventory**](IntegrationApi.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
[**GetCustomerSession**](IntegrationApi.md#GetCustomerSession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
[**GetLoyaltyBalances**](IntegrationApi.md#GetLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/balances | Get customer&#39;s loyalty balances
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

> Audience CreateAudienceV2(ctx, body)

Create audience

Create an audience. The audience can be created directly from scratch or can come from third party platforms.  **Note:** Audiences can also be created from scratch via the Campaign Manager. See the [docs](https://docs.talon.one/docs/product/audiences/creating-audiences).  To create an audience from an existing audience from a [technology partner](https://docs.talon.one/docs/dev/technology-partners/overview): 1. Set the `integration` property to `mparticle`, `segment` etc., depending on a third-party platform. 1. Set `integrationId` to the ID of this audience in a third-party platform.  To create an audience from an existing audience in another platform: 1. Do not use the `integration` property. 1. Set `integrationId` to the ID of this audience in the 3rd-party platform.  To create an audience from scratch: 1. Only set the `name` property.  Once you create your first audience, audience-specific rule conditions are enabled in the Rule Builder. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewAudience**](NewAudience.md)| body | 

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

> Coupon CreateCouponReservation(ctx, couponValue, body)

Create coupon reservation

Create a coupon reservation for the specified customer profiles on the specified coupon. You can also create a reservation via the Campaign Manager using the [Create coupon code reservation](https://docs.talon.one/docs/product/rules/effects/using-effects#reserving-a-coupon-code) effect.  **Note:**  - If the **Reservation mandatory** option was selected when creating the specified coupon, the endpoint creates a **hard** reservation, meaning only users who have this coupon code reserved can redeem it. Otherwise, the endpoint creates a **soft** reservation, meaning the coupon is associated with the specified customer profiles (they show up when using the [List customer data](https://docs.talon.one/integration-api#operation/getCustomerInventory) endpoint), but any user can redeem it. This can be useful, for example, to display a _coupon wallet_ for customers when they visit your store.  - If the **Coupon visibility** option was selected when creating the specified coupon, the coupon code is implicitly soft-reserved for all customers, and the code will be returned for all customer profiles in the [List customer data](https://docs.talon.one/integration-api#operation/getCustomerInventory) endpoint.  - This endpoint overrides the coupon reservation limit set when [the coupon is created](https://docs.talon.one/docs/product/campaigns/coupons/creating-coupons). To ensure that coupons cannot be reserved after the reservation limit is reached, use the [Create coupon code reservation](https://docs.talon.one/docs/product/rules/effects/using-effects#reserving-a-coupon-code) effect in the Rule Builder and the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint.  To delete a reservation, use the [Delete reservation](https://docs.talon.one/integration-api#tag/Coupons/operation/deleteCouponReservation) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string**| The code of the coupon. | 
**body** | [**CouponReservations**](CouponReservations.md)| body | 

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

> Referral CreateReferral(ctx, body)

Create referral code for an advocate

Creates a referral code for an advocate. The code will be valid for the referral campaign for which is created, indicated in the `campaignId` parameter, and will be associated with the profile specified in the `advocateProfileIntegrationId` parameter as the advocate's profile.  **Note:** Any [referral limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets#referral-limits) set are ignored when you use this endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewReferral**](NewReferral.md)| body | 

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

> InlineResponse201 CreateReferralsForMultipleAdvocates(ctx, body, optional)

Create referral codes for multiple advocates

Creates unique referral codes for multiple advocates. The code will be valid for the referral campaign for which it is created, indicated in the `campaignId` parameter, and one referral code will be associated with one advocate using the profile specified in the `advocateProfileIntegrationId` parameter as the advocate's profile.  **Note:** Any [referral limits](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-budgets#referral-limits) set are ignored when you use this endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**NewReferralsForMultipleAdvocates**](NewReferralsForMultipleAdvocates.md)| body | 
 **optional** | ***CreateReferralsForMultipleAdvocatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateReferralsForMultipleAdvocatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **silent** | **optional.**| Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceMembershipsV2

> DeleteAudienceMembershipsV2(ctx, audienceId)

Delete audience memberships

Remove all members from this audience. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64**| The ID of the audience. | 

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

> DeleteAudienceV2(ctx, audienceId)

Delete audience

Delete an audience created by a third-party integration.  **Warning:** This endpoint also removes any associations recorded between a customer profile and this audience.  **Note:** Audiences can also be deleted via the Campaign Manager. See the [docs](https://docs.talon.one/docs/product/audiences/managing-audiences#deleting-an-audience). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64**| The ID of the audience. | 

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

> DeleteCouponReservation(ctx, couponValue, body)

Delete coupon reservations

Remove all the coupon reservations from the provided customer profile integration IDs and the provided coupon code. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string**| The code of the coupon. | 
**body** | [**CouponReservations**](CouponReservations.md)| body | 

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

> DeleteCustomerData(ctx, integrationId)

Delete customer's personal data

Delete all attributes on the customer profile and on entities that reference this customer profile.  **Important:** - Customer data is deleted from all Applications in the [environment](https://docs.talon.one/docs/product/applications/overview#application-environments)   that the API key belongs to. For example, if you use this endpoint with an API key that belongs to a sandbox Application,   customer data will be deleted from all sandbox Applications. This is because customer data is shared   between Applications from the same environment. - To preserve performance, we recommend avoiding deleting customer data during peak-traffic hours. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string**| The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

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

> LoyaltyCard GenerateLoyaltyCard(ctx, loyaltyProgramId, body)

Generate loyalty card

Generate a loyalty card in a specified [card-based loyalty program](https://docs.talon.one/docs/product/loyalty-programs/card-based/card-based-overview).  To link the card to one or more customer profiles, use the `customerProfileIds` parameter in the request body.  **Note:** - The number of customer profiles linked to the loyalty card cannot exceed the loyalty program's `usersPerCardLimit`. To find the program's limit, use the [Get loyalty program](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyProgram) endpoint. - If the loyalty program has a defined code format, it will be used for the loyalty card identifier. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**body** | [**GenerateLoyaltyCard**](GenerateLoyaltyCard.md)| body | 

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


## GetCustomerAchievementHistory

> InlineResponse2002 GetCustomerAchievementHistory(ctx, integrationId, achievementId, optional)

List customer's achievement history

Retrieve all progress history of a given customer in the given achievement. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
**achievementId** | **int64**| The achievement identifier.  | 
 **optional** | ***GetCustomerAchievementHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerAchievementHistoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **progressStatus** | [**optional.Interface of []string**](string.md)| Filter by customer progress status in the achievement.  | 
 **startDate** | **optional.**| Timestamp that filters the results to only contain achievements created on or after the start date. | 
 **endDate** | **optional.**| Timestamp that filters the results to only contain achievements created before or on the end date. | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAchievements

> InlineResponse2001 GetCustomerAchievements(ctx, integrationId, optional)

List customer's available achievements

Retrieve all the achievements available to a given customer and their progress in them. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
 **optional** | ***GetCustomerAchievementsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerAchievementsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignIds** | [**optional.Interface of []string**](string.md)| Filter by one or more Campaign IDs, separated by a comma.  **Note:** If no campaigns are specified, data for all the campaigns in the Application is returned.  | 
 **achievementIds** | [**optional.Interface of []string**](string.md)| Filter by one or more Achievement IDs, separated by a comma.  **Note:** If no achievements are specified, data for all the achievements in the Application is returned.  | 
 **achievementStatus** | [**optional.Interface of []string**](string.md)| Filter by status of the achievement.  **Note:** If the achievement status is not specified, only data for all active achievements in the Application is returned.  | 
 **currentProgressStatus** | [**optional.Interface of []string**](string.md)| Filter by customer progress status in the achievement.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 1000]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId, optional)

List customer data

Return the customer inventory regarding entities referencing this customer profile's `integrationId`.  Typical entities returned are: customer profile information, referral codes, loyalty points, loyalty cards and reserved coupons. Reserved coupons also include redeemed coupons. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string**| The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 
 **optional** | ***GetCustomerInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCustomerInventoryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **optional.**| Set to &#x60;true&#x60; to include customer profile information in the response. | 
 **referrals** | **optional.**| Set to &#x60;true&#x60; to include referral information in the response. | 
 **coupons** | **optional.**| Set to &#x60;true&#x60; to include coupon information in the response. | 
 **loyalty** | **optional.**| Set to &#x60;true&#x60; to include loyalty information in the response. | 
 **giveaways** | **optional.**| Set to &#x60;true&#x60; to include giveaways information in the response. | 
 **achievements** | **optional.**| Set to &#x60;true&#x60; to include achievement information in the response. | 

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

> IntegrationCustomerSessionResponse GetCustomerSession(ctx, customerSessionId)

Get customer session

Get the details of the given customer session.  You can get the same data via other endpoints that also apply changes, which can help you save requests and increase performance. See:  - [Update customer session](#tag/Customer-sessions/operation/updateCustomerSessionV2) - [Update customer profile](#tag/Customer-profiles/operation/updateCustomerProfileV2) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string**| The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

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

> LoyaltyBalancesWithTiers GetLoyaltyBalances(ctx, loyaltyProgramId, integrationId, optional)

Get customer's loyalty balances

Retrieve loyalty ledger balances for the given Integration ID in the specified loyalty program. You can filter balances by date and subledger ID, and include tier-related information in the response.  **Note**: If no filtering options are applied, you retrieve all loyalty balances on the current date for the given integration ID.  Loyalty balances are calculated when Talon.One receives your request using the points stored in our database, so retrieving a large number of balances at once can impact performance.  For more information, see: - [Managing card-based loyalty program data](https://docs.talon.one/docs/product/loyalty-programs/card-based/managing-loyalty-cards) - [Managing profile-based loyalty program data](https://docs.talon.one/docs/product/loyalty-programs/profile-based/managing-pb-lp-data) 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
 **optional** | ***GetLoyaltyBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyBalancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **optional.**| Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **optional.**| The ID of the subledger by which we filter the data. | 
 **includeTiers** | **optional.**| Indicates whether tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the current tier and the number of points required to move to next tier.  | [default to false]
 **includeProjectedTier** | **optional.**| Indicates whether the customer&#39;s projected tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the customer&#39;s active points and the name of the projected tier.  **Note** We recommend filtering by &#x60;subledgerId&#x60; for better performance.  | [default to false]

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

> LoyaltyCardBalances GetLoyaltyCardBalances(ctx, loyaltyProgramId, loyaltyCardId, optional)

Get card's point balances

Retrieve loyalty balances for the given loyalty card in the specified loyalty program with filtering options applied. If no filtering options are applied, all loyalty balances for the given loyalty card are returned. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
 **optional** | ***GetLoyaltyCardBalancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyCardBalancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **optional.**| Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | [**optional.Interface of []string**](string.md)| Filter results by one or more subledger IDs. Must be exact match. | 

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

> InlineResponse2005 GetLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId, optional)

List card's unused loyalty points

Get paginated results of loyalty points for a given loyalty card identifier in a card-based loyalty program. This endpoint returns only the balances of unused points on a loyalty card.  You can filter points by status: - `active`: Points ready to be redeemed. - `pending`: Points with a start date in the future. - `expired`: Points with an expiration date in the past. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
 **optional** | ***GetLoyaltyCardPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyCardPointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.**| Filter points based on their status. | [default to active]
 **subledgerId** | [**optional.Interface of []string**](string.md)| Filter results by one or more subledger IDs. Must be exact match. | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactions

> InlineResponse2003 GetLoyaltyCardTransactions(ctx, loyaltyProgramId, loyaltyCardId, optional)

List card's transactions

Retrieve loyalty transaction logs for the given loyalty card in the specified loyalty program with filtering options applied. If no filtering options are applied, the last 50 loyalty transactions for the given loyalty card are returned. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
 **optional** | ***GetLoyaltyCardTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyCardTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subledgerId** | [**optional.Interface of []string**](string.md)| Filter results by one or more subledger IDs. Must be exact match. | 
 **loyaltyTransactionType** | **optional.**| Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **optional.**| Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **optional.**| Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfilePoints

> InlineResponse2006 GetLoyaltyProgramProfilePoints(ctx, loyaltyProgramId, integrationId, optional)

List customer's unused loyalty points

Get paginated results of loyalty points for a given Integration ID in the specified profile-based loyalty program. This endpoint returns only the balances of unused points linked to a customer profile.  You can filter points by status: - `active`: Points ready to be redeemed. - `pending`: Points with a start date in the future. - `expired`: Points with an expiration date in the past. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
 **optional** | ***GetLoyaltyProgramProfilePointsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyProgramProfilePointsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.**| Filter points based on their status. | [default to active]
 **subledgerId** | **optional.**| The ID of the subledger by which we filter the data. | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfileTransactions

> InlineResponse2004 GetLoyaltyProgramProfileTransactions(ctx, loyaltyProgramId, integrationId, optional)

List customer's loyalty transactions

Retrieve paginated results of loyalty transaction logs for the given Integration ID in the specified loyalty program.  You can filter transactions by date. If no filters are applied, the last 50 loyalty transactions for the given integration ID are returned.  **Note:** To retrieve all loyalty program transaction logs in a given loyalty program, use the [List loyalty program transactions](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyProgramTransactions) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
 **optional** | ***GetLoyaltyProgramProfileTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLoyaltyProgramProfileTransactionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subledgerId** | **optional.**| The ID of the subledger by which we filter the data. | 
 **loyaltyTransactionType** | **optional.**| Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **optional.**| Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **optional.**| Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **optional.**| The number of items in the response. | [default to 50]
 **skip** | **optional.**| The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCustomers

> InlineResponse200 GetReservedCustomers(ctx, couponValue)

List customers that have this coupon reserved

Return all customers that have this coupon marked as reserved. This includes hard and soft reservations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string**| The code of the coupon. | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkLoyaltyCardToProfile

> LoyaltyCard LinkLoyaltyCardToProfile(ctx, loyaltyProgramId, loyaltyCardId, body)

Link customer profile to card

[Loyalty cards](https://docs.talon.one/docs/product/loyalty-programs/card-based/card-based-overview) allow customers to collect and spend loyalty points within a [card-based loyalty program](https://docs.talon.one/docs/product/loyalty-programs/overview#loyalty-program-types). They are useful to gamify loyalty programs and can be used with or without customer profiles linked to them.  Link a customer profile to a given loyalty card for the card to be set as **Registered**. This affects how it can be used. See the [docs](https://docs.talon.one/docs/product/loyalty-programs/card-based/managing-loyalty-cards#linking-customer-profiles-to-a-loyalty-card).  **Note:** You can link as many customer profiles to a given loyalty card as the [**card user limit**](https://docs.talon.one/docs/product/loyalty-programs/card-based/creating-cb-programs) allows. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64**| Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string**| Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 
**body** | [**LoyaltyCardRegistration**](LoyaltyCardRegistration.md)| body | 

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

> ReopenSessionResponse ReopenCustomerSession(ctx, customerSessionId)

Reopen customer session

Reopen a closed [customer session](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions). For example, if a session has been completed but still needs to be edited, you can reopen it with this endpoint. A reopen session is treated like a standard open session.  When reopening a session: - The `talon_session_reopened` event is triggered. You can see it in the **Events** view in the Campaign Manager. - The session state is updated to `open`. - Modified budgets and triggered effects when the session was closed are rolled back except for the list below.  <details>   <summary><strong>Effects and budgets unimpacted by a session reopening</strong></summary>   <div>     <p>The following effects and budgets are left the way they were once the session was originally closed:</p>     <ul>       <li>Add free item effect</li>       <li>Any <strong>non-pending</strong> loyalty points</li>       <li>Award giveaway</li>       <li>Coupon and referral creation</li>       <li>Coupon reservation</li>       <li>Custom effect</li>       <li>Update attribute value</li>       <li>Update cart item attribute value</li>     </ul>   </div> <p>To see an example of roll back, see the <a href=\"https://docs.talon.one/docs/dev/tutorials/rolling-back-effects\">Cancelling a session with campaign budgets tutorial</a>.</p> </details>  **Note:** If your order workflow requires you to create a new session instead of reopening a session, use the [Update customer session](https://docs.talon.one/integration-api#tag/Customer-sessions/operation/updateCustomerSessionV2) endpoint to cancel a closed session and create a new one. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string**| The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

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

> IntegrationStateV2 ReturnCartItems(ctx, customerSessionId, body, optional)

Return cart items

Create a new return request for the specified cart items.  This endpoint automatically changes the session state from `closed` to `partially_returned`.  **Note:** This will roll back any effects associated with these cart items. For more information, see [our documentation on session states](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions#customer-session-states) and [this tutorial](https://docs.talon.one/docs/dev/tutorials/partially-returning-a-session). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string**| The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 
**body** | [**ReturnIntegrationRequest**](ReturnIntegrationRequest.md)| body | 
 **optional** | ***ReturnCartItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReturnCartItemsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dry** | **optional.**| Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 

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

> Catalog SyncCatalog(ctx, catalogId, body)

Sync cart item catalog

Perform the following actions for a given cart item catalog:  - Add an item to the catalog. - Add multiple items to the catalog. - Update the attributes of an item in the catalog. - Update the attributes of multiple items in the catalog. - Remove an item from the catalog. - Remove multiple items from the catalog.  You can either add, update, or delete up to 1000 cart items in a single request. Each item synced to a catalog must have a unique `SKU`.  **Important**: You can perform only one type of action in a single sync request. Syncing items with duplicate `SKU` values in a single request returns an error message with a `400` status code.  For more information, read [managing cart item catalogs](https://docs.talon.one/docs/product/account/dev-tools/managing-cart-item-catalogs).  ### Filtering cart items  Use [cart item attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes) to filter items and select the ones you want to edit or delete when editing or deleting more than one item at a time.  The `filters` array contains an object with the following properties:  - `attr`: A [cart item attribute](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes)   connected to the catalog. It is applied to all items in the catalog. - `op`: The filtering operator indicating the relationship between the value of each   cart item in the catalog and the value of the `value` property for the attribute selected   in `attr`.    The value of `op` can be one of the following:    - `EQ`: Equal to `value`   - `LT`: Less than `value`   - `LE`: Less than or equal to `value`   - `GT`: Greater than `value`   - `GE`: Greater than or equal to `value`   - `IN`: One of the comma-separated values that `value` is set to.    **Note:** `GE`, `LE`, `GT`, `LT` are for numeric values only.  - `value`: The value of the attribute selected in `attr`.  ### Payload examples  Synchronization actions are sent as `PUT` requests. See the structure for each action:  <details>   <summary><strong>Adding an item to the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"attributes\": {             \"color\": \"Navy blue\",             \"type\": \"shoes\"           },           \"replaceIfExists\": true,           \"sku\": \"SKU1241028\",           \"price\": 100,           \"product\": {             \"name\": \"sneakers\"           }         },         \"type\": \"ADD\"       }     ]   }   ```   </div> </details>  <details>   <summary><strong>Adding multiple items to the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"attributes\": {             \"color\": \"Navy blue\",             \"type\": \"shoes\"           },           \"replaceIfExists\": true,           \"sku\": \"SKU1241027\",           \"price\": 100,           \"product\": {             \"name\": \"sneakers\"           }         },         \"type\": \"ADD\"       },       {         \"payload\": {           \"attributes\": {             \"color\": \"Navy blue\",             \"type\": \"shoes\"           },           \"replaceIfExists\": true,           \"sku\": \"SKU1241028\",           \"price\": 100,           \"product\": {             \"name\": \"sneakers\"           }         },         \"type\": \"ADD\"       }     ]   }   ```   </div> </details>  <details>   <summary><strong>Updating the attributes of an item in the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"attributes\": {             \"age\": 11,             \"origin\": \"germany\"           },           \"createIfNotExists\": false,           \"sku\": \"SKU1241028\",           \"product\": {             \"name\": \"sneakers\"           }         },         \"type\": \"PATCH\"       }     ]   }   ```   </div> </details>  <details>   <summary><strong>Updating the attributes of multiple items in the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"attributes\": {             \"color\": \"red\"           },           \"filters\": [             {               \"attr\": \"color\",               \"op\": \"EQ\",               \"value\": \"blue\"             }           ]         },         \"type\": \"PATCH_MANY\"       }     ]   }   ```    </div> </details>  <details>   <summary><strong>Removing an item from the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"sku\": \"SKU1241028\"         },         \"type\": \"REMOVE\"       }     ]   }   ```    </div> </details>  <details>   <summary><strong>Removing multiple items from the catalog</strong></summary>   <div>    ```json   {     \"actions\": [       {         \"payload\": {           \"filters\": [             {               \"attr\": \"color\",               \"op\": \"EQ\",               \"value\": \"blue\"             }           ]         },         \"type\": \"REMOVE_MANY\"       }     ]   }   ```   </div> </details>  <details>   <summary><strong>Removing shoes of sizes above 45 from the catalog</strong></summary>   <div>   <p>   Let's imagine that we have a shoe store and we have decided to stop selling   shoes larger than size 45. We can remove from the catalog all the shoes of sizes above 45   with a single action:</p>    ```json   {     \"actions\": [       {         \"payload\": {           \"filters\": [             {               \"attr\": \"size\",               \"op\": \"GT\",               \"value\": \"45\"             }           ]         },         \"type\": \"REMOVE_MANY\"       }     ]   }   ```   </div> </details> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int64**| The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 
**body** | [**CatalogSyncRequest**](CatalogSyncRequest.md)| body | 

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

> TrackEventV2Response TrackEventV2(ctx, body, optional)

Track event

Triggers a custom event.  To use this endpoint: 1. Define a [custom event](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) in the Campaign Manager. 1. Update or create a rule to check for this event. 1. Trigger the event with this endpoint. After you have successfully sent an event to Talon.One, you can list the received events in the **Events** view in the Campaign Manager.  Talon.One also offers a set of [built-in events](https://docs.talon.one/docs/dev/concepts/entities/events). Ensure you do not create a custom event when you can use a built-in event.  For example, use this endpoint to trigger an event when a customer shares a link to a product. See the [tutorial](https://docs.talon.one/docs/product/tutorials/referrals/incentivizing-product-link-sharing).  <div class=\"redoc-section\">    <p class=\"title\">Important</p>    1. `profileId` is required even though the schema does not specify it.   1. If the customer profile ID is new, a new profile is automatically created but the `customer_profile_created` [built-in event ](https://docs.talon.one/docs/dev/concepts/entities/events) is **not** triggered.   1. We recommend sending requests sequentially. See [Managing parallel requests](https://docs.talon.one/docs/dev/getting-started/integration-tutorial#managing-parallel-requests).   1. [Archived campaigns](https://docs.talon.one/docs/product/campaigns/managing-campaigns#archiving-a-campaign) are not considered in rule evaluation.  </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**IntegrationEventV2Request**](IntegrationEventV2Request.md)| body | 
 **optional** | ***TrackEventV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TrackEventV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **silent** | **optional.**| Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]
 **dry** | **optional.**| Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 
 **forceCompleteEvaluation** | **optional.**| Forces evaluation for all matching campaigns regardless of the [campaign evaluation mode](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation#setting-campaign-evaluation-mode). Requires &#x60;dry&#x3D;true&#x60;.  | [default to false]

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

> UpdateAudienceCustomersAttributes(ctx, audienceId, body)

Update profile attributes for all customers in audience

Update the specified profile attributes to the provided values for all customers in the specified audience. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64**| The ID of the audience. | 
**body** | **map[string]interface{}**| body | 

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

> Audience UpdateAudienceV2(ctx, audienceId, body)

Update audience name

Update the name of the given audience created by a third-party integration. Sending a request to this endpoint does **not** trigger the Rule Engine.  To update the audience's members, use the [Update customer profile](#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64**| The ID of the audience. | 
**body** | [**UpdateAudience**](UpdateAudience.md)| body | 

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

> UpdateCustomerProfileAudiences(ctx, body)

Update multiple customer profiles' audiences

Add customer profiles to or remove them from an audience.  The endpoint supports 1000 audience actions (`add` or `remove`) per request.  **Note:** You can also do this using the [Update audience](https://docs.talon.one/docs/product/rules/effects/using-effects#updating-an-audience) effect. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**CustomerProfileAudienceRequest**](CustomerProfileAudienceRequest.md)| body | 

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

> CustomerProfileIntegrationResponseV2 UpdateCustomerProfileV2(ctx, integrationId, body, optional)

Update customer profile

Update or create a [Customer Profile](https://docs.talon.one/docs/dev/concepts/entities/customer-profiles). This endpoint triggers the Rule Builder.  You can use this endpoint to: - Set attributes on the given customer profile. Ensure you create the attributes in the Campaign Manager, first. - Modify the audience the customer profile is a member of. **Note:** [Archived campaigns](https://docs.talon.one/docs/product/campaigns/managing-campaigns#archiving-a-campaign) are not considered in rule evaluation when `runRuleEngine` is `true`. <div class=\"redoc-section\">   <p class=\"title\">Performance tips</p>    - Updating a customer profile returns a response with the requested integration state.   - You can use the `responseContent` property to save yourself extra API calls. For example, you can get     the customer profile details directly without extra requests.   - We recommend sending requests sequentially.     See [Managing parallel requests](https://docs.talon.one/docs/dev/getting-started/integration-tutorial#managing-parallel-requests). </div> 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string**| The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
**body** | [**CustomerProfileIntegrationRequestV2**](CustomerProfileIntegrationRequestV2.md)| body | 
 **optional** | ***UpdateCustomerProfileV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomerProfileV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **runRuleEngine** | **optional.**| Indicates whether to run the Rule Engine.  If &#x60;true&#x60;, the response includes: - The effects generated by the triggered campaigns are returned in the &#x60;effects&#x60; property. - The created coupons and referral objects.  If &#x60;false&#x60;: - The rules are not executed and the &#x60;effects&#x60; property is always empty. - The response time improves. - You cannot use &#x60;responseContent&#x60; in the body.  | [default to false]
 **dry** | **optional.**| (Only works when &#x60;runRuleEngine&#x3D;true&#x60;) Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;, you can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  | 

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

> MultipleCustomerProfileIntegrationResponseV2 UpdateCustomerProfilesV2(ctx, body, optional)

Update multiple customer profiles

Update (or create) up to 1000 [customer profiles](https://docs.talon.one/docs/dev/concepts/entities/customer-profiles) in 1 request.  The `integrationId` must be any identifier that remains stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  A customer profile [can be linked to one or more sessions](https://docs.talon.one/integration-api#tag/Customer-sessions).  **Note:** This endpoint does not trigger the Rule Engine. To trigger the Rule Engine for customer profile updates, use the [Update customer profile](#tag/Customer-profiles/operation/updateCustomerProfileV2) endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**MultipleCustomerProfileIntegrationRequest**](MultipleCustomerProfileIntegrationRequest.md)| body | 
 **optional** | ***UpdateCustomerProfilesV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomerProfilesV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **silent** | **optional.**| Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

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

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId, body, optional)

Update customer session

Update or create a [customer session](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions). The endpoint responds with the potential promotion rule [effects](https://docs.talon.one/docs/dev/integration-api/api-effects) that match the current cart. For example, use this endpoint to share the contents of a customer's cart with Talon.One.  **Note:**  - The currency for the session and the cart items in it is the currency set for the Application linked to this session. - [Archived campaigns](https://docs.talon.one/docs/product/campaigns/managing-campaigns#archiving-a-campaign) are not considered for rule evaluation.  ### Session management  To use this endpoint, start by learning about [customer sessions](https://docs.talon.one/docs/dev/concepts/entities/customer-sessions) and their states and refer to the `state` parameter documentation the request body schema docs below.  ### Sessions and customer profiles  - To link a session to a customer profile, set the `profileId` parameter in the request body to a customer profile's `integrationId`. - While you can create an anonymous session with `profileId=\"\"`, we recommend you use a guest ID instead. - A profile can be linked to simultaneous sessions in different Applications. Either:   - Use unique session integration IDs or,   - Use the same session integration ID across all of the Applications.  **Note:** If the specified profile does not exist, an empty profile is **created automatically**.   You can update it with [Update customer profile](https://docs.talon.one/integration-api#tag/Customer-profiles/operation/updateCustomerProfileV2).  <div class=\"redoc-section\">   <p class=\"title\">Performance tips</p>    - Updating a customer session returns a response with the new integration state. Use the `responseContent` property to save yourself extra API calls.     For example, you can get the customer profile details directly without extra requests.   - We recommend sending requests sequentially. See [Managing parallel requests](https://docs.talon.one/docs/dev/getting-started/integration-tutorial#managing-parallel-requests). </div>  For more information, see: - The introductory video in [Getting started](https://docs.talon.one/docs/dev/getting-started/overview). - The [integration tutorial](https://docs.talon.one/docs/dev/tutorials/integrating-talon-one). 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string**| The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 
**body** | [**IntegrationRequest**](IntegrationRequest.md)| body | 
 **optional** | ***UpdateCustomerSessionV2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCustomerSessionV2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dry** | **optional.**| Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;: - The endpoint considers **only** the payload that you pass when **closing** the session.   When you do not use the &#x60;dry&#x60; parameter, the endpoint behaves as a typical PUT endpoint. Each update builds upon the previous ones. - You can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  [See the docs](https://docs.talon.one/docs/dev/integration-api/dry-requests).  | 
 **now** | **optional.**| A timestamp value of a future date that acts as a current date when included in the query.  Use this parameter, for example, to test campaigns that would be evaluated for this customer session in the future (say, [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule)).  **Note:**  - It must be an RFC3339 timestamp string. - It can **only** be a date in the future. - It can **only** be used if the &#x60;dry&#x60; parameter in the query is set to &#x60;true&#x60;.  | 

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

