export const pageTitle = 'Imperial Splendour: Rise of the Republic';
export const etwTitle = 'Empire: Total War';

export const versionPrefix = 'v';

const pleaseContact = "Please let us know about the issue and we'll look into it.";

export const apiErrors = {
  startup: 'There was an error on startup.' + pleaseContact,
  play:
    "Couldn't launch the game. Do you have Empire: Total War installed? If yes, please let us know about the error and we'll look into it.",
  website:
    "Couldn't open the website. Do you have a browser installed? If yes, please let us know about the error and we'll look into it.",
  uninstall: "Couldn't completely uninstall the game. Please delete any remaining files manually.",

  //switch errors
  switchToIS:
    "We couldn't cleanly switch to Imperial Splendour and therefore stayed on Empire: Total War. Some things might be a bit off when you play now though. Please let us know about the issue and we'll look into it.",
  switchToETW:
    "We couldn't cleanly switch to Empire: Total War. Some things might be a bit off when you play now. Please let us know about the issue and we'll look into it.",

  unexpected:
    "We ran into an error that shouldn't happen... Please let us know how you got here...",
};

export const switchErrors = {
  cleanFail: {
    toIS: '',
    toETW: '',
    uninstall: '',
  },

  dirtyFail: {
    toIS: '',
    toETW: '',
    uninstall: '',
  },
};

export const newVersion = (vNr: string): string =>
  `We released a new Version (${vNr}) of Imperial Splendour. Go to our Website to download it.`;

export const modalButtonText = 'OK';

/*

// Switch Errors
var FileList = errors.New("FileListError")
var Activation = errors.New("ActivationError")
var Rollback = errors.New("Rollb packError")
var Deactivation = errors.New("DeactivationError")
var StatusUpdate = errors.New("StatusUpdateError")

var Uninstall = errors.New("UninstallError")
var Website = errors.New("WebsiteError")
var Play = errors.New("PlayError")






*/

export const switchErrorCodes = {
  FileListError: (mode: Mode) => switchErrors.cleanFail[mode],
  ActivationError: (mode: Mode) => switchErrors.cleanFail[mode],
  RollbackError: (mode: Mode) => switchErrors.dirtyFail[mode],
  DeactivationError: (mode: Mode) => switchErrors.dirtyFail[mode],
  StatusUpdateError: (mode: Mode) => switchErrors.cleanFail[mode], // or dirty if deactivating
};

type Mode = keyof typeof switchErrors['cleanFail'];

export type ErrorCode = keyof typeof apiErrorCodes;

export const apiErrorCodes = {
  UninstallError: apiErrors.uninstall,
  WebsiteError: apiErrors.website,
  PlayError: apiErrors.play,

  UnexpectedError: apiErrors.unexpected,
};

// type APII = { callEndpoint: () => Promise<void>; expectedErr?: keyof typeof apiErrorCodes };
// const APIMap: { play: APII; website: APII; exit: APII } = {
//   play: { callEndpoint: API.Play, expectedErr: 'PlayError' },
//   website: { callEndpoint: API.GoToWebsite, expectedErr: 'WebsiteError' },
//   exit: { callEndpoint: API.Exit },
// };

// export const callSimpleAPI = async (endpointKey: keyof typeof APIMap): Promise<string> => {
//   const { callEndpoint, expectedErr } = APIMap[endpointKey];

//   try {
//     await callEndpoint();
//   } catch (e: unknown) {
//     return (
//       ((e as Error).message == expectedErr && apiErrorCodes[expectedErr]) || apiErrors.unexpected
//     );
//   }
//   return '';
// };
