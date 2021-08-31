# \ManagementApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLoyaltyPoints**](ManagementApi.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/add_points | Add points in loyalty program for given customer
[**CopyCampaignToApplications**](ManagementApi.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into every specified application
[**CreateAdditionalCost**](ManagementApi.md#CreateAdditionalCost) | **Post** /v1/additional_costs | Define a new additional cost
[**CreateAttribute**](ManagementApi.md#CreateAttribute) | **Post** /v1/attributes | Create custom attribute
[**CreateCampaign**](ManagementApi.md#CreateCampaign) | **Post** /v1/applications/{applicationId}/campaigns | Create campaign
[**CreateCoupons**](ManagementApi.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
[**CreateCouponsForMultipleRecipients**](ManagementApi.md#CreateCouponsForMultipleRecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
[**CreatePasswordRecoveryEmail**](ManagementApi.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateRuleset**](ManagementApi.md#CreateRuleset) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | Create ruleset
[**CreateSession**](ManagementApi.md#CreateSession) | **Post** /v1/sessions | Create session
[**DeleteCampaign**](ManagementApi.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
[**DeleteCoupon**](ManagementApi.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
[**DeleteCoupons**](ManagementApi.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
[**DeleteReferral**](ManagementApi.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
[**DestroySession**](ManagementApi.md#DestroySession) | **Delete** /v1/sessions | Destroy session
[**ExportCoupons**](ManagementApi.md#ExportCoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons to CSV file
[**ExportCustomerSessions**](ManagementApi.md#ExportCustomerSessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions to CSV file
[**ExportEffects**](ManagementApi.md#ExportEffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects to CSV file
[**ExportLoyaltyBalance**](ManagementApi.md#ExportLoyaltyBalance) | **Get** /v1/loyalty_programs/{programID}/export_customer_balance | Export customer loyalty balance to a CSV file
[**ExportLoyaltyLedger**](ManagementApi.md#ExportLoyaltyLedger) | **Get** /v1/loyalty_programs/{programID}/profile/{integrationID}/export_log | Export a customer&#39;s loyalty ledger log to CSV file
[**ExportReferrals**](ManagementApi.md#ExportReferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals to CSV file
[**GetAccessLogsWithoutTotalCount**](ManagementApi.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
[**GetAccount**](ManagementApi.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get account details
[**GetAccountAnalytics**](ManagementApi.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
[**GetAdditionalCost**](ManagementApi.md#GetAdditionalCost) | **Get** /v1/additional_costs/{additionalCostId} | Get an additional cost
[**GetAdditionalCosts**](ManagementApi.md#GetAdditionalCosts) | **Get** /v1/additional_costs | List additional costs
[**GetAllAccessLogs**](ManagementApi.md#GetAllAccessLogs) | **Get** /v1/access_logs | List access logs
[**GetAllRoles**](ManagementApi.md#GetAllRoles) | **Get** /v1/roles | List roles
[**GetApplication**](ManagementApi.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementApi.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get report of health of application API
[**GetApplicationCustomer**](ManagementApi.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
[**GetApplicationCustomerFriends**](ManagementApi.md#GetApplicationCustomerFriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
[**GetApplicationCustomers**](ManagementApi.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
[**GetApplicationCustomersByAttributes**](ManagementApi.md#GetApplicationCustomersByAttributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
[**GetApplicationEventTypes**](ManagementApi.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
[**GetApplicationEventsWithoutTotalCount**](ManagementApi.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
[**GetApplicationSession**](ManagementApi.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
[**GetApplicationSessions**](ManagementApi.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
[**GetApplications**](ManagementApi.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementApi.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get a custom attribute
[**GetAttributes**](ManagementApi.md#GetAttributes) | **Get** /v1/attributes | List custom attributes
[**GetAudiences**](ManagementApi.md#GetAudiences) | **Get** /v1/audiences | List audiences
[**GetCampaign**](ManagementApi.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
[**GetCampaignAnalytics**](ManagementApi.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementApi.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
[**GetCampaigns**](ManagementApi.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
[**GetChanges**](ManagementApi.md#GetChanges) | **Get** /v1/changes | Get audit log for an account
[**GetCouponsByAttributes**](ManagementApi.md#GetCouponsByAttributes) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search | List coupons that match the given attributes
[**GetCouponsWithoutTotalCount**](ManagementApi.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
[**GetCustomerActivityReport**](ManagementApi.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementApi.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementApi.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
[**GetCustomerProfile**](ManagementApi.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get customer profile
[**GetCustomerProfiles**](ManagementApi.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List customer profiles
[**GetCustomersByAttributes**](ManagementApi.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
[**GetEventTypes**](ManagementApi.md#GetEventTypes) | **Get** /v1/event_types | List Event Types
[**GetExports**](ManagementApi.md#GetExports) | **Get** /v1/exports | Get Exports
[**GetLoyaltyPoints**](ManagementApi.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{programID}/profile/{integrationID} | Get the Loyalty Ledger for this integrationID
[**GetLoyaltyProgram**](ManagementApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{programID} | Get loyalty program
[**GetLoyaltyPrograms**](ManagementApi.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List loyalty programs
[**GetLoyaltyStatistics**](ManagementApi.md#GetLoyaltyStatistics) | **Get** /v1/loyalty_programs/{programID}/statistics | Get loyalty program statistics by loyalty program ID
[**GetReferralsWithoutTotalCount**](ManagementApi.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
[**GetRole**](ManagementApi.md#GetRole) | **Get** /v1/roles/{roleId} | Get role
[**GetRuleset**](ManagementApi.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
[**GetRulesets**](ManagementApi.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
[**GetUser**](ManagementApi.md#GetUser) | **Get** /v1/users/{userId} | Get a single user
[**GetUsers**](ManagementApi.md#GetUsers) | **Get** /v1/users | List users in account
[**GetWebhook**](ManagementApi.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get Webhook
[**GetWebhookActivationLogs**](ManagementApi.md#GetWebhookActivationLogs) | **Get** /v1/webhook_activation_logs | List webhook activation log entries
[**GetWebhookLogs**](ManagementApi.md#GetWebhookLogs) | **Get** /v1/webhook_logs | List webhook log entries
[**GetWebhooks**](ManagementApi.md#GetWebhooks) | **Get** /v1/webhooks | List webhooks
[**ImportCollection**](ManagementApi.md#ImportCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import collection via CSV file
[**ImportCoupons**](ManagementApi.md#ImportCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons via CSV file
[**ImportLoyaltyPoints**](ManagementApi.md#ImportLoyaltyPoints) | **Post** /v1/loyalty_programs/{programID}/import_points | Import loyalty points via CSV file
[**ImportPoolGiveaways**](ManagementApi.md#ImportPoolGiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
[**ImportReferrals**](ManagementApi.md#ImportReferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals via CSV file
[**RemoveLoyaltyPoints**](ManagementApi.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{programID}/profile/{integrationID}/deduct_points | Deduct points in loyalty program for given customer
[**ResetPassword**](ManagementApi.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes in all active campaigns of an application
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes
[**UpdateAdditionalCost**](ManagementApi.md#UpdateAdditionalCost) | **Put** /v1/additional_costs/{additionalCostId} | Update an additional cost
[**UpdateAttribute**](ManagementApi.md#UpdateAttribute) | **Put** /v1/attributes/{attributeId} | Update a custom attribute
[**UpdateCampaign**](ManagementApi.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
[**UpdateCoupon**](ManagementApi.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
[**UpdateCouponBatch**](ManagementApi.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update a batch of coupons
[**UpdateReferral**](ManagementApi.md#UpdateReferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
[**UpdateRuleset**](ManagementApi.md#UpdateRuleset) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Update ruleset



## AddLoyaltyPoints

> AddLoyaltyPoints(ctx, programID, integrationID).Body(body).Execute()

Add points in loyalty program for given customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**LoyaltyPoints**](LoyaltyPoints.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyCampaignToApplications

> InlineResponse2002 CopyCampaignToApplications(ctx, applicationId, campaignId).Body(body).Execute()

Copy the campaign into every specified application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyCampaignToApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CampaignCopy**](CampaignCopy.md) |  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalCost

> AccountAdditionalCost CreateAdditionalCost(ctx).Body(body).Execute()

Define a new additional cost



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) |  | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttribute

> Attribute CreateAttribute(ctx).Body(body).Execute()

Create custom attribute



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAttribute**](NewAttribute.md) |  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaign

> Campaign CreateCampaign(ctx, applicationId).Body(body).Execute()

Create campaign



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewCampaign**](NewCampaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupons

> InlineResponse2004 CreateCoupons(ctx, applicationId, campaignId).Body(body).Silent(silent).Execute()

Create coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCoupons**](NewCoupons.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to yes]

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsForMultipleRecipients

> InlineResponse2004 CreateCouponsForMultipleRecipients(ctx, applicationId, campaignId).Body(body).Silent(silent).Execute()

Create coupons for multiple recipients



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsForMultipleRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCouponsForMultipleRecipients**](NewCouponsForMultipleRecipients.md) |  | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains essential data such as the updated customer profiles and session-related information.  | [default to yes]

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePasswordRecoveryEmail

> NewPasswordEmail CreatePasswordRecoveryEmail(ctx).Body(body).Execute()

Request a password reset



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordRecoveryEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewPasswordEmail**](NewPasswordEmail.md) |  | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleset

> Ruleset CreateRuleset(ctx, applicationId, campaignId).Body(body).Execute()

Create ruleset



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewRuleset**](NewRuleset.md) |  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> Session CreateSession(ctx).Body(body).Execute()

Create session



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LoginParams**](LoginParams.md) |  | 

### Return type

[**Session**](Session.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, applicationId, campaignId).Execute()

Delete campaign



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


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


## DeleteCoupon

> DeleteCoupon(ctx, applicationId, campaignId, couponId).Execute()

Delete coupon



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**couponId** | **string** | The ID of the coupon code to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


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


## DeleteCoupons

> DeleteCoupons(ctx, applicationId, campaignId).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartsAfter(startsAfter).StartsBefore(startsBefore).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).Valid(valid).BatchId(batchId).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).Execute()

Delete coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

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


## DeleteReferral

> DeleteReferral(ctx, applicationId, campaignId, referralId).Execute()

Delete referral



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**referralId** | **string** | The ID of the referral code to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferralRequest struct via the builder pattern


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


## DestroySession

> DestroySession(ctx).Execute()

Destroy session



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySessionRequest struct via the builder pattern


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


## ExportCoupons

> string ExportCoupons(ctx, applicationId).CampaignId(campaignId).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).DateFormat(dateFormat).CampaignState(campaignState).Execute()

Export coupons to CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **dateFormat** | **string** | Determines the format of dates in the export document. | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Capmaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomerSessions

> string ExportCustomerSessions(ctx, applicationId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ProfileIntegrationId(profileIntegrationId).DateFormat(dateFormat).CustomerSessionState(customerSessionState).Execute()

Export customer sessions to CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCustomerSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **profileIntegrationId** | **string** | Only return sessions for the customer that matches this customer integration ID. | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 
 **customerSessionState** | **string** | Filter results by state. | 

### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEffects

> string ExportEffects(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()

Export triggered effects to CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportEffectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalance

> string ExportLoyaltyBalance(ctx, programID).Execute()

Export customer loyalty balance to a CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyLedger

> string ExportLoyaltyLedger(ctx, programID, integrationID).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()

Export a customer's loyalty ledger log to CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 


 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportReferrals

> string ExportReferrals(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).BatchId(batchId).DateFormat(dateFormat).Execute()

Export referrals to CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | If &#x60;true&#x60;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. If &#x60;false&#x60;, only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **batchId** | **string** | Filter results by batches of referrals | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogsWithoutTotalCount

> InlineResponse2008 GetAccessLogsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get access logs for Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get account details



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAnalytics

> AccountAnalytics GetAccountAnalytics(ctx, accountId).Execute()

Get account analytics



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCost

> AccountAdditionalCost GetAdditionalCost(ctx, additionalCostId).Execute()

Get an additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCosts

> InlineResponse20021 GetAdditionalCosts(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List additional costs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessLogs

> InlineResponse2009 GetAllAccessLogs(ctx).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List access logs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRoles

> InlineResponse20029 GetAllRoles(ctx).Execute()

List roles



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesRequest struct via the builder pattern


### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, applicationId).Execute()

Get Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApiHealth

> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId).Execute()

Get report of health of application API



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApiHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomer

> ApplicationCustomer GetApplicationCustomer(ctx, applicationId, customerId).Execute()

Get application's customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomerFriends

> InlineResponse20019 GetApplicationCustomerFriends(ctx, applicationId, integrationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List friends referred by customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**integrationId** | **string** | The Integration ID of the Advocate&#39;s Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerFriendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, &#x60;hasMore&#x60; will be true whenever there is a next page. &#x60;totalResultSize&#x60; will always be zero. With this flag set to false, &#x60;hasMore&#x60; will always be set to false. &#x60;totalResultSize&#x60; will contain the total number of results for this query.  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomers

> InlineResponse20011 GetApplicationCustomers(ctx, applicationId).IntegrationId(integrationId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application's customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **string** | Filter results performing an exact matching against the profile integration identifier. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, &#x60;hasMore&#x60; will be true whenever there is a next page. &#x60;totalResultSize&#x60; will always be zero. With this flag set to false, &#x60;hasMore&#x60; will always be set to false. &#x60;totalResultSize&#x60; will contain the total number of results for this query.  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomersByAttributes

> InlineResponse20012 GetApplicationCustomersByAttributes(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application customers matching the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, &#x60;hasMore&#x60; will be true whenever there is a next page. &#x60;totalResultSize&#x60; will always be zero. With this flag set to false, &#x60;hasMore&#x60; will always be set to false. &#x60;totalResultSize&#x60; will contain the total number of results for this query.  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventTypes

> InlineResponse20017 GetApplicationEventTypes(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications event types



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventsWithoutTotalCount

> InlineResponse20016 GetApplicationEventsWithoutTotalCount(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).Execute()

List Applications events



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **type_** | **string** | Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any timezone. Talon.One will convert to UTC internally. | 
 **session** | **string** | Session integration ID filter for events. Must be exact match. | 
 **profile** | **string** | Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **string** | Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **string** | Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **string** | Coupon code | 
 **referralCode** | **string** | Referral code | 
 **ruleQuery** | **string** | Rule name filter for events | 
 **campaignQuery** | **string** | Campaign name filter for events | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSession

> ApplicationSession GetApplicationSession(ctx, applicationId, sessionId).Execute()

Get Application session



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**sessionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSessions

> InlineResponse20015 GetApplicationSessions(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Profile(profile).State(state).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Coupon(coupon).Referral(referral).IntegrationId(integrationId).Execute()

List Application sessions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **profile** | **string** | Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **string** | Filter by sessions with this state. Must be exact match. | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any timezone. Talon.One will convert to UTC internally. | 
 **coupon** | **string** | Filter by sessions with this coupon. Must be exact match. | 
 **referral** | **string** | Filter by sessions with this referral. Must be exact match. | 
 **integrationId** | **string** | Filter by sessions with this integrationId. Must be exact match. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> InlineResponse2001 GetApplications(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttribute

> Attribute GetAttribute(ctx, attributeId).Execute()

Get a custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributes

> InlineResponse20020 GetAttributes(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Entity(entity).Execute()

List custom attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **entity** | **string** | Returned attributes will be filtered by supplied entity | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiences

> InlineResponse20018 GetAudiences(ctx).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List audiences



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, &#x60;hasMore&#x60; will be true whenever there is a next page. &#x60;totalResultSize&#x60; will always be zero. With this flag set to false, &#x60;hasMore&#x60; will always be set to false. &#x60;totalResultSize&#x60; will contain the total number of results for this query.  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, applicationId, campaignId).Execute()

Get campaign



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAnalytics

> InlineResponse20010 GetCampaignAnalytics(ctx, applicationId, campaignId).RangeStart(rangeStart).RangeEnd(rangeEnd).Granularity(granularity).Execute()

Get analytics of campaigns



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 
 **granularity** | **string** | The time interval between the results in the returned time-series. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByAttributes

> InlineResponse2002 GetCampaignByAttributes(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Execute()

List campaigns that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignSearch**](CampaignSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Capmaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> InlineResponse2002 GetCampaigns(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Name(name).Tags(tags).CreatedBefore(createdBefore).CreatedAfter(createdAfter).CampaignGroupId(campaignGroupId).TemplateId(templateId).Execute()

List campaigns



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Capmaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **campaignGroupId** | **int32** | Filter results to campaigns owned by the specified campaign group ID. | 
 **templateId** | **int32** | The ID of the Campaign Template this Campaign was created from. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> InlineResponse20027 GetChanges(ctx).PageSize(pageSize).Skip(skip).Sort(sort).ApplicationId(applicationId).EntityPath(entityPath).UserId(userId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).WithTotalResultSize(withTotalResultSize).IncludeOld(includeOld).Execute()

Get audit log for an account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **applicationId** | **int32** |  | 
 **entityPath** | **string** | Filter results on a case insensitive matching of the url path of the entity | 
 **userId** | **int32** |  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result will include the total size of the result, across all pages. This might decrease performance on large data sets. With this flag set to true, &#x60;hasMore&#x60; will be true whenever there is a next page. &#x60;totalResultSize&#x60; will always be zero. With this flag set to false, &#x60;hasMore&#x60; will always be set to false. &#x60;totalResultSize&#x60; will contain the total number of results for this query.  | 
 **includeOld** | **bool** | When this flag is set to false, the state without the change will not be returned. The default value is true. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsByAttributes

> InlineResponse2004 GetCouponsByAttributes(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

List coupons that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CouponSearch**](CouponSearch.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsWithoutTotalCount

> InlineResponse2005 GetCouponsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).Execute()

List coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReport

> CustomerActivityReport GetCustomerActivityReport(ctx, applicationId, customerId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Execute()

Get customer's activity report



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReportsWithoutTotalCount

> InlineResponse20014 GetCustomerActivityReportsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()

Get Activity Reports for Application Customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp, must be an RFC3339 timestamp string | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp, must be an RFC3339 timestamp string | 

 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **name** | **string** | Only return reports matching the customer name | 
 **integrationId** | **string** | Only return reports matching the integrationId | 
 **campaignName** | **string** | Only return reports matching the campaignName | 
 **advocateName** | **string** | Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAnalytics

> CustomerAnalytics GetCustomerAnalytics(ctx, applicationId, customerId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get customer's analytics report



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfile

> CustomerProfile GetCustomerProfile(ctx, customerId).Execute()

Get customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerProfile**](CustomerProfile.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfiles

> InlineResponse20013 GetCustomerProfiles(ctx).PageSize(pageSize).Skip(skip).Execute()

List customer profiles



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersByAttributes

> InlineResponse20013 GetCustomersByAttributes(ctx).Body(body).PageSize(pageSize).Skip(skip).Execute()

List customer profiles matching the given attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> InlineResponse20025 GetEventTypes(ctx).ApplicationIds(applicationIds).Name(name).IncludeOldVersions(includeOldVersions).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Event Types



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Filter by one or more application ids separated by comma | 
 **name** | **string** | Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **bool** | Include all versions of every event type. | [default to false]
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> InlineResponse20028 GetExports(ctx).PageSize(pageSize).Skip(skip).ApplicationId(applicationId).CampaignId(campaignId).Entity(entity).Execute()

Get Exports



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **applicationId** | **int32** |  | 
 **campaignId** | **int32** |  | 
 **entity** | **string** | The name of the entity type that was exported. | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPoints

> LoyaltyLedger GetLoyaltyPoints(ctx, programID, integrationID).Execute()

Get the Loyalty Ledger for this integrationID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** | The identifier for the application, must be unique within the account. | 
**integrationID** | **string** | The identifier for the application, must be unique within the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> LoyaltyProgram GetLoyaltyProgram(ctx, programID).Execute()

Get loyalty program



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPrograms

> InlineResponse2007 GetLoyaltyPrograms(ctx).Execute()

List loyalty programs



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramsRequest struct via the builder pattern


### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyStatistics

> LoyaltyStatistics GetLoyaltyStatistics(ctx, programID).Execute()

Get loyalty program statistics by loyalty program ID



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyStatistics**](LoyaltyStatistics.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralsWithoutTotalCount

> InlineResponse2006 GetReferralsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()

List referrals



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **code** | **string** | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiry date is set and in the past. The second matches referrals in which start date is null or in the past and expiry date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **string** | Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> Role GetRole(ctx, roleId).Execute()

Get role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleset

> Ruleset GetRuleset(ctx, applicationId, campaignId, rulesetId).Execute()

Get ruleset



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**rulesetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesets

> InlineResponse2003 GetRulesets(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List campaign rulesets



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get a single user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse20026 GetUsers(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List users in account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get Webhook



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookActivationLogs

> InlineResponse20023 GetWebhookActivationLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).IntegrationRequestUuid(integrationRequestUuid).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List webhook activation log entries



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookActivationLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **integrationRequestUuid** | **string** | Filter results by integration request UUID. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** |  | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. You can use any timezone. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogs

> InlineResponse20024 GetWebhookLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Status(status).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).RequestUuid(requestUuid).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List webhook log entries



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** |  | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **requestUuid** | **string** | Filter results by request UUID. | 
 **createdBefore** | **time.Time** | Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. You can use any timezone. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> InlineResponse20022 GetWebhooks(ctx).ApplicationIds(applicationIds).Sort(sort).PageSize(pageSize).Skip(skip).Execute()

List webhooks



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Filter by one or more application ids separated by comma | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCollection

> Import ImportCollection(ctx, applicationId, campaignId, collectionId).UpFile(upFile).Execute()

Import collection via CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**collectionId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upFile** | **string** | The file with the information about the data that should be imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCoupons

> Import ImportCoupons(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import coupons via CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file with the information about the data that should be imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyPoints

> Import ImportLoyaltyPoints(ctx, programID).UpFile(upFile).Execute()

Import loyalty points via CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file with the information about the data that should be imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportPoolGiveaways

> Import ImportPoolGiveaways(ctx, poolId).UpFile(upFile).Execute()

Import giveaway codes into a giveaway pool



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportPoolGiveawaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file with the information about the data that should be imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportReferrals

> Import ImportReferrals(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import referrals via CSV file



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file with the information about the data that should be imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoyaltyPoints

> RemoveLoyaltyPoints(ctx, programID, integrationID).Body(body).Execute()

Deduct points in loyalty program for given customer



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**programID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**LoyaltyPoints**](LoyaltyPoints.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> NewPassword ResetPassword(ctx).Body(body).Execute()

Reset password



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewPassword**](NewPassword.md) |  | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWideWithoutTotalCount

> InlineResponse2005 SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

List coupons that match the given attributes in all active campaigns of an application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedApplicationWideWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Capmaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedWithoutTotalCount

> InlineResponse2005 SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

List coupons that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pageSize** | **int32** | The number of items to include in this response. When omitted, the maximum value of 1000 will be used. | 
 **skip** | **int32** | Skips the given number of items when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. Sorting defaults to ascending order, prefix the field name with &#x60;-&#x60; to sort in descending order. | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any timezone. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiry date is set and in the past. The second matches coupons in which start date is null or in the past and expiry date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the Id of a referral, that meaning the coupons that had been created as an effect of the usage of a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalCost

> AccountAdditionalCost UpdateAdditionalCost(ctx, additionalCostId).Body(body).Execute()

Update an additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) |  | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> Attribute UpdateAttribute(ctx, attributeId).Body(body).Execute()

Update a custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAttribute**](NewAttribute.md) |  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, applicationId, campaignId).Body(body).Execute()

Update campaign



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCampaign**](UpdateCampaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> Coupon UpdateCoupon(ctx, applicationId, campaignId, couponId).Body(body).Execute()

Update coupon



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**couponId** | **string** | The ID of the coupon code to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateCoupon**](UpdateCoupon.md) |  | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponBatch

> UpdateCouponBatch(ctx, applicationId, campaignId).Body(body).Execute()

Update a batch of coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCouponBatch**](UpdateCouponBatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferral

> Referral UpdateReferral(ctx, applicationId, campaignId, referralId).Body(body).Execute()

Update referral



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**referralId** | **string** | The ID of the referral code to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateReferral**](UpdateReferral.md) |  | 

### Return type

[**Referral**](Referral.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRuleset

> Ruleset UpdateRuleset(ctx, applicationId, campaignId, rulesetId).Body(body).Execute()

Update ruleset



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** |  | 
**campaignId** | **int32** |  | 
**rulesetId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**NewRuleset**](NewRuleset.md) |  | 

### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

