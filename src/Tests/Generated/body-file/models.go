package filegroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "io"
)    

// Error is
type Error struct {
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

// ReadCloser is
type ReadCloser struct {
    autorest.Response `json:"-"`
    Value *io.ReadCloser `json:"value,omitempty"`
}
