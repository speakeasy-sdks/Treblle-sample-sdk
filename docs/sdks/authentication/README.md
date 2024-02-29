# Authentication


## Overview

The authentication endpoints.

### Available Operations

* [login](#login) - Authenticate with the API by providing a username and password.

## login

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
        $request = new Operations\LoginRequestBody();
    $request->type = Operations\Type::ApiKey;;

    $requestSecurity = new Operations\LoginSecurity();
    $requestSecurity->password = '<PASSWORD>';
    $requestSecurity->username = '<USERNAME>';

    $response = $sdk->authentication->login($request, $requestSecurity);

    if ($response->object !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `$request`                                                                                         | [\OpenAPI\OpenAPI\Models\Operations\LoginRequestBody](../../Models/Operations/LoginRequestBody.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `security`                                                                                         | [\OpenAPI\OpenAPI\Models\Operations\LoginSecurity](../../Models/Operations/LoginSecurity.md)       | :heavy_check_mark:                                                                                 | The security requirements to use for the request.                                                  |


### Response

**[?\OpenAPI\OpenAPI\Models\Operations\LoginResponse](../../Models/Operations/LoginResponse.md)**

