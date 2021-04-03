import 'isomorphic-fetch';

import { apiErrors } from './strings';
import type { APIType } from './types';

export let API: APIType;

export const callAPI = async (
  callback: () => Promise<void>,
  errorCode: keyof typeof apiErrors
): Promise<string> => {
  try {
    await callback();
  } catch (e) {
    e.message();
    return apiErrors[errorCode];
  }
  return '';
};

export const getNewestVersion = async (): Promise<string> =>
  (await fetch('https://imperialsplendour.com/version')).json();
