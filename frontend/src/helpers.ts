import 'isomorphic-fetch';

import { apiErrorCodes, ErrorCode } from './strings';
import type { APIType } from './types';

export const getNewestVersion = async (): Promise<string> =>
  (await fetch('https://imperialsplendour.com/version')).json();

type APII = { callEndpoint: () => Promise<void>; expectedErr?: ErrorCode };
type APIMap = { play: APII; website: APII; exit: APII };

export type EndpointKeys = keyof APIMap;

export const callAPI = (API: APIType) => async (key: EndpointKeys): Promise<void> => {
  const apiMap: APIMap = {
    play: { callEndpoint: API.Play, expectedErr: 'PlayError' },
    website: { callEndpoint: API.GoToWebsite, expectedErr: 'WebsiteError' },
    exit: { callEndpoint: API.Exit },
  };
  const { callEndpoint, expectedErr } = apiMap[key];

  try {
    await callEndpoint();
  } catch (e: unknown) {
    throw new Error(mapError(e as Error, expectedErr));
  }
};

const mapError = ({ message }: Error, ...expectedErr: (ErrorCode | undefined)[]) =>
  apiErrorCodes[expectedErr.find((code) => code === message) || 'UnexpectedError'];

const mapErrorz = ({ message }: Error) =>
  apiErrorCodes[message as ErrorCode] || apiErrorCodes.UnexpectedError;
