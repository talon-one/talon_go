# Migration from older versions to v3.0.0

## Background and a short introduction to Cart Item Catalogs

Since July last year we are developing the [Cart Item Catalogs][1] feature and incorporating it into our platform and the Integration API.

Part of the Cart Item Catalogs feature core offerings is the ability to pre-sync a _catalog of items_ to Talon.One's platform and rely on them when sending integration requests to update Customer Sessions.

To support this flexibility, the formerly required `Name` and `Price` fields of the [CartItem][2] model are now flagged as optional. \
In the SDK, the meaning of this was changing the struct field types to pointers to a string and a float32 for the `Name` and `Price` fields respectively:

```diff
type CartItem struct {
	// Name of item.
-	Name string `json:"name"`
+	Name *string `json:"name,omitempty"`
	
	// ...
-	// Price of item
-	Price float32 `json:"price"`
+	// Price of the item in the currency defined by your Application. This field is required if this item is not part of a [catalog](https://docs.talon.one/docs/product/account/dev-tools/managing-cart-item-catalogs). If it is part of a catalog, setting a price here overrides the price from the catalog.
+	Price *float32 `json:"price,omitempty"`
}
```

## Required code changes

The main change upgrading to the SDK's new version ([v3.0.0][3]) is to update the references to the `Name` and `Price` of the `CartItem` model, upon initialization and accessing them.

### `CartItem` struct initialization

For example, let's assume we have an instruction in our code to initialize a list of cart items with their properties, and assign it to a customer session model.

For prior versions of the SDKs that looks like:
```go
newCustomerSession.SetCartItems([]talon.CartItem{
	{
		Name:     "Pad Thai - Veggie",
		Sku:      "pad-332",
		Quantity: 1,
		Price:    5.5,
		Category: talon.PtrString("Noodles"),
	},
	{
		Name:     "Chang",
		Sku:      "chang-br-42",
		Quantity: 1,
		Price:    2.3,
		Category: talon.PtrString("Beverages"),
	},
})
```

However, from version 3.0.0 on, that will look like:
```go
newCustomerSession.SetCartItems([]talon.CartItem{
	{
		Name:     talon.PtrString("Pad Thai - Veggie"),
		Sku:      "pad-332",
		Quantity: 1,
		Price:    talon.PtrFloat32(5.5),
		Category: talon.PtrString("Noodles"),
	},
	{
		Name:     talon.PtrString("Chang"),
		Sku:      "chang-br-42",
		Quantity: 1,
		Price:    talon.PtrFloat32(2.3),
		Category: talon.PtrString("Beverages"),
	},
})
```

Note the usage of the wrapping of the `Name` and `Price` fields' values with the provided [util functions][4] for pointer wrapping:
```diff
-	Name:     "Pad Thai - Veggie",
+	Name:     talon.PtrString("Pad Thai - Veggie"),

-	Price:    5.5,
+	Price:    talon.PtrFloat32(5.5),
```

If the item you are sending as part of the customer session update is part of a Cart Item Catalog, you can simply omit the `Name` field, `Price` field or both. These will be propagated automatically by the pre-synced Cart Item Catalog the item belongs to by the platform.

For items that are part of a Cart Item Catalog, sending a value in one or both of the fields **will override the values that were pre-synced as part of the Cart Item Catalog**.

For the above example, let's assume the _"Chang"_ beer item is part of a catalog already presence and synced to Talon.One, the code could be simplified to:
```go
newCustomerSession.SetCartItems([]talon.CartItem{
	{
		Name:     talon.PtrString("Pad Thai - Veggie"),
		Sku:      "pad-332",
		Quantity: 1,
		Price:    talon.PtrFloat32(5.5),
		Category: talon.PtrString("Noodles"),
	},
	{
		Sku:      "chang-br-42",
		Quantity: 1,
		Category: talon.PtrString("Beverages"),
	},
})
```

### Accessing returned `CartItem` fields

When your code needs to use the returned values for these fields of the `CartItem` model, you would need to update these references to also dereference the pointer to the value that is now returned from our APIs.

For example, if our code is accessing the name and price of a cart item upon a successful request to the `UpdateCustomerSession` endpoint (don't forget to request the `"customerSession"` as part of the [`ResponseContent` field][5]), then we'd need to update these references.

For prior versions of the SDKs that looks like:
```go
integrationState, _, err := api.
	UpdateCustomerSessionV2(context, "my_customer_session_id").
	Body(integrationRequest).
	Execute()

var firstCartItemName string
var secondCartItemPrice float32

firstCartItemName = integrationState.GetCustomerSession().CartItems[0].Name
secondCartItemPrice = integrationState.GetCustomerSession().CartItems[1].Price
```

However, from version 3.0.0 on, that will look like:
```go
integrationState, _, err := api.
	UpdateCustomerSessionV2(context, "my_customer_session_id").
	Body(integrationRequest).
	Execute()

var firstCartItemName string
var secondCartItemPrice float32

firstCartItemName = integrationState.GetCustomerSession().CartItems[0].GetName()
secondCartItemPrice = integrationState.GetCustomerSession().CartItems[1].GetPrice()
```

If we used the Getter functions already, our code won't require any changes. However, accessing the fields directly will impose the below changes:
```diff
integrationState, _, err := api.
	UpdateCustomerSessionV2(context, "my_customer_session_id").
	Body(integrationRequest).
	Execute()

var firstCartItemName string
var secondCartItemPrice float32

- firstCartItemName = integrationState.GetCustomerSession().CartItems[0].Name
+ firstCartItemName = integrationState.GetCustomerSession().CartItems[0].GetName()
- secondCartItemPrice = integrationState.GetCustomerSession().CartItems[1].Price
+ secondCartItemPrice = integrationState.GetCustomerSession().CartItems[1].GetPrice()
```

[1]: https://docs.talon.one/docs/product/account/dev-tools/managing-cart-item-catalogs "Cart Item Catalog"
[2]: ../model_cart_item.go "CartItem Model" 
[3]: https://github.com/talon-one/talon_go/releases/tag/v3.0.0 "v3.0.0"
[4]: ../utils.go "util functions"
[5]: https://github.com/talon-one/talon_go/blob/master/model_integration_request.go#L20-L21 "IntegrationRequest's response content"
