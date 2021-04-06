export const pageTitle = 'Imperial Splendour: Rise of the Republic';
export const etwTitle = 'Empire: Total War';

export const versionPrefix = 'v';

enum Error {
  'FileListError',
  'ActivationError',
  'RollbackError',
  'DeactivationError',
  'StatusUpdateError',

  'UninstallError',
  'WebsiteError',
  'PlayError',
}

const pleaseContact = " Please let us know about the issue and we'll look into it.";

export const apiErrors = {
  startup: 'There was an error on startup.' + pleaseContact,
  play:
    "Couldn't launch the game. Do you have Empire: Total War installed? If yes, please let us know about the error and we'll look into it.",
  website:
    "Couldn't open the website. Do you have a browser installed? If yes, please let us know about the error and we'll look into it.",
  deactivationErrorOnUninstall:
    "We couldn't deactivate Imperial Splendour to prepare for uninstalling." + pleaseContact,
  uninstallError: "We couldn't delete all the files. Please go in and delete them manually.", // TODO: Adjust for additional uninstall capabilities
  deactivationErrorWhenDeactivating:
    'We ran into issues switching to the base game. Your files and status might be out of sync.' +
    pleaseContact,
  fileListErrorInGeneral:
    'We ran into an issue when trying to switch and aborted before doing anything.' + pleaseContact,
  rollbackSuccessfullError:
    "There was an error switching to Imperial Splendour. We rolled back everything and you're still on the base game." +
    pleaseContact,
  rollbackErrorError:
    'We ran into issues switching to Imperial Splendour. Your files and status might be out of sync.' +
    pleaseContact,
  unexpectedOnSwitch:
    "We ran into an error we didn't expect. Your files and status might be out of sync." +
    pleaseContact,

  unexpected:
    "We ran into an error that shouldn't happen... Please let us know how you got here...",
};

export const newVersion = (vNr: string): string =>
  `We released a new Version (${vNr}) of Imperial Splendour. Go to our Website to download it.`;

export const modalButtonText = 'OK';

export type ErrorCode = keyof typeof apiErrorCodes;

export const apiErrorCodes = {
  UninstallError: apiErrors.uninstallError,
  WebsiteError: apiErrors.website,
  PlayError: apiErrors.play,

  UnexpectedError: apiErrors.unexpected,
};
