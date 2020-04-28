import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:args/args.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:meta/meta.dart';
import 'package:shelf/shelf.dart' as shelf;
import 'package:shelf/shelf_io.dart' as io;

part 'client.freezed.dart';
part 'client.g.dart';

// example v0.0.1 ebe26b74a56342a2fb2bf54ac0c0b43300a257b0
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/dart
// Do not edit by hand. Update your webrpc schema and re-generate.

// WebRPC description and code-gen version
String webRPCVersion() {
  return "v1";
}

// Schema version of your RIDL schema
String WebRPCSchemaVersion() {
  return "v0.0.1";
}

// Schema hash generated from your RIDL schema
String WebRPCSchemaHash() {
  return "ebe26b74a56342a2fb2bf54ac0c0b43300a257b0";
}

// **********************************************************************
// MESSAGE TYPES.
// **********************************************************************
@freezed
abstract class Kind with _$Kind {
  const factory Kind.user() = USER;
  const factory Kind.admin() = ADMIN;
  factory Kind.fromJson(Map<String, dynamic> json) => _$KindFromJson(json);
}

@freezed
abstract class Empty with _$Empty {
  factory Empty() = _Empty;
  factory Empty.fromJson(Map<String, dynamic> json) => _$EmptyFromJson(json);
}

@freezed
abstract class User with _$User {
  factory User({
    @required int id,
    @JsonKey(name: 'USERNAME') @required String username,
    @required String role,
  }) = _User;
  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}

@freezed
abstract class SearchFilter with _$SearchFilter {
  factory SearchFilter({
    @required String q,
  }) = _SearchFilter;
  factory SearchFilter.fromJson(Map<String, dynamic> json) =>
      _$SearchFilterFromJson(json);
}

@freezed
abstract class Version with _$Version {
  factory Version({
    @required String webrpcVersion,
    @required String schemaVersion,
    @required String schemaHash,
  }) = _Version;
  factory Version.fromJson(Map<String, dynamic> json) =>
      _$VersionFromJson(json);
}

@freezed
abstract class ComplexType with _$ComplexType {
  factory ComplexType({
    @required Map<String, dynamic> meta,
    @required Map<String, Map<String, int>> metaNestedExample,
    @required List<String> namesList,
    @required List<int> numsList,
    @required List<List<String>> doubleArray,
    @required List<User> listOfUsers,
    @required Map<String, User> mapOfUsers,
    @required User user,
    List<Map<String, int>> listOfMaps,
  }) = _ComplexType;
  factory ComplexType.fromJson(Map<String, dynamic> json) =>
      _$ComplexTypeFromJson(json);
}

// *********************************************************************
// ExampleService METHOD ARGUMENT TYPES.
// *********************************************************************
@freezed
abstract class ExampleServiceGetUserArgs with _$ExampleServiceGetUserArgs {
  factory ExampleServiceGetUserArgs({
    @required Map<String, String> header,
    @required int userID,
  }) = _ExampleServiceGetUserArgs;
  factory ExampleServiceGetUserArgs.fromJson(Map<String, dynamic> json) =>
      _$ExampleServiceGetUserArgsFromJson(json);
}

@freezed
abstract class ExampleServiceFindUserArgs with _$ExampleServiceFindUserArgs {
  factory ExampleServiceFindUserArgs({
    @required SearchFilter s,
  }) = _ExampleServiceFindUserArgs;
  factory ExampleServiceFindUserArgs.fromJson(Map<String, dynamic> json) =>
      _$ExampleServiceFindUserArgsFromJson(json);
}

// *********************************************************************
// ExampleService METHOD RETURN TYPES.
// *********************************************************************

@freezed
abstract class ExampleServiceStatusReturn with _$ExampleServiceStatusReturn {
  factory ExampleServiceStatusReturn({
    @required bool status,
  }) = _ExampleServiceStatusReturn;
}

@freezed
abstract class ExampleServiceVersionReturn with _$ExampleServiceVersionReturn {
  factory ExampleServiceVersionReturn({
    @required Version version,
  }) = _ExampleServiceVersionReturn;
}

@freezed
abstract class ExampleServiceGetUserReturn with _$ExampleServiceGetUserReturn {
  factory ExampleServiceGetUserReturn({
    @required int code,
    @required User user,
  }) = _ExampleServiceGetUserReturn;
}

@freezed
abstract class ExampleServiceFindUserReturn
    with _$ExampleServiceFindUserReturn {
  factory ExampleServiceFindUserReturn({
    @required String name,
    @required User user,
  }) = _ExampleServiceFindUserReturn;
}

// *********************************************************************
// AnotherExampleService METHOD ARGUMENT TYPES.
// *********************************************************************
@freezed
abstract class AnotherExampleServiceGetUserArgs
    with _$AnotherExampleServiceGetUserArgs {
  factory AnotherExampleServiceGetUserArgs({
    @required Map<String, String> header,
    @required int userID,
  }) = _AnotherExampleServiceGetUserArgs;
  factory AnotherExampleServiceGetUserArgs.fromJson(
          Map<String, dynamic> json) =>
      _$AnotherExampleServiceGetUserArgsFromJson(json);
}

@freezed
abstract class AnotherExampleServiceFindUserArgs
    with _$AnotherExampleServiceFindUserArgs {
  factory AnotherExampleServiceFindUserArgs({
    @required SearchFilter s,
  }) = _AnotherExampleServiceFindUserArgs;
  factory AnotherExampleServiceFindUserArgs.fromJson(
          Map<String, dynamic> json) =>
      _$AnotherExampleServiceFindUserArgsFromJson(json);
}

// *********************************************************************
// AnotherExampleService METHOD RETURN TYPES.
// *********************************************************************

@freezed
abstract class AnotherExampleServiceStatusReturn
    with _$AnotherExampleServiceStatusReturn {
  factory AnotherExampleServiceStatusReturn({
    @required bool status,
  }) = _AnotherExampleServiceStatusReturn;
}

@freezed
abstract class AnotherExampleServiceVersionReturn
    with _$AnotherExampleServiceVersionReturn {
  factory AnotherExampleServiceVersionReturn({
    @required Version version,
  }) = _AnotherExampleServiceVersionReturn;
}

@freezed
abstract class AnotherExampleServiceGetUserReturn
    with _$AnotherExampleServiceGetUserReturn {
  factory AnotherExampleServiceGetUserReturn({
    @required int code,
    @required User user,
  }) = _AnotherExampleServiceGetUserReturn;
}

@freezed
abstract class AnotherExampleServiceFindUserReturn
    with _$AnotherExampleServiceFindUserReturn {
  factory AnotherExampleServiceFindUserReturn({
    @required String name,
    @required User user,
  }) = _AnotherExampleServiceFindUserReturn;
}

// *********************************************************************
// SERVICE INTERFACES.
// *********************************************************************
// TODO implement ExampleService.
abstract class ExampleService {
  FutureOr<void> ping();
  FutureOr<ExampleServiceStatusReturn> status();
  FutureOr<ExampleServiceVersionReturn> version();
  FutureOr<ExampleServiceGetUserReturn> getUser(
      {@required ExampleServiceGetUserArgs params});
  FutureOr<ExampleServiceFindUserReturn> findUser(
      {@required ExampleServiceFindUserArgs params});
}

// TODO implement AnotherExampleService.
abstract class AnotherExampleService {
  FutureOr<void> ping();
  FutureOr<AnotherExampleServiceStatusReturn> status();
  FutureOr<AnotherExampleServiceVersionReturn> version();
  FutureOr<AnotherExampleServiceGetUserReturn> getUser(
      {@required AnotherExampleServiceGetUserArgs params});
  FutureOr<AnotherExampleServiceFindUserReturn> findUser(
      {@required AnotherExampleServiceFindUserArgs params});
}

// *********************************************************************
// SERVER IMPLEMENTATION.
// *********************************************************************
class WebRpcServer {
  // For Google Cloud Run, set _hostname to '0.0.0.0'.
  String _hostname;
  // Provide a {Logger} implementation to log exceptions.
  RpcLogger _log;
  // Provide a preconfigured shelf.Pipeline with desired middleware.
  Set<shelf.Middleware> _middleware;
  // Shelf Pipeline.
  final shelf.Pipeline _pipeline = const shelf.Pipeline();
  // A reference to the http server.
  HttpServer _server;
  // Expose internal server for user customization.
  HttpServer get server => _server;
  final ExampleService exampleService;
  final AnotherExampleService anotherExampleService;
  WebRpcServer({
    @required this.exampleService,
    @required this.anotherExampleService,
    RpcLogger logger,
    String hostName = 'localhost',
    List<shelf.Middleware> middleware,
  }) {
    _hostname = hostName;
    _log = logger ?? _rpcLogger;
    _middleware = middleware?.toSet() ?? [shelf.logRequests()];
  }

  bool _jsonFriendly(shelf.Request r) =>
      r.headers['Content-Type'].contains('application/json') &&
      r.headers['Accept'].contains('application/json');

  FutureOr<shelf.Response> _requestHandler(shelf.Request r) async {
    final route = r.url.path;
    if (r.method != 'POST') {
      final info =
          'unsupported method: ${r.method}, (only POST is allowed. path: $route';
      _log.info(info);
      return rpcResp.BadRoute(route, msg: info);
    }

    if (!_jsonFriendly(r)) {
      final info =
          'unexpected Content-Type: ${r.headers['Content-Type']} or Accept: ${r.headers['Accept']}. path: $route';
      _log.info(info);
      return rpcResp.BadRoute(route, msg: info);
    }

    switch (r.url.path) {
      case '/rpc/ExampleService/Ping':
        {
          return _handleExampleServicePing(r);
        }
        break;

      case '/rpc/ExampleService/Status':
        {
          return _handleExampleServiceStatus(r);
        }
        break;

      case '/rpc/ExampleService/Version':
        {
          return _handleExampleServiceVersion(r);
        }
        break;

      case '/rpc/ExampleService/GetUser':
        {
          return _handleExampleServiceGetUser(r);
        }
        break;

      case '/rpc/ExampleService/FindUser':
        {
          return _handleExampleServiceFindUser(r);
        }
        break;

      case '/rpc/AnotherExampleService/Ping':
        {
          return _handleAnotherExampleServicePing(r);
        }
        break;

      case '/rpc/AnotherExampleService/Status':
        {
          return _handleAnotherExampleServiceStatus(r);
        }
        break;

      case '/rpc/AnotherExampleService/Version':
        {
          return _handleAnotherExampleServiceVersion(r);
        }
        break;

      case '/rpc/AnotherExampleService/GetUser':
        {
          return _handleAnotherExampleServiceGetUser(r);
        }
        break;

      case '/rpc/AnotherExampleService/FindUser':
        {
          return _handleAnotherExampleServiceFindUser(r);
        }
        break;

      default:
        {
          final info = 'no handler for path: $route';
          _log.info(info);
          return rpcResp.BadRoute(route, msg: info);
        }
        break;
    }
  }

  FutureOr<shelf.Response> _handleExampleServicePing(shelf.Request r) async {
    try {
      // Attempt to call service method.

      await exampleService.ping();
      return rpcResp.Ok();
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Ping');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Ping');
    }
  }

  FutureOr<shelf.Response> _handleExampleServiceStatus(shelf.Request r) async {
    try {
      // Attempt to call service method.

      final ExampleServiceStatusReturn result = await exampleService.status();
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Status');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Status');
    }
  }

  FutureOr<shelf.Response> _handleExampleServiceVersion(shelf.Request r) async {
    try {
      // Attempt to call service method.

      final ExampleServiceVersionReturn result = await exampleService.version();
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Version');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/Version');
    }
  }

  FutureOr<shelf.Response> _handleExampleServiceGetUser(shelf.Request r) async {
    try {
      // Attempt to call service method.
      final json = await r.readAsString();
      final ExampleServiceGetUserArgs args =
          ExampleServiceGetUserArgs.fromJson(jsonDecode(json));
      final ExampleServiceGetUserReturn result =
          await exampleService.getUser(params: args);
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/GetUser');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/GetUser');
    }
  }

  FutureOr<shelf.Response> _handleExampleServiceFindUser(
      shelf.Request r) async {
    try {
      // Attempt to call service method.
      final json = await r.readAsString();
      final ExampleServiceFindUserArgs args =
          ExampleServiceFindUserArgs.fromJson(jsonDecode(json));
      final ExampleServiceFindUserReturn result =
          await exampleService.findUser(params: args);
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/FindUser');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/ExampleService/FindUser');
    }
  }

  FutureOr<shelf.Response> _handleAnotherExampleServicePing(
      shelf.Request r) async {
    try {
      // Attempt to call service method.

      await anotherExampleService.ping();
      return rpcResp.Ok();
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Ping');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Ping');
    }
  }

  FutureOr<shelf.Response> _handleAnotherExampleServiceStatus(
      shelf.Request r) async {
    try {
      // Attempt to call service method.

      final AnotherExampleServiceStatusReturn result =
          await anotherExampleService.status();
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Status');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Status');
    }
  }

  FutureOr<shelf.Response> _handleAnotherExampleServiceVersion(
      shelf.Request r) async {
    try {
      // Attempt to call service method.

      final AnotherExampleServiceVersionReturn result =
          await anotherExampleService.version();
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Version');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/Version');
    }
  }

  FutureOr<shelf.Response> _handleAnotherExampleServiceGetUser(
      shelf.Request r) async {
    try {
      // Attempt to call service method.
      final json = await r.readAsString();
      final AnotherExampleServiceGetUserArgs args =
          AnotherExampleServiceGetUserArgs.fromJson(jsonDecode(json));
      final AnotherExampleServiceGetUserReturn result =
          await anotherExampleService.getUser(params: args);
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/GetUser');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/GetUser');
    }
  }

  FutureOr<shelf.Response> _handleAnotherExampleServiceFindUser(
      shelf.Request r) async {
    try {
      // Attempt to call service method.
      final json = await r.readAsString();
      final AnotherExampleServiceFindUserArgs args =
          AnotherExampleServiceFindUserArgs.fromJson(jsonDecode(json));
      final AnotherExampleServiceFindUserReturn result =
          await anotherExampleService.findUser(params: args);
      return rpcResp.Ok(json: jsonEncode(result.toJson()));
    }
    // Catch WebRPCExceptions.
    on WebRPCException catch (e, stackTrace) {
      _logWebRpcExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/FindUser');
    }
    // Catch all other exceptions.
    on Exception catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      return rpcResp.Fail('/rpc/AnotherExampleService/FindUser');
    }
  }

  ArgResults _parseArgs(List<String> args) {
    var parser = ArgParser()..addOption('port', abbr: 'p');
    try {
      return parser.parse(args);
    } on ArgParserException catch (e, stackTrace) {
      _logExc(_log, e, null, stackTrace);
      print('arg parsing error occured: $e');
      rethrow;
    }
  }

  // For Google Cloud Run, we respect the PORT environment variable
  int _getPort(ArgResults args) =>
      int.tryParse(args['port'] ?? Platform.environment['PORT'] ?? '8080');

  void _configurePipeline() =>
      _middleware.forEach((mddlwr) => _pipeline.addMiddleware(mddlwr));

  Future<void> serve(List<String> args,
      {SecurityContext securityContext,
      int backlog,
      bool shared = false}) async {
    final result = _parseArgs(args);
    final port = _getPort(result);

    if (port == null) {
      stdout.writeln(
          'Could not parse port value "${port.toString()}" into a number.');
      // 64: command line usage error
      exitCode = 64;
      return;
    }

    _configurePipeline();
    final handler = _pipeline.addHandler(_requestHandler);
    _server = await io.serve(handler, _hostname, port,
        securityContext: securityContext, backlog: backlog, shared: shared);
    print('Serving at http://${_server.address.host}:${_server.port}');
  }
}

// *********************************************************************
// SERVER-SIDE HELPER CODE.
// *********************************************************************

enum RpcLogLevel {
  Info,
  Fine,
  Finer,
  Finest,
  Config,
  Warning,
  Severe,
  Shout,
}

// This exception should be thrown from all WEBRPC-DART service method implementations.
// Throwing this exception and providing an [RpcLogLevel] allows the rpc logging mechanism to log all caught excetpions at the correct level.
class WebRPCException extends HttpException {
  @override
  final String message;
  final RpcLogLevel level;
  WebRPCException(
      {this.message = 'webrpc error', this.level = RpcLogLevel.Info})
      : super('$message');
}

String _rpcLogMsg(WebRPCException exc, [Object error, StackTrace stackTrace]) =>
    '{message: ${exc.message}, level: ${exc.level}, timeStamp: ${DateTime.now().toString()}, error: $error, stackTrace: $stackTrace}';

// Helper Method for logging WebRPCExceptions.
void _logWebRpcExc(RpcLogger log, WebRPCException exc,
    [Object error, StackTrace stackTrace]) {
  switch (exc.level) {
    case RpcLogLevel.Config:
      {
        log.config(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Fine:
      {
        log.fine(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Finer:
      {
        log.finer(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Finest:
      {
        log.finest(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Info:
      {
        log.info(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Warning:
      {
        log.warning(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Severe:
      {
        log.severe(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
      break;
    case RpcLogLevel.Shout:
      {
        log.shout(_rpcLogMsg(exc, error, stackTrace), error, stackTrace);
      }
  }
}

// Contains static method helpers for handling requests.
class rpcResp {
  static String _message(String status, {String info}) =>
      'webrpc error: $status, details: $info';

  static const _jsonHeader = {
    'Content-Type': 'application/json',
    'X-Content-Type-Options': 'nosniff'
  };

  static shelf.Response _jsonResp(int code, {dynamic json}) => shelf.Response(
        code,
        headers: _jsonHeader,
        body: json,
      );

  static shelf.Response Ok({String json}) => shelf.Response.ok(
        json,
        headers: _jsonHeader,
      );

  static shelf.Response Found(String route, {String msg = ''}) =>
      shelf.Response.found(route, headers: _jsonHeader, body: msg);

  static shelf.Response MovedPerm(String route, {String msg = ''}) =>
      shelf.Response.movedPermanently(
        route,
        headers: _jsonHeader,
      );

  static shelf.Response NotModified() =>
      shelf.Response.notModified(headers: _jsonHeader);

  static shelf.Response SeeOther(String route, {String msg = ''}) =>
      shelf.Response.seeOther(
        route,
        headers: _jsonHeader,
        body: msg,
      );

  static shelf.Response Unknown(String route, {String msg = ''}) => _jsonResp(
        err.Unknown.code,
        json: RpcErr(
                message: _message(err.Unknown.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Unknown)
            .toJson(),
      );

  static shelf.Response Fail(String route, {String msg = ''}) => _jsonResp(
        err.Fail.code,
        json: RpcErr(
                message: _message(err.Fail.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Fail)
            .toJson(),
      );

  static shelf.Response Canceled(String route, {String msg = ''}) => _jsonResp(
        err.Canceled.code,
        json: RpcErr(
                message: _message(err.Canceled.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Canceled)
            .toJson(),
      );

  static shelf.Response InvalidArgument(String route, {String msg = ''}) =>
      _jsonResp(
        err.InvalidArgument.code,
        json: RpcErr(
                message: _message(err.InvalidArgument.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.InvalidArgument)
            .toJson(),
      );

  static shelf.Response DeadlineExceeded(String route, {String msg = ''}) =>
      _jsonResp(
        err.DeadlineExceeded.code,
        json: RpcErr(
                message: _message(err.DeadlineExceeded.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.DeadlineExceeded)
            .toJson(),
      );

  static shelf.Response NotFound(String route, {String msg = ''}) => _jsonResp(
        err.NotFound.code,
        json: RpcErr(
                message: _message(err.NotFound.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.NotFound)
            .toJson(),
      );

  static shelf.Response BadRoute(String route, {String msg = ''}) => _jsonResp(
        err.BadRoute.code,
        json: RpcErr(
                message: _message(err.BadRoute.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.BadRoute)
            .toJson(),
      );

  static shelf.Response AlreadyExists(String route, {String msg = ''}) =>
      _jsonResp(
        err.AlreadyExists.code,
        json: RpcErr(
                message: _message(err.AlreadyExists.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.AlreadyExists)
            .toJson(),
      );

  static shelf.Response PermissionDenied(String route, {String msg = ''}) =>
      _jsonResp(
        err.PermissionDenied.code,
        json: RpcErr(
                message: _message(err.PermissionDenied.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.PermissionDenied)
            .toJson(),
      );

  static shelf.Response Unauthenticated(String route, {String msg = ''}) =>
      _jsonResp(
        err.Unauthenticated.code,
        json: RpcErr(
                message: _message(err.Unauthenticated.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Unauthenticated)
            .toJson(),
      );

  static shelf.Response ResourceExhausted(String route, {String msg = ''}) =>
      _jsonResp(
        err.ResourceExhausted.code,
        json: RpcErr(
                message: _message(err.ResourceExhausted.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.ResourceExhausted)
            .toJson(),
      );

  static shelf.Response FailedPrecondition(String route, {String msg = ''}) =>
      _jsonResp(
        err.FailedPrecondition.code,
        json: RpcErr(
                message: _message(err.FailedPrecondition.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.FailedPrecondition)
            .toJson(),
      );

  static shelf.Response Aborted(String route, {String msg = ''}) => _jsonResp(
        err.Aborted.code,
        json: RpcErr(
                message: _message(err.Aborted.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Aborted)
            .toJson(),
      );

  static shelf.Response OutOfRange(String route, {String msg = ''}) =>
      _jsonResp(
        err.OutOfRange.code,
        json: RpcErr(
                message: _message(err.OutOfRange.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.OutOfRange)
            .toJson(),
      );

  static shelf.Response Unimplemented(String route, {String msg = ''}) =>
      _jsonResp(
        err.Unimplemented.code,
        json: RpcErr(
                message: _message(err.Unimplemented.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Unimplemented)
            .toJson(),
      );

  static shelf.Response Internal(String route, {String msg = ''}) => _jsonResp(
        err.Internal.code,
        json: RpcErr(
                message: _message(err.Internal.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Internal)
            .toJson(),
      );

  static shelf.Response Unavailable(String route, {String msg = ''}) =>
      _jsonResp(
        err.Unavailable.code,
        json: RpcErr(
                message: _message(err.Unavailable.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.Unavailable)
            .toJson(),
      );

  static shelf.Response DataLoss(String route, {String msg = ''}) => _jsonResp(
        err.DataLoss.code,
        json: RpcErr(
                message: _message(err.DataLoss.status, info: msg),
                path: route,
                time: DateTime.now(),
                httpErr: err.DataLoss)
            .toJson(),
      );
}

// *********************************************************************
// WEBRPC-DART HELPER CODE.
// *********************************************************************

String _logMsg(
  Exception exc, [
  Object error,
  StackTrace stackTrace,
]) =>
    '{message: ${exc.toString()}, timeStamp: ${DateTime.now().toString()}, error: $error, stackTrace: $stackTrace}';

void _logExc(
  RpcLogger log,
  Exception exc, [
  Object error,
  StackTrace stackTrace,
]) =>
    log.warning(_logMsg(exc, error, stackTrace));

abstract class RpcLogger {
  void _log(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      print(
          '{message: $message}, error: $error, stackTrace: $stackTrace, time: ${DateTime.now()}');
  void finest(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void finer(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void fine(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void config(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void info(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void warning(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void severe(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
  void shout(
    message, [
    Object error,
    StackTrace stackTrace,
  ]) =>
      _log(message, error, stackTrace);
}

class _Logger extends RpcLogger {
  _Logger();
}

final _rpcLogger = _Logger();

// An error in the http stack.
class HttpErr {
  final String status;
  final int code;
  const HttpErr(this.status, this.code);

  Map<String, dynamic> toMap() => {'status': status, 'code': code};
  String toJson() => jsonEncode(toMap());
  static HttpErr fromMap(Map<String, dynamic> map) =>
      HttpErr(map['status'] as String, map['code'] as int);

  static HttpErr fromJson(json) => fromMap(jsonDecode(json));
}

// An error created by the rpc server.
class RpcErr {
  final String message;
  final String path;
  final DateTime time;
  final HttpErr httpErr;
  const RpcErr({this.message, this.path, this.httpErr, this.time});

  Map<String, dynamic> toMap() => {
        'message': message,
        'path': path,
        'httpErr': httpErr.toMap(),
        'time-stamp': time.toString()
      };
  String toJson() => jsonEncode(toMap());
  static RpcErr fromMap(Map<String, dynamic> map) => RpcErr(
      message: map['message'] as String,
      path: map['path'] as String,
      time: DateTime.parse(map['time-stamp']),
      httpErr: HttpErr.fromMap(map['httpErr']));
  static RpcErr fromJson(json) => fromMap(jsonDecode(json));
}

// Contains static fields for creating and identifying http errors.
class err {
  // Unknown error. For example when handling errors raised by APIs that do not
  // return enough error information.
  static const HttpErr Unknown = HttpErr('unknown', 400);
  // 422 (Unprocessable Entity) Fail error. General failure error type.
  static const HttpErr Fail = HttpErr('fail', 422);
  // RequestTimeout Canceled indicates the operation was cancelled (typically by the caller).
  static const HttpErr Canceled = HttpErr('canceled', 408);
  // InvalidArgument indicates client specified an invalid argument. It
  // indicates arguments that are problematic regardless of the state of the
  // system (i.e. a malformed file name, required argument, number out of range,
  // etc.).
  static const HttpErr InvalidArgument = HttpErr('invalid argument', 422);
  // RequestTimeOut. DeadlineExceeded means operation expired before completion. For operations
  // that change the state of the system, this error may be returned even if the
  // operation has completed successfully (timeout).
  static const HttpErr DeadlineExceeded = HttpErr('deadline exceeded', 408);
  // NotFound means some requested entity was not found.
  static const HttpErr NotFound = HttpErr('not found', 404);
  // BadRoute means that the requested URL path wasn't routable to a webrpc
  // service and method. This is returned by the generated server, and usually
  // shouldn't be returned by applications. Instead, applications should use
  // NotFound or Unimplemented.
  static const HttpErr BadRoute = HttpErr('bad route', 404);
  // AlreadyExists means an attempt to create an entity failed because one
  // already exists. Conflict.
  static const HttpErr AlreadyExists = HttpErr('already exists', 409);
  // PermissionDenied indicates the caller does not have permission to execute
  // the specified operation. It must not be used if the caller cannot be
  // identified (Unauthenticated).
  static const HttpErr PermissionDenied = HttpErr('permission denied', 403);
  // Unauthenticated indicates the request does not have valid authentication
  // credentials for the operation. Unauthorized.
  static const HttpErr Unauthenticated = HttpErr('unauthenticated', 401);
  // ResourceExhausted indicates some resource has been exhausted, perhaps a
  // per-user quota, or perhaps the entire file system is out of space. Forbidden.
  static const HttpErr ResourceExhausted = HttpErr('resource exhausted', 403);
  // FailedPrecondition indicates operation was rejected because the system is
  // not in a state required for the operation's execution. For example, doing
  // an rmdir operation on a directory that is non-empty, or on a non-directory
  // object, or when having conflicting read-modify-write on the same resource. Precondition failed.
  static const HttpErr FailedPrecondition = HttpErr('failed precondition', 412);
  // Aborted indicates the operation was aborted, typically due to a concurrency
  // issue like sequencer check failures, transaction aborts, etc.
  static const HttpErr Aborted = HttpErr('aborted', 409);
  // OutOfRange means operation was attempted past the valid range. For example,
  // seeking or reading past end of a paginated collection.
  //
  // Unlike InvalidArgument, this error indicates a problem that may be fixed if
  // the system state changes (i.e. adding more items to the collection).
  //
  // There is a fair bit of overlap between FailedPrecondition and OutOfRange.
  // We recommend using OutOfRange (the more specific error) when it applies so
  // that callers who are iterating through a space can easily look for an
  // OutOfRange error to detect when they are done.
  static const HttpErr OutOfRange = HttpErr('out of range', 400);
  // Unimplemented indicates operation is not implemented or not
  // supported/enabled in this service.
  static const HttpErr Unimplemented = HttpErr('unimplemented', 501);
  // Internal errors. When some invariants expected by the underlying system
  // have been broken. In other words, something bad happened in the library or
  // backend service. Do not confuse with HTTP Internal Server Error; an
  // Internal error could also happen on the client code, i.e. when parsing a
  // server response.
  static const HttpErr Internal = HttpErr('internal', 500);
  // Unavailable indicates the service is currently unavailable. This is a most
  // likely a transient condition and may be corrected by retrying with a
  // backoff. Service Unavailable.
  static const HttpErr Unavailable = HttpErr('unavailable', 503);
  // DataLoss indicates unrecoverable data loss or corruption.
  static const HttpErr DataLoss = HttpErr('data loss', 500);
}
