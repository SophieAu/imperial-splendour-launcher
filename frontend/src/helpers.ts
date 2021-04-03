import 'isomorphic-fetch';

import { apiErrors } from './strings';

export const callAPI = async (
  callback: () => Promise<void>,
  errorCode: keyof typeof apiErrors
): Promise<string> => {
  try {
    await callback();
  } catch (e: unknown) {
    (e as Error).message;
    return apiErrors[errorCode];
  }
  return '';
};

export const getNewestVersion = async (): Promise<string> =>
  (await fetch('https://imperialsplendour.com/version')).json();
