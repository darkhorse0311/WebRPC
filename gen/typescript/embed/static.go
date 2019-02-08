// Code generated by statik. DO NOT EDIT.

// Package contains static assets.
package embed

var Asset = "PK\x03\x04\x14\x00\x08\x00\x00\x00\xc4mHN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00client.ts.tmplUT\x05\x00\x01!\x88]\\{{define \"client\"}}\n\n{{- if .Services}}\n// Client\n{{range .Services}}\nconst {{.Name | constPathPrefix}} = \"/rpc/{{.Name}}/\"\n{{end}}\n\n{{- range .Services}}\nexport class {{.Name}} implements {{.Name | serviceInterfaceName}} {\n  private hostname: string\n  private fetch: Fetch\n  private path = '/rpc/{{.Name}}/'\n\n  constructor(hostname: string, fetch: Fetch) {\n    this.hostname = hostname\n    this.fetch = fetch\n  }\n\n  private url(name: string): string {\n    return this.hostname + this.path + name\n  }\n\n  {{range .Methods}}\n  {{.Name}}({{.Inputs | methodInputs}}): {{.Outputs | methodOutputs}} {\n    return this.fetch(\n      this.url('{{.Name}}'),\n      {{if .Inputs | len}}\n      createHTTPRequest(params, headers)\n      {{else}}\n      createHTTPRequest({}, headers)\n      {{end}}\n    ).then((res) => {\n      if (!res.ok) {\n        return throwHTTPError(res)\n      }\n      {{range $output := .Outputs}}\n      return res.json().then((_data) => {return {{$output | newResponseConcreteType}}(_data)})\n      {{end}}\n    })\n  }\n  {{end}}\n}\n{{end -}}\n\n{{end -}}\n{{end}}\nPK\x07\x08\x1bchh(\x04\x00\x00(\x04\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xd3\xb1GN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0f\x00	\x00helpers.ts.tmplUT\x05\x00\x01\xcf\xad\\\\{{define \"helpers\"}}\n\nexport interface WebRPCErrorJSON {\n  code: string\n  msg: string\n  meta: {\n    [index: string]: string\n  }\n}\n\nexport class WebRPCError extends Error {\n  code: string\n  meta: {\n    [index: string]: string\n  }\n\n  constructor(te: WebRPCErrorJSON) {\n    super(te.msg)\n\n    this.code = te.code\n    this.meta = te.meta\n  }\n}\n\nexport const throwHTTPError = (resp: Response) => {\n  return resp.json().then((err: WebRPCErrorJSON) => { throw new WebRPCError(err) })\n}\n\nexport const createHTTPRequest = (body: object = {}, headers: object = {}): object => {\n  return {\n    method: 'POST',\n    headers: { ...headers, 'Content-Type': 'application/json' },\n    body: JSON.stringify(body || {})\n  }\n}\n\nexport type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>\n{{end}}\nPK\x07\x08d\x1eg	\x1e\x03\x00\x00\x1e\x03\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xcdmHN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x11\x00	\x00proto.gen.ts.tmplUT\x05\x00\x013\x88]\\{{define \"proto\"}}\n/* tslint:disable */\n\n// This file has been generated by https://github.com/webrpc/webrpc\n// Do not edit.\n\n{{template \"types\" .}}\n{{template \"client\" .}}\n{{template \"server\" .}}\n{{template \"helpers\" .}}\n{{end}}\nPK\x07\x08/\x83\x00\x8d\xe6\x00\x00\x00\xe6\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xb1mHN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0e\x00	\x00server.ts.tmplUT\x05\x00\x01\xfe\x87]\\{{define \"server\"}}\n{{- if .Services}}\n// TODO: Server\n{{end -}}\n{{end}}\nPK\x07\x08\x8a@[\xefI\x00\x00\x00I\x00\x00\x00PK\x03\x04\x14\x00\x08\x00\x00\x00\xdcmHN\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x0d\x00	\x00types.ts.tmplUT\x05\x00\x01P\x88]\\{{define \"types\"}}\n\n{{- if .Messages -}}\n{{range .Messages -}}\n\n{{if .Type | isEnum -}}\n{{$enumName := .Name}}\n{{range .Fields -}}\n  // {{$enumName}}_{{.Name}} = {{.Type}} {{.Value}}\n{{end -}}\n{{end -}}\n\n{{- if .Type | isStruct  }}\nexport interface {{.Name | interfaceName}} {\n  {{range .Fields -}}\n  {{.Name | exportedField}}?: {{.Type | fieldType}}\n  {{end -}}\n\n  toJSON?(): object\n}\n\nexport class {{.Name}} implements {{.Name | interfaceName}} {\n  private _data: {{.Name | interfaceName}}\n  constructor(_data?: {{.Name | interfaceName}}) {\n    this._data = {}\n    if (_data) {\n      {{range .Fields -}}\n      this._data['{{.Name | exportedField}}'] = _data['{{.Name | exportedField}}']!\n      {{end}}\n    }\n  }\n  {{ range .Fields -}}\n  public get {{.Name | exportedField}}(): {{.Type | fieldType}} {\n    return this._data['{{.Name | exportedField }}']!\n  }\n  public set {{.Name | exportedField}}(value: {{.Type | fieldType}}) {\n    this._data['{{.Name | exportedField}}'] = value\n  }\n  {{end}}\n  public toJSON(): object {\n    return this._data\n  }\n}\n{{end}}\n{{end}}\n{{end}}\n\n{{if .Services}}\n{{range .Services -}}\nexport interface {{.Name | serviceInterfaceName}} {\n  {{range .Methods -}}\n    {{.Name}}({{.Inputs | methodInputs}}): {{.Outputs | methodOutputs}}\n  {{- end}}\n}\n{{- end}}\n{{end -}}\n\n{{end}}\nPK\x07\x08\xe0%`\x8a\x1b\x05\x00\x00\x1b\x05\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xc4mHN\x1bchh(\x04\x00\x00(\x04\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x00\x00\x00\x00client.ts.tmplUT\x05\x00\x01!\x88]\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xd3\xb1GNd\x1eg	\x1e\x03\x00\x00\x1e\x03\x00\x00\x0f\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81m\x04\x00\x00helpers.ts.tmplUT\x05\x00\x01\xcf\xad\\\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xcdmHN/\x83\x00\x8d\xe6\x00\x00\x00\xe6\x00\x00\x00\x11\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\xd1\x07\x00\x00proto.gen.ts.tmplUT\x05\x00\x013\x88]\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xb1mHN\x8a@[\xefI\x00\x00\x00I\x00\x00\x00\x0e\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\xff\x08\x00\x00server.ts.tmplUT\x05\x00\x01\xfe\x87]\\PK\x01\x02\x14\x03\x14\x00\x08\x00\x00\x00\xdcmHN\xe0%`\x8a\x1b\x05\x00\x00\x1b\x05\x00\x00\x0d\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xb4\x81\x8d	\x00\x00types.ts.tmplUT\x05\x00\x01P\x88]\\PK\x05\x06\x00\x00\x00\x00\x05\x00\x05\x00\\\x01\x00\x00\xec\x0e\x00\x00\x00\x00"
