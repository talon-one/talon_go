# \ManagementApi

All URIs are relative to *https://yourbaseurl.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUserByEmail**](ManagementApi.md#ActivateUserByEmail) | **Post** /v1/users/activate | Enable user by email address
[**AddLoyaltyCardPoints**](ManagementApi.md#AddLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/add_points | Add points to card
[**AddLoyaltyPoints**](ManagementApi.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/add_points | Add points to customer profile
[**CopyCampaignToApplications**](ManagementApi.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into the specified Application
[**CreateAccountCollection**](ManagementApi.md#CreateAccountCollection) | **Post** /v1/collections | Create account-level collection
[**CreateAchievement**](ManagementApi.md#CreateAchievement) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | Create achievement
[**CreateAdditionalCost**](ManagementApi.md#CreateAdditionalCost) | **Post** /v1/additional_costs | Create additional cost
[**CreateAttribute**](ManagementApi.md#CreateAttribute) | **Post** /v1/attributes | Create custom attribute
[**CreateBatchLoyaltyCards**](ManagementApi.md#CreateBatchLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards/batch | Create loyalty cards
[**CreateCampaignFromTemplate**](ManagementApi.md#CreateCampaignFromTemplate) | **Post** /v1/applications/{applicationId}/create_campaign_from_template | Create campaign from campaign template
[**CreateCollection**](ManagementApi.md#CreateCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | Create campaign-level collection
[**CreateCoupons**](ManagementApi.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
[**CreateCouponsAsync**](ManagementApi.md#CreateCouponsAsync) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_async | Create coupons asynchronously
[**CreateCouponsDeletionJob**](ManagementApi.md#CreateCouponsDeletionJob) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_deletion_jobs | Creates a coupon deletion job
[**CreateCouponsForMultipleRecipients**](ManagementApi.md#CreateCouponsForMultipleRecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
[**CreateInviteEmail**](ManagementApi.md#CreateInviteEmail) | **Post** /v1/invite_emails | Resend invitation email
[**CreateInviteV2**](ManagementApi.md#CreateInviteV2) | **Post** /v2/invites | Invite user
[**CreatePasswordRecoveryEmail**](ManagementApi.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateSession**](ManagementApi.md#CreateSession) | **Post** /v1/sessions | Create session
[**CreateStore**](ManagementApi.md#CreateStore) | **Post** /v1/applications/{applicationId}/stores | Create store
[**DeactivateUserByEmail**](ManagementApi.md#DeactivateUserByEmail) | **Post** /v1/users/deactivate | Disable user by email address
[**DeductLoyaltyCardPoints**](ManagementApi.md#DeductLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/deduct_points | Deduct points from card
[**DeleteAccountCollection**](ManagementApi.md#DeleteAccountCollection) | **Delete** /v1/collections/{collectionId} | Delete account-level collection
[**DeleteAchievement**](ManagementApi.md#DeleteAchievement) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Delete achievement
[**DeleteCampaign**](ManagementApi.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
[**DeleteCollection**](ManagementApi.md#DeleteCollection) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Delete campaign-level collection
[**DeleteCoupon**](ManagementApi.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
[**DeleteCoupons**](ManagementApi.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
[**DeleteLoyaltyCard**](ManagementApi.md#DeleteLoyaltyCard) | **Delete** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Delete loyalty card
[**DeleteReferral**](ManagementApi.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
[**DeleteStore**](ManagementApi.md#DeleteStore) | **Delete** /v1/applications/{applicationId}/stores/{storeId} | Delete store
[**DeleteUser**](ManagementApi.md#DeleteUser) | **Delete** /v1/users/{userId} | Delete user
[**DeleteUserByEmail**](ManagementApi.md#DeleteUserByEmail) | **Post** /v1/users/delete | Delete user by email address
[**DestroySession**](ManagementApi.md#DestroySession) | **Delete** /v1/sessions | Destroy session
[**DisconnectCampaignStores**](ManagementApi.md#DisconnectCampaignStores) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores | Disconnect stores
[**ExportAccountCollectionItems**](ManagementApi.md#ExportAccountCollectionItems) | **Get** /v1/collections/{collectionId}/export | Export account-level collection&#39;s items
[**ExportAchievements**](ManagementApi.md#ExportAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId}/export | Export achievement customer data
[**ExportAudiencesMemberships**](ManagementApi.md#ExportAudiencesMemberships) | **Get** /v1/audiences/{audienceId}/memberships/export | Export audience members
[**ExportCampaignStores**](ManagementApi.md#ExportCampaignStores) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/export | Export stores
[**ExportCollectionItems**](ManagementApi.md#ExportCollectionItems) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/export | Export campaign-level collection&#39;s items
[**ExportCoupons**](ManagementApi.md#ExportCoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons
[**ExportCustomerSessions**](ManagementApi.md#ExportCustomerSessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions
[**ExportCustomersTiers**](ManagementApi.md#ExportCustomersTiers) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customers_tiers | Export customers&#39; tier data
[**ExportEffects**](ManagementApi.md#ExportEffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects
[**ExportLoyaltyBalance**](ManagementApi.md#ExportLoyaltyBalance) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balance | Export customer loyalty balance to CSV
[**ExportLoyaltyBalances**](ManagementApi.md#ExportLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balances | Export customer loyalty balances
[**ExportLoyaltyCardBalances**](ManagementApi.md#ExportLoyaltyCardBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_card_balances | Export all card transaction logs
[**ExportLoyaltyCardLedger**](ManagementApi.md#ExportLoyaltyCardLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/export_log | Export card&#39;s ledger log
[**ExportLoyaltyCards**](ManagementApi.md#ExportLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/export | Export loyalty cards
[**ExportLoyaltyLedger**](ManagementApi.md#ExportLoyaltyLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/export_log | Export customer&#39;s transaction logs
[**ExportPoolGiveaways**](ManagementApi.md#ExportPoolGiveaways) | **Get** /v1/giveaways/pools/{poolId}/export | Export giveaway codes of a giveaway pool
[**ExportReferrals**](ManagementApi.md#ExportReferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals
[**GetAccessLogsWithoutTotalCount**](ManagementApi.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
[**GetAccount**](ManagementApi.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get account details
[**GetAccountAnalytics**](ManagementApi.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
[**GetAccountCollection**](ManagementApi.md#GetAccountCollection) | **Get** /v1/collections/{collectionId} | Get account-level collection
[**GetAchievement**](ManagementApi.md#GetAchievement) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Get achievement
[**GetAdditionalCost**](ManagementApi.md#GetAdditionalCost) | **Get** /v1/additional_costs/{additionalCostId} | Get additional cost
[**GetAdditionalCosts**](ManagementApi.md#GetAdditionalCosts) | **Get** /v1/additional_costs | List additional costs
[**GetAllAccessLogs**](ManagementApi.md#GetAllAccessLogs) | **Get** /v1/access_logs | List access logs
[**GetApplication**](ManagementApi.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementApi.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get Application health
[**GetApplicationCustomer**](ManagementApi.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
[**GetApplicationCustomerFriends**](ManagementApi.md#GetApplicationCustomerFriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
[**GetApplicationCustomers**](ManagementApi.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
[**GetApplicationCustomersByAttributes**](ManagementApi.md#GetApplicationCustomersByAttributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
[**GetApplicationEventTypes**](ManagementApi.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
[**GetApplicationEventsWithoutTotalCount**](ManagementApi.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
[**GetApplicationSession**](ManagementApi.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
[**GetApplicationSessions**](ManagementApi.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
[**GetApplications**](ManagementApi.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementApi.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get custom attribute
[**GetAttributes**](ManagementApi.md#GetAttributes) | **Get** /v1/attributes | List custom attributes
[**GetAudienceMemberships**](ManagementApi.md#GetAudienceMemberships) | **Get** /v1/audiences/{audienceId}/memberships | List audience members
[**GetAudiences**](ManagementApi.md#GetAudiences) | **Get** /v1/audiences | List audiences
[**GetAudiencesAnalytics**](ManagementApi.md#GetAudiencesAnalytics) | **Get** /v1/audiences/analytics | List audience analytics
[**GetCampaign**](ManagementApi.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
[**GetCampaignAnalytics**](ManagementApi.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementApi.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
[**GetCampaignGroup**](ManagementApi.md#GetCampaignGroup) | **Get** /v1/campaign_groups/{campaignGroupId} | Get campaign access group
[**GetCampaignGroups**](ManagementApi.md#GetCampaignGroups) | **Get** /v1/campaign_groups | List campaign access groups
[**GetCampaignTemplates**](ManagementApi.md#GetCampaignTemplates) | **Get** /v1/campaign_templates | List campaign templates
[**GetCampaigns**](ManagementApi.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
[**GetChanges**](ManagementApi.md#GetChanges) | **Get** /v1/changes | Get audit logs for an account
[**GetCollection**](ManagementApi.md#GetCollection) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Get campaign-level collection
[**GetCollectionItems**](ManagementApi.md#GetCollectionItems) | **Get** /v1/collections/{collectionId}/items | Get collection items
[**GetCouponsWithoutTotalCount**](ManagementApi.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
[**GetCustomerActivityReport**](ManagementApi.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementApi.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementApi.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
[**GetCustomerProfile**](ManagementApi.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get customer profile
[**GetCustomerProfileAchievementProgress**](ManagementApi.md#GetCustomerProfileAchievementProgress) | **Get** /v1/applications/{applicationId}/achievement_progress/{integrationId} | List customer achievements
[**GetCustomerProfiles**](ManagementApi.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List customer profiles
[**GetCustomersByAttributes**](ManagementApi.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
[**GetEventTypes**](ManagementApi.md#GetEventTypes) | **Get** /v1/event_types | List event types
[**GetExports**](ManagementApi.md#GetExports) | **Get** /v1/exports | Get exports
[**GetLoyaltyCard**](ManagementApi.md#GetLoyaltyCard) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Get loyalty card
[**GetLoyaltyCardTransactionLogs**](ManagementApi.md#GetLoyaltyCardTransactionLogs) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/logs | List card&#39;s transactions
[**GetLoyaltyCards**](ManagementApi.md#GetLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards | List loyalty cards
[**GetLoyaltyPoints**](ManagementApi.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId} | Get customer&#39;s full loyalty ledger
[**GetLoyaltyProgram**](ManagementApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{loyaltyProgramId} | Get loyalty program
[**GetLoyaltyProgramTransactions**](ManagementApi.md#GetLoyaltyProgramTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/transactions | List loyalty program transactions
[**GetLoyaltyPrograms**](ManagementApi.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List loyalty programs
[**GetLoyaltyStatistics**](ManagementApi.md#GetLoyaltyStatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/statistics | Get loyalty program statistics
[**GetReferralsWithoutTotalCount**](ManagementApi.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
[**GetRoleV2**](ManagementApi.md#GetRoleV2) | **Get** /v2/roles/{roleId} | Get role
[**GetRuleset**](ManagementApi.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
[**GetRulesets**](ManagementApi.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
[**GetStore**](ManagementApi.md#GetStore) | **Get** /v1/applications/{applicationId}/stores/{storeId} | Get store
[**GetUser**](ManagementApi.md#GetUser) | **Get** /v1/users/{userId} | Get user
[**GetUsers**](ManagementApi.md#GetUsers) | **Get** /v1/users | List users in account
[**GetWebhook**](ManagementApi.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get webhook
[**GetWebhookActivationLogs**](ManagementApi.md#GetWebhookActivationLogs) | **Get** /v1/webhook_activation_logs | List webhook activation log entries
[**GetWebhookLogs**](ManagementApi.md#GetWebhookLogs) | **Get** /v1/webhook_logs | List webhook log entries
[**GetWebhooks**](ManagementApi.md#GetWebhooks) | **Get** /v1/webhooks | List webhooks
[**ImportAccountCollection**](ManagementApi.md#ImportAccountCollection) | **Post** /v1/collections/{collectionId}/import | Import data into existing account-level collection
[**ImportAllowedList**](ManagementApi.md#ImportAllowedList) | **Post** /v1/attributes/{attributeId}/allowed_list/import | Import allowed values for attribute
[**ImportAudiencesMemberships**](ManagementApi.md#ImportAudiencesMemberships) | **Post** /v1/audiences/{audienceId}/memberships/import | Import audience members
[**ImportCampaignStores**](ManagementApi.md#ImportCampaignStores) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/import | Import stores
[**ImportCollection**](ManagementApi.md#ImportCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import data into existing campaign-level collection
[**ImportCoupons**](ManagementApi.md#ImportCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons
[**ImportLoyaltyCards**](ManagementApi.md#ImportLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_cards | Import loyalty cards
[**ImportLoyaltyCustomersTiers**](ManagementApi.md#ImportLoyaltyCustomersTiers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_customers_tiers | Import customers into loyalty tiers
[**ImportLoyaltyPoints**](ManagementApi.md#ImportLoyaltyPoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_points | Import loyalty points
[**ImportPoolGiveaways**](ManagementApi.md#ImportPoolGiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
[**ImportReferrals**](ManagementApi.md#ImportReferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals
[**InviteUserExternal**](ManagementApi.md#InviteUserExternal) | **Post** /v1/users/invite | Invite user from identity provider
[**ListAccountCollections**](ManagementApi.md#ListAccountCollections) | **Get** /v1/collections | List collections in account
[**ListAchievements**](ManagementApi.md#ListAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | List achievements
[**ListAllRolesV2**](ManagementApi.md#ListAllRolesV2) | **Get** /v2/roles | List roles
[**ListCatalogItems**](ManagementApi.md#ListCatalogItems) | **Get** /v1/catalogs/{catalogId}/items | List items in a catalog
[**ListCollections**](ManagementApi.md#ListCollections) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | List collections in campaign
[**ListCollectionsInApplication**](ManagementApi.md#ListCollectionsInApplication) | **Get** /v1/applications/{applicationId}/collections | List collections in Application
[**ListStores**](ManagementApi.md#ListStores) | **Get** /v1/applications/{applicationId}/stores | List stores
[**NotificationActivation**](ManagementApi.md#NotificationActivation) | **Put** /v1/notifications/{notificationId}/activation | Activate or deactivate notification
[**OktaEventHandlerChallenge**](ManagementApi.md#OktaEventHandlerChallenge) | **Get** /v1/provisioning/okta | Validate Okta API ownership
[**PostAddedDeductedPointsNotification**](ManagementApi.md#PostAddedDeductedPointsNotification) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/notifications/added_deducted_points | Create notification about added or deducted loyalty points
[**PostCatalogsStrikethroughNotification**](ManagementApi.md#PostCatalogsStrikethroughNotification) | **Post** /v1/applications/{applicationId}/catalogs/notifications/strikethrough | Create strikethrough notification
[**PostPendingPointsNotification**](ManagementApi.md#PostPendingPointsNotification) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/notifications/pending_points | Create notification about pending loyalty points
[**RemoveLoyaltyPoints**](ManagementApi.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/deduct_points | Deduct points from customer profile
[**ResetPassword**](ManagementApi.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**ScimCreateUser**](ManagementApi.md#ScimCreateUser) | **Post** /v1/provisioning/scim/Users | Create SCIM user
[**ScimDeleteUser**](ManagementApi.md#ScimDeleteUser) | **Delete** /v1/provisioning/scim/Users/{userId} | Delete SCIM user
[**ScimGetResourceTypes**](ManagementApi.md#ScimGetResourceTypes) | **Get** /v1/provisioning/scim/ResourceTypes | List supported SCIM resource types
[**ScimGetSchemas**](ManagementApi.md#ScimGetSchemas) | **Get** /v1/provisioning/scim/Schemas | List supported SCIM schemas
[**ScimGetServiceProviderConfig**](ManagementApi.md#ScimGetServiceProviderConfig) | **Get** /v1/provisioning/scim/ServiceProviderConfig | Get SCIM service provider configuration
[**ScimGetUser**](ManagementApi.md#ScimGetUser) | **Get** /v1/provisioning/scim/Users/{userId} | Get SCIM user
[**ScimGetUsers**](ManagementApi.md#ScimGetUsers) | **Get** /v1/provisioning/scim/Users | List SCIM users
[**ScimPatchUser**](ManagementApi.md#ScimPatchUser) | **Patch** /v1/provisioning/scim/Users/{userId} | Update SCIM user attributes
[**ScimReplaceUserAttributes**](ManagementApi.md#ScimReplaceUserAttributes) | **Put** /v1/provisioning/scim/Users/{userId} | Update SCIM user
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes (without total count)
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementApi.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes in campaign (without total count)
[**TransferLoyaltyCard**](ManagementApi.md#TransferLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transfer | Transfer card data
[**UpdateAccountCollection**](ManagementApi.md#UpdateAccountCollection) | **Put** /v1/collections/{collectionId} | Update account-level collection
[**UpdateAchievement**](ManagementApi.md#UpdateAchievement) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Update achievement
[**UpdateAdditionalCost**](ManagementApi.md#UpdateAdditionalCost) | **Put** /v1/additional_costs/{additionalCostId} | Update additional cost
[**UpdateAttribute**](ManagementApi.md#UpdateAttribute) | **Put** /v1/attributes/{attributeId} | Update custom attribute
[**UpdateCampaign**](ManagementApi.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
[**UpdateCollection**](ManagementApi.md#UpdateCollection) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Update campaign-level collection&#39;s description
[**UpdateCoupon**](ManagementApi.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
[**UpdateCouponBatch**](ManagementApi.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update coupons
[**UpdateLoyaltyCard**](ManagementApi.md#UpdateLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Update loyalty card status
[**UpdateReferral**](ManagementApi.md#UpdateReferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
[**UpdateRoleV2**](ManagementApi.md#UpdateRoleV2) | **Put** /v2/roles/{roleId} | Update role
[**UpdateStore**](ManagementApi.md#UpdateStore) | **Put** /v1/applications/{applicationId}/stores/{storeId} | Update store
[**UpdateUser**](ManagementApi.md#UpdateUser) | **Put** /v1/users/{userId} | Update user



## ActivateUserByEmail

> ActivateUserByEmail(ctx).Body(body).Execute()

Enable user by email address



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ActivateUserRequest**](ActivateUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyCardPoints

> AddLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).Body(body).Execute()

Add points to card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyPoints

> AddLoyaltyPoints(ctx, loyaltyProgramId, integrationId).Body(body).Execute()

Add points to customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The identifier of the profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyCampaignToApplications

> InlineResponse2006 CopyCampaignToApplications(ctx, applicationId, campaignId).Body(body).Execute()

Copy the campaign into the specified Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyCampaignToApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CampaignCopy**](CampaignCopy.md) | body | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCollection

> Collection CreateAccountCollection(ctx).Body(body).Execute()

Create account-level collection



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewCollection**](NewCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAchievement

> Achievement CreateAchievement(ctx, applicationId, campaignId).Body(body).Execute()

Create achievement



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateAchievement**](CreateAchievement.md) | body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalCost

> AccountAdditionalCost CreateAdditionalCost(ctx).Body(body).Execute()

Create additional cost



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) | body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
 **body** | [**NewAttribute**](NewAttribute.md) | body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchLoyaltyCards

> LoyaltyCardBatchResponse CreateBatchLoyaltyCards(ctx, loyaltyProgramId).Body(body).Execute()

Create loyalty cards



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**LoyaltyCardBatch**](LoyaltyCardBatch.md) | body | 

### Return type

[**LoyaltyCardBatchResponse**](LoyaltyCardBatchResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignFromTemplate

> CreateTemplateCampaignResponse CreateCampaignFromTemplate(ctx, applicationId).Body(body).Execute()

Create campaign from campaign template



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateTemplateCampaign**](CreateTemplateCampaign.md) | body | 

### Return type

[**CreateTemplateCampaignResponse**](CreateTemplateCampaignResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx, applicationId, campaignId).Body(body).Execute()

Create campaign-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCampaignCollection**](NewCampaignCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupons

> InlineResponse2008 CreateCoupons(ctx, applicationId, campaignId).Body(body).Silent(silent).Execute()

Create coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCoupons**](NewCoupons.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsAsync

> AsyncCouponCreationResponse CreateCouponsAsync(ctx, applicationId, campaignId).Body(body).Execute()

Create coupons asynchronously



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCouponCreationJob**](NewCouponCreationJob.md) | body | 

### Return type

[**AsyncCouponCreationResponse**](AsyncCouponCreationResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsDeletionJob

> AsyncCouponDeletionJobResponse CreateCouponsDeletionJob(ctx, applicationId, campaignId).Body(body).Execute()

Creates a coupon deletion job



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsDeletionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCouponDeletionJob**](NewCouponDeletionJob.md) | body | 

### Return type

[**AsyncCouponDeletionJobResponse**](AsyncCouponDeletionJobResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsForMultipleRecipients

> InlineResponse2008 CreateCouponsForMultipleRecipients(ctx, applicationId, campaignId).Body(body).Silent(silent).Execute()

Create coupons for multiple recipients



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsForMultipleRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewCouponsForMultipleRecipients**](NewCouponsForMultipleRecipients.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the perfomance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to yes]

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteEmail

> NewInviteEmail CreateInviteEmail(ctx).Body(body).Execute()

Resend invitation email



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewInviteEmail**](NewInviteEmail.md) | body | 

### Return type

[**NewInviteEmail**](NewInviteEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteV2

> User CreateInviteV2(ctx).Body(body).Execute()

Invite user



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewInvitation**](NewInvitation.md) | body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
 **body** | [**NewPasswordEmail**](NewPasswordEmail.md) | body | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
 **body** | [**LoginParams**](LoginParams.md) | body | 

### Return type

[**Session**](Session.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStore

> Store CreateStore(ctx, applicationId).Body(body).Execute()

Create store



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewStore**](NewStore.md) | body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUserByEmail

> DeactivateUserByEmail(ctx).Body(body).Execute()

Disable user by email address



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeactivateUserRequest**](DeactivateUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeductLoyaltyCardPoints

> DeductLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).Body(body).Execute()

Deduct points from card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeductLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountCollection

> DeleteAccountCollection(ctx, collectionId).Execute()

Delete account-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAchievement

> DeleteAchievement(ctx, applicationId, campaignId, achievementId).Execute()

Delete achievement



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, applicationId, campaignId, collectionId).Execute()

Delete campaign-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string** | The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | - &#x60;expired&#x60;: Matches coupons in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches coupons in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches coupons in which start date is set and in the future.  | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **usable** | **string** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **referralId** | **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s &#x60;RecipientIntegrationId&#x60; field.  | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoyaltyCard

> DeleteLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Execute()

Delete loyalty card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string** | The ID of the referral code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStore(ctx, applicationId, storeId).Execute()

Delete store



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserByEmail

> DeleteUserByEmail(ctx).Body(body).Execute()

Delete user by email address



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteUserRequest**](DeleteUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
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

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectCampaignStores

> DisconnectCampaignStores(ctx, applicationId, campaignId).Execute()

Disconnect stores



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAccountCollectionItems

> string ExportAccountCollectionItems(ctx, collectionId).Execute()

Export account-level collection's items



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAchievements

> string ExportAchievements(ctx, applicationId, campaignId, achievementId).Execute()

Export achievement customer data



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAchievementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAudiencesMemberships

> string ExportAudiencesMemberships(ctx, audienceId).Execute()

Export audience members



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAudiencesMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCampaignStores

> string ExportCampaignStores(ctx, applicationId, campaignId).Execute()

Export stores



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCollectionItems

> string ExportCollectionItems(ctx, applicationId, campaignId, collectionId).Execute()

Export campaign-level collection's items



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCoupons

> string ExportCoupons(ctx, applicationId).CampaignId(campaignId).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).DateFormat(dateFormat).CampaignState(campaignState).ValuesOnly(valuesOnly).Execute()

Export coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **dateFormat** | **string** | Determines the format of dates in the export document. | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **valuesOnly** | **bool** | Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomerSessions

> string ExportCustomerSessions(ctx, applicationId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ProfileIntegrationId(profileIntegrationId).DateFormat(dateFormat).CustomerSessionState(customerSessionState).Execute()

Export customer sessions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

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

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomersTiers

> string ExportCustomersTiers(ctx, loyaltyProgramId).SubledgerIds(subledgerIds).TierNames(tierNames).Execute()

Export customers' tier data



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCustomersTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subledgerIds** | [**[]string**](string.md) | An array of subledgers IDs to filter the export by. | 
 **tierNames** | [**[]string**](string.md) | An array of tier names to filter the export by. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEffects

> string ExportEffects(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()

Export triggered effects



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportEffectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalance

> string ExportLoyaltyBalance(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export customer loyalty balance to CSV



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalances

> string ExportLoyaltyBalances(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export customer loyalty balances



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardBalances

> string ExportLoyaltyCardBalances(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export all card transaction logs



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardLedger

> string ExportLoyaltyCardLedger(ctx, loyaltyProgramId, loyaltyCardId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()

Export card's ledger log



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCards

> string ExportLoyaltyCards(ctx, loyaltyProgramId).BatchId(batchId).Execute()

Export loyalty cards



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchId** | **string** | Filter results by loyalty card batch ID. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyLedger

> string ExportLoyaltyLedger(ctx, loyaltyProgramId, integrationId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()

Export customer's transaction logs



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The identifier of the profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 


 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPoolGiveaways

> string ExportPoolGiveaways(ctx, poolId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

Export giveaway codes of a giveaway pool



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int32** | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPoolGiveawaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **time.Time** | Timestamp that filters the results to only contain giveaways created before this date. Must be an RFC3339 timestamp string. | 
 **createdAfter** | **time.Time** | Timestamp that filters the results to only contain giveaways created after this date. Must be an RFC3339 timestamp string. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportReferrals

> string ExportReferrals(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).BatchId(batchId).DateFormat(dateFormat).Execute()

Export referrals



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | - &#x60;expired&#x60;: Matches referrals in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches referrals in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | - &#x60;true&#x60;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;, only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **batchId** | **string** | Filter results by batches of referrals | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogsWithoutTotalCount

> InlineResponse20019 GetAccessLogsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get access logs for Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**accountId** | **int32** | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**accountId** | **int32** | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCollection

> Collection GetAccountCollection(ctx, collectionId).Execute()

Get account-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAchievement

> Achievement GetAchievement(ctx, applicationId, campaignId, achievementId).Execute()

Get achievement



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCost

> AccountAdditionalCost GetAdditionalCost(ctx, additionalCostId).Execute()

Get additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** | The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCosts

> InlineResponse20036 GetAdditionalCosts(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List additional costs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessLogs

> InlineResponse20020 GetAllAccessLogs(ctx).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List access logs



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApiHealth

> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId).Execute()

Get Application health



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApiHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomerFriends

> InlineResponse20033 GetApplicationCustomerFriends(ctx, applicationId, integrationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List friends referred by customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string** | The Integration ID of the Advocate&#39;s Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerFriendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomers

> InlineResponse20022 GetApplicationCustomers(ctx, applicationId).IntegrationId(integrationId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application's customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **string** | Filter results performing an exact matching against the profile integration identifier. | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomersByAttributes

> InlineResponse20023 GetApplicationCustomersByAttributes(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application customers matching the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) | body | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventTypes

> InlineResponse20029 GetApplicationEventTypes(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications event types



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventsWithoutTotalCount

> InlineResponse20028 GetApplicationEventsWithoutTotalCount(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).Execute()

List Applications events



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **type_** | **string** | Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **session** | **string** | Session integration ID filter for events. Must be exact match. | 
 **profile** | **string** | Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **string** | Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **string** | Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **string** | Coupon code | 
 **referralCode** | **string** | Referral code | 
 **ruleQuery** | **string** | Rule name filter for events | 
 **campaignQuery** | **string** | Campaign name filter for events | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**sessionId** | **int32** | The **internal** ID of the session. You can get the ID with the [List Application sessions](https://docs.talon.one/management-api#tag/Customer-data/operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSessions

> InlineResponse20027 GetApplicationSessions(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Profile(profile).State(state).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Coupon(coupon).Referral(referral).IntegrationId(integrationId).StoreIntegrationId(storeIntegrationId).Execute()

List Application sessions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **profile** | **string** | Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **string** | Filter by sessions with this state. Must be exact match. | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **coupon** | **string** | Filter by sessions with this coupon. Must be exact match. | 
 **referral** | **string** | Filter by sessions with this referral. Must be exact match. | 
 **integrationId** | **string** | Filter by sessions with this integrationId. Must be exact match. | 
 **storeIntegrationId** | **string** | The integration ID of the store. You choose this ID when you create a store. | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> InlineResponse2005 GetApplications(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttribute

> Attribute GetAttribute(ctx, attributeId).Execute()

Get custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributes

> InlineResponse20034 GetAttributes(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Entity(entity).Execute()

List custom attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **entity** | **string** | Returned attributes will be filtered by supplied entity. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceMemberships

> InlineResponse20032 GetAudienceMemberships(ctx, audienceId).PageSize(pageSize).Skip(skip).Sort(sort).ProfileQuery(profileQuery).Execute()

List audience members



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **profileQuery** | **string** | The filter to select a profile. | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiences

> InlineResponse20030 GetAudiences(ctx).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List audiences



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiencesAnalytics

> InlineResponse20031 GetAudiencesAnalytics(ctx).AudienceIds(audienceIds).Sort(sort).Execute()

List audience analytics



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencesAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceIds** | **string** | The IDs of one or more audiences, separated by commas, by which to filter results. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAnalytics

> InlineResponse20021 GetCampaignAnalytics(ctx, applicationId, campaignId).RangeStart(rangeStart).RangeEnd(rangeEnd).Granularity(granularity).Execute()

Get analytics of campaigns



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **granularity** | **string** | The time interval between the results in the returned time-series. | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByAttributes

> InlineResponse2006 GetCampaignByAttributes(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Execute()

List campaigns that match the given attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignSearch**](CampaignSearch.md) | body | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroup

> CampaignGroup GetCampaignGroup(ctx, campaignGroupId).Execute()

Get campaign access group



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGroupId** | **int32** | The ID of the campaign access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignGroup**](CampaignGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroups

> InlineResponse20011 GetCampaignGroups(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List campaign access groups



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplates

> InlineResponse20012 GetCampaignTemplates(ctx).PageSize(pageSize).Skip(skip).Sort(sort).State(state).Name(name).Tags(tags).UserId(userId).Execute()

List campaign templates



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **state** | **string** | Filter results by the state of the campaign template. | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign template. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign template. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values.  | 
 **userId** | **int32** | Filter results by user ID. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> InlineResponse2006 GetCampaigns(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Name(name).Tags(tags).CreatedBefore(createdBefore).CreatedAfter(createdAfter).CampaignGroupId(campaignGroupId).TemplateId(templateId).StoreId(storeId).Execute()

List campaigns



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **campaignGroupId** | **int32** | Filter results to campaigns owned by the specified campaign access group ID. | 
 **templateId** | **int32** | The ID of the Campaign Template this Campaign was created from. | 
 **storeId** | **int32** | Filter results to campaigns linked to the specified store ID. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> InlineResponse20042 GetChanges(ctx).PageSize(pageSize).Skip(skip).Sort(sort).ApplicationId(applicationId).EntityPath(entityPath).UserId(userId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).WithTotalResultSize(withTotalResultSize).ManagementKeyId(managementKeyId).IncludeOld(includeOld).Execute()

Get audit logs for an account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **entityPath** | **string** | Filter results on a case insensitive matching of the url path of the entity | 
 **userId** | **int32** | Filter results by user ID. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **managementKeyId** | **int32** | Filter results that match the given management key ID. | 
 **includeOld** | **bool** | When this flag is set to false, the state without the change will not be returned. The default value is true. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> Collection GetCollection(ctx, applicationId, campaignId, collectionId).Execute()

Get campaign-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItems

> InlineResponse20018 GetCollectionItems(ctx, collectionId).PageSize(pageSize).Skip(skip).Execute()

Get collection items



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsWithoutTotalCount

> InlineResponse2009 GetCouponsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Redeemed(redeemed).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).ExpiresBefore(expiresBefore).ExpiresAfter(expiresAfter).StartsBefore(startsBefore).StartsAfter(startsAfter).ValuesOnly(valuesOnly).Execute()

List coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **redeemed** | **string** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &gt; 0&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &#x3D; 0&#x60; will be returned. - This field cannot be used in conjunction with the &#x60;usable&#x60; query parameter.  | 
 **referralId** | **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valuesOnly** | **bool** | Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReportsWithoutTotalCount

> InlineResponse20026 GetCustomerActivityReportsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()

Get Activity Reports for Application Customers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **name** | **string** | Only return reports matching the customer name | 
 **integrationId** | **string** | Filter results performing an exact matching against the profile integration identifier. | 
 **campaignName** | **string** | Only return reports matching the campaignName | 
 **advocateName** | **string** | Only return reports matching the current customer referrer name | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int32** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**customerId** | **int32** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerProfile**](CustomerProfile.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfileAchievementProgress

> InlineResponse20047 GetCustomerProfileAchievementProgress(ctx, applicationId, integrationId).PageSize(pageSize).Skip(skip).AchievementId(achievementId).Title(title).Execute()

List customer achievements



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string** | The identifier of the profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileAchievementProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **achievementId** | **int32** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 
 **title** | **string** | Filter results by the &#x60;title&#x60; of an achievement. | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfiles

> InlineResponse20025 GetCustomerProfiles(ctx).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()

List customer profiles



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sandbox** | **bool** | Indicates whether you are pointing to a sandbox or Live customer. | [default to false]

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersByAttributes

> InlineResponse20024 GetCustomersByAttributes(ctx).Body(body).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()

List customer profiles matching the given attributes



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) | body | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sandbox** | **bool** | Indicates whether you are pointing to a sandbox or Live customer. | [default to false]

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> InlineResponse20040 GetEventTypes(ctx).Name(name).IncludeOldVersions(includeOldVersions).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List event types



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **bool** | Include all versions of every event type. | [default to false]
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> InlineResponse20043 GetExports(ctx).PageSize(pageSize).Skip(skip).ApplicationId(applicationId).CampaignId(campaignId).Entity(entity).Execute()

Get exports



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **campaignId** | **int32** | Filter by the campaign ID on which the limit counters are used. | 
 **entity** | **string** | The name of the entity type that was exported. | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCard

> LoyaltyCard GetLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Execute()

Get loyalty card



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactionLogs

> InlineResponse20016 GetLoyaltyCardTransactionLogs(ctx, loyaltyProgramId, loyaltyCardId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).SubledgerId(subledgerId).Execute()

List card's transactions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardTransactionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCards

> InlineResponse20015 GetLoyaltyCards(ctx, loyaltyProgramId).PageSize(pageSize).Skip(skip).Sort(sort).Identifier(identifier).ProfileId(profileId).BatchId(batchId).Execute()

List loyalty cards



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **identifier** | **string** | The card code by which to filter loyalty cards in the response. | 
 **profileId** | **int32** | Filter results by customer profile ID. | 
 **batchId** | **string** | Filter results by loyalty card batch ID. | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPoints

> LoyaltyLedger GetLoyaltyPoints(ctx, loyaltyProgramId, integrationId).Execute()

Get customer's full loyalty ledger



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The identifier of the profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> LoyaltyProgram GetLoyaltyProgram(ctx, loyaltyProgramId).Execute()

Get loyalty program



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramTransactions

> InlineResponse20014 GetLoyaltyProgramTransactions(ctx, loyaltyProgramId).LoyaltyTransactionType(loyaltyTransactionType).SubledgerId(subledgerId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).Execute()

List loyalty program transactions



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPrograms

> InlineResponse20013 GetLoyaltyPrograms(ctx).Execute()

List loyalty programs



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramsRequest struct via the builder pattern


### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyStatistics

> LoyaltyDashboardData GetLoyaltyStatistics(ctx, loyaltyProgramId).Execute()

Get loyalty program statistics



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyDashboardData**](LoyaltyDashboardData.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralsWithoutTotalCount

> InlineResponse20010 GetReferralsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()

List referrals



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **code** | **string** | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiration date is set and in the past. The second matches referrals in which start date is null or in the past and expiration date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **string** | Filter results by match with a profile id specified in the referral&#39;s AdvocateProfileIntegrationId field | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleV2

> RoleV2 GetRoleV2(ctx, roleId).Execute()

Get role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**rulesetId** | **int32** | The ID of the ruleset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesets

> InlineResponse2007 GetRulesets(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List campaign rulesets



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> Store GetStore(ctx, applicationId, storeId).Execute()

Get store



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> InlineResponse20041 GetUsers(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List users in account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get webhook



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | The ID of the webhook. You can find the ID in the Campaign Manager&#39;s URL when you display the details of the webhook in **Account** &gt; **Webhooks**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookActivationLogs

> InlineResponse20038 GetWebhookActivationLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).IntegrationRequestUuid(integrationRequestUuid).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List webhook activation log entries



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookActivationLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **integrationRequestUuid** | **string** | Filter results by integration request UUID. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookLogs

> InlineResponse20039 GetWebhookLogs(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Status(status).WebhookId(webhookId).ApplicationId(applicationId).CampaignId(campaignId).RequestUuid(requestUuid).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

List webhook log entries



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **webhookId** | **float32** | Filter results by Webhook. | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **requestUuid** | **string** | Filter results by request UUID. | 
 **createdBefore** | **time.Time** | Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> InlineResponse20037 GetWebhooks(ctx).ApplicationIds(applicationIds).Sort(sort).PageSize(pageSize).Skip(skip).CreationType(creationType).Visibility(visibility).OutgoingIntegrationsTypeId(outgoingIntegrationsTypeId).Title(title).Execute()

List webhooks



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Checks if the given catalog or its attributes are referenced in the specified Application ID.  **Note**: If no Application ID is provided, we check for all connected Applications.  | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **creationType** | **string** | Filter results by creation type. | 
 **visibility** | **string** | Filter results by visibility. | 
 **outgoingIntegrationsTypeId** | **int32** | Filter results by outgoing integration type ID. | 
 **title** | **string** | Filter results performing case-insensitive matching against the webhook title. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccountCollection

> Import ImportAccountCollection(ctx, collectionId).UpFile(upFile).Execute()

Import data into existing account-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAllowedList

> Import ImportAllowedList(ctx, attributeId).UpFile(upFile).Execute()

Import allowed values for attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAudiencesMemberships

> Import ImportAudiencesMemberships(ctx, audienceId).UpFile(upFile).Execute()

Import audience members



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int32** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAudiencesMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCampaignStores

> Import ImportCampaignStores(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import stores



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCollection

> Import ImportCollection(ctx, applicationId, campaignId, collectionId).UpFile(upFile).Execute()

Import data into existing campaign-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCoupons

> Import ImportCoupons(ctx, applicationId, campaignId).SkipDuplicates(skipDuplicates).UpFile(upFile).Execute()

Import coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipDuplicates** | **bool** | An indicator of whether to skip duplicate coupon values instead of causing an error. Duplicate values are ignored when &#x60;skipDuplicates&#x3D;true&#x60;.  | 
 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCards

> Import ImportLoyaltyCards(ctx, loyaltyProgramId).UpFile(upFile).Execute()

Import loyalty cards



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCustomersTiers

> Import ImportLoyaltyCustomersTiers(ctx, loyaltyProgramId).UpFile(upFile).Execute()

Import customers into loyalty tiers



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyCustomersTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyPoints

> Import ImportLoyaltyPoints(ctx, loyaltyProgramId).UpFile(upFile).Execute()

Import loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**poolId** | **int32** | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportPoolGiveawaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportReferrals

> Import ImportReferrals(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import referrals



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserExternal

> InviteUserExternal(ctx).Body(body).Execute()

Invite user from identity provider



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NewExternalInvitation**](NewExternalInvitation.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountCollections

> InlineResponse20017 ListAccountCollections(ctx).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAchievements

> InlineResponse20046 ListAchievements(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Title(title).Execute()

List achievements



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAchievementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 50]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **title** | **string** | Filter by the display name for the achievement in the campaign manager.  **Note**: If no &#x60;title&#x60; is provided, all the achievements from the campaign are returned.  | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRolesV2

> InlineResponse20044 ListAllRolesV2(ctx).Execute()

List roles



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllRolesV2Request struct via the builder pattern


### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItems

> InlineResponse20035 ListCatalogItems(ctx, catalogId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Sku(sku).ProductNames(productNames).Execute()

List items in a catalog



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int32** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **sku** | [**[]string**](string.md) | Filter results by one or more SKUs. Must be exact match. | 
 **productNames** | [**[]string**](string.md) | Filter results by one or more product names. Must be exact match. | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> InlineResponse20017 ListCollections(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in campaign



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsInApplication

> InlineResponse20017 ListCollectionsInApplication(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in Application



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsInApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> InlineResponse20045 ListStores(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).CampaignId(campaignId).Name(name).IntegrationId(integrationId).Query(query).Execute()

List stores



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **campaignId** | **float32** | Filter results by campaign. | 
 **name** | **string** | The name of the store. | 
 **integrationId** | **string** | The integration ID of the store. | 
 **query** | **string** | Filter results by &#x60;name&#x60; or &#x60;integrationId&#x60;. | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationActivation

> NotificationActivation(ctx, notificationId).Body(body).Execute()

Activate or deactivate notification



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **int32** | The ID of the notification. Get it with the appropriate _List notifications_ endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationActivationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NotificationActivation**](NotificationActivation.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OktaEventHandlerChallenge

> OktaEventHandlerChallenge(ctx).Execute()

Validate Okta API ownership



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOktaEventHandlerChallengeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddedDeductedPointsNotification

> BaseNotification PostAddedDeductedPointsNotification(ctx, loyaltyProgramId).Body(body).Execute()

Create notification about added or deducted loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddedDeductedPointsNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewBaseNotification**](NewBaseNotification.md) | body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCatalogsStrikethroughNotification

> BaseNotification PostCatalogsStrikethroughNotification(ctx, applicationId).Body(body).Execute()

Create strikethrough notification



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCatalogsStrikethroughNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewBaseNotification**](NewBaseNotification.md) | body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPendingPointsNotification

> BaseNotification PostPendingPointsNotification(ctx, loyaltyProgramId).Body(body).Execute()

Create notification about pending loyalty points



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPendingPointsNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewBaseNotification**](NewBaseNotification.md) | body | 

### Return type

[**BaseNotification**](BaseNotification.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoyaltyPoints

> RemoveLoyaltyPoints(ctx, loyaltyProgramId, integrationId).Body(body).Execute()

Deduct points from customer profile



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The identifier of the profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
 **body** | [**NewPassword**](NewPassword.md) | body | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimCreateUser

> ScimUser ScimCreateUser(ctx).Body(body).Execute()

Create SCIM user



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScimCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScimNewUser**](ScimNewUser.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimDeleteUser

> ScimDeleteUser(ctx, userId).Execute()

Delete SCIM user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetResourceTypes

> ScimResourceTypesListResponse ScimGetResourceTypes(ctx).Execute()

List supported SCIM resource types



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetResourceTypesRequest struct via the builder pattern


### Return type

[**ScimResourceTypesListResponse**](ScimResourceTypesListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetSchemas

> ScimSchemasListResponse ScimGetSchemas(ctx).Execute()

List supported SCIM schemas



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetSchemasRequest struct via the builder pattern


### Return type

[**ScimSchemasListResponse**](ScimSchemasListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetServiceProviderConfig

> ScimServiceProviderConfigResponse ScimGetServiceProviderConfig(ctx).Execute()

Get SCIM service provider configuration



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetServiceProviderConfigRequest struct via the builder pattern


### Return type

[**ScimServiceProviderConfigResponse**](ScimServiceProviderConfigResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUser

> ScimUser ScimGetUser(ctx, userId).Execute()

Get SCIM user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUsers

> ScimUsersListResponse ScimGetUsers(ctx).Execute()

List SCIM users



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetUsersRequest struct via the builder pattern


### Return type

[**ScimUsersListResponse**](ScimUsersListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimPatchUser

> ScimUser ScimPatchUser(ctx, userId).Body(body).Execute()

Update SCIM user attributes



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimPatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScimPatchRequest**](ScimPatchRequest.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimReplaceUserAttributes

> ScimUser ScimReplaceUserAttributes(ctx, userId).Body(body).Execute()

Update SCIM user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimReplaceUserAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ScimNewUser**](ScimNewUser.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWideWithoutTotalCount

> InlineResponse2009 SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

List coupons that match the given attributes (without total count)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedApplicationWideWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | body | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedWithoutTotalCount

> InlineResponse2009 SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

List coupons that match the given attributes in campaign (without total count)



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | body | 
 **pageSize** | **int32** | The number of items in the response. | [default to 1000]
 **skip** | **int32** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** This parameter works only with numeric fields.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int32** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferLoyaltyCard

> TransferLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Body(body).Execute()

Transfer card data



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TransferLoyaltyCard**](TransferLoyaltyCard.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountCollection

> Collection UpdateAccountCollection(ctx, collectionId).Body(body).Execute()

Update account-level collection



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateCollection**](UpdateCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAchievement

> Achievement UpdateAchievement(ctx, applicationId, campaignId, achievementId).Body(body).Execute()

Update achievement



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int32** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateAchievement**](UpdateAchievement.md) | body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalCost

> AccountAdditionalCost UpdateAdditionalCost(ctx, additionalCostId).Body(body).Execute()

Update additional cost



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int32** | The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAdditionalCost**](NewAdditionalCost.md) | body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> Attribute UpdateAttribute(ctx, attributeId).Body(body).Execute()

Update custom attribute



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NewAttribute**](NewAttribute.md) | body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCampaign**](UpdateCampaign.md) | body | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, applicationId, campaignId, collectionId).Body(body).Execute()

Update campaign-level collection's description



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int32** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateCampaignCollection**](UpdateCampaignCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string** | The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateCoupon**](UpdateCoupon.md) | body | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponBatch

> UpdateCouponBatch(ctx, applicationId, campaignId).Body(body).Execute()

Update coupons



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateCouponBatch**](UpdateCouponBatch.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoyaltyCard

> LoyaltyCard UpdateLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Body(body).Execute()

Update loyalty card status



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int32** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateLoyaltyCard**](UpdateLoyaltyCard.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int32** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string** | The ID of the referral code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateReferral**](UpdateReferral.md) | body | 

### Return type

[**Referral**](Referral.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleV2

> RoleV2 UpdateRoleV2(ctx, roleId).Body(body).Execute()

Update role



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int32** | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **RoleV2Base** | body | 

### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStore

> Store UpdateStore(ctx, applicationId, storeId).Body(body).Execute()

Update store



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**NewStore**](NewStore.md) | body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).Body(body).Execute()

Update user



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateUser**](UpdateUser.md) | body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

