package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Pet5 represents the Pet5 schema from the OpenAPI specification
type Pet5 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet6 represents the Pet6 schema from the OpenAPI specification
type Pet6 struct {
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// Pet16 represents the Pet16 schema from the OpenAPI specification
type Pet16 struct {
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// Pet1 represents the Pet1 schema from the OpenAPI specification
type Pet1 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet3 represents the Pet3 schema from the OpenAPI specification
type Pet3 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet8 represents the Pet8 schema from the OpenAPI specification
type Pet8 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet15 represents the Pet15 schema from the OpenAPI specification
type Pet15 struct {
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// Pet17 represents the Pet17 schema from the OpenAPI specification
type Pet17 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet4 represents the Pet4 schema from the OpenAPI specification
type Pet4 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet18 represents the Pet18 schema from the OpenAPI specification
type Pet18 struct {
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// Pet10 represents the Pet10 schema from the OpenAPI specification
type Pet10 struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// Pet11 represents the Pet11 schema from the OpenAPI specification
type Pet11 struct {
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}

// Pet14 represents the Pet14 schema from the OpenAPI specification
type Pet14 struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// Pet19 represents the Pet19 schema from the OpenAPI specification
type Pet19 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet25 represents the Pet25 schema from the OpenAPI specification
type Pet25 struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// Pet20 represents the Pet20 schema from the OpenAPI specification
type Pet20 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet21 represents the Pet21 schema from the OpenAPI specification
type Pet21 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet23 represents the Pet23 schema from the OpenAPI specification
type Pet23 struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// Pet22 represents the Pet22 schema from the OpenAPI specification
type Pet22 struct {
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// Pet represents the Pet schema from the OpenAPI specification
type Pet struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet2 represents the Pet2 schema from the OpenAPI specification
type Pet2 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet9 represents the Pet9 schema from the OpenAPI specification
type Pet9 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet13 represents the Pet13 schema from the OpenAPI specification
type Pet13 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet7 represents the Pet7 schema from the OpenAPI specification
type Pet7 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet12 represents the Pet12 schema from the OpenAPI specification
type Pet12 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}

// Pet24 represents the Pet24 schema from the OpenAPI specification
type Pet24 struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Tag string `json:"tag,omitempty"`
}
