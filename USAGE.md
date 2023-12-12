<!-- Start SDK Example Usage [usage] -->
```php
<?php

declare(strict_types=1);
require_once 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$sdk = OpenAPI\Speakeasy::builder()->build();

try {
    $response = $sdk->drinks->listDrinks(Components\DrinkType::Spirit);

    if ($response->classes !== null) {
        // handle response
    }
} catch (Exception $e) {
    // handle exception
}

```
<!-- End SDK Example Usage [usage] -->