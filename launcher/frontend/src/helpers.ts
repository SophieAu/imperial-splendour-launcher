import 'isomorphic-fetch';

import { apiErrors, Errors } from './strings';
import type { APIType } from './types';

export const getNewestVersion = async (): Promise<string> =>
  (await fetch('https://imperialsplendour.com/version')).json();

type APII = { callEndpoint: () => Promise<void>; expectedErr?: Errors };
type APIMap = { play: APII; website: APII; exit: APII };

export type EndpointKeys = keyof APIMap;

export const callAPI = (API: APIType) => async (key: EndpointKeys): Promise<void> => {
  const apiMap: APIMap = {
    play: { callEndpoint: API.Play, expectedErr: Errors.PLAY },
    website: { callEndpoint: API.GoToWebsite, expectedErr: Errors.WEBSITE },
    exit: { callEndpoint: API.Exit },
  };
  const { callEndpoint, expectedErr } = apiMap[key];

  try {
    await callEndpoint();
  } catch (e: unknown) {
    throw new Error(mapError(e as Error, expectedErr));
  }
};

const mapError = ({ message }: Error, expectedErr?: Errors): string => {
  if (expectedErr !== message) return apiErrors.unexpected;

  if (message === Errors.FILELIST) return apiErrors.fileList;
  else if (message === Errors.ACTIVATION) return apiErrors.fileList;
  else if (message === Errors.ROLLBACK) return apiErrors.rollbackError;
  else if (message === Errors.DEACTIVATION) return apiErrors.fileList;
  //TODO:  NO
  else if (message === Errors.STATUS) return apiErrors.fileList;
  //TODO: NO
  else if (message === Errors.UNINSTALL) return apiErrors.uninstall;
  else if (message === Errors.WEBSITE) return apiErrors.website;
  else if (message === Errors.PLAY) return apiErrors.play;

  return apiErrors.unexpected;
};

export const mapSwitchError = (isDeactivating: boolean, { message }: Error): string => {
  if (message === Errors.FILELIST) return apiErrors.fileList;
  else if (isDeactivating && (message === Errors.DEACTIVATION || message === Errors.STATUS))
    return apiErrors.deactivationOnDeactivation;
  else if (!isDeactivating && message === Errors.ROLLBACK) return apiErrors.rollbackError;
  else if (!isDeactivating && (message === Errors.ACTIVATION || message === Errors.STATUS))
    return apiErrors.rollbackSuccessfull;

  return apiErrors.unexpectedOnSwitch;
};

export const mapUninstallError = ({ message }: Error): string => {
  if (message === Errors.UNINSTALL) return apiErrors.uninstall;
  else if (message === Errors.DEACTIVATION) return apiErrors.deactivationOnUninstall;

  return apiErrors.unexpected;
};
