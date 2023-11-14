<!-- Start SDK Example Usage -->
```php
<?php

declare(strict_types=1);
require_once 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '';

$sdk = OpenAPI\Speakeasy::builder()
    ->setSecurity($security)
    ->build();

try {
    $response = $sdk->drinks->listDrinks(Components\DrinkType::Spirit);

    if ($response->classes !== null) {
        // handle response
    }
} catch (Exception $e) {
    // handle exception
}

```
<!-- End SDK Example Usage -->