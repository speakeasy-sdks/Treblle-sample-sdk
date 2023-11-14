# Orders


## Overview

The orders endpoints.

### Available Operations

* [createOrder](#createorder) - Create an order.

## createOrder

Create an order for a drink.

### Example Usage

```php
<?php

declare(strict_types=1);
require_once 'vendor/autoload.php';

use \OpenAPI\OpenAPI;
use \OpenAPI\OpenAPI\Models\Components;
use \OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '';

$sdk = OpenAPI\Speakeasy::builder()
    ->setSecurity($security)
    ->build();

try {


    $response = $sdk->orders->createOrder([
    new Components\OrderInput(),
], 'string');

    if ($response->order !== null) {
        // handle response
    }
} catch (Exception $e) {
    // handle exception
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `requestBody`                                                                                 | array<[\OpenAPI\OpenAPI\Models\Components\OrderInput](../../Models/Components/OrderInput.md)> | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `callbackUrl`                                                                                 | *string*                                                                                      | :heavy_minus_sign:                                                                            | The url to call when the order is updated.                                                    |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\CreateOrderResponse](../../Models/Operations/CreateOrderResponse.md)**

