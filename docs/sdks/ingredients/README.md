# Ingredients


## Overview

The ingredients endpoints.

### Available Operations

* [listIngredients](#listingredients) - Get a list of ingredients.

## listIngredients

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```php
<?php

declare(strict_types=1);
require_once 'vendor/autoload.php';

use \OpenAPI\OpenAPI;
use \OpenAPI\OpenAPI\Models\Components;
use \OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '<YOUR_API_KEY_HERE>';

$sdk = OpenAPI\Speakeasy::builder()->setSecurity($security)->build();

try {
    

    $response = $sdk->ingredients->listIngredients([
    '<value>',
]);

    if ($response->classes !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ingredients`                                                                         | array<*string*>                                                                       | :heavy_minus_sign:                                                                    | A list of ingredients to filter by. If not provided all ingredients will be returned. |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\ListIngredientsResponse](../../Models/Operations/ListIngredientsResponse.md)**

