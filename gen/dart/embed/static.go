// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00n\x95\x9dP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00client.dart.tmplUT\x05\x00\x01\xd1\xca\xa9^{{define \"client\"}}\n\n{{if .Services}}\n// *********************************************************************\n// RpcResponse TYPE.\n// *********************************************************************\n\n// This class provides type safe access to the state of an RpcRequest \n// and it's Response data. For more info See https://www.azavea.com/blog/2019/12/12/modeling-state-with-typescript/\n// See https://pub.dev/packages/freezed to learn how to use this type.\n@freezed\nabstract class RpcResponse<T> with _$RpcResponse<T> {\n  factory RpcResponse.ok({@required T data,}) = _RpcResponseOk<T>;\n  const factory RpcResponse.err({@required String reason, @required int statusCode, String stackTrace,}) = _RpcResonseErr<T>;\n  const factory RpcResponse.pending() = _RpcResponsePending<T>;\n}\n// ***********************************************************************\n// WEBRPC-DART SERVICE CLIENTS.\n// ***********************************************************************\n\n{{range .Services}}\nclass {{.Name}} {\n  final String host;\n  String _srvcPath = '/rpc/{{.Name}}/';\n  {{.Name}}({\n    this.host = 'localhost',\n  }) {\n      _srvcPath = '${_removeSlash(host)}/rpc/{{.Name}}/';\n  } \n\n  Future<http.Response> _makeRequest(String route,\n      {dynamic json = \"{}\", Map<String, String> headers,}) {\n    final path = '$_srvcPath/$route';\n    return http.post(path,\n        headers: {\n          ...?headers,\n          'Content-Type': 'application/json',\n          'Accept': 'application/json',\n        },\n        body: jsonEncode(json),);\n  }\n\n  _RpcErr _getErr(http.Response r) {\n    try {\n      return _RpcErr.fromJson(jsonDecode(r.body)); \n    } on Exception catch (_) {\n      return _RpcErr.unknown;\n    }\n  }\n \n  {{range .Methods}}\n  {{. | methodOutputsClient}} {{.Name | methodName}}({\n  {{if .Inputs}}\n  {{- range .Inputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Inputs */}}\n  {{- range .Inputs}}\n  {{if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n{{- end}} {{/* end of range .Inputs */}}\n{{- end}} {{/* end of if .Inputs */}}\n  Map<String, String> headers,\n  }) async* {\n    const num = 0;\n    while (num == 0) {\n      yield const RpcResponse.pending();\n      try {\n        {{- if .Inputs | len}} {{/* if method has args */}}\n          final _{{. | methodArgumentInputClassName}} args = _{{. | methodArgumentInputClassName}}({{- range .Inputs}}{{.Name}}:{{.Name}},{{- end}});\n          {{end}} {{/* end of if .Inputs */}}\n          final http.Response response = await _makeRequest(\n            '{{.Name}}',\n            {{- if .Inputs | len}}json: jsonEncode(args.toJson(),),{{end}}\n            headers: headers,\n          );\n\n          if (!_nonErrorcodes.contains(response.statusCode)) {\n            final _RpcErr err = _getErr(response);\n            yield RpcResponse.err(\n                reason: err.message,\n                statusCode: err.httpErr.code,\n                );\n                break;\n          }\n          yield RpcResponse.ok(data:{{if .Outputs | len}} {{. | methodArgumentOutputClassName}}.fromJson(\n              jsonDecode(response.body,),),{{else}} response.statusCode,{{end}});\n          break;    \n      } on Exception catch (e, stackTrace) {\n        yield RpcResponse.err(\n          statusCode: 400,\n          reason: e.toString(),\n          stackTrace: stackTrace.toString(),\n        );\n        break;\n      }\n    } \n \n  }\n  {{end}} {{/* end of range .Methods */}}\n\n}\n{{end}} {{/* end of range .Services */}}\n\n\n{{end}} {{/* end of if .Services */}}\n{{end}} {{/* end of top level Define */}}\nPK\x07\x08]Y\x0f6^\x0e\x00\x00^\x0e\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\x14<\x9dP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00helpers.dart.tmplUT\x05\x00\x01\x98-\xa9^{{define \"helpers\"}}\n// *********************************************************************\n// WEBRPC-DART HELPER CODE.\n// *********************************************************************\n{{if .TargetOpts.Client}}\n\nString _removeSlash(String host) => host.endsWith('/')\n    ? host.replaceRange(host.length - 1, host.length, '')\n    : host;\n\nconst Set<int> _nonErrorcodes = {\n  200,\n  201,\n  202,\n  203,\n  204,\n  205,\n  206,\n  300,\n  301,\n  302,\n  303,\n  304,\n  307,\n  308,\n};\n\n// An http error.\nclass _HttpErr {\n  final String status;\n  final int code;\n  const _HttpErr(this.status, this.code);\n\n  Map<String, dynamic> toMap() => {'status': status, 'code': code};\n  String toJson() => jsonEncode(toMap());\n  static _HttpErr fromMap(Map<String, dynamic> map) =>\n      _HttpErr(map['status'] as String, map['code'] as int);\n\n  static _HttpErr fromJson(dynamic json) => fromMap(jsonDecode(json));\n}\n\n// An error created by the rpc server.\nclass _RpcErr {\n  final String message;\n  final String path;\n  final DateTime time;\n  final _HttpErr httpErr;\n  const _RpcErr({this.message, this.path, this.httpErr, this.time});\n  static const _RpcErr unknown = _RpcErr(message: 'an unknown error has occured', path: 'unknown', httpErr: _HttpErr('unknown', 400), time: null );\n  Map<String, dynamic> toMap() => {\n        'message': message,\n        'path': path,\n        'httpErr': httpErr.toMap(),\n        'time-stamp': time.toString()\n      };\n  String toJson() => jsonEncode(toMap());\n  static _RpcErr fromMap(Map<String, dynamic> map) => _RpcErr(\n      message: map['message'] as String,\n      path: map['path'] as String,\n      time: DateTime.parse(map['time-stamp']),\n      httpErr: _HttpErr.fromMap(map['httpErr']));\n  static _RpcErr fromJson(dynamic json) => fromMap(jsonDecode(json));\n}\n{{end}}\n\n{{if .TargetOpts.Server}}\nString _logMsg(Exception exc, [Object error, StackTrace stackTrace,]) =>\n    '{message: ${exc.toString()}, timeStamp: ${DateTime.now().toString()}, error: $error, stackTrace: $stackTrace}';\n\nvoid _logExc(RpcLogger log, Exception exc,\n        [Object error, StackTrace stackTrace,]) =>\n    log.warning(_logMsg(exc, error, stackTrace));\n\nabstract class RpcLogger {\n  void _log(String message, [Object error, StackTrace stackTrace,]) => print(\n      '{message: $message}, error: $error, stackTrace: $stackTrace, time: ${DateTime.now()}');\n  void finest(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void finer(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void fine(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void config(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void info(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void warning(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void severe(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n  void shout(String message, [Object error, StackTrace stackTrace,]) =>\n      _log(message, error, stackTrace);\n}\n\nclass _Logger extends RpcLogger {\n  _Logger();\n}\n\nfinal _rpcLogger = _Logger();\n\n// Contains static fields for creating and identifying http errors.\nclass _err {\n  // Unknown error. For example when handling errors raised by APIs that do not\n  // return enough error information.\n  static const HttpErr Unknown = HttpErr('unknown', 400);\n  // 422 (Unprocessable Entity) Fail error. General failure error type.\n  static const HttpErr Fail = HttpErr('fail', 422);\n  // RequestTimeout Canceled indicates the operation was cancelled (typically by the caller).\n  static const HttpErr Canceled = HttpErr('canceled', 408);\n  // InvalidArgument indicates client specified an invalid argument. It\n  // indicates arguments that are problematic regardless of the state of the\n  // system (i.e. a malformed file name, required argument, number out of range,\n  // etc.).\n  static const HttpErr InvalidArgument = HttpErr('invalid argument', 422);\n  // RequestTimeOut. DeadlineExceeded means operation expired before completion. For operations\n  // that change the state of the system, this error may be returned even if the\n  // operation has completed successfully (timeout).\n  static const HttpErr DeadlineExceeded = HttpErr('deadline exceeded', 408);\n  // NotFound means some requested entity was not found.\n  static const HttpErr NotFound = HttpErr('not found', 404);\n  // BadRoute means that the requested URL path wasn't routable to a webrpc\n  // service and method. This is returned by the generated server, and usually\n  // shouldn't be returned by applications. Instead, applications should use\n  // NotFound or Unimplemented.\n  static const HttpErr BadRoute = HttpErr('bad route', 404);\n  // AlreadyExists means an attempt to create an entity failed because one\n  // already exists. Conflict.\n  static const HttpErr AlreadyExists = HttpErr('already exists', 409);\n  // PermissionDenied indicates the caller does not have permission to execute\n  // the specified operation. It must not be used if the caller cannot be\n  // identified (Unauthenticated).\n  static const HttpErr PermissionDenied = HttpErr('permission denied', 403);\n  // Unauthenticated indicates the request does not have valid authentication\n  // credentials for the operation. Unauthorized.\n  static const HttpErr Unauthenticated = HttpErr('unauthenticated', 401);\n  // ResourceExhausted indicates some resource has been exhausted, perhaps a\n  // per-user quota, or perhaps the entire file system is out of space. Forbidden.\n  static const HttpErr ResourceExhausted = HttpErr('resource exhausted', 403);\n  // FailedPrecondition indicates operation was rejected because the system is\n  // not in a state required for the operation's execution. For example, doing\n  // an rmdir operation on a directory that is non-empty, or on a non-directory\n  // object, or when having conflicting read-modify-write on the same resource. Precondition failed.\n  static const HttpErr FailedPrecondition = HttpErr('failed precondition', 412);\n  // Aborted indicates the operation was aborted, typically due to a concurrency\n  // issue like sequencer check failures, transaction aborts, etc.\n  static const HttpErr Aborted = HttpErr('aborted', 409);\n  // OutOfRange means operation was attempted past the valid range. For example,\n  // seeking or reading past end of a paginated collection.\n  //\n  // Unlike InvalidArgument, this error indicates a problem that may be fixed if\n  // the system state changes (i.e. adding more items to the collection).\n  //\n  // There is a fair bit of overlap between FailedPrecondition and OutOfRange.\n  // We recommend using OutOfRange (the more specific error) when it applies so\n  // that callers who are iterating through a space can easily look for an\n  // OutOfRange error to detect when they are done.\n  static const HttpErr OutOfRange = HttpErr('out of range', 400);\n  // Unimplemented indicates operation is not implemented or not\n  // supported/enabled in this service.\n  static const HttpErr Unimplemented = HttpErr('unimplemented', 501);\n  // Internal errors. When some invariants expected by the underlying system\n  // have been broken. In other words, something bad happened in the library or\n  // backend service. Do not confuse with HTTP Internal Server Error; an\n  // Internal error could also happen on the client code, i.e. when parsing a\n  // server response.\n  static const HttpErr Internal = HttpErr('internal', 500);\n  // Unavailable indicates the service is currently unavailable. This is a most\n  // likely a transient condition and may be corrected by retrying with a\n  // backoff. Service Unavailable.\n  static const HttpErr Unavailable = HttpErr('unavailable', 503);\n  // DataLoss indicates unrecoverable data loss or corruption.\n  static const HttpErr DataLoss = HttpErr('data loss', 500);\n}\n{{end}}\n{{end}}PK\x07\x08h\x02\xcb\xaf\x89\x1f\x00\x00\x89\x1f\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xd4\x1d\x9cP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x13\x00	\x00proto.gen.dart.tmplUT\x05\x00\x01!\xa7\xa7^{{- define \"proto\" -}}\nimport 'dart:async';\nimport 'dart:convert';\n{{if .TargetOpts.Server}}\nimport 'dart:io';\n{{end}}\nimport 'package:meta/meta.dart';\nimport 'package:freezed_annotation/freezed_annotation.dart';\n{{if .TargetOpts.Client}}\nimport 'package:http/http.dart' as http;\n{{if .TargetOpts.Extra | useFlutter}}\nimport 'package:flutter/foundation.dart';\n{{end}}\n{{end}}\n{{if .TargetOpts.Server}}\nimport 'package:args/args.dart';\nimport 'package:shelf/shelf.dart' as shelf;\nimport 'package:shelf/shelf_io.dart' as io;\n{{end}}\n\n{{if .TargetOpts.PkgName}}\npart '{{.TargetOpts.PkgName}}.freezed.dart';\npart '{{.TargetOpts.PkgName}}.g.dart';\n{{end}}\n\n// {{.Name}} {{.SchemaVersion}} {{.SchemaHash}}\n// --\n// This file has been generated by https://github.com/webrpc/webrpc using gen/dart\n// Do not edit by hand. Update your webrpc schema and re-generate.\n\n// WebRPC description and code-gen version\nString webRPCVersion()  {\n  return \"{{.WebRPCVersion}}\";\n}\n\n// Schema version of your RIDL schema\nString WebRPCSchemaVersion() {\n  return \"{{.SchemaVersion}}\";\n}\n\n// Schema hash generated from your RIDL schema\nString WebRPCSchemaHash() {\n  return \"{{.SchemaHash}}\";\n}\n\n{{template \"types\" .}}\n\n{{if .TargetOpts.Server}}\n  {{template \"server\" .}}\n  {{template \"server_helpers\" .}}\n{{end -}}\n\n{{if .TargetOpts.Client}}\n  {{template \"client\" .}}\n{{end -}}\n\n{{template \"helpers\" .}}\n\n{{- end}}\nPK\x07\x08\x10\xad\xa7\xcdm\x05\x00\x00m\x05\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x006\xa1\x9cP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00server.dart.tmplUT\x05\x00\x01\x89\x8d\xa8^{{define \"server\"}}\n{{- if .Services}}\n// *********************************************************************\n// SERVICE INTERFACES.\n// *********************************************************************\n{{range .Services -}}\n// TODO implement {{.Name}}.\nabstract class {{.Name}} {\n  {{- range .Methods}}\n  {{. | methodOutputs}} {{.Name | methodName}}({{if .Inputs}} {  \n  {{- range .Inputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Inputs */}}\n  {{- range .Inputs}}\n  {{if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n{{- end}}\n  } {{- end}});\n  {{- end}}\n}\n{{end}}  {{/* end of range .Services for INTERFACES */}}\n\n// *********************************************************************\n// SERVER IMPLEMENTATION.\n// *********************************************************************\nclass WebRpcServer {\n  // For Google Cloud Run, set _hostname to '0.0.0.0'.\n  String _hostname;\n  // Provide a {Logger} implementation to log exceptions.\n  RpcLogger _log;\n  // Provide a preconfigured shelf.Pipeline with desired middleware.\n  Set<shelf.Middleware> _middleware;\n  // Shelf Pipeline.\n  final shelf.Pipeline _pipeline = const shelf.Pipeline();\n  // A reference to the http server.\n  HttpServer _server;\n  // Expose internal server for user customization.\n  HttpServer get server => _server;\n{{- range .Services}}\n  final {{.Name}} {{.Name | serviceImplName}};\n{{- end}}  \n  WebRpcServer(\n    { {{ range .Services }}@required this.{{.Name | serviceImplName}},\n      {{end -}}\n      RpcLogger logger,\n      String hostName = 'localhost',\n      List<shelf.Middleware> middleware,}) {\n      _hostname = hostName;\n      _log = logger ?? _rpcLogger;\n      _middleware = middleware?.toSet() ?? [shelf.logRequests()];\n  }\n\n  bool _jsonFriendly(shelf.Request r) =>\n      r.headers['Content-Type'].contains('application/json') &&\n      r.headers['Accept'].contains('application/json');  \n\n  FutureOr<shelf.Response> _requestHandler(shelf.Request r) async {\n    final route = r.url.path;\n    if (r.method != 'POST') {\n      final info =\n          'unsupported method: ${r.method}, (only POST is allowed. path: $route';\n      _log.info(info);\n      return rpcResp.BadRoute(route, msg: info);\n    }\n\n    if (!_jsonFriendly(r)) {\n      final info =\n          'unexpected Content-Type: ${r.headers['Content-Type']} or Accept: ${r.headers['Accept']}. path: $route';\n      _log.info(info);\n      return rpcResp.BadRoute(route, msg: info);\n    }\n\n    switch (r.url.path) {\n      {{ range .Services }}\n      {{$name := .Name}}\n      {{- range .Methods}}\n      case '/rpc/{{$name}}/{{.Name}}': {\n         return _handle{{$name}}{{.Name}}(r);\n      }\n      break;\n      {{end -}} {{/* end of range .Methods */}}\n      {{end -}} {{/* end of range .Services */}}\n      default:\n        {\n          final info = 'no handler for path: $route';\n          _log.info(info);\n          return rpcResp.BadRoute(route, msg: info);\n        }\n        break;\n    }\n  }\n  {{- range .Services}}\n  {{$name := .Name}}\n  {{- range .Methods}}\n  FutureOr<shelf.Response> _handle{{$name}}{{.Name}}(shelf.Request r) async {\n    try {\n      // Attempt to call service method.\n    {{ if .Inputs|len -}}\n      final json = await r.readAsString();\n      final _{{. | methodArgumentInputClassName}} args = _{{. | methodArgumentInputClassName}}.fromJson(jsonDecode(json));\n      {{ if .Outputs | len -}}\n      final {{. | methodArgumentOutputClassName}} result = await {{$name | serviceImplName}}.{{.Name | methodName}}({{- range .Inputs}}{{.Name}}:args.{{.Name}},{{- end}});  \n      return rpcResp.Ok(json: jsonEncode(result.toJson()));\n      {{else}}\n      await {{$name | serviceImplName}}.{{.Name | methodName}}({{- range .Inputs}}{{.Name}}:args.{{.Name}},{{- end}});\n      return rpcResp.Ok();\n      {{- end -}} {{/* end if .Outputs */}}\n    {{else}}\n      {{ if .Outputs | len -}}\n      final {{. | methodArgumentOutputClassName}} result = await {{$name | serviceImplName}}.{{.Name | methodName}}();  \n      return rpcResp.Ok(json: jsonEncode(result.toJson()));\n      {{else}}\n      await {{$name | serviceImplName}}.{{.Name | methodName}}();\n      return rpcResp.Ok();\n      {{- end -}}\n    {{- end -}} {{/* end if .Inputs */}}\n    }\n    // Catch WebRPCExceptions.\n    on WebRPCException catch (e, stackTrace) {\n      _logWebRpcExc(_log, e, null, stackTrace);\n      return rpcResp.Fail('/rpc/{{$name}}/{{.Name}}');\n    }\n    // Catch all other exceptions. \n    on Exception catch (e, stackTrace) {\n      _logExc(_log, e, null, stackTrace);\n      return rpcResp.Fail('/rpc/{{$name}}/{{.Name}}');\n    }\n  }\n  {{end}}{{/* end of range .Methods */}}\n  {{- end -}} {{/* end of range .Services for Methods */}}\n  ArgResults _parseArgs(List<String> args) {\n  final parser = ArgParser()..addOption('port', abbr: 'p');\n  try {\n    return parser.parse(args);\n  } on ArgParserException catch (e, stackTrace) {\n    _logExc(_log, e, null, stackTrace);\n    print('arg parsing error occured: $e');\n    rethrow;\n  }\n}\n\n  // For Google Cloud Run, we respect the PORT environment variable\n  int _getPort(ArgResults args) =>\n      int.tryParse(args['port'] ?? Platform.environment['PORT'] ?? '8080');\n\n  void _configurePipeline() =>\n    _middleware.forEach((mddlwr) => _pipeline.addMiddleware(mddlwr));\n\n  Future<void> serve(List<String> args,\n      {SecurityContext securityContext,\n      int backlog,\n      bool shared = false}) async {\n    final result = _parseArgs(args);\n    final port = _getPort(result);\n\n    if (port == null) {\n      stdout.writeln(\n          'Could not parse port value \"${port.toString()}\" into a number.');\n      // 64: command line usage error\n      exitCode = 64;\n      return;\n    }\n\n    _configurePipeline();\n    final handler = _pipeline.addHandler(_requestHandler);\n    _server = await io.serve(handler, _hostname, port,\n        securityContext: securityContext, backlog: backlog, shared: shared);\n    print('Serving at http://${_server.address.host}:${_server.port}');\n  }\n\n\n}\n{{end}} {{/* end of if .Services */}}\n{{end}} {{/* end of top level define */}}\nPK\x07\x08R \x1aP\x11\x18\x00\x00\x11\x18\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xc2.\x9bP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x18\x00	\x00server_helpers.dart.tmplUT\x05\x00\x01}s\xa6^{{ define \"server_helpers\" }}\n// *********************************************************************\n// SERVER-SIDE HELPER CODE.\n// *********************************************************************\n\nenum RpcLogLevel {\n  Info,\n  Fine,\n  Finer,\n  Finest,\n  Config,\n  Warning,\n  Severe,\n  Shout,\n}\n\n// This exception should be thrown from all WEBRPC-DART service method implementations.\n// Throwing this exception and providing an [RpcLogLevel] allows the rpc logging mechanism to log all caught excetpions at the correct level.\nclass WebRPCException extends HttpException {\n  @override\n  final String message;\n  final RpcLogLevel level;\n  WebRPCException(\n      {this.message = 'webrpc error', this.level = RpcLogLevel.Info})\n      : super('$message');\n}\n\nString _rpcLogMsg(WebRPCException exc, [Object error, StackTrace stackTrace]) =>\n    '{message: ${exc.message}, level: ${exc.level}, timeStamp: ${DateTime.now().toString()}, error: $error, stackTrace: $stackTrace}';\n\n// Helper Method for logging WebRPCExceptions.\nvoid _logWebRpcExc(RpcLogger log, WebRPCException exc,\n    [Object error, StackTrace stackTrace]) {\n  switch (exc.level) {\n    case RpcLogLevel.Config:\n      {\n        log.config(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Fine:\n      {\n        log.fine(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Finer:\n      {\n        log.finer(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Finest:\n      {\n        log.finest(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Info:\n      {\n        log.info(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Warning:\n      {\n        log.warning(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Severe:\n      {\n        log.severe(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n      break;\n    case RpcLogLevel.Shout:\n      {\n        log.shout(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);\n      }\n  }\n}\n\n// Contains static method helpers for handling requests.\nclass rpcResp {\n  static String _message(String status, {String info}) =>\n      'webrpc error: $status, details: $info';\n\n  static const _jsonHeader = {\n    'Content-Type': 'application/json',\n    'X-Content-Type-Options': 'nosniff'\n  };\n\n  static shelf.Response _jsonResp(int code, {dynamic json}) => shelf.Response(\n        code,\n        headers: _jsonHeader,\n        body: json,\n      );\n\n  static shelf.Response Ok({String json}) => shelf.Response.ok(\n        json,\n        headers: _jsonHeader,\n      );\n\n  static shelf.Response Found(String route, {String msg = ''}) =>\n      shelf.Response.found(route, headers: _jsonHeader, body: msg);\n\n  static shelf.Response MovedPerm(String route, {String msg = ''}) =>\n      shelf.Response.movedPermanently(\n        route,\n        headers: _jsonHeader,\n      );\n\n  static shelf.Response NotModified() =>\n      shelf.Response.notModified(headers: _jsonHeader);\n\n  static shelf.Response SeeOther(String route, {String msg = ''}) =>\n      shelf.Response.seeOther(\n        route,\n        headers: _jsonHeader,\n        body: msg,\n      );\n\n  static shelf.Response Unknown(String route, {String msg = ''}) => _jsonResp(\n        err.Unknown.code,\n        json: RpcErr(\n                message: _message(err.Unknown.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unknown)\n            .toJson(),\n      );\n\n  static shelf.Response Fail(String route, {String msg = ''}) => _jsonResp(\n        err.Fail.code,\n        json: RpcErr(\n                message: _message(err.Fail.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Fail)\n            .toJson(),\n      );\n\n  static shelf.Response Canceled(String route, {String msg = ''}) => _jsonResp(\n        err.Canceled.code,\n        json: RpcErr(\n                message: _message(err.Canceled.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Canceled)\n            .toJson(),\n      );\n\n  static shelf.Response InvalidArgument(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.InvalidArgument.code,\n        json: RpcErr(\n                message: _message(err.InvalidArgument.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.InvalidArgument)\n            .toJson(),\n      );\n\n  static shelf.Response DeadlineExceeded(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.DeadlineExceeded.code,\n        json: RpcErr(\n                message: _message(err.DeadlineExceeded.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.DeadlineExceeded)\n            .toJson(),\n      );\n\n  static shelf.Response NotFound(String route, {String msg = ''}) => _jsonResp(\n        err.NotFound.code,\n        json: RpcErr(\n                message: _message(err.NotFound.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.NotFound)\n            .toJson(),\n      );\n\n  static shelf.Response BadRoute(String route, {String msg = ''}) => _jsonResp(\n        err.BadRoute.code,\n        json: RpcErr(\n                message: _message(err.BadRoute.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.BadRoute)\n            .toJson(),\n      );\n\n  static shelf.Response AlreadyExists(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.AlreadyExists.code,\n        json: RpcErr(\n                message: _message(err.AlreadyExists.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.AlreadyExists)\n            .toJson(),\n      );\n\n  static shelf.Response PermissionDenied(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.PermissionDenied.code,\n        json: RpcErr(\n                message: _message(err.PermissionDenied.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.PermissionDenied)\n            .toJson(),\n      );\n\n  static shelf.Response Unauthenticated(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unauthenticated.code,\n        json: RpcErr(\n                message: _message(err.Unauthenticated.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unauthenticated)\n            .toJson(),\n      );\n\n  static shelf.Response ResourceExhausted(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.ResourceExhausted.code,\n        json: RpcErr(\n                message: _message(err.ResourceExhausted.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.ResourceExhausted)\n            .toJson(),\n      );\n\n  static shelf.Response FailedPrecondition(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.FailedPrecondition.code,\n        json: RpcErr(\n                message: _message(err.FailedPrecondition.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.FailedPrecondition)\n            .toJson(),\n      );\n\n  static shelf.Response Aborted(String route, {String msg = ''}) => _jsonResp(\n        err.Aborted.code,\n        json: RpcErr(\n                message: _message(err.Aborted.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Aborted)\n            .toJson(),\n      );\n\n  static shelf.Response OutOfRange(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.OutOfRange.code,\n        json: RpcErr(\n                message: _message(err.OutOfRange.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.OutOfRange)\n            .toJson(),\n      );\n\n  static shelf.Response Unimplemented(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unimplemented.code,\n        json: RpcErr(\n                message: _message(err.Unimplemented.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unimplemented)\n            .toJson(),\n      );\n\n  static shelf.Response Internal(String route, {String msg = ''}) => _jsonResp(\n        err.Internal.code,\n        json: RpcErr(\n                message: _message(err.Internal.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Internal)\n            .toJson(),\n      );\n\n  static shelf.Response Unavailable(String route, {String msg = ''}) =>\n      _jsonResp(\n        err.Unavailable.code,\n        json: RpcErr(\n                message: _message(err.Unavailable.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.Unavailable)\n            .toJson(),\n      );\n\n  static shelf.Response DataLoss(String route, {String msg = ''}) => _jsonResp(\n        err.DataLoss.code,\n        json: RpcErr(\n                message: _message(err.DataLoss.status, info: msg),\n                path: route,\n                time: DateTime.now(),\n                httpErr: err.DataLoss)\n            .toJson(),\n      );\n}\n\n{{ end -}}PK\x07\x08?\xad\xdd\xc5\xc3%\x00\x00\xc3%\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xca\x94\x9dP\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00types.dart.tmplUT\x05\x00\x01\x9c\xc9\xa9^{{define \"types\"}}\n// **********************************************************************\n// MESSAGE TYPES.\n// **********************************************************************\n{{- if .Messages -}}\n{{range .Messages -}}\n\n{{if .Type | isEnum -}}\n{{$enumName := .Name}}\n@freezed\nabstract class {{$enumName}} with _${{$enumName}} {\n{{- range .Fields}}\n  const factory {{$enumName}}.{{. | makeLowerCase}}() = {{.Name}};\n{{- end}}\n  factory {{$enumName}}.fromJson(Map<String, dynamic> json) => _${{$enumName}}FromJson(json);\n}\n{{end -}}\n\n{{- if .Type | isStruct  }}\n@freezed\nabstract class {{.Name}} with _${{.Name}} {\n  @JsonSerializable(explicitToJson: true)\n  factory {{.Name}}({{if .Fields}}{\n  {{- range .Fields}}\n  {{- if not .Optional}}\n    {{if . | exportableField -}}{{. | jsonKey}} {{end}}@required {{.Type | fieldType}} {{.Name}},\n    {{- end -}}\n    {{- end -}}\n    {{- range .Fields}}\n  {{- if .Optional}}\n   {{if . | exportableField -}}{{. | jsonKey}}{{end}} {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n  {{- end }}\n  }{{ end }}) = _{{.Name}};\n  factory {{.Name}}.fromJson(Map<String, dynamic> json) => _${{.Name}}FromJson(json);\n}\n{{- end -}}\n\n{{- end -}}\n{{- end}}\n\n\n{{if .Services -}}\n{{$isClient := .TargetOpts.Client}}\n{{range .Services -}}\n// *********************************************************************\n// {{.Name}} METHOD ARGUMENT TYPES.\n// *********************************************************************\n{{- range .Methods -}}\n{{- if .Inputs }}\n@freezed\nabstract class _{{. | methodArgumentInputClassName}} with _$_{{. | methodArgumentInputClassName}} {\n  @JsonSerializable(explicitToJson: true)\n  factory _{{. | methodArgumentInputClassName}} ({\n  {{- range .Inputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Inputs */}}\n  {{- range .Inputs}}\n  {{if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n{{- end}}\n  }) = _{{. | methodArgumentInputClassName}}_Freezed;\n  factory _{{. | methodArgumentInputClassName}}.fromJson(Map<String, dynamic> json) => _$_{{. | methodArgumentInputClassName}}FromJson(json);\n}\n{{- end -}} {{/* end of if .Inputs */}}\n{{- end -}} {{/* end of range .Methods */}}\n\n// *********************************************************************\n// {{.Name}} METHOD RETURN TYPES.\n// *********************************************************************\n{{range .Methods -}}\n{{if .Outputs}}\n\n@freezed\nabstract class {{. | methodArgumentOutputClassName}} with _${{. | methodArgumentOutputClassName}} {\n  @JsonSerializable(explicitToJson: true)\n  factory {{. | methodArgumentOutputClassName}}({\n    {{- range .Outputs}}\n  {{- if not .Optional}}@required {{.Type | fieldType}} {{.Name}},\n  {{- end -}} {{/* end of if not .Optional */}}\n  {{- end -}} {{/* end of range .Outputs */}}\n  {{- range .Outputs }}\n  {{- if .Optional}}\n  {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n  {{- end -}}\n  }) = _{{. | methodArgumentOutputClassName}};\n  factory {{. | methodArgumentOutputClassName}}.fromJson(Map<String, dynamic> json) => _${{. | methodArgumentOutputClassName}}FromJson(json);\n}\n{{end}} {{/* end of if .Outputs */}}\n{{- end -}} {{/* end of range .Methods */}}\n{{- end -}}  {{/* end of range .Services For Inputs and Outputs */}}\n\n{{- end -}} {{/* end of if .Services */}}\n\n{{- end -}} {{/* end of top level define */}}\nPK\x07\x08\x04h\xd9\x16H\x0d\x00\x00H\x0d\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00n\x95\x9dP]Y\x0f6^\x0e\x00\x00^\x0e\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.dart.tmplUT\x05\x00\x01\xd1\xca\xa9^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\x14<\x9dPh\x02\xcb\xaf\x89\x1f\x00\x00\x89\x1f\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xa5\x0e\x00\x00helpers.dart.tmplUT\x05\x00\x01\x98-\xa9^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xd4\x1d\x9cP\x10\xad\xa7\xcdm\x05\x00\x00m\x05\x00\x00\x13\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81v.\x00\x00proto.gen.dart.tmplUT\x05\x00\x01!\xa7\xa7^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x006\xa1\x9cPR \x1aP\x11\x18\x00\x00\x11\x18\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81-4\x00\x00server.dart.tmplUT\x05\x00\x01\x89\x8d\xa8^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xc2.\x9bP?\xad\xdd\xc5\xc3%\x00\x00\xc3%\x00\x00\x18\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x85L\x00\x00server_helpers.dart.tmplUT\x05\x00\x01}s\xa6^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xca\x94\x9dP\x04h\xd9\x16H\x0d\x00\x00H\x0d\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x97r\x00\x00types.dart.tmplUT\x05\x00\x01\x9c\xc9\xa9^PK\x05\x06\x00\x00\x00\x00\x06\x00\x06\x00\xb5\x01\x00\x00%\x80\x00\x00\x00\x00"
