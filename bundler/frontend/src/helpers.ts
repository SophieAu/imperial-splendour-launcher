import { apiErrors } from './strings';

export const mapError = (error: Error): string => {
  const [errCode, ...args] = ((error as unknown) as string).split(' ');

  if (errCode === 'InnoSetupError') return apiErrors.innoSetup;
  if (errCode === 'EmptySourceError') return apiErrors.emptySourceError;
  if (errCode === 'InvalidVersionError') return apiErrors.invalidVersionError;

  if (errCode === 'TempFolderCreationError') return apiErrors.tempFolderCreationError;
  if (errCode === 'ReadSourceDirError') return apiErrors.readSourceDirError;
  if (errCode === 'ReadFileListError') return apiErrors.readFileListError;
  if (errCode === 'FileMissingError') return apiErrors.fileMissingError(args);
  if (errCode === 'MoveFileError') return apiErrors.moveFileError(args);
  if (errCode === 'SaveFileListError') return apiErrors.saveFileListError;

  if (errCode === 'UserScriptMissingError') return apiErrors.userScriptMissingError;
  if (errCode === 'MoveUserScriptError') return apiErrors.moveUserScriptError;

  if (errCode === 'DownloadError') return apiErrors.downloadError;
  if (errCode === 'InfoFileError') return apiErrors.infoFileError;
  if (errCode === 'VersionUpdateError') return apiErrors.versionUpdateError;
  if (errCode === 'TempFolderUpdateError') return apiErrors.tempFolderUpdateError;

  if (errCode === 'NoFileListError') return apiErrors.noFileListError;

  if (errCode === 'CompileSetupError') return apiErrors.compileSetupError;
  if (errCode === 'ZipFilesError') return apiErrors.zipFilesError;

  if (errCode === 'InvalidDirError') return apiErrors.invalidDirError;
  if (errCode === 'InvalidFileError') return apiErrors.invalidFileError;

  return apiErrors.unexpected;
};
