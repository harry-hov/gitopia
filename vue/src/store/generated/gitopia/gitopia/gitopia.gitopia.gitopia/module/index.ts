// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateIssue } from "./types/gitopia/tx";
import { MsgCreateUser } from "./types/gitopia/tx";
import { MsgDeleteRepository } from "./types/gitopia/tx";
import { MsgDeleteWhois } from "./types/gitopia/tx";
import { MsgDeleteIssue } from "./types/gitopia/tx";
import { MsgUpdateWhois } from "./types/gitopia/tx";
import { MsgUpdateRepository } from "./types/gitopia/tx";
import { MsgDeleteUser } from "./types/gitopia/tx";
import { MsgSetWhois } from "./types/gitopia/tx";
import { MsgCreateRepository } from "./types/gitopia/tx";
import { MsgUpdateUser } from "./types/gitopia/tx";
import { MsgCreateIssue } from "./types/gitopia/tx";


const types = [
  ["/gitopia.gitopia.gitopia.MsgUpdateIssue", MsgUpdateIssue],
  ["/gitopia.gitopia.gitopia.MsgCreateUser", MsgCreateUser],
  ["/gitopia.gitopia.gitopia.MsgDeleteRepository", MsgDeleteRepository],
  ["/gitopia.gitopia.gitopia.MsgDeleteWhois", MsgDeleteWhois],
  ["/gitopia.gitopia.gitopia.MsgDeleteIssue", MsgDeleteIssue],
  ["/gitopia.gitopia.gitopia.MsgUpdateWhois", MsgUpdateWhois],
  ["/gitopia.gitopia.gitopia.MsgUpdateRepository", MsgUpdateRepository],
  ["/gitopia.gitopia.gitopia.MsgDeleteUser", MsgDeleteUser],
  ["/gitopia.gitopia.gitopia.MsgSetWhois", MsgSetWhois],
  ["/gitopia.gitopia.gitopia.MsgCreateRepository", MsgCreateRepository],
  ["/gitopia.gitopia.gitopia.MsgUpdateUser", MsgUpdateUser],
  ["/gitopia.gitopia.gitopia.MsgCreateIssue", MsgCreateIssue],
  
];

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgUpdateIssue: (data: MsgUpdateIssue): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateIssue", value: data }),
    msgCreateUser: (data: MsgCreateUser): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateUser", value: data }),
    msgDeleteRepository: (data: MsgDeleteRepository): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteRepository", value: data }),
    msgDeleteWhois: (data: MsgDeleteWhois): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteWhois", value: data }),
    msgDeleteIssue: (data: MsgDeleteIssue): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteIssue", value: data }),
    msgUpdateWhois: (data: MsgUpdateWhois): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateWhois", value: data }),
    msgUpdateRepository: (data: MsgUpdateRepository): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateRepository", value: data }),
    msgDeleteUser: (data: MsgDeleteUser): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgDeleteUser", value: data }),
    msgSetWhois: (data: MsgSetWhois): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgSetWhois", value: data }),
    msgCreateRepository: (data: MsgCreateRepository): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateRepository", value: data }),
    msgUpdateUser: (data: MsgUpdateUser): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgUpdateUser", value: data }),
    msgCreateIssue: (data: MsgCreateIssue): EncodeObject => ({ typeUrl: "/gitopia.gitopia.gitopia.MsgCreateIssue", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
