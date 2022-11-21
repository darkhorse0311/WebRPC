/* eslint-disable */
// node-ts v1.0.0 ad568d12a684eb48ae2b38f98a2933cd65f7a4fc
// --
// Code generated by webrpc-gen@v0.7.x-dev with typescript generator. DO NOT EDIT.
//
// webrpc-gen -schema=service.ridl -target=typescript -client -out=./webapp/client.gen.ts

// WebRPC description and code-gen version
export const WebRPCVersion = "v1"

// Schema version of your RIDL schema
export const WebRPCSchemaVersion = "v1.0.0"

// Schema hash generated from your RIDL schema
export const WebRPCSchemaHash = "ad568d12a684eb48ae2b38f98a2933cd65f7a4fc"

//
// Types
//
export enum Kind {
  USER = 'USER',
  ADMIN = 'ADMIN'
}

export interface User {
  id: number
  USERNAME: string
  role: Kind
  meta: {[key: string]: any}
  
  createdAt?: string
}

export interface Page {
  num: number
}

export interface ExampleService {
  ping(headers?: object): Promise<PingReturn>
  getUser(args: GetUserArgs, headers?: object): Promise<GetUserReturn>
}

export interface PingArgs {
}

export interface PingReturn {  
}
export interface GetUserArgs {
  userID: number
}

export interface GetUserReturn {
  code: number
  user: User  
}


  
//
// Client
//
export class ExampleService implements ExampleService {
  protected hostname: string
  protected fetch: Fetch
  protected path = '/rpc/ExampleService/'

  constructor(hostname: string, fetch: Fetch) {
    this.hostname = hostname
    this.fetch = (input: RequestInfo, init?: RequestInit) => fetch(input, init)
  }

  private url(name: string): string {
    return this.hostname + this.path + name
  }
  
  ping = (headers?: object): Promise<PingReturn> => {
    return this.fetch(
      this.url('Ping'),
      createHTTPRequest({}, headers)
      ).then((res) => {
      return buildResponse(res).then(_data => {
        return {
        }
      })
    })
  }
  
  getUser = (args: GetUserArgs, headers?: object): Promise<GetUserReturn> => {
    return this.fetch(
      this.url('GetUser'),
      createHTTPRequest(args, headers)).then((res) => {
      return buildResponse(res).then(_data => {
        return {
          code: <number>(_data.code), 
          user: <User>(_data.user)
        }
      })
    })
  }
  
}

  
export interface WebRPCError extends Error {
  code: string
  msg: string
	status: number
}

const createHTTPRequest = (body: object = {}, headers: object = {}): object => {
  return {
    method: 'POST',
    headers: { ...headers, 'Content-Type': 'application/json' },
    body: JSON.stringify(body || {})
  }
}

const buildResponse = (res: Response): Promise<any> => {
  return res.text().then(text => {
    let data
    try {
      data = JSON.parse(text)
    } catch(err) {
      throw { code: 'unknown', msg: `expecting JSON, got: ${text}`, status: res.status } as WebRPCError
    }
    if (!res.ok) {
      throw data // webrpc error response
    }
    return data
  })
}

export type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>
