package datetimerfc1123group

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/date"
)    

// Error is
type Error struct {
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

// TimeRFC1123 is
type TimeRFC1123 struct {
    autorest.Response `json:"-"`
    Value *date.TimeRFC1123 `json:"value,omitempty"`
}

