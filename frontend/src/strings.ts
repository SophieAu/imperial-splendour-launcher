export const pageTitle = 'Imperial Splendour: Rise of the Republic';
export const etwTitle = 'Empire: Total War';

export const versionPrefix = 'v';

export enum Errors {
  FILELIST = 'FileListError',
  ACTIVATION = 'ActivationError',
  ROLLBACK = 'RollbackError',
  DEACTIVATION = 'DeactivationError',
  STATUS = 'StatusUpdateError',

  UNINSTALL = 'UninstallError',
  WEBSITE = 'WebsiteError',
  PLAY = 'PlayError',
}

const pleaseContact = " Please let us know about the issue and we'll look into it.";

export const apiErrors = {
  startup: 'There was an error on startup.' + pleaseContact,
  play:
    "Couldn't launch the game. Do you have Empire: Total War installed? If yes, please let us know about the error and we'll look into it.",
  website:
    "Couldn't open the website. Do you have a browser installed? If yes, please let us know about the error and we'll look into it.",
  deactivationOnUninstall:
    "We couldn't deactivate Imperial Splendour to prepare for uninstalling." + pleaseContact,
  uninstall: "We couldn't delete all the files. Please go in and delete them manually.", // TODO: Adjust for additional uninstall capabilities
  deactivationOnDeactivation:
    'We ran into issues switching to the base game. Your files and status might be out of sync.' +
    pleaseContact,
  fileList:
    'We ran into an issue when trying to switch and aborted before doing anything.' + pleaseContact,
  rollbackSuccessfull:
    "There was an error switching to Imperial Splendour. We rolled back everything and you're still on the base game." +
    pleaseContact,
  rollbackError:
    'We ran into issues switching to Imperial Splendour. Your files and status might be out of sync.' +
    pleaseContact,
  unexpectedOnSwitch:
    "We ran into an error we didn't expect. Your files and status might be out of sync." +
    pleaseContact,

  unexpected: "We ran into an error that shouldn't happen..." + pleaseContact,
};

export const newVersion = (vNr: string): string =>
  `We released a new Version (${vNr}) of Imperial Splendour. Go to our Website to download it.`;

export const modalButtonText = 'OK';
