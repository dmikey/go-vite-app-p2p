/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "main";

export interface API {
  EndPoints: API_EndPoints | undefined;
}

export interface API_EndPoints {
  GetMetaData: string;
}

export interface AppMeta {
  name: string;
}

function createBaseAPI(): API {
  return { EndPoints: undefined };
}

export const API = {
  encode(message: API, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.EndPoints !== undefined) {
      API_EndPoints.encode(message.EndPoints, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): API {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAPI();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.EndPoints = API_EndPoints.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): API {
    return { EndPoints: isSet(object.EndPoints) ? API_EndPoints.fromJSON(object.EndPoints) : undefined };
  },

  toJSON(message: API): unknown {
    const obj: any = {};
    if (message.EndPoints !== undefined) {
      obj.EndPoints = API_EndPoints.toJSON(message.EndPoints);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<API>, I>>(base?: I): API {
    return API.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<API>, I>>(object: I): API {
    const message = createBaseAPI();
    message.EndPoints = (object.EndPoints !== undefined && object.EndPoints !== null)
      ? API_EndPoints.fromPartial(object.EndPoints)
      : undefined;
    return message;
  },
};

function createBaseAPI_EndPoints(): API_EndPoints {
  return { GetMetaData: "" };
}

export const API_EndPoints = {
  encode(message: API_EndPoints, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.GetMetaData !== "") {
      writer.uint32(10).string(message.GetMetaData);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): API_EndPoints {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAPI_EndPoints();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.GetMetaData = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): API_EndPoints {
    return { GetMetaData: isSet(object.GetMetaData) ? globalThis.String(object.GetMetaData) : "" };
  },

  toJSON(message: API_EndPoints): unknown {
    const obj: any = {};
    if (message.GetMetaData !== "") {
      obj.GetMetaData = message.GetMetaData;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<API_EndPoints>, I>>(base?: I): API_EndPoints {
    return API_EndPoints.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<API_EndPoints>, I>>(object: I): API_EndPoints {
    const message = createBaseAPI_EndPoints();
    message.GetMetaData = object.GetMetaData ?? "";
    return message;
  },
};

function createBaseAppMeta(): AppMeta {
  return { name: "" };
}

export const AppMeta = {
  encode(message: AppMeta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AppMeta {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAppMeta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AppMeta {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: AppMeta): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AppMeta>, I>>(base?: I): AppMeta {
    return AppMeta.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AppMeta>, I>>(object: I): AppMeta {
    const message = createBaseAppMeta();
    message.name = object.name ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
