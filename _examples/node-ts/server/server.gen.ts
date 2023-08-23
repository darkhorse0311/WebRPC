/* eslint-disable */
// node-ts v1.0.0 2167bb6fb0a186ff099cd8a309c6ffbbd7d1b9c0
// --
// Code generated by webrpc-gen@v0.14.0-dev with typescript generator. DO NOT EDIT.
//
// webrpc-gen -schema=service.ridl -target=typescript -server -out=./server/server.gen.ts

// WebRPC description and code-gen version
export const WebRPCVersion = "v1"

// Schema version of your RIDL schema
export const WebRPCSchemaVersion = "v1.0.0"

// Schema hash generated from your RIDL schema
export const WebRPCSchemaHash = "2167bb6fb0a186ff099cd8a309c6ffbbd7d1b9c0"

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
  ping(headers?: object, signal?: AbortSignal): Promise<PingReturn>
  getUser(args: GetUserArgs, headers?: object, signal?: AbortSignal): Promise<GetUserReturn>
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
// Server
//

export class WebRPCError extends Error {
    statusCode?: number

    constructor(msg: string = "error", statusCode?: number) {
        super("webrpc error: " + msg);

        Object.setPrototypeOf(this, WebRPCError.prototype);

        this.statusCode = statusCode;
    }
}

import express from 'express'

        export type ExampleServiceService = {
            
                Ping: (args: PingArgs) => PingReturn | Promise<PingReturn>
            
                GetUser: (args: GetUserArgs) => GetUserReturn | Promise<GetUserReturn>
            
        }

        export const createExampleServiceApp = (serviceImplementation: ExampleServiceService) => {
            const app = express();

            app.use(express.json())

            app.post('/*', async (req, res) => {
                const requestPath = req.baseUrl + req.path

                if (!req.body) {
                    res.status(400).send("webrpc error: missing body");

                    return
                }

                switch(requestPath) {
                    

                    case "/rpc/ExampleService/Ping": {                        
                        try {
                            

                            const response = await serviceImplementation["Ping"](req.body);

                            

                            res.status(200).json(response);
                        } catch (err) {
                            if (err instanceof WebRPCError) {
                                const statusCode = err.statusCode || 400
                                const message = err.message

                                res.status(statusCode).json({
                                    msg: message,
                                    status: statusCode,
                                    code: ""
                                });

                                return
                            }

                            if (err.message) {
                                res.status(400).send(err.message);

                                return;
                            }

                            res.status(400).end();
                        }
                    }

                    return;
                    

                    case "/rpc/ExampleService/GetUser": {                        
                        try {
                            
                                    if (!("userID" in req.body)) {
                                        throw new WebRPCError("Missing Argument `userID`")
                                    }
                                if ("userID" in req.body && !validateType(req.body["userID"], "number")) {
                                    throw new WebRPCError("Invalid Argument: userID")
                                }
                            

                            const response = await serviceImplementation["GetUser"](req.body);

                            
                                if (!("code" in response)) {
                                    throw new WebRPCError("internal", 500);
                                }
                            
                                if (!("user" in response)) {
                                    throw new WebRPCError("internal", 500);
                                }
                            

                            res.status(200).json(response);
                        } catch (err) {
                            if (err instanceof WebRPCError) {
                                const statusCode = err.statusCode || 400
                                const message = err.message

                                res.status(statusCode).json({
                                    msg: message,
                                    status: statusCode,
                                    code: ""
                                });

                                return
                            }

                            if (err.message) {
                                res.status(400).send(err.message);

                                return;
                            }

                            res.status(400).end();
                        }
                    }

                    return;
                    

                    default: {
                        res.status(404).end()
                    }
                }
            });

            return app;
        };

  

const JS_TYPES = [
    "bigint",
    "boolean",
    "function",
    "number",
    "object",
    "string",
    "symbol",
    "undefined"
]


    const validateKind = (value: any) => {
        
            
                if (!("USER" in value) || !validateType(value["USER"], "number")) {
                    return false
                }
            
        
            
                if (!("ADMIN" in value) || !validateType(value["ADMIN"], "number")) {
                    return false
                }
            
        

        return true
    }

    const validateUser = (value: any) => {
        
            
                if (!("id" in value) || !validateType(value["id"], "number")) {
                    return false
                }
            
        
            
                if (!("USERNAME" in value) || !validateType(value["USERNAME"], "string")) {
                    return false
                }
            
        
            
                if (!("role" in value) || !validateType(value["role"], "Kind")) {
                    return false
                }
            
        
            
                if (!("meta" in value) || !validateType(value["meta"], "object")) {
                    return false
                }
            
        
            
                if (!("-" in value) || !validateType(value["-"], "number")) {
                    return false
                }
            
        
            
                if ("createdAt" in value && !validateType(value["createdAt"], "string")) {
                    return false
                }
            
        

        return true
    }

    const validatePage = (value: any) => {
        
            
                if (!("num" in value) || !validateType(value["num"], "number")) {
                    return false
                }
            
        

        return true
    }


const TYPE_VALIDATORS: { [type: string]: (value: any) => boolean } = {
    
        Kind: validateKind,
    
        User: validateUser,
    
        Page: validatePage,
    
}

const validateType = (value: any, type: string) => {
    if (JS_TYPES.indexOf(type) > -1) {
        return typeof value === type;
    }

    const validator = TYPE_VALIDATORS[type];

    if (!validator) {
        return false;
    }

    return validator(value);
}

