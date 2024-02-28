package function

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/appengine/v1"
	"google.golang.org/api/option"
)

func init() {
	functions.HTTP("getInstances", getInstances)
}

func getInstances(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Get the "appId" query parameter from the URL
	var appID = r.URL.Query().Get("appId")

	// Create a new App Engine service client
	appengineService, err := appengine.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not create App Engine service client: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the list of services for the app
	servicesList, err := appengineService.Apps.Services.List(appID).Do()
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not list services: %v", err), http.StatusInternalServerError)
		return
	}

	// Prepare a slice to hold all instances
	var allInstances []*appengine.Instance

	// Iterate over each service
	for _, service := range servicesList.Services {
		// Get the list of versions for the service
		versionsList, err := appengineService.Apps.Services.Versions.List(appID, service.Id).Do()
		if err != nil {
			fmt.Printf("Could not list versions for service %s: %v\n", service.Id, err)
			continue
		}

		// Iterate over each version
		for _, version := range versionsList.Versions {
			// Get the list of instances for the version
			instancesList, err := appengineService.Apps.Services.Versions.Instances.List(appID, service.Id, version.Id).Do()
			if err != nil {
				fmt.Printf("Could not list instances for version %s of service %s: %v\n", version.Id, service.Id, err)
				continue
			}

			// Append the instances to the allInstances slice
			allInstances = append(allInstances, instancesList.Instances...)
		}
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode the allInstances slice to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(allInstances); err != nil {
		http.Error(w, fmt.Sprintf("Could not encode instances to JSON: %v", err), http.StatusInternalServerError)
	}
}