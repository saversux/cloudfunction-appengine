# App Engine Instances Fetcher

This Go application is designed to fetch and list instances of services running on Google App Engine.

## Overview

The main function, `getInstances`, is an HTTP-triggered function that retrieves the instances of all services and versions in a specified Google App Engine application.

## How it works

1. The function first retrieves the "appId" from the query parameters of the incoming HTTP request.
2. It then creates a new App Engine service client without authentication.
3. Using this client, it fetches the list of services for the specified app.
4. For each service, it fetches the list of versions.
5. For each version, it fetches the list of instances.
6. Finally, it writes the details of each instance to the HTTP response.

## Dependencies

This application uses the following packages:

- `context` for managing multiple request contexts.
- `fmt` for formatting text for output.
- `net/http` for handling HTTP requests.
- `github.com/GoogleCloudPlatform/functions-framework-go/functions` for creating HTTP-triggered functions.
- `google.golang.org/api/appengine/v1` for interacting with the App Engine Admin API.
- `google.golang.org/api/option` for client options like authentication.

## Usage

To use this function, deploy it to a Google Cloud Function and trigger it with an HTTP request. The "appId" query parameter should be set to the ID of the App Engine application you want to inspect.

For example:
https://REGION-PROJECT_ID.cloudfunctions.net/getInstances?appId=YOUR_APP_ID

Please replace REGION, PROJECT_ID, and YOUR_APP_ID with your actual Google Cloud region, project ID, and App Engine application ID respectively.

## Disclaimer

This README.md was generated with the assistance of GitHub Copilot, an AI programming assistant. While every effort has been made to ensure the accuracy and completeness of this documentation, it is provided "as is" without warranty of any kind, express or implied.