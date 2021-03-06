package types

// SchemaConverter is available for those types which resemble
// a Schema object, such as Parameter and Items. The resulting
// Schema will not be 100% compatible as the original object as
// available fields differ slightly, but should be close enough
// to be used interchangeably.
type SchemaConverter interface {
	ConvertToSchema() (Schema, error)
}

// Validator objects can validate themselves.
type Validator interface {
	Validate(bool) error
}

// ResolveError is returned when Resolver fails to resolve a reference (`$ref`)
type ResolveError interface {
	Fatal() bool
}

// Resolver resolves JSON references (`$ref`)
type Resolver interface {
	// Resolves a JSON reference. In very rare occasions where
	// you do not want to actually resolve the value but allow
	// processing to continue, you may return an error value
	// that implements ResolveError and return false for the
	// `Fatal` method.
	Resolve(string) (interface{}, error)
}

const defaultSwaggerVersion = "2.0"
const refOnlyTmpl = `{"$ref":%s}`

type CollectionFormat string

const (
	CSV   CollectionFormat = "csv"
	SSV   CollectionFormat = "ssv"
	TSV   CollectionFormat = "tsv"
	Pipes CollectionFormat = "pipes"
	Multi CollectionFormat = "multi"
)

type Scheme = string
type MediaType = string
type MIMEType = string

const (
	WWWFormURLEncoded MIMEType = "application/x-www-form-urlencoded"
	MultiparFormDaata MIMEType = "multipart/form-data"
)

type Location string

const (
	InPath   Location = "path"
	InQuery  Location = "query"
	InHeader Location = "header"
	InBody   Location = "body"
	InForm   Location = "formData"
)

type PrimitiveType string

const (
	Invalid PrimitiveType = "invalid"
	Integer PrimitiveType = "integer"
	Number  PrimitiveType = "number"
	String  PrimitiveType = "string"
	Boolean PrimitiveType = "boolean"
	Object  PrimitiveType = "object"
	Array   PrimitiveType = "array"
	File    PrimitiveType = "file"
	Null    PrimitiveType = "null"
)

type Extensions map[string]interface{}

type OpenAPI = Swagger
type Swagger interface {
	QueryJSON(string) (interface{}, bool)
}

type swagger struct {
	version             string                  `json:"swagger" builder:"-" default:"defaultSwaggerVersion"`
	info                Info                    `json:"info" builder:"required"`
	host                string                  `json:"host,omitempty"`
	basePath            string                  `json:"basePath,omitempty"`
	schemes             SchemeList              `json:"schemes,omitempty"`
	consumes            MIMETypeList            `json:"consumes,omitempty"`
	produces            MIMETypeList            `json:"produces,omitempty"`
	paths               Paths                   `json:"paths" builder:"required"`
	definitions         InterfaceMap            `json:"definitions,omitempty"`
	parameters          ParameterMap            `json:"parameters,omitempty"`
	responses           ResponseMap             `json:"responses,omitempty"`
	securityDefinitions SecuritySchemeMap       `json:"securityDefinitions,omitempty"`
	security            SecurityRequirementList `json:"security,omitempty"`
	tags                TagList                 `json:"tags,omitempty"`
	externalDocs        ExternalDocumentation   `json:"externalDocs,omitempty"`
}

type Info interface {
}

type info struct {
	title          string  `json:"title" builder:"required"`
	version        string  `json:"version" builder:"required"`
	description    string  `json:"description,omitempty"`
	termsOfService string  `json:"termsOfService,omitempty"`
	contact        Contact `json:"contact,omitempty"`
	license        License `json:"license,omitempty"`
}

type Contact interface {
}

type contact struct {
	name  string `json:"name"`
	url   string `json:"url"`
	email string `json:"email"`
}

type License interface {
}

type license struct {
	name string `json:"name" builder:"required"`
	url  string `json:"url"`
}

type Paths interface {
}

type paths struct {
	paths PathItemMap `json:"-" mutator:"-"`
}

type PathItem interface {
	//gen:lazy Operations() *OperationListIterator
	//gen:lazy setName(string)
	//gen:lazy setPath(string)
}

type pathItem struct {
	name string `json:"-" builder:"-" mutator:"-" resolve:"-"`
	path string `json:"-" builder:"-" mutator:"-" resolve:"-"`
	// reference string `json:"$ref"`
	get        Operation     `json:"get,omitempty" builder:"-" mutator:"-"`
	put        Operation     `json:"put,omitempty" builder:"-" mutator:"-"`
	post       Operation     `json:"post,omitempty" builder:"-" mutator:"-"`
	delete     Operation     `json:"delete,omitempty" builder:"-" mutator:"-"`
	options    Operation     `json:"options,omitempty" builder:"-" mutator:"-"`
	head       Operation     `json:"head,omitempty" builder:"-" mutator:"-"`
	patch      Operation     `json:"patch,omitempty" builder:"-" mutator:"-"`
	parameters ParameterList `json:"parameters,omitempty"`
}

type Operation interface {
	//gen:lazy setPathItem(PathItem)
	//gen:lazy setVerb(string)
}

type operation struct {
	verb         string                  `json:"-" builder:"-" mutator:"-" resolve:"-"`
	pathItem     PathItem                `json:"-" builder:"-" mutator:"-" resolve:"-"`
	tags         StringList              `json:"tags,omitempty"`
	summary      string                  `json:"summary,omitempty"`
	description  string                  `json:"description,omitempty"`
	externalDocs ExternalDocumentation   `json:"externalDocs,omitempty"`
	operationID  string                  `json:"operationId,omitempty"`
	consumes     MIMETypeList            `json:"consumes,omitempty"`
	produces     MIMETypeList            `json:"produces,omitempty"`
	parameters   ParameterList           `json:"parameters,omitempty"`
	responses    Responses               `json:"responses" builder:"required"`
	schemes      SchemeList              `json:"schemes,omitempty"`
	deprecated   bool                    `json:"deprecated,omitempty"`
	security     SecurityRequirementList `json:"security,omitempty"`
}

type ExternalDocumentation interface {
}

type externalDocumentation struct {
	url         string `json:"url" builder:"required"`
	description string `json:"description,omitempty"`
}

type Parameter interface {
	SchemaConverter
}

type parameter struct {
	name        string   `json:"name" builder:"required"`
	description string   `json:"description,omitempty"`
	required    bool     `json:"required,omitempty"`
	in          Location `json:"in" builder:"required"`
	// Only applicable if when in == body
	schema Schema `json:"schema,omitempty"`
	// Only applicable if when in != body
	typ              PrimitiveType    `json:"type,omitempty"`
	format           string           `json:"format,omitempty"`
	title            string           `json:"title,omitempty"`
	allowEmptyValue  bool             `json:"allowEmptyValue,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          *float64         `json:"maximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMaximum *float64         `json:"exclusiveMaximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minimum          *float64         `json:"minimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMinimum *float64         `json:"exclusiveMinimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	maxLength        *int             `json:"maxLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minLength        *int             `json:"minLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         *int             `json:"maxItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minItems         *int             `json:"minItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	uniqueItems      *bool            `json:"uniqueItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       *float64         `json:"multipleOf,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
}

type Items interface {
	SchemaConverter
}

type items struct {
	typ              PrimitiveType    `json:"type,omitempty"`
	format           string           `json:"format,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          *float64         `json:"maximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMaximum *float64         `json:"exclusiveMaximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minimum          *float64         `json:"minimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMinimum *float64         `json:"exclusiveMinimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	maxLength        *int             `json:"maxLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minLength        *int             `json:"minLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         *int             `json:"maxItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minItems         *int             `json:"minItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	uniqueItems      *bool            `json:"uniqueItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       *float64         `json:"multipleOf,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
}

type Responses interface {
}

type responses struct {
	defaultValue Response    `json:"default,omitempty"`
	responses    ResponseMap `json:"-" mutator:"-" visit:"true"`
}

type Response interface {
	//gen:lazy setStatusCode(string)
}

type response struct {
	name        string     `json:"-"`
	statusCode  string     `json:"-"`
	description string     `json:"description" builder:"required"`
	schema      Schema     `json:"schema,omitempty"`
	headers     HeaderMap  `json:"headers,omitempty"`
	examples    ExampleMap `json:"example,omitempty"`
}

type Header interface {
}

type header struct {
	name             string           `json:"-"`
	description      string           `json:"description,omitempty"`
	typ              string           `json:"type" builder:"required"`
	format           string           `json:"format,omitempty"`
	items            Items            `json:"items,omitempty"`
	collectionFormat CollectionFormat `json:"collectionFormat,omitempty"`
	defaultValue     interface{}      `json:"default,omitempty"`
	maximum          float64          `json:"maximum,omitempty"`
	exclusiveMaximum float64          `json:"exclusiveMaximum,omitempty"`
	minimum          float64          `json:"minimum,omitempty"`
	exclusiveMinimum float64          `json:"exclusiveMinimum,omitempty"`
	maxLength        int              `json:"maxLength,omitempty"`
	minLength        int              `json:"minLength,omitempty"`
	pattern          string           `json:"pattern,omitempty"`
	maxItems         int              `json:"maxItems,omitempty"`
	minItems         int              `json:"minItems,omitempty"`
	uniqueItems      bool             `json:"uniqueItems,omitempty"`
	enum             InterfaceList    `json:"enum,omitempty"`
	multipleOf       float64          `json:"multipleOf,omitempty"`
}

type Schema interface {
	IsRequiredProperty(string) bool
	IsValid() bool
	SchemaConverter
}

type schema struct {
	name                string                `json:"-" builder:"-"` // This is only populated when applicable
	typ                 PrimitiveType         `json:"type,omitempty"`
	format              string                `json:"format,omitempty"`
	title               string                `json:"title,omitempty"`
	multipleOf          *float64              `json:"multipleOf,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	maximum             *float64              `json:"maximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMaximum    *float64              `json:"exclusiveMaximum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minimum             *float64              `json:"minimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	exclusiveMinimum    *float64              `json:"exclusiveMinimum,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	maxLength           *int                  `json:"maxLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minLength           *int                  `json:"minLength,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	pattern             string                `json:"pattern,omitempty"`
	maxItems            *int                  `json:"maxItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minItems            *int                  `json:"minItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	uniqueItems         *bool                 `json:"uniqueItems,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	maxProperties       *int                  `json:"maxProperties,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	minProperties       *int                  `json:"minProperties,omitempty" accessor:"indirect" builder:"indirect" mutator:"indirect"`
	required            StringList            `json:"required,omitempty"`
	enum                InterfaceList         `json:"enum,omitempty"`
	allOf               SchemaList            `json:"allOf,omitempty"`
	items               Schema                `json:"items,omitempty"`
	properties          SchemaMap             `json:"properties,omitempty" builder:"-"`
	additionaProperties SchemaMap             `json:"additionalProperties,omitempty"`
	defaultValue        interface{}           `json:"default,omitempty"`
	discriminator       string                `json:"discriminator,omitempty"`
	readOnly            bool                  `json:"readOnly,omitempty"`
	externalDocs        ExternalDocumentation `json:"externalDocs,omitempty"`
	example             interface{}           `json:"example,omitempty"`
	deprecated          bool                  `json:"deprecated,omitempty"`
	xml                 XML                   `json:"xml,omitempty"`
}

type XML interface {
}

type xml struct {
	name      string `json:"name,omitempty"`
	namespace string `json:"namespace,omitempty"`
	prefix    string `json:"prefix,omitempty"`
	attribute bool   `json:"attribute,omitempty"`
	wrapped   bool   `json:"wrapped,omitempty"`
}

type SecurityScheme interface {
}

type securityScheme struct {
	typ              string    `json:"type" builder:"required"`
	description      string    `json:"description,omitempty"`
	name             string    `json:"name,omitempty"`
	in               string    `json:"in,omitempty"`
	flow             string    `json:"flow,omitempty"`
	authorizationURL string    `json:"authorizationUrl,omitempty"`
	tokenURL         string    `json:"tokenUrl,omitempty"`
	scopes           StringMap `json:"scopes,omitempty"`
}

type SecurityRequirement interface {
}

type securityRequirement struct {
	name   string    `json:"-"`
	scopes ScopesMap `json:"-" mutator:"-" builder:"-"`
}

type Tag interface {
}

type tag struct {
	name         string                `json:"name" builder:"required"`
	description  string                `json:"description,omitempty"`
	externalDocs ExternalDocumentation `json:"externalDocs,omitempty"`
}

type ExampleMapKey = string
type ExampleMap map[ExampleMapKey]interface{}
type HeaderMapKey = string
type HeaderMap map[HeaderMapKey]Header
type InterfaceList []interface{}
type InterfaceMapKey = string
type InterfaceMap map[InterfaceMapKey]interface{}
type MIMETypeList []MIMEType
type ParameterList []Parameter
type ParameterMapKey = string
type ParameterMap map[ParameterMapKey]Parameter
type PathItemMapKey = string
type PathItemMap map[PathItemMapKey]PathItem
type ResponseMapKey = string
type ResponseMap map[ResponseMapKey]Response
type SchemeList []string
type SchemaList []Schema
type SchemaMapKey = string
type SchemaMap map[SchemaMapKey]Schema
type ScopesMapKey = string
type ScopesMap map[ScopesMapKey][]string
type SecurityRequirementList []SecurityRequirement
type SecuritySchemeMapKey = string
type SecuritySchemeMap map[SecuritySchemeMapKey]SecurityScheme
type StringList []string
type StringMapKey = string
type StringMap map[StringMapKey]string
type TagList []Tag
