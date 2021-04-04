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
    const actualErr = (e as Error).message;

    throw new Error(
      (actualErr == expectedErr && apiErrorCodes[expectedErr]) || apiErrorCodes['UnexpectedError']
    );
  }
};
