package booleangroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)    

// Bool is
type Bool struct {
    autorest.Response `json:"-"`
    Value *bool `json:"value,omitempty"`
}

// Error is
type Error struct {
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

