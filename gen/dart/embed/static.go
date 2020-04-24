// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var	Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00client.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^{{define \"client\"}}\n{{if .Services}}\n//\n// Client\n//\n\nString _removeSlash(String host) => host.endsWith('/')\n? host.replaceRange(host.length - 1, host.length, '')\n: host;\n\n{{range .Services}}\nclass {{.Name}} {\n  final http.Client client;\n  final String host;\n  final String path = '/rpc/{{.Name}}/';\n  String url(String name) => '${_removeSlash(host)}$path$name';\n  {{range .Methods}}\n\n  {{end}}\n  {{.Name}}(this.client, this.host);\n}\n{{end}}\n\n\n{{end}}\n{{end}}\nPK\x07\x08\xee$m\xf8\xcd\x01\x00\x00\xcd\x01\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00helpers.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^{{define \"helpers\"}}\n\n//\n// Helpers\n//\n\n\n{{end}}PK\x07\x08EC\xf4k0\x00\x00\x000\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x13\x00	\x00proto.gen.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^{{- define \"proto\" -}}\n\nimport 'package:meta/meta.dart';\n{{if .TargetOpts.Client}}\nimport 'package:http/http.dart' as http;\n{{end}}\n\n// {{.Name}} {{.SchemaVersion}} {{.SchemaHash}}\n// --\n// This file has been generated by https://github.com/webrpc/webrpc using gen/golang\n// Do not edit by hand. Update your webrpc schema and re-generate.\n\n// WebRPC description and code-gen version\nString webRPCVersion()  {\n  return \"{{.WebRPCVersion}}\";\n}\n\n// Schema version of your RIDL schema\nString WebRPCSchemaVersion() {\n  return \"{{.SchemaVersion}}\";\n}\n\n// Schema hash generated from your RIDL schema\nString WebRPCSchemaHash() {\n  return \"{{.SchemaHash}}\";\n}\n\n{{template \"types\" .}}\n\n{{if .TargetOpts.Server}}\n  {{template \"server\" .}}\n{{end}}\n\n{{if .TargetOpts.Client}}\n  {{template \"client\" .}}\n{{end}}\n\n{{template \"helpers\" .}}\n\n{{- end}}\nPK\x07\x086*\xacaB\x03\x00\x00B\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x10\x00	\x00server.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^{{define \"server\"}}\n{{if .Services}}\n//\n// Server\n//\n\n{{end}}\n{{end}}\nPK\x07\x08)B\x17\xc4F\x00\x00\x00F\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00y\x0c\x98P\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00types.dart.tmplUT\x05\x00\x01vB\xa2^{{define \"types\"}}\n//\n// Types\n//\n\n{{- if .Messages -}}\n{{range .Messages -}}\n\n\n\n\n{{if .Type | isEnum -}}\n{{$enumName := .Name}}\nenum {{$enumName}} {\n{{- range $i, $field := .Fields}}\n  {{- if $i}},{{end}}\n  {{$field.Name}}\n{{- end}}\n}\n{{end -}}\n\n{{- if .Type | isStruct  }}\n@freezed\nabstract class {{.Name}} with _${{.Name}} {\n  const factory {{.Name}}({{if .Fields}}{\n  {{- range .Fields}}\n  {{- if not .Optional}}\n    {{if . | exportableField -}}{{. | jsonKey}} {{end}}@required {{.Type | fieldType}} {{.Name}},\n    {{- end -}}\n    {{- end -}}\n    {{- range .Fields}}\n  {{- if .Optional}}\n   {{if . | exportableField -}}{{. | jsonKey}}{{end}} {{.Type | fieldType}} {{.Name}},\n  {{- end -}}\n  {{- end }}\n  }{{ end }}) = _{{.Name}};\n}\n{{- end -}}\n\n\n{{end -}}\n{{end -}}\n{{end -}}\nPK\x07\x08\xa2\xa38\x0e\x0c\x03\x00\x00\x0c\x03\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P\xee$m\xf8\xcd\x01\x00\x00\xcd\x01\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00client.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xa1\x1e\x96PEC\xf4k0\x00\x00\x000\x00\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x14\x02\x00\x00helpers.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P6*\xacaB\x03\x00\x00B\x03\x00\x00\x13\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x8c\x02\x00\x00proto.gen.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xa1\x1e\x96P)B\x17\xc4F\x00\x00\x00F\x00\x00\x00\x10\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x18\x06\x00\x00server.dart.tmplUT\x05\x00\x01\x9e\xbf\x9f^PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00y\x0c\x98P\xa2\xa38\x0e\x0c\x03\x00\x00\x0c\x03\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\xa5\x06\x00\x00types.dart.tmplUT\x05\x00\x01vB\xa2^PK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00f\x01\x00\x00\xf7	\x00\x00\x00\x00"
