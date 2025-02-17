// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file proto/counter.proto (package counter.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * You can leave this empty if you just always increment by 1
 *
 * @generated from message counter.v1.IncrementRequest
 */
export class IncrementRequest extends Message<IncrementRequest> {
  constructor(data?: PartialMessage<IncrementRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "counter.v1.IncrementRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IncrementRequest {
    return new IncrementRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IncrementRequest {
    return new IncrementRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IncrementRequest {
    return new IncrementRequest().fromJsonString(jsonString, options);
  }

  static equals(a: IncrementRequest | PlainMessage<IncrementRequest> | undefined, b: IncrementRequest | PlainMessage<IncrementRequest> | undefined): boolean {
    return proto3.util.equals(IncrementRequest, a, b);
  }
}

/**
 * @generated from message counter.v1.IncrementResponse
 */
export class IncrementResponse extends Message<IncrementResponse> {
  /**
   * @generated from field: int64 new_value = 1;
   */
  newValue = protoInt64.zero;

  constructor(data?: PartialMessage<IncrementResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "counter.v1.IncrementResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "new_value", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IncrementResponse {
    return new IncrementResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IncrementResponse {
    return new IncrementResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IncrementResponse {
    return new IncrementResponse().fromJsonString(jsonString, options);
  }

  static equals(a: IncrementResponse | PlainMessage<IncrementResponse> | undefined, b: IncrementResponse | PlainMessage<IncrementResponse> | undefined): boolean {
    return proto3.util.equals(IncrementResponse, a, b);
  }
}

