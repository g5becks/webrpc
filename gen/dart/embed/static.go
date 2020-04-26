// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x005\n\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00client.dart.tmplUT\x05\x00\x017\xe1\xa4^{{define \"client\"}}\n\n{{if .Services}}\n//\n// Client\n//\n\nString _removeSlash(String host) => host.endsWith('/')\n? host.replaceRange(host.length - 1, host.length, '')\n: host;\n\n{{range .Services}}\nclass {{.Name}} {\n  final http.Client client;\n  final String host;\n  final String path = '/rpc/{{.Name}}/';\n  String url(String name) => '${_removeSlash(host)}$path$name';\n  {{range .Methods}}\n\n  {{end}}\n  {{.Name}}(this.client, this.host);\n}\n{{end}}\n\n\n{{end}}\n{{end}}\nPK\x07\x08\xd8n\xd9\xec\xce\x01\x00\x00\xce\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00C\x0b\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00helpers.dart.tmplUT\x05\x00\x01.\xe3\xa4^{{define \"helpers\"}}\n\n//\n// HELPER CODE.\n//\n\nabstract class Logger {\n  void _log(message, [Object error, StackTrace stackTrace]) => print(\n      '{message: $message}, error: $error, stackTrace: $stackTrace, time: ${DateTime.now()}');\n  void finest(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void finer(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void fine(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void config(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void info(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void warning(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void severe(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n  void shout(message, [Object error, StackTrace stackTrace]) =>\n      _log(message, error, stackTrace);\n}\n\nclass _Logger extends Logger {\n  _Logger();\n}\n\nfinal _rpcLogger = _Logger();\n\n// An error in the http stack.\nclass HttpErr {\n  final String status;\n  final int code;\n  const HttpErr(this.status, this.code);\n\n  Map<String, dynamic> toMap() => {'status': status, 'code': code};\n  String toJson() => jsonEncode(toMap());\n  static HttpErr fromMap(Map<String, dynamic> map) =>\n      HttpErr(map['status'] as String, map['code'] as int);\n\n  static HttpErr fromJson(json) => fromMap(jsonDecode(json));\n}\n\n// An error created by the rpc server.\nclass RpcErr {\n  final String message;\n  final String path;\n  final DateTime time;\n  final HttpErr httpErr;\n  const RpcErr({this.message, this.path, this.httpErr, this.time});\n\n  Map<String, dynamic> toMap() => {\n        'message': message,\n        'path': path,\n        'httpErr': httpErr.toMap(),\n        'time-stamp': time.toString()\n      };\n  String toJson() => jsonEncode(toMap());\n  static RpcErr fromMap(Map<String, dynamic> map) => RpcErr(\n      message: map['message'] as String,\n      path: map['path'] as String,\n      time: DateTime.parse(map['time-stamp']),\n      httpErr: HttpErr.fromMap(map['httpErr']));\n  static RpcErr fromJson(json) => fromMap(jsonDecode(json));\n}\n\n// Contains static fields for creating and identifying http errors.\nclass err {\n  // Unknown error. For example when handling errors raised by APIs that do not\n  // return enough error information.\n  static HttpErr Unknown = HttpErr('unknown', 400);\n  // 422 (Unprocessable Entity) Fail error. General failure error type.\n  static HttpErr Fail = HttpErr('fail', 422);\n  // RequestTimeout Canceled indicates the operation was cancelled (typically by the caller).\n  static HttpErr Canceled = HttpErr('canceled', 408);\n  // InvalidArgument indicates client specified an invalid argument. It\n  // indicates arguments that are problematic regardless of the state of the\n  // system (i.e. a malformed file name, required argument, number out of range,\n  // etc.).\n  static HttpErr InvalidArgument = HttpErr('invalid argument', 422);\n  // RequestTimeOut. DeadlineExceeded means operation expired before completion. For operations\n  // that change the state of the system, this error may be returned even if the\n  // operation has completed successfully (timeout).\n  static HttpErr DeadlineExceeded = HttpErr('deadline exceeded', 408);\n  // NotFound means some requested entity was not found.\n  static HttpErr NotFound = HttpErr('not found', 404);\n  // BadRoute means that the requested URL path wasn't routable to a webrpc\n  // service and method. This is returned by the generated server, and usually\n  // shouldn't be returned by applications. Instead, applications should use\n  // NotFound or Unimplemented.\n  static HttpErr BadRoute = HttpErr('bad route', 404);\n  // AlreadyExists means an attempt to create an entity failed because one\n  // already exists. Conflict.\n  static HttpErr AlreadyExists = HttpErr('already exists', 409);\n  // PermissionDenied indicates the caller does not have permission to execute\n  // the specified operation. It must not be used if the caller cannot be\n  // identified (Unauthenticated).\n  static HttpErr PermissionDenied = HttpErr('permission denied', 403);\n  // Unauthenticated indicates the request does not have valid authentication\n  // credentials for the operation. Unauthorized.\n  static HttpErr Unauthenticated = HttpErr('unauthenticated', 401);\n  // ResourceExhausted indicates some resource has been exhausted, perhaps a\n  // per-user quota, or perhaps the entire file system is out of space. Forbidden.\n  static HttpErr ResourceExhausted = HttpErr('resource exhausted', 403);\n  // FailedPrecondition indicates operation was rejected because the system is\n  // not in a state required for the operation's execution. For example, doing\n  // an rmdir operation on a directory that is non-empty, or on a non-directory\n  // object, or when having conflicting read-modify-write on the same resource. Precondition failed.\n  static HttpErr FailedPrecondition = HttpErr('failed precondition', 412);\n  // Aborted indicates the operation was aborted, typically due to a concurrency\n  // issue like sequencer check failures, transaction aborts, etc.\n  static HttpErr Aborted = HttpErr('aborted', 409);\n  // OutOfRange means operation was attempted past the valid range. For example,\n  // seeking or reading past end of a paginated collection.\n  //\n  // Unlike InvalidArgument, this error indicates a problem that may be fixed if\n  // the system state changes (i.e. adding more items to the collection).\n  //\n  // There is a fair bit of overlap between FailedPrecondition and OutOfRange.\n  // We recommend using OutOfRange (the more specific error) when it applies so\n  // that callers who are iterating through a space can easily look for an\n  // OutOfRange error to detect when they are done.\n  static HttpErr OutOfRange = HttpErr('out of range', 400);\n  // Unimplemented indicates operation is not implemented or not\n  // supported/enabled in this service.\n  static HttpErr Unimplemented = HttpErr('unimplemented', 501);\n  // Internal errors. When some invariants expected by the underlying system\n  // have been broken. In other words, something bad happened in the library or\n  // backend service. Do not confuse with HTTP Internal Server Error; an\n  // Internal error could also happen on the client code, i.e. when parsing a\n  // server response.\n  static HttpErr Internal = HttpErr('internal', 500);\n  // Unavailable indicates the service is currently unavailable. This is a most\n  // likely a transient condition and may be corrected by retrying with a\n  // backoff. Service Unavailable.\n  static HttpErr Unavailable = HttpErr('unavailable', 503);\n  // DataLoss indicates unrecoverable data loss or corruption.\n  static HttpErr DataLoss = HttpErr('data loss', 500);\n}\n{{end}}PK\x07\x08Tv\"\xad\xfb\x1a\x00\x00\xfb\x1a\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xd8\x0b\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x13\x00	\x00proto.gen.dart.tmplUT\x05\x00\x01H\xe4\xa4^{{- define \"proto\" -}}\nimport 'dart:async';\nimport 'dart:convert';\nimport 'dart:io';\n\nimport 'package:meta/meta.dart';\nimport 'package:freezed_annotation/freezed_annotation.dart';\n{{if .TargetOpts.Client}}\nimport 'package:http/http.dart' as http;\n{{end}}\n{{if .TargetOpts.Server}}\nimport 'package:shelf/shelf.dart' as shelf;\nimport 'package:shelf/shelf_io.dart' as io;\n{{end}}\n\n// {{.Name}} {{.SchemaVersion}} {{.SchemaHash}}\n// --\n// This file has been generated by https://github.com/webrpc/webrpc using gen/dart\n// Do not edit by hand. Update your webrpc schema and re-generate.\n\n// WebRPC description and code-gen version\nString webRPCVersion()  {\n  return \"{{.WebRPCVersion}}\";\n}\n\n// Schema version of your RIDL schema\nString WebRPCSchemaVersion() {\n  return \"{{.SchemaVersion}}\";\n}\n\n// Schema hash generated from your RIDL schema\nString WebRPCSchemaHash() {\n  return \"{{.SchemaHash}}\";\n}\n\n{{template \"types\" .}}\n\n{{if .TargetOpts.Server}}\n  {{template \"server\" .}}\n  {{template \"server_helpers\" .}}\n{{end}}\n\n{{if .TargetOpts.Client}}\n  {{template \"client\" .}}\n{{end}}\n\n{{template \"helpers\" .}}\n\n{{- end}}\nPK\x07\x08Wd\xc4FW\x04\x00\x00W\x04\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x005\n\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00server.dart.tmplUT\x05\x00\x017\xe1\xa4^{{define \"server\"}}\n\n{{if .Services}}\n// *********************************************************************\n// SERVICE INTERFACES.\n// *********************************************************************\n{{range .Services -}}\nabstract class {{.Name}} {\n  {{- range .Methods}}\n  {{. | methodOutputs}} {{.Name | methodName}}({{. | serverMethodInputs}});\n  {{- end}}\n}\n{{end}}  {{/* end of range .Services for INTERFACES */}}\n\n// *********************************************************************\n// SERVER IMPLEMENTATION.\n// *********************************************************************\nclass WebRpcServer {\n  // For Google Cloud Run, set _hostname to '0.0.0.0'.\n  String _hostname;\n  // Provide a {Logger} implementation to log failed requests.\n  Logger _log;\n  // Provide a preconfigured shelf.Pipeline with desired middleware.\n  Set<shelf.Middleware> _middleware;\n{{- range .Services}}\n  final {{.Name}} {{.Name | serviceImplName}};\n{{- end}}  \n  WebRpcServer(\n    { {{ range .Services }}@required this.{{.Name | serviceImplName}},\n      {{end -}}\n      Logger log,\n      String hostName = 'localhost',\n      List<shelf.Middleware> middleware}) {\n      _hostname = hostName;\n      _log = log ?? _rpcLogger;\n      _middleware = middleware?.toSet() ?? [shelf.logRequests()];\n  }\n\n  bool _jsonFriendly(shelf.Request r) =>\n      r.headers['Content-Type'].contains('application/json') &&\n      r.headers['Accept'].contains('application/json');  \n\n  FutureOr<shelf.Response> _requestHandler(shelf.Request r) {\n    final route = r.url.path;\n    if (r.method != 'POST') {\n      final info =\n          'unsupported method: ${r.method}, (only POST is allowed. path: $route';\n      _log.info(info);\n      return rpcResp.BadRoute(route, msg: info);\n    }\n\n    if (!_jsonFriendly(r)) {\n      final info =\n          'unexpected Content-Type: ${r.headers['Content-Type']} or Accept: ${r.headers['Accept']}. path: $route';\n      _log.info(info);\n      return rpcResp.BadRoute(route, msg: info);\n    }\n    \n    switch (r.url.path) {\n      {{ range .Services }}\n      {{$name := .Name}}\n      {{- range .Methods}}\n      case '/rpc/{{$name}}/{{.Name}}': {\n         return _handle{{$name}}{{.Name}}(r);\n      }\n      break;\n      {{end -}} {{/* end of range .Methods */}}\n      {{end -}} {{/* end of range .Services */}}\n      default:\n        {\n          return rpcResp.BadRoute(route, msg: 'no handler for path: $route');\n        }\n        break;\n    }\n  }\n  {{- range .Services}}\n  {{$name := .Name}}\n  {{- range .Methods}}\n  FutureOr<shelf.Response> _handle{{$name}}{{.Name}}(shelf.Request r) {\n    \n  }\n  {{end}}{{/* end of range .Methods */}}\n  {{- end -}} {{/* end of range .Services for Methods */}}\n  FutureOr<void> run() async {}\n  //// SERVICES SHOULD END HERE DEBUG!!!!!!\n}\n\n\n{{end}} {{/* end of if .Services */}}\n{{end}} {{/* end of top level define */}}\nPK\x07\x08$\xadW\xe7/\x0b\x00\x00/\x0b\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xc3\x0b\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x18\x00	\x00server_helpers.dart.tmplUT\x05\x00\x01\x1f\xe4\xa4^{{ define \"server_helpers\" }}\n// Contains static method helpers for handling requests.\nclass rpcResp {\n  static String _message(String status, {String info}) =>\n      'webrpc error: $status, details: $info';\n\n  static const _jsonHeader = {\n    'Content-Type': 'application/json',\n    'X-Content-Type-Options': 'nosniff'\n  };\n\n  static shelf.Response _jsonResp(int code, {dynamic json}) => shelf.Response(\n        code,\n        headers: _jsonHeader,\n        body: json,\n      );\n\n  static shelf.Response Ok({String json}) => shelf.Response.ok(\n        json,\n        headers: _jsonHeader,\n      );\n\n  static shelf.Response Found(String route, {String msg = ''}) =>\n      shelf.Response.found(route, headers: _jsonHeader, body: msg);\n\n  static shelf.Response MovedPerm(String route, {String msg = ''}) =>\n      shelf.Response.movedPermanently(\n        route,\n        headers: _jsonHeader,\n      );\n\n  static shelf.Response NotModified() =>\n      shelf.Response.notModified(headers: _jsonHeader);\n\n  static shelf.Response SeeOther(String route, {String msg = ''}) =>\n      shelf.Response.seeOther(\n        route,\n        headers: _jsonHeader,\n        body: msg,\n      );\n\n  static shelf.Response Unknown(String route, {String msg = ''}) => _jsonResp(\n        err.Unknown.code,\n        json: RpcErr(\n                message: _message(err.Unknown.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unknown)\n            .toJson(),\n      );\n\n  static shelf.Response Fail(String route, {String msg = ''}) => _jsonResp(\n        err.Fail.code,\n        json: RpcErr(\n                message: _message(err.Fail.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Fail)\n            .toJson(),\n      );\n\n  static shelf.Response Canceled(String route, {String msg = ''}) => _jsonResp(\n        err.Canceled.code,\n        json: RpcErr(\n                message: _message(err.Canceled.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Canceled)\n            .toJson(),\n      );\n\n  static shelf.Response InvalidArgument(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.InvalidArgument.code,\n        json: RpcErr(\n                message: _message(err.InvalidArgument.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.InvalidArgument)\n            .toJson(),\n      );\n\n  static shelf.Response DeadlineExceeded(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.DeadlineExceeded.code,\n        json: RpcErr(\n                message: _message(err.DeadlineExceeded.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.DeadlineExceeded)\n            .toJson(),\n      );\n\n  static shelf.Response NotFound(String route, {String msg = ''}) => _jsonResp(\n        err.NotFound.code,\n        json: RpcErr(\n                message: _message(err.NotFound.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.NotFound)\n            .toJson(),\n      );\n\n  static shelf.Response BadRoute(String route, {String msg = ''}) => _jsonResp(\n        err.BadRoute.code,\n        json: RpcErr(\n                message: _message(err.BadRoute.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.BadRoute)\n            .toJson(),\n      );\n\n  static shelf.Response AlreadyExists(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.AlreadyExists.code,\n        json: RpcErr(\n                message: _message(err.AlreadyExists.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.AlreadyExists)\n            .toJson(),\n      );\n\n  static shelf.Response PermissionDenied(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.PermissionDenied.code,\n        json: RpcErr(\n                message: _message(err.PermissionDenied.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.PermissionDenied)\n            .toJson(),\n      );\n\n  static shelf.Response Unauthenticated(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unauthenticated.code,\n        json: RpcErr(\n                message: _message(err.Unauthenticated.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unauthenticated)\n            .toJson(),\n      );\n\n  static shelf.Response ResourceExhausted(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.ResourceExhausted.code,\n        json: RpcErr(\n                message: _message(err.ResourceExhausted.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.ResourceExhausted)\n            .toJson(),\n      );\n\n  static shelf.Response FailedPrecondition(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.FailedPrecondition.code,\n        json: RpcErr(\n                message: _message(err.FailedPrecondition.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.FailedPrecondition)\n            .toJson(),\n      );\n\n  static shelf.Response Aborted(String route, {String msg = ''}) => _jsonResp(\n        err.Aborted.code,\n        json: RpcErr(\n                message: _message(err.Aborted.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Aborted)\n            .toJson(),\n      );\n\n  static shelf.Response OutOfRange(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.OutOfRange.code,\n        json: RpcErr(\n                message: _message(err.OutOfRange.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.OutOfRange)\n            .toJson(),\n      );\n\n  static shelf.Response Unimplemented(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unimplemented.code,\n        json: RpcErr(\n                message: _message(err.Unimplemented.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unimplemented)\n            .toJson(),\n      );\n\n  static shelf.Response Internal(String route, {String msg = ''}) => _jsonResp(\n        err.Internal.code,\n        json: RpcErr(\n                message: _message(err.Internal.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Internal)\n            .toJson(),\n      );\n\n  static shelf.Response Unavailable(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unavailable.code,\n        json: RpcErr(\n                message: _message(err.Unavailable.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unavailable)\n            .toJson(),\n      );\n\n  static shelf.Response DataLoss(String route, {String msg = ''}) => _jsonResp(\n        err.DataLoss.code,\n        json: RpcErr(\n                message: _message(err.DataLoss.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.DataLoss)\n            .toJson(),\n      );\n}\n\n{{ end -}}PK\x07\x08R\x87\x9e\xa8N\x1d\x00\x00N\x1d\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x005\n\x9aP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00types.dart.tmplUT\x05\x00\x017\xe1\xa4^{{define \"types\"}}\n// **********************************************************************\n// MESSAGE TYPES.\n// **********************************************************************\n\n{{- if .Messages -}}\n{{range .Messages -}}\n\n\n\n\n{{if .Type | isEnum -}}\n{{$enumName := .Name}}\n@freezed\nabstract class {{$enumName}} with _${{$enumName}} {\n{{- range .Fields}}\n  const factory {{$enumName}}.{{. | makeLowerCase}}() = {{.Name}};\n{{- end}}\n  factory {{$enumName}}.fromJson(Map<String, dynamic> json) => _${{$enumName}}FromJson(json);\n}\n{{end -}}\n\n{{- if .Type | isStruct  }}\n@freezed\nabstract class {{.Name}} with _${{.Name}} {\n  const factory {{.Name}}({{if .Fields}}{\n  {{- range .Fields}}\n  {{- if not .Optional}}\n    {{if . | exportableField -}}{{. | jsonKey}} {{end}}@required {{.Type | fieldType}} {{.Name}},\n    {{- end -}}\n    {{- end -}}\n    {{- range .Fields}}\n  {{- if .Optional}}\n   {{if . | exportableField -}}{{. | jsonKey}}{{end}} {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n  {{- end }}\n  }{{ end }}) = _{{.Name}};\n  factory {{.Name}}.fromJson(Map<String, dynamic> json) => _${{.Name}}FromJson(json);\n}\n{{- end -}}\n\n{{- end -}}\n{{- end}}\n\n\n{{if .Services -}}\n{{range .Services -}}\n// *********************************************************************\n// METHOD ARGUMENT TYPES.\n// *********************************************************************\n{{- range .Methods -}}\n{{- if .Inputs }}\n@freezed\nabstract class {{. | methodArgumentInputClassName}} with _${{. | methodArgumentInputClassName}} {\n  const factory {{. | methodArgumentInputClassName}} ({\n  {{- range .Inputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Inputs */}}\n  {{- range .Inputs}}\n  {{if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n{{- end}}\n  }) = _{{. | methodArgumentInputClassName}};\n  factory {{. | methodArgumentInputClassName}}.fromJson(Map<String, dynamic> json) => _${{. | methodArgumentInputClassName}}FromJson(json);\n}\n{{- end -}} {{/* end of if .Inputs */}}\n{{- end -}} {{/* end of range .Methods */}}\n\n// *********************************************************************\n// METHOD RETURN TYPES.\n// *********************************************************************\n{{range .Methods -}}\n{{if .Outputs}}\n\n@freezed\nabstract class {{. | methodArgumentOutputClassName}} with _${{. | methodArgumentOutputClassName}} {\n  const factory {{. | methodArgumentOutputClassName}}({\n    {{- range .Outputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Outputs */}}\n  {{- range .Outputs }}\n  {{- if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n  {{- end -}}\n  }) = _{{. | methodArgumentOutputClassName}};\n  factory {{. | methodArgumentOutputClassName}}.fromJson(Map<String, dynamic> json) => _${{. | methodArgumentOutputClassName}}FromJson(json);\n}\n{{end}} {{/* end of if .Outputs */}}\n{{- end -}} {{/* end of range .Methods */}}\n{{- end -}}  {{/* end of range .Services For Inputs and Outputs */}}\n\n{{- end -}} {{/* end of if .Services */}}\n\n{{end -}} {{/* end of top level define */}}\nPK\x07\x08u]\xa42\x99\x0c\x00\x00\x99\x0c\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x005\n\x9aP\xd8n\xd9\xec\xce\x01\x00\x00\xce\x01\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.dart.tmplUT\x05\x00\x017\xe1\xa4^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00C\x0b\x9aPTv\"\xad\xfb\x1a\x00\x00\xfb\x1a\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x15\x02\x00\x00helpers.dart.tmplUT\x05\x00\x01.\xe3\xa4^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xd8\x0b\x9aPWd\xc4FW\x04\x00\x00W\x04\x00\x00\x13\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81X\x1d\x00\x00proto.gen.dart.tmplUT\x05\x00\x01H\xe4\xa4^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x005\n\x9aP$\xadW\xe7/\x0b\x00\x00/\x0b\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xf9!\x00\x00server.dart.tmplUT\x05\x00\x017\xe1\xa4^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xc3\x0b\x9aPR\x87\x9e\xa8N\x1d\x00\x00N\x1d\x00\x00\x18\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81o-\x00\x00server_helpers.dart.tmplUT\x05\x00\x01\x1f\xe4\xa4^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x005\n\x9aPu]\xa42\x99\x0c\x00\x00\x99\x0c\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x0cK\x00\x00types.dart.tmplUT\x05\x00\x017\xe1\xa4^PK\x05\x06\x00\x00\x00\x00\x06\x00\x06\x00\xb5\x01\x00\x00\xebW\x00\x00\x00\x00"
