<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace OpenAPI\OpenAPI;

class SDKConfiguration
{
	public ?\GuzzleHttp\ClientInterface $defaultClient = null;
	public ?\GuzzleHttp\ClientInterface $securityClient = null;
	public ?Models\Components\Security $security = null;
	public string $serverUrl = '';
	public string $server = '';
	/** @var array<string, array<string, string>> */
	public ?array $serverDefaults = [
		'prod' => [
		],
		'staging' => [
		],
		'customer' => [
			'environment' => 'prod',
			'organization' => 'api',
		],
	];
	public string $language = 'php';
	public string $openapiDocVersion = '1.0.0';
	public string $sdkVersion = '0.1.2';
	public string $genVersion = '2.253.0';
	public string $userAgent = 'speakeasy-sdk/php 0.1.2 2.253.0 1.0.0 openapi/openapi';
	

	public function getServerUrl(): string
	{
		
		if ($this->serverUrl !== '') {
			return $this->serverUrl;
		}
		if ($this->server === '') {
			$this->server = Speakeasy::SERVER_PROD;
		}

		return Speakeasy::SERVERS[$this->server];
	}
	
	/**
	 * @return array<string, string>
	 */
	public function getServerDefaults(): ?array
	{
		if ($this->server === '') {
			$this->server = Speakeasy::SERVER_PROD;
		}

		return $this->serverDefaults[$this->server];
	}
}