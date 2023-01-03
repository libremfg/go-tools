package graphql

import (
	"sort"
	"time"
)

// The Response from the GraphQL Server
type Response struct {
	Errors     []GraphQLError `json:"errors,omitempty"`
	Data       Data           `json:"data"`
	Extensions Extensions     `json:"extensions"`
}

// GraphQL Error
type GraphQLError struct {
	Message   string     `json:"message"`
	Locations []Location `json:"location,omitempty"`
	Path      []string   `json:"path,omitempty"`
}

// Location of the graphQL Error in the requested document
type Location struct {
	Line   int64 `json:"line"`
	Column int64 `json:"column"`
}

// Sort the Response
func (r *Response) Sort() {
	r.Data.Schema.Sort()
}

// The Data response
type Data struct {
	Schema  __Schema `json:"__schema,omitempty"`
	Service _Service `json:"_service,omitempty"`
}

type _Service struct {
	SDL *string `json:"sdl,omitempty"`
}

// The schema introspection system is accessible from the meta‐fields __schema
// and __type which are accessible from the type of the root of a query
// operation.
// See https://spec.graphql.org/June2018/#sec-Schema-Introspection
type __Schema struct {
	Types            []__ObjectType `json:"types"`
	QueryType        SchemaType     `json:"queryType"`
	MutationType     SchemaType     `json:"mutationType"`
	SubscriptionType SchemaType     `json:"subscriptionType"`
	Directives       []__Directive  `json:"directives"`
}

// Sort the child arrays of Schema
func (s *__Schema) Sort() {

	sort.Slice(s.Directives, func(a, b int) bool {
		return s.Directives[a].Name < s.Directives[b].Name
	})
	for i := range s.Directives {
		sort.Slice(s.Directives[i].Args, func(j, k int) bool {
			return s.Directives[i].Args[j].Name < s.Directives[i].Args[k].Name
		})

		sort.Slice(s.Directives[i].Locations, func(j, k int) bool {
			return string(s.Directives[i].Locations[j]) < string(s.Directives[i].Locations[k])
		})
	}

	sort.Slice(s.Types, func(a, b int) bool {
		return s.Types[a].Name < s.Types[b].Name
	})

	for i := range s.Types {
		s.Types[i].Sort()
	}
}

// __Type is at the core of the type introspection system. It represents
// scalars, interfaces, object types, unions, enums in the system.
//
// __Type also represents type modifiers, which are used to modify a type that
// it refers to (ofType: __Type). This is how we represent lists, non‐nullable
// types, and the combinations thereof.
//
// See https://spec.graphql.org/June2018/#sec-The-__Type-Type
type __Type struct {
	Kind        __TypeKind `json:"kind"`
	Name        string     `json:"name"`
	Description string     `json:"description"`

	// OBJECT and INTERFACE only
	Fields []__Field `json:"fields"`

	// OBJECT only
	Interfaces []__Type `json:"interfaces"`

	// INTERFACE and UNION only
	PossibleTypes []__Type `json:"possibleTypes"`

	// ENUM only
	EnumValues []EnumValue `json:"enumValues"`

	// INPUT_OBJECT only
	InputFields []__InputValue `json:"inputFields"`

	// NON_NULL and LIST only
	OfType *__Type `json:"ofType,omitempty"`
}

// Sort the child arrays of __Type
func (t *__Type) Sort() {
	sort.Slice(t.Fields, func(i, j int) bool {
		return t.Fields[i].Name < t.Fields[j].Name
	})
	for i := range t.Fields {
		sort.Slice(t.Fields[i].Args, func(j, k int) bool {
			return t.Fields[i].Args[j].Name < t.Fields[i].Args[k].Name
		})
	}

	sort.Slice(t.Interfaces, func(i, j int) bool {
		return t.Interfaces[i].Name < t.Interfaces[j].Name
	})
	for i := range t.Interfaces {
		t.Interfaces[i].Sort()
	}

	sort.Slice(t.PossibleTypes, func(i, j int) bool {
		return t.PossibleTypes[i].Name < t.PossibleTypes[j].Name
	})
	for i := range t.PossibleTypes {
		t.PossibleTypes[i].Sort()
	}

	sort.Slice(t.EnumValues, func(i, j int) bool {
		return t.EnumValues[i].Name < t.EnumValues[j].Name
	})

	sort.Slice(t.InputFields, func(i, j int) bool {
		return t.InputFields[i].Name < t.InputFields[j].Name
	})
}

// Similar to __Type but without OfType
type __ObjectType struct {
	Kind        __TypeKind `json:"kind"`
	Name        string     `json:"name"`
	Description string     `json:"description"`

	// OBJECT and INTERFACE only
	Fields []__Field `json:"fields"`

	// OBJECT only
	Interfaces []__Type `json:"interfaces"`

	// INTERFACE and UNION only
	PossibleTypes []__Type `json:"possibleTypes"`

	// ENUM only
	EnumValues []EnumValue `json:"enumValues"`

	// INPUT_OBJECT only
	InputFields []__InputValue `json:"inputFields"`
}

// Sort the child arrays of __ObjectType
func (t *__ObjectType) Sort() {
	sort.Slice(t.Fields, func(i, j int) bool {
		return t.Fields[i].Name < t.Fields[j].Name
	})
	for i := range t.Fields {
		sort.Slice(t.Fields[i].Args, func(j, k int) bool {
			return t.Fields[i].Args[j].Name < t.Fields[i].Args[k].Name
		})
	}

	sort.Slice(t.Interfaces, func(i, j int) bool {
		return t.Interfaces[i].Name < t.Interfaces[j].Name
	})
	for i := range t.Interfaces {
		t.Interfaces[i].Sort()
	}

	sort.Slice(t.PossibleTypes, func(i, j int) bool {
		return t.PossibleTypes[i].Name < t.PossibleTypes[j].Name
	})
	for i := range t.PossibleTypes {
		t.PossibleTypes[i].Sort()
	}

	sort.Slice(t.EnumValues, func(i, j int) bool {
		return t.EnumValues[i].Name < t.EnumValues[j].Name
	})

	sort.Slice(t.InputFields, func(i, j int) bool {
		return t.InputFields[i].Name < t.InputFields[j].Name
	})
}

// There are several different kinds of type. In each kind, different fields
// are actually valid. These kinds are listed in the __TypeKind enumeration.
//
// See https://spec.graphql.org/June2018/#sec-Type-Kinds
type __TypeKind string

const (
	TypeKindSCALAR       __TypeKind = "SCALAR"
	TypeKindOBJECT                  = "OBJECT"
	TypeKindINTERFACE               = "INTERFACE"
	TypeKindUNION                   = "UNION"
	TypeKindENUM                    = "ENUM"
	TypeKindINPUT_OBJECT            = "INPUT_OBJECT"
	TypeKindLIST                    = "LIST"
	TypeKindNON_NULL                = "NOT_NULL"
)

// The __Field type represents each field in an Object or Interface type.
// See https://spec.graphql.org/June2018/#sec-The-__Field-Type
type __Field struct {
	Name              string         `json:"name,omitempty"`
	Description       string         `json:"description"`
	Args              []__InputValue `json:"args"`
	Type              __Type         `json:"type"`
	IsDeprecated      bool           `json:"isDeprecated"`
	DeprecationReason *string        `json:"deprecationReason"`
}

// The __InputValue type represents field and directive arguments as well as
// the inputFields of an input object.
//
// See https://spec.graphql.org/June2018/#sec-The-__InputValue-Type
type __InputValue struct {
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Type         __Type  `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

// The __EnumValue type represents one of possible values of an enum.
// See https://spec.graphql.org/June2018/#sec-The-__EnumValue-Type
type EnumValue struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason,omitempty"`
}

// The SchemaType is a helper class that just contains Name.
type SchemaType struct {
	Name string `json:"name"`
}

// The __Directive type represents a Directive that a server supports.
//
// See https://spec.graphql.org/June2018/#sec-The-__Directive-Type
type __Directive struct {
	Args        []__InputValue        `json:"args"`
	Description string                `json:"description,omitempty"`
	Locations   []__DirectiveLocation `json:"locations"`
	Name        string                `json:"name"`
}

type __DirectiveLocation string

const (
	DirectiveLocationQUERY                  __DirectiveLocation = "QUERY"
	DirectiveLocationMUTATION                                   = "MUTATION"
	DirectiveLocationSUBSCRIPTION                               = "SUBSCRIPTION"
	DirectiveLocationFIELD                                      = "FIELD"
	DirectiveLocationFRAGMENT_DEFINITION                        = "FRAGMENT_DEFINITION"
	DirectiveLocationFRAGMENT_SPREAD                            = "FRAGMENT_SPREAD"
	DirectiveLocationINLINE_FRAGMENT                            = "INLINE_FRAGMENT"
	DirectiveLocationSCHEMA                                     = "SCHEMA"
	DirectiveLocationSCALAR                                     = "SCALAR"
	DirectiveLocationOBJECT                                     = "OBJECT"
	DirectiveLocationFIELD_DEFINITION                           = "FIELD_DEFINITION"
	DirectiveLocationARGUMENT_DEFINITION                        = "ARGUMENT_DEFINITION"
	DirectiveLocationINTERFACE                                  = "INTERFACE"
	DirectiveLocationUNION                                      = "UNION"
	DirectiveLocationENUM                                       = "ENUM"
	DirectiveLocationENUM_VALUE                                 = "ENUM_VALUE"
	DirectiveLocationINPUT_OBJECT                               = "INPUT_OBJECT"
	DirectiveLocationINPUT_FIELD_DEFINITION                     = "INPUT_FIELD_DEFINITION"
)

// Extensions
type Extensions struct {
	Tracing *Tracing `json:"tracing,omitempty"`
}

// Tracing object
type Tracing struct {
	Version   uint      `json:"version"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Duration  uint64    `json:"duration"`
}
