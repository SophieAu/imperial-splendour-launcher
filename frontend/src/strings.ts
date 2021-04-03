export const pageTitle = 'Imperial Splendour: Rise of the Republic';
export const etwTitle = 'Empire: Total War';

export const versionPrefix = 'v';

export const apiErrors = {
  startup: "There was an error on startup. Please let us know and we'll look into it",
  play:
    "Couldn't launch the game. Do you have Empire: Total War installed? If yes, please let us know about the error and we'll look into it.",
  website:
    "Couldn't open the website. Do you have a browser installed? If yes, please let us know about the error and we'll look into it.",
  uninstall: "Couldn't completely uninstall the game. Please delete any remaining files manually.",
  exit: "Couldn't exit the launcher. This should be impossible. Please tell us how you did this...",

  //switch errors
  switchToIS:
    "We couldn't cleanly switch to Imperial Splendour and therefore stayed on Empire: Total War. Some things might be a bit off when you play now though. Please let us know about the issue and we'll look into it.",
  switchToETW:
    "We couldn't cleanly switch to Empire: Total War. Some things might be a bit off when you play now. Please let us know about the issue and we'll look into it.",
};

export const newVersion = (vNr: string): string =>
  `We released a new Version (${vNr}) of Imperial Splendour. Go to our Website to download it.`;

export const modalButtonText = 'OK';

/*

// Switch Errors
var FileList = errors.New("FileListError")
var Activation = errors.New("ActivationError")
var Rollback = errors.New("RollbackError")
var Deactivation = errors.New("DeactivationError")
var StatusUpdate = errors.New("StatusUpdateError")

var Uninstall = errors.New("UninstallError")
var Website = errors.New("WebsiteError")
var Play = errors.New("PlayError")






*/

// const apiErrorCodes = {
//   FileListError: apiErrors.switch.cleanFail,
//   ActivationError: apiErrors.switch.cleanFail,
//   RollbackError: apiErrors.switch.dirtyFail,
//   DeactivationError: apiErrors.switch.dirtyFail,
//   StatusUpdateError: apiErrors.switch.cleanFail, // or dirty if deactivating

//   UninstallError: apiErrors.uninstall,
//   WebsiteError: apiErrors.website,
//   PlayError: apiErrors.play,
// };
