import 'dart:async';

import 'package:meta/meta.dart';
import 'package:freezed_annotation/freezed_annotation.dart';


import 'package:shelf/shelf.dart' as shelf;
import 'package:shelf/shelf_io.dart' as io;


// example v0.0.1 ebe26b74a56342a2fb2bf54ac0c0b43300a257b0
// --
// This file has been generated by https://github.com/webrpc/webrpc using gen/dart
// Do not edit by hand. Update your webrpc schema and re-generate.

// WebRPC description and code-gen version
String webRPCVersion()  {
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
  const factory Empty() = _Empty;
  factory Empty.fromJson(Map<String, dynamic> json) => _$EmptyFromJson(json);
}
@freezed
abstract class User with _$User {
  const factory User({
     @required int id,
    @Jsonkey(name: 'USERNAME') @required String username,
     @required String role,
  }) = _User;
  factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);
}
@freezed
abstract class SearchFilter with _$SearchFilter {
  const factory SearchFilter({
     @required String q,
  }) = _SearchFilter;
  factory SearchFilter.fromJson(Map<String, dynamic> json) => _$SearchFilterFromJson(json);
}
@freezed
abstract class Version with _$Version {
  const factory Version({
     @required String webrpcVersion,
     @required String schemaVersion,
     @required String schemaHash,
  }) = _Version;
  factory Version.fromJson(Map<String, dynamic> json) => _$VersionFromJson(json);
}
@freezed
abstract class ComplexType with _$ComplexType {
  const factory ComplexType({
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
  factory ComplexType.fromJson(Map<String, dynamic> json) => _$ComplexTypeFromJson(json);
}


// *********************************************************************
// METHOD ARGUMENT TYPES.
// *********************************************************************
@freezed
abstract class ExampleServiceGetUserArgs with _$ExampleServiceGetUserArgs {
  const factory ExampleServiceGetUserArgs ({@required Map<String, String> header,@required int userID,
  
  
  }) = _ExampleServiceGetUserArgs;
  factory ExampleServiceGetUserArgs.fromJson(Map<String, dynamic> json) => _$ExampleServiceGetUserArgsFromJson(json);
}
@freezed
abstract class ExampleServiceFindUserArgs with _$ExampleServiceFindUserArgs {
  const factory ExampleServiceFindUserArgs ({@required SearchFilter s,
  
  }) = _ExampleServiceFindUserArgs;
  factory ExampleServiceFindUserArgs.fromJson(Map<String, dynamic> json) => _$ExampleServiceFindUserArgsFromJson(json);
}

// *********************************************************************
// METHOD RETURN TYPES.
// *********************************************************************
 

@freezed
abstract class ExampleServiceStatusReturn with _$ExampleServiceStatusReturn {
  const factory ExampleServiceStatusReturn({@required bool status,}) = _ExampleServiceStatusReturn;
  factory ExampleServiceStatusReturn.fromJson(Map<String, dynamic> json) => _$ExampleServiceStatusReturnFromJson(json);
}
 

@freezed
abstract class ExampleServiceVersionReturn with _$ExampleServiceVersionReturn {
  const factory ExampleServiceVersionReturn({@required Version version,}) = _ExampleServiceVersionReturn;
  factory ExampleServiceVersionReturn.fromJson(Map<String, dynamic> json) => _$ExampleServiceVersionReturnFromJson(json);
}
 

@freezed
abstract class ExampleServiceGetUserReturn with _$ExampleServiceGetUserReturn {
  const factory ExampleServiceGetUserReturn({@required int code,@required User user,}) = _ExampleServiceGetUserReturn;
  factory ExampleServiceGetUserReturn.fromJson(Map<String, dynamic> json) => _$ExampleServiceGetUserReturnFromJson(json);
}
 

@freezed
abstract class ExampleServiceFindUserReturn with _$ExampleServiceFindUserReturn {
  const factory ExampleServiceFindUserReturn({@required String name,@required User user,}) = _ExampleServiceFindUserReturn;
  factory ExampleServiceFindUserReturn.fromJson(Map<String, dynamic> json) => _$ExampleServiceFindUserReturnFromJson(json);
}
 // *********************************************************************
// METHOD ARGUMENT TYPES.
// *********************************************************************
@freezed
abstract class AnotherExampleServiceGetUserArgs with _$AnotherExampleServiceGetUserArgs {
  const factory AnotherExampleServiceGetUserArgs ({@required Map<String, String> header,@required int userID,
  
  
  }) = _AnotherExampleServiceGetUserArgs;
  factory AnotherExampleServiceGetUserArgs.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceGetUserArgsFromJson(json);
}
@freezed
abstract class AnotherExampleServiceFindUserArgs with _$AnotherExampleServiceFindUserArgs {
  const factory AnotherExampleServiceFindUserArgs ({@required SearchFilter s,
  
  }) = _AnotherExampleServiceFindUserArgs;
  factory AnotherExampleServiceFindUserArgs.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceFindUserArgsFromJson(json);
}

// *********************************************************************
// METHOD RETURN TYPES.
// *********************************************************************
 

@freezed
abstract class AnotherExampleServiceStatusReturn with _$AnotherExampleServiceStatusReturn {
  const factory AnotherExampleServiceStatusReturn({@required bool status,}) = _AnotherExampleServiceStatusReturn;
  factory AnotherExampleServiceStatusReturn.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceStatusReturnFromJson(json);
}
 

@freezed
abstract class AnotherExampleServiceVersionReturn with _$AnotherExampleServiceVersionReturn {
  const factory AnotherExampleServiceVersionReturn({@required Version version,}) = _AnotherExampleServiceVersionReturn;
  factory AnotherExampleServiceVersionReturn.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceVersionReturnFromJson(json);
}
 

@freezed
abstract class AnotherExampleServiceGetUserReturn with _$AnotherExampleServiceGetUserReturn {
  const factory AnotherExampleServiceGetUserReturn({@required int code,@required User user,}) = _AnotherExampleServiceGetUserReturn;
  factory AnotherExampleServiceGetUserReturn.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceGetUserReturnFromJson(json);
}
 

@freezed
abstract class AnotherExampleServiceFindUserReturn with _$AnotherExampleServiceFindUserReturn {
  const factory AnotherExampleServiceFindUserReturn({@required String name,@required User user,}) = _AnotherExampleServiceFindUserReturn;
  factory AnotherExampleServiceFindUserReturn.fromJson(Map<String, dynamic> json) => _$AnotherExampleServiceFindUserReturnFromJson(json);
}
 




  


// *********************************************************************
// SERVICE INTERFACES.
// *********************************************************************
abstract class ExampleService {
  FutureOr<void> ping();
  FutureOr<ExampleServiceStatusReturn> status();
  FutureOr<ExampleServiceVersionReturn> version();
  FutureOr<ExampleServiceGetUserReturn> getUser({@required ExampleServiceGetUserArgs args});
  FutureOr<ExampleServiceFindUserReturn> findUser({@required ExampleServiceFindUserArgs args});
}
abstract class AnotherExampleService {
  FutureOr<void> ping();
  FutureOr<AnotherExampleServiceStatusReturn> status();
  FutureOr<AnotherExampleServiceVersionReturn> version();
  FutureOr<AnotherExampleServiceGetUserReturn> getUser({@required AnotherExampleServiceGetUserArgs args});
  FutureOr<AnotherExampleServiceFindUserReturn> findUser({@required AnotherExampleServiceFindUserArgs args});
}
  

// *********************************************************************
// SERVER IMPLEMENTATION.
// *********************************************************************
class WebRpcServer {
  // For Google Cloud Run, set _hostname to '0.0.0.0'.
  String _hostname;
  // Provide a {Logger} implementation to log failed requests.
  Logger _log;
  // Provide a preconfigured shelf.Pipeline with desired middleware.
  Set<shelf.Middleware> _middleware;
  final ExampleService exampleService;
  final AnotherExampleService anotherExampleService;  
  WebRpcServer(
    { @required this.exampleService,
      @required this.anotherExampleService,
      Logger log,
      String hostName = 'localhost',
      List<shelf.Middleware> middleware}) {
      _hostname = hostName;
      _log = log ?? _rpcLogger;
      _middleware = middleware?.toSet() ?? [shelf.logRequests()];
  }

  bool _jsonFriendly(shelf.Request r) =>
      r.headers['Content-Type'].contains('application/json') &&
      r.headers['Accept'].contains('application/json');  

  FutureOr<shelf.Response> _requestHandler(shelf.Request r) {
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
      
      
      case '/rpc/ExampleService/Ping': {
         return _handleExampleServicePing(r);
      }
      break;
      
      case '/rpc/ExampleService/Status': {
         return _handleExampleServiceStatus(r);
      }
      break;
      
      case '/rpc/ExampleService/Version': {
         return _handleExampleServiceVersion(r);
      }
      break;
      
      case '/rpc/ExampleService/GetUser': {
         return _handleExampleServiceGetUser(r);
      }
      break;
      
      case '/rpc/ExampleService/FindUser': {
         return _handleExampleServiceFindUser(r);
      }
      break;
      
      
      
      case '/rpc/AnotherExampleService/Ping': {
         return _handleAnotherExampleServicePing(r);
      }
      break;
      
      case '/rpc/AnotherExampleService/Status': {
         return _handleAnotherExampleServiceStatus(r);
      }
      break;
      
      case '/rpc/AnotherExampleService/Version': {
         return _handleAnotherExampleServiceVersion(r);
      }
      break;
      
      case '/rpc/AnotherExampleService/GetUser': {
         return _handleAnotherExampleServiceGetUser(r);
      }
      break;
      
      case '/rpc/AnotherExampleService/FindUser': {
         return _handleAnotherExampleServiceFindUser(r);
      }
      break;
      
      
      default:
        {
          return rpcResp.BadRoute(route, msg: 'no handler for path: $route');
        }
        break;
    }
  }
  
  FutureOr<shelf.Response> _handleExampleServicePing(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleExampleServiceStatus(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleExampleServiceVersion(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleExampleServiceGetUser(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleExampleServiceFindUser(shelf.Request r) {
    
  }
  
  
  FutureOr<shelf.Response> _handleAnotherExampleServicePing(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleAnotherExampleServiceStatus(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleAnotherExampleServiceVersion(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleAnotherExampleServiceGetUser(shelf.Request r) {
    
  }
  
  FutureOr<shelf.Response> _handleAnotherExampleServiceFindUser(shelf.Request r) {
    
  }
  
  FutureOr<void> run() async {}
  //// SERVICES SHOULD END HERE DEBUG!!!!!!
}


 







//
// HELPER CODE.
//

abstract class Logger {
  void _log(message, [Object error, StackTrace stackTrace]) => print(
      '{message: $message}, error: $error, stackTrace: $stackTrace, time: ${DateTime.now()}');
  void finest(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void finer(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void fine(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void config(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void info(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void warning(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void severe(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
  void shout(message, [Object error, StackTrace stackTrace]) =>
      _log(message, error, stackTrace);
}

class _Logger extends Logger {
  _Logger();
}

final _rpcLogger = _Logger();


class HttpErr {
  final String status;
  final int code;
  const HttpErr(this.status, this.code);

  Map<String, dynamic> toMap() => {'status': status, 'code': code};
  String toJson() => jsonEncode(toMap());
  static HttpErr fromMap(Map<String, dynamic> map) =>
      HttpErr(map['status'], map['code']);

  static HttpErr fromJson(json) => fromMap(jsonDecode(json));
}

class RpcErr {
  final String message;
  final String path;
  final HttpErr httpErr;
  const RpcErr({this.message, this.path, this.httpErr});

  Map<String, dynamic> toMap() =>
      {'message': message, 'path': path, 'httpErr': httpErr.toMap()};
  String toJson() => jsonEncode(toMap());
  static RpcErr fromMap(Map<String, dynamic> map) => RpcErr(
      message: map['message'],
      path: map['path'],
      httpErr: HttpErr.fromMap(map['httpErr']));
  static RpcErr fromJson(json) => fromMap(jsonDecode(json));    
}

class Err {
  // Unknown error. For example when handling errors raised by APIs that do not
  // return enough error information.
  static HttpErr Unknown = HttpErr('unknown', 400);
  // 422 (Unprocessable Entity) Fail error. General failure error type.
  static HttpErr Fail = HttpErr('fail', 422);
  // RequestTimeout Canceled indicates the operation was cancelled (typically by the caller).
  static HttpErr Canceled = HttpErr('canceled', 408);
  // InvalidArgument indicates client specified an invalid argument. It
  // indicates arguments that are problematic regardless of the state of the
  // system (i.e. a malformed file name, required argument, number out of range,
  // etc.).
  static HttpErr InvalidArgument = HttpErr('invalid argument', 422);
  // RequestTimeOut. DeadlineExceeded means operation expired before completion. For operations
  // that change the state of the system, this error may be returned even if the
  // operation has completed successfully (timeout).
  static HttpErr DeadlineExceeded = HttpErr('deadline exceeded', 408);
  // NotFound means some requested entity was not found.
  static HttpErr NotFound = HttpErr('not found', 404);
  // BadRoute means that the requested URL path wasn't routable to a webrpc
  // service and method. This is returned by the generated server, and usually
  // shouldn't be returned by applications. Instead, applications should use
  // NotFound or Unimplemented.
  static HttpErr BadRoute = HttpErr('bad route', 404);
  // AlreadyExists means an attempt to create an entity failed because one
  // already exists. Conflict.
  static HttpErr AlreadyExists = HttpErr('already exists', 409);
  // PermissionDenied indicates the caller does not have permission to execute
  // the specified operation. It must not be used if the caller cannot be
  // identified (Unauthenticated).
  static HttpErr PermissionDenied = HttpErr('permission denied', 403);
  // Unauthenticated indicates the request does not have valid authentication
  // credentials for the operation. Unauthorized.
  static HttpErr Unauthenticated = HttpErr('unauthenticated', 401);
  // ResourceExhausted indicates some resource has been exhausted, perhaps a
  // per-user quota, or perhaps the entire file system is out of space. Forbidden.
  static HttpErr ResourceExhausted = HttpErr('resource exhausted', 403);
  // FailedPrecondition indicates operation was rejected because the system is
  // not in a state required for the operation's execution. For example, doing
  // an rmdir operation on a directory that is non-empty, or on a non-directory
  // object, or when having conflicting read-modify-write on the same resource. Precondition failed.
  static HttpErr FailedPrecondition = HttpErr('failed precondition', 412);
  // Aborted indicates the operation was aborted, typically due to a concurrency
  // issue like sequencer check failures, transaction aborts, etc.
  static HttpErr Aborted = HttpErr('aborted', 409);
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
  static HttpErr OutOfRange = HttpErr('out of range', 400);
  // Unimplemented indicates operation is not implemented or not
  // supported/enabled in this service.
  static HttpErr Unimplemented = HttpErr('unimplemented', 501);
  // Internal errors. When some invariants expected by the underlying system
  // have been broken. In other words, something bad happened in the library or
  // backend service. Do not confuse with HTTP Internal Server Error; an
  // Internal error could also happen on the client code, i.e. when parsing a
  // server response.
  static HttpErr Internal = HttpErr('internal', 500);
  // Unavailable indicates the service is currently unavailable. This is a most
  // likely a transient condition and may be corrected by retrying with a
  // backoff. Service Unavailable.
  static HttpErr Unavailable = HttpErr('unavailable', 503);
  // DataLoss indicates unrecoverable data loss or corruption.
  static HttpErr DataLoss = HttpErr('data loss', 500);
}



