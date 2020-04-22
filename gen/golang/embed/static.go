// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00f\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00client.go.tmplUT\x05\x00\x010\xbf\x9f^{{define \"client\"}}\n{{if .Services}}\n//\n// Client\n//\n\n{{range .Services}}\nconst {{.Name | constPathPrefix}} = \"/rpc/{{.Name}}/\"\n{{end}}\n\n{{range .Services}}\n  {{ $serviceName := .Name | clientServiceName}}\n  type {{$serviceName}} struct {\n    client HTTPClient\n    urls   [{{.Methods | countMethods}}]string\n  }\n\n  func {{.Name | newClientServiceName }}(addr string, client HTTPClient) {{.Name}} {\n    prefix := urlBase(addr) + {{.Name | constPathPrefix}}\n    urls := [{{.Methods | countMethods}}]string{\n      {{- range .Methods}}\n      prefix + \"{{.Name}}\",\n      {{- end}}\n    }\n    return &{{$serviceName}}{\n      client: client,\n      urls:   urls,\n    }\n  }\n\n  {{range $i, $method := .Methods}}\n    func (c *{{$serviceName}}) {{.Name}}({{.Inputs | methodInputs}}) ({{.Outputs | methodOutputs }}) {\n      {{- $inputVar := \"nil\" -}}\n      {{- $outputVar := \"nil\" -}}\n      {{- if .Inputs | len}}\n      {{- $inputVar = \"in\"}}\n      in := struct {\n        {{- range $i, $input := .Inputs}}\n          Arg{{$i}} {{$input | methodArgType}} `json:\"{{$input.Name | downcaseName}}\"`\n        {{- end}}          \n      }{ {{.Inputs | methodArgNames}} }\n      {{- end}}\n      {{- if .Outputs | len}}\n      {{- $outputVar = \"&out\"}}\n      out := struct {\n        {{- range $i, $output := .Outputs}}\n          Ret{{$i}} {{$output | methodArgType}} `json:\"{{$output.Name | downcaseName}}\"`\n        {{- end}}          \n      }{}\n      {{- end}}\n\n      err := doJSONRequest(ctx, c.client, c.urls[{{$i}}], {{$inputVar}}, {{$outputVar}})\n      return {{argsList .Outputs \"out.Ret\"}}{{commaIfLen .Outputs}} err\n    }\n  {{end}}\n{{end}}\n\n// HTTPClient is the interface used by generated clients to send HTTP requests.\n// It is fulfilled by *(net/http).Client, which is sufficient for most users.\n// Users can provide their own implementation for special retry policies.\ntype HTTPClient interface {\n  Do(req *http.Request) (*http.Response, error)\n}\n\n// urlBase helps ensure that addr specifies a scheme. If it is unparsable\n// as a URL, it returns addr unchanged.\nfunc urlBase(addr string) string {\n  // If the addr specifies a scheme, use it. If not, default to\n  // http. If url.Parse fails on it, return it unchanged.\n  url, err := url.Parse(addr)\n  if err != nil {\n    return addr\n  }\n  if url.Scheme == \"\" {\n    url.Scheme = \"http\"\n  }\n  return url.String()\n}\n\n// newRequest makes an http.Request from a client, adding common headers.\nfunc newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {\n  req, err := http.NewRequest(\"POST\", url, reqBody)\n  if err != nil {\n    return nil, err\n  }\n  req.Header.Set(\"Accept\", contentType)\n  req.Header.Set(\"Content-Type\", contentType)\n	if headers, ok := HTTPRequestHeaders(ctx); ok {\n		for k := range headers {\n			for _, v := range headers[k] {\n				req.Header.Add(k, v)\n			}\n		}\n	}\n  return req, nil\n}\n\n// doJSONRequest is common code to make a request to the remote service.\nfunc doJSONRequest(ctx context.Context, client HTTPClient, url string, in, out interface{}) error {\n	reqBody, err := json.Marshal(in)\n	if err != nil {\n		return clientError(\"failed to marshal json request\", err)\n	}\n	if err = ctx.Err(); err != nil {\n		return clientError(\"aborted because context was done\", err)\n	}\n\n	req, err := newRequest(ctx, url, bytes.NewBuffer(reqBody), \"application/json\")\n	if err != nil {\n		return clientError(\"could not build request\", err)\n	}\n	resp, err := client.Do(req)\n	if err != nil {\n		return clientError(\"request failed\", err)\n	}\n\n	defer func() {\n		cerr := resp.Body.Close()\n		if err == nil && cerr != nil {\n			err = clientError(\"failed to close response body\", cerr)\n		}\n	}()\n\n	if err = ctx.Err(); err != nil {\n		return clientError(\"aborted because context was done\", err)\n	}\n\n	if resp.StatusCode != 200 {\n		return errorFromResponse(resp)\n	}\n\n	if out != nil {\n		respBody, err := ioutil.ReadAll(resp.Body)\n		if err != nil {\n			return clientError(\"failed to read response body\", err)\n		}\n\n		err = json.Unmarshal(respBody, &out)\n		if err != nil {\n			return clientError(\"failed to unmarshal json response body\", err)\n		}\n		if err = ctx.Err(); err != nil {\n			return clientError(\"aborted because context was done\", err)\n		}\n	}\n\n	return nil\n}\n\n// errorFromResponse builds a webrpc Error from a non-200 HTTP response.\nfunc errorFromResponse(resp *http.Response) Error {\n	respBody, err := ioutil.ReadAll(resp.Body)\n	if err != nil {\n		return clientError(\"failed to read server error response body\", err)\n	}\n\n	var respErr ErrorPayload\n	if err := json.Unmarshal(respBody, &respErr); err != nil {\n		return clientError(\"failed unmarshal error response\", err)\n	}\n\n	errCode := ErrorCode(respErr.Code)\n\n	if HTTPStatusFromErrorCode(errCode) == 0 {\n		return ErrorInternal(\"invalid code returned from server error response: %s\", respErr.Code)\n	}\n\n	return &rpcErr{\n		code:  errCode,\n		msg:   respErr.Msg,\n		cause: errors.New(respErr.Cause),\n	}\n}\n\nfunc clientError(desc string, err error) Error {\n	return WrapError(ErrInternal, err, desc)\n}\n\nfunc WithHTTPRequestHeaders(ctx context.Context, h http.Header) (context.Context, error) {\n	if _, ok := h[\"Accept\"]; ok {\n		return nil, errors.New(\"provided header cannot set Accept\")\n	}\n	if _, ok := h[\"Content-Type\"]; ok {\n		return nil, errors.New(\"provided header cannot set Content-Type\")\n	}\n\n	copied := make(http.Header, len(h))\n	for k, vv := range h {\n		if vv == nil {\n			copied[k] = nil\n			continue\n		}\n		copied[k] = make([]string, len(vv))\n		copy(copied[k], vv)\n	}\n\n	return context.WithValue(ctx, HTTPClientRequestHeadersCtxKey, copied), nil\n}\n\nfunc HTTPRequestHeaders(ctx context.Context) (http.Header, bool) {\n	h, ok := ctx.Value(HTTPClientRequestHeadersCtxKey).(http.Header)\n	return h, ok\n}\n{{end}}\n{{end}}\nPK\x07\x08\xc5\xc9w\xb8]\x16\x00\x00]\x16\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00f\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00helpers.go.tmplUT\x05\x00\x010\xbf\x9f^{{define \"helpers\"}}\n\n//\n// Helpers\n//\n\ntype ErrorPayload struct {\n	Status int    `json:\"status\"`\n	Code   string `json:\"code\"`\n	Cause  string `json:\"cause,omitempty\"`\n	Msg    string `json:\"msg\"`\n	Error  string `json:\"error\"`\n}\n\ntype Error interface {\n	// Code is of the valid error codes\n	Code() ErrorCode\n\n	// Msg returns a human-readable, unstructured messages describing the error\n	Msg() string\n\n	// Cause is reason for the error\n	Cause() error\n\n	// Error returns a string of the form \"webrpc error <Code>: <Msg>\"\n	Error() string\n\n	// Error response payload\n	Payload() ErrorPayload\n}\n\nfunc Errorf(code ErrorCode, msgf string, args ...interface{}) Error {\n	msg := fmt.Sprintf(msgf, args...)\n	if IsValidErrorCode(code) {\n		return &rpcErr{code: code, msg: msg}\n	}\n	return &rpcErr{code: ErrInternal, msg: \"invalid error type \" + string(code)}\n}\n\nfunc WrapError(code ErrorCode, cause error, format string, args ...interface{}) Error {\n	msg := fmt.Sprintf(format, args...)\n	if IsValidErrorCode(code) {\n		return &rpcErr{code: code, msg: msg, cause: cause}\n	}\n	return &rpcErr{code: ErrInternal, msg: \"invalid error type \" + string(code), cause: cause}\n}\n\nfunc Failf(format string, args ...interface{}) Error {\n	return Errorf(ErrFail, format, args...)\n}\n\nfunc WrapFailf(cause error, format string, args ...interface{}) Error {\n	return WrapError(ErrFail, cause, format, args...)\n}\n\nfunc ErrorNotFound(format string, args ...interface{}) Error {\n	return Errorf(ErrNotFound, format, args...)\n}\n\nfunc ErrorInvalidArgument(argument string, validationMsg string) Error {\n	return Errorf(ErrInvalidArgument, argument+\" \"+validationMsg)\n}\n\nfunc ErrorRequiredArgument(argument string) Error {\n	return ErrorInvalidArgument(argument, \"is required\")\n}\n\nfunc ErrorInternal(format string, args ...interface{}) Error {\n	return Errorf(ErrInternal, format, args...)\n}\n\ntype ErrorCode string\n\nconst (\n	// Unknown error. For example when handling errors raised by APIs that do not\n	// return enough error information.\n	ErrUnknown ErrorCode = \"unknown\"\n\n	// Fail error. General failure error type.\n	ErrFail ErrorCode = \"fail\"\n\n	// Canceled indicates the operation was cancelled (typically by the caller).\n	ErrCanceled ErrorCode = \"canceled\"\n\n	// InvalidArgument indicates client specified an invalid argument. It\n	// indicates arguments that are problematic regardless of the state of the\n	// system (i.e. a malformed file name, required argument, number out of range,\n	// etc.).\n	ErrInvalidArgument ErrorCode = \"invalid argument\"\n\n	// DeadlineExceeded means operation expired before completion. For operations\n	// that change the state of the system, this error may be returned even if the\n	// operation has completed successfully (timeout).\n	ErrDeadlineExceeded ErrorCode = \"deadline exceeded\"\n\n	// NotFound means some requested entity was not found.\n	ErrNotFound ErrorCode = \"not found\"\n\n	// BadRoute means that the requested URL path wasn't routable to a webrpc\n	// service and method. This is returned by the generated server, and usually\n	// shouldn't be returned by applications. Instead, applications should use\n	// NotFound or Unimplemented.\n	ErrBadRoute ErrorCode = \"bad route\"\n\n	// AlreadyExists means an attempt to create an entity failed because one\n	// already exists.\n	ErrAlreadyExists ErrorCode = \"already exists\"\n\n	// PermissionDenied indicates the caller does not have permission to execute\n	// the specified operation. It must not be used if the caller cannot be\n	// identified (Unauthenticated).\n	ErrPermissionDenied ErrorCode = \"permission denied\"\n\n	// Unauthenticated indicates the request does not have valid authentication\n	// credentials for the operation.\n	ErrUnauthenticated ErrorCode = \"unauthenticated\"\n\n	// ResourceExhausted indicates some resource has been exhausted, perhaps a\n	// per-user quota, or perhaps the entire file system is out of space.\n	ErrResourceExhausted ErrorCode = \"resource exhausted\"\n\n	// FailedPrecondition indicates operation was rejected because the system is\n	// not in a state required for the operation's execution. For example, doing\n	// an rmdir operation on a directory that is non-empty, or on a non-directory\n	// object, or when having conflicting read-modify-write on the same resource.\n	ErrFailedPrecondition ErrorCode = \"failed precondition\"\n\n	// Aborted indicates the operation was aborted, typically due to a concurrency\n	// issue like sequencer check failures, transaction aborts, etc.\n	ErrAborted ErrorCode = \"aborted\"\n\n	// OutOfRange means operation was attempted past the valid range. For example,\n	// seeking or reading past end of a paginated collection.\n	//\n	// Unlike InvalidArgument, this error indicates a problem that may be fixed if\n	// the system state changes (i.e. adding more items to the collection).\n	//\n	// There is a fair bit of overlap between FailedPrecondition and OutOfRange.\n	// We recommend using OutOfRange (the more specific error) when it applies so\n	// that callers who are iterating through a space can easily look for an\n	// OutOfRange error to detect when they are done.\n	ErrOutOfRange ErrorCode = \"out of range\"\n\n	// Unimplemented indicates operation is not implemented or not\n	// supported/enabled in this service.\n	ErrUnimplemented ErrorCode = \"unimplemented\"\n\n	// Internal errors. When some invariants expected by the underlying system\n	// have been broken. In other words, something bad happened in the library or\n	// backend service. Do not confuse with HTTP Internal Server Error; an\n	// Internal error could also happen on the client code, i.e. when parsing a\n	// server response.\n	ErrInternal ErrorCode = \"internal\"\n\n	// Unavailable indicates the service is currently unavailable. This is a most\n	// likely a transient condition and may be corrected by retrying with a\n	// backoff.\n	ErrUnavailable ErrorCode = \"unavailable\"\n\n	// DataLoss indicates unrecoverable data loss or corruption.\n	ErrDataLoss ErrorCode = \"data loss\"\n\n	// ErrNone is the zero-value, is considered an empty error and should not be\n	// used.\n	ErrNone ErrorCode = \"\"\n)\n\nfunc HTTPStatusFromErrorCode(code ErrorCode) int {\n	switch code {\n	case ErrCanceled:\n		return 408 // RequestTimeout\n	case ErrUnknown:\n		return 400 // Bad Request\n	case ErrFail:\n		return 422 // Unprocessable Entity\n	case ErrInvalidArgument:\n		return 400 // BadRequest\n	case ErrDeadlineExceeded:\n		return 408 // RequestTimeout\n	case ErrNotFound:\n		return 404 // Not Found\n	case ErrBadRoute:\n		return 404 // Not Found\n	case ErrAlreadyExists:\n		return 409 // Conflict\n	case ErrPermissionDenied:\n		return 403 // Forbidden\n	case ErrUnauthenticated:\n		return 401 // Unauthorized\n	case ErrResourceExhausted:\n		return 403 // Forbidden\n	case ErrFailedPrecondition:\n		return 412 // Precondition Failed\n	case ErrAborted:\n		return 409 // Conflict\n	case ErrOutOfRange:\n		return 400 // Bad Request\n	case ErrUnimplemented:\n		return 501 // Not Implemented\n	case ErrInternal:\n		return 500 // Internal Server Error\n	case ErrUnavailable:\n		return 503 // Service Unavailable\n	case ErrDataLoss:\n		return 500 // Internal Server Error\n	case ErrNone:\n		return 200 // OK\n	default:\n		return 0 // Invalid!\n	}\n}\n\nfunc IsErrorCode(err error, code ErrorCode) bool {\n	if rpcErr, ok := err.(Error); ok {\n		if rpcErr.Code() == code {\n			return true\n		}\n	}\n	return false\n}\n\nfunc IsValidErrorCode(code ErrorCode) bool {\n	return HTTPStatusFromErrorCode(code) != 0\n}\n\ntype rpcErr struct {\n	code  ErrorCode\n	msg   string\n	cause error\n}\n\nfunc (e *rpcErr) Code() ErrorCode {\n	return e.code\n}\n\nfunc (e *rpcErr) Msg() string {\n	return e.msg\n}\n\nfunc (e *rpcErr) Cause() error {\n	return e.cause\n}\n\nfunc (e *rpcErr) Error() string {\n	if e.cause != nil && e.cause.Error() != \"\" {\n		if e.msg != \"\" {\n			return fmt.Sprintf(\"webrpc %s error: %s -- %s\", e.code, e.cause.Error(), e.msg)\n		} else {\n			return fmt.Sprintf(\"webrpc %s error: %s\", e.code, e.cause.Error())\n		}\n	} else {\n		return fmt.Sprintf(\"webrpc %s error: %s\", e.code, e.msg)\n	}\n}\n\nfunc (e *rpcErr) Payload() ErrorPayload {\n	statusCode := HTTPStatusFromErrorCode(e.Code())\n	errPayload := ErrorPayload{\n		Status: statusCode,\n		Code:   string(e.Code()),\n		Msg:    e.Msg(),\n		Error:  e.Error(),\n	}\n	if e.Cause() != nil {\n		errPayload.Cause = e.Cause().Error()\n	}\n	return errPayload\n}\n\ntype contextKey struct {\n	name string\n}\n\nfunc (k *contextKey) String() string {\n	return \"webrpc context value \" + k.name\n}\n\nvar (\n	// For Client\n	HTTPClientRequestHeadersCtxKey = &contextKey{\"HTTPClientRequestHeaders\"}\n\n	// For Server\n	HTTPResponseWriterCtxKey = &contextKey{\"HTTPResponseWriter\"}\n\n	HTTPRequestCtxKey = &contextKey{\"HTTPRequest\"}\n\n	ServiceNameCtxKey = &contextKey{\"ServiceName\"}\n\n	MethodNameCtxKey = &contextKey{\"MethodName\"}\n)\n\n{{end}}\nPK\x07\x08\x83\xea\x053\xde!\x00\x00\xde!\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00f\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00proto.gen.go.tmplUT\x05\x00\x010\xbf\x9f^{{- define \"proto\" -}}\n// {{.Name}} {{.SchemaVersion}} {{.SchemaHash}}\n// --\n// This file has been generated by https://github.com/webrpc/webrpc using gen/golang\n// Do not edit by hand. Update your webrpc schema and re-generate.\npackage {{.TargetOpts.PkgName}}\n\nimport (\n  \"context\"\n  \"encoding/json\"\n  \"fmt\"\n  \"io/ioutil\"\n  \"net/http\"\n  \"time\"\n  \"strings\"\n  \"bytes\"\n  \"errors\"\n  \"io\"\n  \"net/url\"\n)\n\n// WebRPC description and code-gen version\nfunc WebRPCVersion() string {\n  return \"{{.WebRPCVersion}}\"\n}\n\n// Schema version of your RIDL schema\nfunc WebRPCSchemaVersion() string {\n  return \"{{.SchemaVersion}}\"\n}\n\n// Schema hash generated from your RIDL schema\nfunc WebRPCSchemaHash() string {\n  return \"{{.SchemaHash}}\"\n}\n\n{{template \"types\" .}}\n\n{{if .TargetOpts.Server}}\n  {{template \"server\" .}}\n{{end}}\n\n{{if .TargetOpts.Client}}\n  {{template \"client\" .}}\n{{end}}\n\n{{template \"helpers\" .}}\n\n{{- end}}\nPK\x07\x08g4\x9a/\x89\x03\x00\x00\x89\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00f\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00server.go.tmplUT\x05\x00\x010\xbf\x9f^{{define \"server\"}}\n{{if .Services}}\n//\n// Server\n//\n\ntype WebRPCServer interface {\n  http.Handler\n}\n\n{{- range .Services}}\n  {{$name := .Name}}\n  {{$serviceName := .Name | serverServiceName}}\n\n  type {{$serviceName}} struct {\n    {{.Name}}\n  }\n\n  func {{ .Name | newServerServiceName }}(svc {{.Name}}) WebRPCServer {\n    return &{{$serviceName}}{\n      {{.Name}}: svc,\n    }\n  }\n\n  func (s *{{$serviceName}}) ServeHTTP(w http.ResponseWriter, r *http.Request) {\n    ctx := r.Context()\n    ctx = context.WithValue(ctx, HTTPResponseWriterCtxKey, w)\n    ctx = context.WithValue(ctx, HTTPRequestCtxKey, r)\n    ctx = context.WithValue(ctx, ServiceNameCtxKey, \"{{.Name}}\")\n\n    if r.Method != \"POST\" {\n      err := Errorf(ErrBadRoute, \"unsupported method %q (only POST is allowed)\", r.Method)\n      RespondWithError(w, err)\n      return\n    }\n\n    switch r.URL.Path {\n    {{- range .Methods}}\n    case \"/rpc/{{$name}}/{{.Name}}\":\n      s.{{.Name | serviceMethodName}}(ctx, w, r)\n      return\n    {{- end}}\n    default:\n      err := Errorf(ErrBadRoute, \"no handler for path %q\", r.URL.Path)\n      RespondWithError(w, err)\n      return\n    }\n  }\n\n  {{range .Methods}}\n    func (s *{{$serviceName}}) {{.Name | serviceMethodName}}(ctx context.Context, w http.ResponseWriter, r *http.Request) {\n      header := r.Header.Get(\"Content-Type\")\n      i := strings.Index(header, \";\")\n      if i == -1 {\n        i = len(header)\n      }\n\n      switch strings.TrimSpace(strings.ToLower(header[:i])) {\n      case \"application/json\":\n        s.{{ .Name | serviceMethodJSONName }}(ctx, w, r)\n      default:\n        err := Errorf(ErrBadRoute, \"unexpected Content-Type: %q\", r.Header.Get(\"Content-Type\"))\n        RespondWithError(w, err)\n      }\n    }\n\n    func (s *{{$serviceName}}) {{.Name | serviceMethodJSONName}}(ctx context.Context, w http.ResponseWriter, r *http.Request) {\n      var err error\n      ctx = context.WithValue(ctx, MethodNameCtxKey, \"{{.Name}}\")\n\n      {{- if .Inputs|len}}\n      reqContent := struct {\n      {{- range $i, $input := .Inputs}}\n        Arg{{$i}} {{. | methodArgType}} `json:\"{{$input.Name | downcaseName}}\"`\n      {{- end}}\n      }{}\n\n      reqBody, err := ioutil.ReadAll(r.Body)\n      if err != nil {\n        err = WrapError(ErrInternal, err, \"failed to read request data\")\n        RespondWithError(w, err)\n        return\n      }\n      defer r.Body.Close()\n\n      err = json.Unmarshal(reqBody, &reqContent)\n      if err != nil {\n        err = WrapError(ErrInvalidArgument, err, \"failed to unmarshal request data\")\n        RespondWithError(w, err)\n        return\n      }\n      {{- end}}\n\n      // Call service method\n      {{- range $i, $output := .Outputs}}\n      var ret{{$i}} {{$output | methodArgType}}\n      {{- end}}\n      func() {\n        defer func() {\n          // In case of a panic, serve a 500 error and then panic.\n          if rr := recover(); rr != nil {\n            RespondWithError(w, ErrorInternal(\"internal service panic\"))\n            panic(rr)\n          }\n        }()\n        {{argsList .Outputs \"ret\"}}{{.Outputs | commaIfLen}} err = s.{{$name}}.{{.Name}}(ctx{{.Inputs | commaIfLen}}{{argsList .Inputs \"reqContent.Arg\"}})\n      }()\n      {{- if .Outputs | len}}\n      respContent := struct {\n      {{- range $i, $output := .Outputs}}\n        Ret{{$i}} {{$output | methodArgType}} `json:\"{{$output.Name | downcaseName}}\"`\n      {{- end}}         \n      }{ {{argsList .Outputs \"ret\"}} }\n      {{- end}}\n\n      if err != nil {\n        RespondWithError(w, err)\n        return\n      }\n\n      {{- if .Outputs | len}}\n      respBody, err := json.Marshal(respContent)\n      if err != nil {\n        err = WrapError(ErrInternal, err, \"failed to marshal json response\")\n        RespondWithError(w, err)\n        return\n      }\n      {{- end}}\n\n      w.Header().Set(\"Content-Type\", \"application/json\")\n      w.WriteHeader(http.StatusOK)\n\n      {{- if .Outputs | len}}\n      w.Write(respBody)\n      {{- end}}\n    }\n  {{end}}\n{{- end}}\n\nfunc RespondWithError(w http.ResponseWriter, err error) {\n	rpcErr, ok := err.(Error)\n	if !ok {\n		rpcErr = WrapError(ErrInternal, err, \"webrpc error\")\n	}\n\n	statusCode := HTTPStatusFromErrorCode(rpcErr.Code())\n\n	w.Header().Set(\"Content-Type\", \"application/json\")\n	w.WriteHeader(statusCode)\n\n	respBody, _ := json.Marshal(rpcErr.Payload())\n	w.Write(respBody)\n}\n{{end}}\n{{end}}\nPK\x07\x08{\x8fd\xdd\xe8\x10\x00\x00\xe8\x10\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00f\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0d\x00	\x00types.go.tmplUT\x05\x00\x010\xbf\x9f^{{define \"types\"}}\n\n{{if .Messages}}\n//\n// Types\n//\n\n{{range .Messages}}\n  {{if .Type | isEnum}}\n    {{$enumName := .Name}}\n    {{$enumType := .EnumType}}\n    type {{$enumName}} {{$enumType}}\n\n    const (\n      {{- range .Fields}}\n        {{$enumName}}_{{.Name}} {{$enumName}} = {{.Value}}\n      {{- end}}\n    )\n\n    var {{$enumName}}_name = map[{{$enumType}}]string {\n      {{- range .Fields}}\n        {{.Value}}: \"{{.Name}}\",\n      {{- end}}\n    }\n\n    var {{$enumName}}_value = map[string]{{$enumType}} {\n      {{- range .Fields}}\n        \"{{.Name}}\": {{.Value}},\n      {{- end}}\n    }\n\n    func (x {{$enumName}}) String() string {\n      return {{$enumName}}_name[{{$enumType}}(x)]\n    }\n\n    func (x {{$enumName}}) MarshalJSON() ([]byte, error) {\n      buf := bytes.NewBufferString(`\"`)\n      buf.WriteString({{$enumName}}_name[{{$enumType}}(x)])\n      buf.WriteString(`\"`)\n      return buf.Bytes(), nil\n    }\n\n    func (x *{{$enumName}}) UnmarshalJSON(b []byte) error {\n      var j string\n      err := json.Unmarshal(b, &j)\n      if err != nil {\n        return err\n      }\n      *x = {{$enumName}}({{$enumName}}_value[j])\n      return nil\n    }\n  {{end}}\n  {{if .Type | isStruct  }}\n    type {{.Name}} struct {\n      {{- range .Fields}}\n        {{. | exportedField}} {{. | fieldOptional}}{{. | fieldTypeDef}} {{. | fieldTags}}\n      {{- end}}\n    }\n  {{end}}\n{{end}}\n{{end}}\n{{if .Services}}\n  {{range .Services}}\n    type {{.Name}} interface {\n      {{- range .Methods}}\n        {{.Name}}({{.Inputs | methodInputs}}) ({{.Outputs | methodOutputs}})\n      {{- end}}\n    }\n  {{end}}\n  var WebRPCServices = map[string][]string{\n    {{- range .Services}}\n      \"{{.Name}}\": {\n        {{- range .Methods}}\n          \"{{.Name}}\",\n        {{- end}}\n      },\n    {{- end}}\n  }\n{{end}}\n\n{{end}}\nPK\x07\x08\xf8\xf7\x1e\xb7\xff\x06\x00\x00\xff\x06\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00f\x1e\x96P\xc5\xc9w\xb8]\x16\x00\x00]\x16\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.go.tmplUT\x05\x00\x010\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00f\x1e\x96P\x83\xea\x053\xde!\x00\x00\xde!\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xa2\x16\x00\x00helpers.go.tmplUT\x05\x00\x010\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00f\x1e\x96Pg4\x9a/\x89\x03\x00\x00\x89\x03\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xc68\x00\x00proto.gen.go.tmplUT\x05\x00\x010\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00f\x1e\x96P{\x8fd\xdd\xe8\x10\x00\x00\xe8\x10\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x97<\x00\x00server.go.tmplUT\x05\x00\x010\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00f\x1e\x96P\xf8\xf7\x1e\xb7\xff\x06\x00\x00\xff\x06\x00\x00\x0d\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xc4M\x00\x00types.go.tmplUT\x05\x00\x010\xbf\x9f^PK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00\\\x01\x00\x00\x07U\x00\x00\x00\x00"
