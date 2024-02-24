# Config


### Available Operations

* [subscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## subscribeToWebhooks

Subscribe to webhooks.

### Example Usage

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use \OpenAPI\OpenAPI;
use \OpenAPI\OpenAPI\Models\Components;
use \OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '<YOUR_API_KEY_HERE>';

$sdk = OpenAPI\Speakeasy::builder()->setSecurity($security)->build();

try {
        $request = [
        new Operations\RequestBody(),
    ];

    $response = $sdk->config->subscribeToWebhooks($request);

    if ($response->statusCode === 200) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter                                  | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `$request`                                 | [array](../../.md)                         | :heavy_check_mark:                         | The request object to use for the request. |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\SubscribeToWebhooksResponse](../../Models/Operations/SubscribeToWebhooksResponse.md)**

