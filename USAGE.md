<!-- Start SDK Example Usage [usage] -->
### Sign in

First you need to send an authentication request to the API by providing your username and password.
In the request body, you should specify the type of token you would like to receive: API key or JSON Web Token.
If your credentials are valid, you will receive a token in the response object: `res.object.token: str`.

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$sdk = OpenAPI\Speakeasy::builder()->build();

try {
    $request = new Operations\LoginRequestBody();
    $request->type = Operations\Type::ApiKey;

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

### Browse available drinks

Once you are authenticated, you can use the token you received for all other authenticated endpoints.
For example, you can filter the list of available drinks by type.

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$sdk = OpenAPI\Speakeasy::builder()->build();

try {
    $requestSecurity = new Operations\ListDrinksSecurity();
    $requestSecurity->bearerAuth = '<YOUR_JWT>';

    $response = $sdk->drinks->listDrinks($requestSecurity, Components\DrinkType::Spirit);

    if ($response->classes !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}

```

### Create an order

When you submit an order, you can include a callback URL along your request.
This URL will get called whenever the supplier updates the status of your order.

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '<YOUR_API_KEY>';

$sdk = OpenAPI\Speakeasy::builder()
    ->setSecurity($security)
    ->build();

try {
    $response = $sdk->orders->createOrder([new Components\OrderInput()], '<value>');

    if ($response->order !== null) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}

```

### Subscribe to webhooks to receive stock updates

```php
<?php

declare(strict_types=1);

require 'vendor/autoload.php';

use OpenAPI\OpenAPI;
use OpenAPI\OpenAPI\Models\Components;
use OpenAPI\OpenAPI\Models\Operations;

$security = new Components\Security();
$security->apiKey = '<YOUR_API_KEY>';

$sdk = OpenAPI\Speakeasy::builder()
    ->setSecurity($security)
    ->build();

try {
    $request = [new Operations\RequestBody()];

    $response = $sdk->config->subscribeToWebhooks($request);

    if ($response->statusCode === 200) {
        // handle response
    }
} catch (Throwable $e) {
    // handle exception
}

```
<!-- End SDK Example Usage [usage] -->