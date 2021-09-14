import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "gitopia.gitopia.gitopia";
export interface Release {
    creator: string;
    id: number;
    repositoryId: string;
    tagName: string;
    target: string;
    name: string;
    description: string;
    attachments: string;
    draft: string;
    preRelease: string;
    isTag: string;
    createdAt: string;
    updatedAt: string;
    publishedAt: string;
}
export declare const Release: {
    encode(message: Release, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Release;
    fromJSON(object: any): Release;
    toJSON(message: Release): unknown;
    fromPartial(object: DeepPartial<Release>): Release;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
