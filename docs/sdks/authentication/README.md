# Authentication


## Overview

The authentication endpoints.

### Available Operations

* [authenticate](#authenticate) - Authenticate with the API by providing a username and password.

## authenticate

Authenticate with the API by providing a username and password.

### Example Usage

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use \OpenAPI\OpenAPI;
use \OpenAPI\OpenAPI\Models\Components;
use \OpenAPI\OpenAPI\Models\Operations;

$sdk = OpenAPI\Speakeasy::builder()->build();

try {
        $request = new Operations\AuthenticateRequestBody();
    $request->password = 'Nxq_X5HXg1lXJa5';
    $request->username = 'Asa_Stamm77';;

    $response = $sdk->authentication->authenticate($request);

    if ($response->object !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `$request`                                                                                                       | [\OpenAPI\OpenAPI\Models\Operations\AuthenticateRequestBody](../../Models/Operations/AuthenticateRequestBody.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\AuthenticateResponse](../../Models/Operations/AuthenticateResponse.md)**

