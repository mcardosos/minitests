package numbergroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/shopspring/decimal"
)    

// Decimal is
type Decimal struct {
    autorest.Response `json:"-"`
    Value *decimal.Decimal `json:"value,omitempty"`
}

// Error is
type Error struct {
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

// Float64 is
type Float64 struct {
    autorest.Response `json:"-"`
    Value *float64 `json:"value,omitempty"`
}

