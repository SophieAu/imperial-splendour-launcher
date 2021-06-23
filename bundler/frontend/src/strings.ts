export const pageTitle = 'Imperial Splendour Bundler';
export const modalButtonText = 'OK';

export const apiErrors = {
  innoSetup: "couldn't find a InnoSetup installation. Please check that it's installed.",
  emptySourceError: 'Please select a directory with the mod files',
  invalidVersionError: "The version needs to be of the format 'MajorVersion.MinorVersion'.",
  noFileListError: "You didn't provide a file list. Please add one as the bundler needs it.",

  tempFolderCreationError: 'Could not create the temporary folders needed for the installation.',
  readSourceDirError:
    "Couldn't read the source dir. Make sure it actually exists and the bundler has read and write rights.",
  readFileListError:
    "Couldn't read the file list. Make sure it actually exists and the bundler has read and write rights.",

  fileMissingError: (files: string[]): string =>
    `The files ${join(
      files
    )} are in the file list but couldn't be found in the folder. Make sure to add them in or remove the from the list`,

  moveFileError: (files: string[]): string =>
    `Couldn't move the file(s) ${join(
      files
    )}. Make sure you haven't moved/renamed them in the last few seconds or changed the folder access. And make sure you have enough disc space.`,
  saveFileListError:
    "Couldn't copy the file list to the mod folder. Make sure you have folder access.",
  userScriptMissingError:
    "There seems to be no user script in the source folder (or the file is corrupted/the bundler doesn't have access rights).",
  moveUserScriptError:
    "Couldn't move the user script to the mod folder. Make sure you have folder access.",
  downloadError:
    "Couldn't download the necessary files from GitHub. Check your internet connection and the access rights for the Mod files folder.",

  infoFileError: "Couldn't create the info file.",

  versionUpdateError: "Couldn't update the version in the InnoSetup script.",
  tempFolderUpdateError: "Couldn't update the temp folder name in the InnoSetup script.",

  compileSetupError: "Couldn't compile the setup. Check the log for further info.",

  zipFilesError: "Couldn't create a .zip from the Mod Files. Check the log for more info",

  invalidDirError: 'The path you selected does not look like a valid directory.',
  invalidFileError: 'The file you selected does not look like a valid file.',

  unexpected: 'I have no idea what happened here... Please let me know how you did this',
};

const join = (strArr: string[]): string => strArr.join(', ');
