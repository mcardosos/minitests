// Package durationgroup implements the Azure ARM Durationgroup service API
// version 1.0.0.
// 
// Test Infrastructure for AutoRest
package durationgroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

const (
    // APIVersion is the version of the Durationgroup
    APIVersion = "1.0.0"

    // DefaultBaseURI is the default URI used for the service Durationgroup
    DefaultBaseURI = "https://localhost"
)

// ManagementClient is the base client for Durationgroup.
type ManagementClient struct {
    autorest.Client
    BaseURI string
    APIVersion string
}

// New creates an instance of the ManagementClient client.
func New() ManagementClient {
    return ManagementClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: DefaultBaseURI,
        APIVersion: APIVersion,
    }
}

