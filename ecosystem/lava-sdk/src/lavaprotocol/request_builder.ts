import {
  EARLIEST_BLOCK,
  FINALIZED_BLOCK,
  LATEST_BLOCK,
  NOT_APPLICABLE,
  SAFE_BLOCK,
} from "../common/common";
import {
  RelayRequest,
  RelaySession,
  RelayPrivateData,
  QualityOfServiceReport,
  RelayReply,
} from "../grpc_web_services/lavanet/lava/pairing/relay_pb";
import { SingleConsumerSession } from "../lavasession/consumerTypes";
import { sha256 } from "@cosmjs/crypto";
import { Decimal } from "@cosmjs/math";
import { Logger } from "../logger/logger";

export interface SendRelayData {
  connectionType: string;
  data: string;
  url: string;
  apiInterface: string;
  chainId: string;
  requestedBlock: number;
}

export function newRelayData(relayData: SendRelayData): RelayPrivateData {
  const { data, url, connectionType } = relayData;
  // create request private data
  const enc = new TextEncoder();
  const requestPrivateData = new RelayPrivateData();
  requestPrivateData.setConnectionType(connectionType);
  requestPrivateData.setApiUrl(url);
  requestPrivateData.setData(enc.encode(data));
  requestPrivateData.setRequestBlock(NOT_APPLICABLE); // TODO: when block parsing is implemented, replace this with the request parsed block.
  requestPrivateData.setApiInterface(relayData.apiInterface);
  requestPrivateData.setSalt(getNewSalt());
  return requestPrivateData;
}

function getNewSalt(): Uint8Array {
  const salt = generateRandomUint();
  const nonceBytes = new Uint8Array(8);
  const dataView = new DataView(nonceBytes.buffer);

  // use LittleEndian
  dataView.setBigUint64(0, BigInt(salt), true);

  return nonceBytes;
}

function generateRandomUint(): number {
  const min = 1;
  const max = Number.MAX_SAFE_INTEGER;
  return Math.floor(Math.random() * (max - min) + min);
}

export function constructRelayRequest(
  lavaChainID: string,
  chainID: string,
  relayData: RelayPrivateData,
  providerAddress: string,
  singleConsumerSession: SingleConsumerSession,
  epoch: number,
  reportedProviders: string
): RelayRequest {
  const relayRequest = new RelayRequest();
  relayRequest.setRelayData(relayData);
  const relaySession = constructRelaySession(
    lavaChainID,
    chainID,
    relayData,
    providerAddress,
    singleConsumerSession,
    epoch,
    reportedProviders
  );
  relayRequest.setRelaySession(relaySession);
  return relayRequest;
}

function constructRelaySession(
  lavaChainID: string,
  chainID: string,
  relayData: RelayPrivateData,
  providerAddress: string,
  singleConsumerSession: SingleConsumerSession,
  epoch: number,
  reportedProviders: string
): RelaySession {
  const lastQos = singleConsumerSession.qoSInfo.lastQoSReport;
  let newQualityOfServiceReport: QualityOfServiceReport | undefined = undefined;

  if (lastQos != undefined) {
    newQualityOfServiceReport = new QualityOfServiceReport();
    // TODO: needs to serialize the QoS report value like a serialized Dec
    newQualityOfServiceReport.setLatency(
      Decimal.fromUserInput(lastQos.getLatency(), 0).toString()
    );
    newQualityOfServiceReport.setAvailability(
      Decimal.fromUserInput(lastQos.getAvailability(), 0).toString()
    );
    newQualityOfServiceReport.setSync(
      Decimal.fromUserInput(lastQos.getSync(), 0).toString()
    );
  }
  const lastQosExcellence =
    singleConsumerSession.qoSInfo.lastExcellenceQoSReport;

  let newQualityOfServiceReportExcellence: QualityOfServiceReport | undefined =
    undefined;

  if (lastQosExcellence != undefined) {
    newQualityOfServiceReportExcellence = new QualityOfServiceReport();

    // TODO: needs to serialize the QoS report value like a serialized Dec
    newQualityOfServiceReportExcellence.setLatency(
      Decimal.fromUserInput(lastQosExcellence.getLatency(), 0).toString()
    );
    newQualityOfServiceReportExcellence.setAvailability(
      Decimal.fromUserInput(lastQosExcellence.getAvailability(), 0).toString()
    );
    newQualityOfServiceReportExcellence.setSync(
      Decimal.fromUserInput(lastQosExcellence.getSync(), 0).toString()
    );
  }

  const relaySession = new RelaySession();
  relaySession.setSpecId(chainID);
  relaySession.setLavaChainId(lavaChainID);
  relaySession.setSessionId(singleConsumerSession.sessionId);
  relaySession.setProvider(providerAddress);
  relaySession.setSig(new Uint8Array());
  relaySession.setContentHash(calculateContentHash(relayData));
  relaySession.setEpoch(epoch);
  relaySession.setRelayNum(singleConsumerSession.relayNum);
  relaySession.setQosReport(newQualityOfServiceReport);
  relaySession.setCuSum(
    singleConsumerSession.cuSum + singleConsumerSession.latestRelayCu
  );
  relaySession.setQosExcellenceReport(newQualityOfServiceReportExcellence);

  relaySession.setUnresponsiveProviders(reportedProviders);
  return relaySession;
}

function calculateContentHash(relayRequestData: RelayPrivateData): Uint8Array {
  const requestBlock = relayRequestData.getRequestBlock();
  const requestBlockBytes = convertRequestedBlockToUint8Array(requestBlock);

  const apiInterfaceBytes = encodeUtf8(relayRequestData.getApiInterface());
  const connectionTypeBytes = encodeUtf8(relayRequestData.getConnectionType());
  const apiUrlBytes = encodeUtf8(relayRequestData.getApiUrl());
  const dataBytes = relayRequestData.getData();
  const dataUint8Array =
    dataBytes instanceof Uint8Array ? dataBytes : encodeUtf8(dataBytes);
  const saltBytes = relayRequestData.getSalt();
  const saltUint8Array =
    saltBytes instanceof Uint8Array ? saltBytes : encodeUtf8(saltBytes);

  const msgData = concatUint8Arrays([
    apiInterfaceBytes,
    connectionTypeBytes,
    apiUrlBytes,
    dataUint8Array,
    requestBlockBytes,
    saltUint8Array,
  ]);

  const hash = sha256(msgData);
  return hash;
}

function encodeUtf8(str: string): Uint8Array {
  return new TextEncoder().encode(str);
}

function concatUint8Arrays(arrays: Uint8Array[]): Uint8Array {
  const totalLength = arrays.reduce((acc, arr) => acc + arr.length, 0);
  const result = new Uint8Array(totalLength);
  let offset = 0;
  arrays.forEach((arr) => {
    result.set(arr, offset);
    offset += arr.length;
  });
  return result;
}

function convertRequestedBlockToUint8Array(requestBlock: number): Uint8Array {
  const requestBlockBytes = new Uint8Array(8);
  let number = BigInt(requestBlock);
  if (requestBlock < 0) {
    // Convert the number to its 64-bit unsigned representation
    const maxUint64 = BigInt(2) ** BigInt(64);
    number = maxUint64 + BigInt(requestBlock);
  }

  // Copy the bytes from the unsigned representation to the byte array
  for (let i = 0; i < 8; i++) {
    requestBlockBytes[i] = Number((number >> BigInt(8 * i)) & BigInt(0xff));
  }

  return requestBlockBytes;
}

export function UpdateRequestedBlock(
  request: RelayPrivateData,
  response: RelayReply
) {
  request.setRequestBlock(
    ReplaceRequestedBlock(request.getRequestBlock(), response.getLatestBlock())
  );
  return;
}

export function ReplaceRequestedBlock(
  requestedBlock: number,
  latestBlock: number
): number {
  switch (requestedBlock) {
    case LATEST_BLOCK:
    case SAFE_BLOCK:
    case FINALIZED_BLOCK:
      return latestBlock;
    case EARLIEST_BLOCK:
      // TODO: add support for earliest block reliability
      return NOT_APPLICABLE;
    default:
      return requestedBlock;
  }
}

export function IsFinalizedBlock(
  requestedBlock: number,
  latestBlock: number,
  finalizationCriteria: number
): boolean {
  switch (requestedBlock) {
    case NOT_APPLICABLE:
      return false;
    default:
      if (requestedBlock < 0) {
        return false;
      }
      if (requestedBlock <= latestBlock - finalizationCriteria) {
        return true;
      }
  }
  return false;
}

export function verifyRelayReply(
  reply: RelayReply,
  relayRequest: RelayRequest,
  providerAddress: string
): Error | undefined {
  // TODO: implement signature verificaion
  return;
}

interface FinalizationData {
  finalizedBlocks: Map<number, string>;
  finalizationConflict: undefined;
}

export function verifyFinalizationData(
  reply: RelayReply,
  relayRequest: RelayRequest,
  providerAddr: string,
  consumerAddress: string,
  latestSessionBlock: number,
  blockDistanceForfinalization: number
): FinalizationData | Error {
  // TODO: implement signature extraction on finalization data

  // relayFinalization := pairingtypes.NewRelayFinalization(pairingtypes.NewRelayExchange(*relayRequest, *reply), consumerAcc)
  // serverKey, err := sigs.RecoverPubKey(relayFinalization)
  // if err != nil {
  // 	return nil, nil, err
  // }

  // serverAddr, err := sdk.AccAddressFromHexUnsafe(serverKey.Address().String())
  // if err != nil {
  // 	return nil, nil, err
  // }

  // if serverAddr.String() != providerAddr {
  // 	return nil, nil, utils.LavaFormatError("reply server address mismatch in finalization data ", ProviderFinzalizationDataError, utils.Attribute{Key: "parsed Address", Value: serverAddr.String()}, utils.Attribute{Key: "expected address", Value: providerAddr})
  // }

  const finalizedBlocks = new Map<number, string>();
  // const finalizedHashes64 = reply.getFinalizedBlocksHashes_asB64();
  const dec = new TextDecoder();
  const decodedResponse = dec.decode(reply.getFinalizedBlocksHashes_asU8());
  const finalizaedBlocksObj = JSON.parse(decodedResponse);
  for (const key in finalizaedBlocksObj) {
    const numericKey = parseInt(key, 10);
    finalizedBlocks.set(numericKey, finalizaedBlocksObj[key]);
  }
  const finalizationConflict = verifyFinalizationDataIntegrity(
    reply,
    latestSessionBlock,
    finalizedBlocks,
    providerAddr,
    blockDistanceForfinalization
  );

  if (finalizationConflict instanceof Error) {
    return finalizationConflict;
  }
  const finalizationData: FinalizationData = {
    finalizedBlocks: finalizedBlocks,
    finalizationConflict: finalizationConflict,
  };
  return finalizationData;
}

function verifyFinalizationDataIntegrity(
  reply: RelayReply,
  existingSessionLatestBlock: number,
  finalizedBlocks: Map<number, string>,
  providerPublicAddress: string,
  blockDistanceForFinalizedData: number
): Error | undefined {
  return undefined;
}
