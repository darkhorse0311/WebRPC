"use strict";
/* tslint:disable */
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
Object.defineProperty(exports, "__esModule", { value: true });
// This file has been generated by https://github.com/webrpc/webrpc
// Do not edit.
var Kind;
(function (Kind) {
    Kind[Kind["Kind_USER"] = 1] = "Kind_USER";
    Kind[Kind["Kind_ADMIN"] = 2] = "Kind_ADMIN";
})(Kind = exports.Kind || (exports.Kind = {}));
exports.Kind_name = {
    '1': 'USER',
    '2': 'ADMIN'
};
var Empty = /** @class */ (function () {
    function Empty(_data) {
        this._data = {};
        if (_data) {
        }
    }
    Empty.prototype.toJSON = function () {
        return this._data;
    };
    return Empty;
}());
exports.Empty = Empty;
var GetUserRequest = /** @class */ (function () {
    function GetUserRequest(_data) {
        this._data = {};
        if (_data) {
            this._data['userID'] = _data['userID'];
        }
    }
    Object.defineProperty(GetUserRequest.prototype, "UserID", {
        get: function () {
            return this._data['userID'];
        },
        set: function (value) {
            this._data['userID'] = value;
        },
        enumerable: true,
        configurable: true
    });
    GetUserRequest.prototype.toJSON = function () {
        return this._data;
    };
    return GetUserRequest;
}());
exports.GetUserRequest = GetUserRequest;
var User = /** @class */ (function () {
    function User(_data) {
        this._data = {};
        if (_data) {
            this._data['id'] = _data['id'];
            this._data['USERNAME'] = _data['USERNAME'];
            this._data['created_at'] = _data['created_at'];
        }
    }
    Object.defineProperty(User.prototype, "ID", {
        get: function () {
            return this._data['id'];
        },
        set: function (value) {
            this._data['id'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(User.prototype, "Username", {
        get: function () {
            return this._data['USERNAME'];
        },
        set: function (value) {
            this._data['USERNAME'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(User.prototype, "CreatedAt", {
        get: function () {
            return this._data['created_at'];
        },
        set: function (value) {
            this._data['created_at'] = value;
        },
        enumerable: true,
        configurable: true
    });
    User.prototype.toJSON = function () {
        return this._data;
    };
    return User;
}());
exports.User = User;
var RandomStuff = /** @class */ (function () {
    function RandomStuff(_data) {
        this._data = {};
        if (_data) {
            this._data['meta'] = _data['meta'];
            this._data['metaNestedExample'] = _data['metaNestedExample'];
            this._data['namesList'] = _data['namesList'];
            this._data['numsList'] = _data['numsList'];
            this._data['doubleArray'] = _data['doubleArray'];
            this._data['listOfMaps'] = _data['listOfMaps'];
            this._data['listOfUsers'] = _data['listOfUsers'];
            this._data['mapOfUsers'] = _data['mapOfUsers'];
            this._data['user'] = _data['user'];
        }
    }
    Object.defineProperty(RandomStuff.prototype, "Meta", {
        get: function () {
            return this._data['meta'];
        },
        set: function (value) {
            this._data['meta'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "MetaNestedExample", {
        get: function () {
            return this._data['metaNestedExample'];
        },
        set: function (value) {
            this._data['metaNestedExample'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "NamesList", {
        get: function () {
            return this._data['namesList'];
        },
        set: function (value) {
            this._data['namesList'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "NumsList", {
        get: function () {
            return this._data['numsList'];
        },
        set: function (value) {
            this._data['numsList'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "DoubleArray", {
        get: function () {
            return this._data['doubleArray'];
        },
        set: function (value) {
            this._data['doubleArray'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "ListOfMaps", {
        get: function () {
            return this._data['listOfMaps'];
        },
        set: function (value) {
            this._data['listOfMaps'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "ListOfUsers", {
        get: function () {
            return this._data['listOfUsers'];
        },
        set: function (value) {
            this._data['listOfUsers'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "MapOfUsers", {
        get: function () {
            return this._data['mapOfUsers'];
        },
        set: function (value) {
            this._data['mapOfUsers'] = value;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(RandomStuff.prototype, "User", {
        get: function () {
            return this._data['user'];
        },
        set: function (value) {
            this._data['user'] = value;
        },
        enumerable: true,
        configurable: true
    });
    RandomStuff.prototype.toJSON = function () {
        return this._data;
    };
    return RandomStuff;
}());
exports.RandomStuff = RandomStuff;
// Client
var ExampleServicePathPrefix = "/rpc/ExampleService/";
var ExampleService = /** @class */ (function () {
    function ExampleService(hostname, fetch) {
        this.path = '/rpc/ExampleService/';
        this.hostname = hostname;
        this.fetch = fetch;
    }
    ExampleService.prototype.url = function (name) {
        return this.hostname + this.path + name;
    };
    ExampleService.prototype.Ping = function (headers) {
        if (headers === void 0) { headers = {}; }
        return this.fetch(this.url('Ping'), exports.createHTTPRequest({}, headers)).then(function (res) {
            if (!res.ok) {
                return exports.throwHTTPError(res);
            }
            return res.json().then(function (_data) { return (_data); });
        });
    };
    ExampleService.prototype.GetUser = function (params, headers) {
        if (headers === void 0) { headers = {}; }
        return this.fetch(this.url('GetUser'), exports.createHTTPRequest(params, headers)).then(function (res) {
            if (!res.ok) {
                return exports.throwHTTPError(res);
            }
            return res.json().then(function (_data) { return new User(_data); });
        });
    };
    return ExampleService;
}());
exports.ExampleService = ExampleService;
exports.throwHTTPError = function (resp) {
    return resp.json().then(function (err) { throw err; });
};
exports.createHTTPRequest = function (body, headers) {
    if (body === void 0) { body = {}; }
    if (headers === void 0) { headers = {}; }
    return {
        method: 'POST',
        headers: __assign({}, headers, { 'Content-Type': 'application/json' }),
        body: JSON.stringify(body || {})
    };
};
