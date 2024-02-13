# Drinks


## Overview

The drinks endpoints.

### Available Operations

* [getDrink](#getdrink) - Get a drink.
* [listDrinks](#listdrinks) - Get a list of drinks.

## getDrink

Get a drink by name, if authenticated this will include stock levels and product codes otherwise it will only include public information.

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
    

    $response = $sdk->drinks->getDrink('string');

    if ($response->drink !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter          | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `name`             | *string*           | :heavy_check_mark: | N/A                |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\GetDrinkResponse](../../Models/Operations/GetDrinkResponse.md)**


## listDrinks

Get a list of drinks, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```php
<?php

declare(strict_types=1);
require_once 'vendor/autoload.php';

use \OpenAPI\OpenAPI;
use \OpenAPI\OpenAPI\Models\Components;
use \OpenAPI\OpenAPI\Models\Operations;

$sdk = OpenAPI\Speakeasy::builder()->build();

try {
    

    $response = $sdk->drinks->listDrinks(Components\DrinkType::Spirit);

    if ($response->classes !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `drinkType`                                                                          | [\OpenAPI\OpenAPI\Models\Components\DrinkType](../../Models/Components/DrinkType.md) | :heavy_minus_sign:                                                                   | The type of drink to filter by. If not provided all drinks will be returned.         |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\ListDrinksResponse](../../Models/Operations/ListDrinksResponse.md)**

