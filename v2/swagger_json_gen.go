package openapi

// This file was automatically generated by gentyeps.go on 2018-05-28T21:36:31+09:00
// DO NOT EDIT MANUALLY. All changes will be lost

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

var _ = json.Unmarshal
var _ = log.Printf
var _ = errors.Cause

type swaggerMarshalProxy struct {
	Reference           string                  `json:"$ref,omitempty"`
	Version             string                  `json:"swagger"`
	Info                Info                    `json:"info"`
	Host                string                  `json:"host,omitempty"`
	BasePath            string                  `json:"basePath,omitempty"`
	Schemes             SchemeList              `json:"schemes,omitempty"`
	Consumes            MIMETypeList            `json:"consumes,omitempty"`
	Produces            MIMETypeList            `json:"produces,omitempty"`
	Paths               Paths                   `json:"paths"`
	Definitions         SchemaMap               `json:"definitions,omitempty"`
	Parameters          ParameterMap            `json:"parameters,omitempty"`
	Responses           ResponseMap             `json:"responses,omitempty"`
	SecurityDefinitions SecuritySchemeMap       `json:"securityDefinitions,omitempty"`
	Security            SecurityRequirementList `json:"security,omitempty"`
	Tags                TagList                 `json:"tags,omitempty"`
	ExternalDocs        ExternalDocumentation   `json:"externalDocs,omitempty"`
}

func (v *swagger) MarshalJSON() ([]byte, error) {
	var proxy swaggerMarshalProxy
	if s := v.reference; len(s) > 0 {
		return []byte(fmt.Sprintf(refOnlyTmpl, strconv.Quote(s))), nil
	}
	proxy.Version = v.version
	proxy.Info = v.info
	proxy.Host = v.host
	proxy.BasePath = v.basePath
	proxy.Schemes = v.schemes
	proxy.Consumes = v.consumes
	proxy.Produces = v.produces
	proxy.Paths = v.paths
	proxy.Definitions = v.definitions
	proxy.Parameters = v.parameters
	proxy.Responses = v.responses
	proxy.SecurityDefinitions = v.securityDefinitions
	proxy.Security = v.security
	proxy.Tags = v.tags
	proxy.ExternalDocs = v.externalDocs
	buf, err := json.Marshal(proxy)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal struct`)
	}
	if len(v.extras) > 0 {
		extrasBuf, err := json.Marshal(v.extras)
		if err != nil || len(extrasBuf) <= 2 {
			return nil, errors.Wrap(err, `failed to marshal struct (extras)`)
		}
		buf = append(append(buf[:len(buf)-1], ','), extrasBuf[1:]...)
	}
	return buf, nil
}

func (v *swagger) UnmarshalJSON(data []byte) error {
	var proxy map[string]json.RawMessage
	if err := json.Unmarshal(data, &proxy); err != nil {
		return err
	}
	if raw := proxy["$ref"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.reference); err != nil {
			return errors.Wrap(err, `failed to unmarshal $ref`)
		}
		return nil
	}

	if raw := proxy["swagger"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.version); err != nil {
			return errors.Wrap(err, `failed to unmarshal field swagger`)
		}
		delete(proxy, "swagger")
	}

	if raw := proxy["info"]; len(raw) > 0 {
		var decoded info
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Info`)
		}

		v.info = &decoded
		delete(proxy, "info")
	}

	if raw := proxy["host"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.host); err != nil {
			return errors.Wrap(err, `failed to unmarshal field host`)
		}
		delete(proxy, "host")
	}

	if raw := proxy["basePath"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.basePath); err != nil {
			return errors.Wrap(err, `failed to unmarshal field basePath`)
		}
		delete(proxy, "basePath")
	}

	if raw := proxy["schemes"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.schemes); err != nil {
			return errors.Wrap(err, `failed to unmarshal field schemes`)
		}
		delete(proxy, "schemes")
	}

	if raw := proxy["consumes"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.consumes); err != nil {
			return errors.Wrap(err, `failed to unmarshal field consumes`)
		}
		delete(proxy, "consumes")
	}

	if raw := proxy["produces"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.produces); err != nil {
			return errors.Wrap(err, `failed to unmarshal field produces`)
		}
		delete(proxy, "produces")
	}

	if raw := proxy["paths"]; len(raw) > 0 {
		var decoded paths
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field Paths`)
		}

		v.paths = &decoded
		delete(proxy, "paths")
	}

	if raw := proxy["definitions"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.definitions); err != nil {
			return errors.Wrap(err, `failed to unmarshal field definitions`)
		}
		delete(proxy, "definitions")
	}

	if raw := proxy["parameters"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.parameters); err != nil {
			return errors.Wrap(err, `failed to unmarshal field parameters`)
		}
		delete(proxy, "parameters")
	}

	if raw := proxy["responses"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.responses); err != nil {
			return errors.Wrap(err, `failed to unmarshal field responses`)
		}
		delete(proxy, "responses")
	}

	if raw := proxy["securityDefinitions"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.securityDefinitions); err != nil {
			return errors.Wrap(err, `failed to unmarshal field securityDefinitions`)
		}
		delete(proxy, "securityDefinitions")
	}

	if raw := proxy["security"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.security); err != nil {
			return errors.Wrap(err, `failed to unmarshal field security`)
		}
		delete(proxy, "security")
	}

	if raw := proxy["tags"]; len(raw) > 0 {
		if err := json.Unmarshal(raw, &v.tags); err != nil {
			return errors.Wrap(err, `failed to unmarshal field tags`)
		}
		delete(proxy, "tags")
	}

	if raw := proxy["externalDocs"]; len(raw) > 0 {
		var decoded externalDocumentation
		if err := json.Unmarshal(raw, &decoded); err != nil {
			return errors.Wrap(err, `failed to unmarshal field ExternalDocs`)
		}

		v.externalDocs = &decoded
		delete(proxy, "externalDocs")
	}

	for name, raw := range proxy {
		if strings.HasPrefix(name, `x-`) {
			if v.extras == nil {
				v.extras = Extensions{}
			}
			v.extras[name] = raw
		}
	}
	return nil
}

func (v *swagger) Resolve(resolver *Resolver) error {
	if v.IsUnresolved() {

		resolved, err := resolver.Resolve(v.Reference())
		if err != nil {
			return errors.Wrapf(err, `failed to resolve reference %s`, v.Reference())
		}
		asserted, ok := resolved.(*swagger)
		if !ok {
			return errors.Wrapf(err, `expected resolved reference to be of type Swagger, but got %T`, resolved)
		}
		mutator := MutateSwagger(v)
		mutator.Version(asserted.Version())
		mutator.Info(asserted.Info())
		mutator.Host(asserted.Host())
		mutator.BasePath(asserted.BasePath())
		for iter := asserted.Schemes(); iter.Next(); {
			item := iter.Item()
			mutator.Scheme(item)
		}
		for iter := asserted.Consumes(); iter.Next(); {
			item := iter.Item()
			mutator.Consume(item)
		}
		for iter := asserted.Produces(); iter.Next(); {
			item := iter.Item()
			mutator.Produce(item)
		}
		mutator.Paths(asserted.Paths())
		for iter := asserted.Definitions(); iter.Next(); {
			key, item := iter.Item()
			mutator.Definition(key, item)
		}
		for iter := asserted.Parameters(); iter.Next(); {
			key, item := iter.Item()
			mutator.Parameter(key, item)
		}
		for iter := asserted.Responses(); iter.Next(); {
			key, item := iter.Item()
			mutator.Response(key, item)
		}
		for iter := asserted.SecurityDefinitions(); iter.Next(); {
			key, item := iter.Item()
			mutator.SecurityDefinition(key, item)
		}
		for iter := asserted.Security(); iter.Next(); {
			item := iter.Item()
			mutator.Security(item)
		}
		for iter := asserted.Tags(); iter.Next(); {
			item := iter.Item()
			mutator.Tag(item)
		}
		mutator.ExternalDocs(asserted.ExternalDocs())
		if err := mutator.Do(); err != nil {
			return errors.Wrap(err, `failed to mutate`)
		}
		v.resolved = true
	}
	if v.info != nil {
		if err := v.info.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Info`)
		}
	}
	if v.schemes != nil {
		if err := v.schemes.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Schemes`)
		}
	}
	if v.consumes != nil {
		if err := v.consumes.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Consumes`)
		}
	}
	if v.produces != nil {
		if err := v.produces.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Produces`)
		}
	}
	if v.paths != nil {
		if err := v.paths.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Paths`)
		}
	}
	if v.definitions != nil {
		if err := v.definitions.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Definitions`)
		}
	}
	if v.parameters != nil {
		if err := v.parameters.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Parameters`)
		}
	}
	if v.responses != nil {
		if err := v.responses.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Responses`)
		}
	}
	if v.securityDefinitions != nil {
		if err := v.securityDefinitions.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve SecurityDefinitions`)
		}
	}
	if v.security != nil {
		if err := v.security.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Security`)
		}
	}
	if v.tags != nil {
		if err := v.tags.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve Tags`)
		}
	}
	if v.externalDocs != nil {
		if err := v.externalDocs.Resolve(resolver); err != nil {
			return errors.Wrap(err, `failed to resolve ExternalDocs`)
		}
	}
	return nil
}

func (v *swagger) QueryJSON(path string) (ret interface{}, ok bool) {
	path = strings.TrimLeftFunc(path, func(r rune) bool { return r == '#' || r == '/' })
	if path == "" {
		return v, true
	}

	var frag string
	if i := strings.Index(path, "/"); i > -1 {
		frag = path[:i]
		path = path[i+1:]
	} else {
		frag = path
		path = ""
	}

	var target interface{}

	switch frag {
	case "swagger":
		target = v.version
	case "info":
		target = v.info
	case "host":
		target = v.host
	case "basePath":
		target = v.basePath
	case "schemes":
		target = v.schemes
	case "consumes":
		target = v.consumes
	case "produces":
		target = v.produces
	case "paths":
		target = v.paths
	case "definitions":
		target = v.definitions
	case "parameters":
		target = v.parameters
	case "responses":
		target = v.responses
	case "securityDefinitions":
		target = v.securityDefinitions
	case "security":
		target = v.security
	case "tags":
		target = v.tags
	case "externalDocs":
		target = v.externalDocs
	default:
		return nil, false
	}

	if qj, ok := target.(QueryJSONer); ok {
		return qj.QueryJSON(path)
	}
	if path == "" {
		return target, true
	}
	return nil, false
}
