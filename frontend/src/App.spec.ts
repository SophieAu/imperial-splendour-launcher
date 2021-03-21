import '@testing-library/jest-dom/extend-expect';

import { fireEvent, render, wait, waitFor } from '@testing-library/svelte';

import App from './App.svelte';
import { apiErrors, etwTitle, pageTitle } from './strings';

const mockVersion = jest.fn();
const mockIsActive = jest.fn();
const mockPlay = jest.fn();
const mockSwitch = jest.fn();
const mockGoToWebsite = jest.fn();
const mockUninstall = jest.fn();
const mockExit = jest.fn();

const mockAPI = {
  Version: mockVersion,
  IsActive: mockIsActive,
  Play: mockPlay,
  Switch: mockSwitch,
  GoToWebsite: mockGoToWebsite,
  Uninstall: mockUninstall,
  Exit: mockExit,
};

afterEach(() => {
  mockVersion.mockReset();
  mockIsActive.mockReset();
  mockPlay.mockReset();
  mockSwitch.mockReset();
  mockGoToWebsite.mockReset();
  mockUninstall.mockReset();
  mockExit.mockReset();
});

test('Show error on startup when loading the version fails', async () => {
  mockVersion.mockRejectedValue(new Error('VersionError'));
  mockIsActive.mockResolvedValue(true);

  const { getByAltText, getByText, queryByText } = render(App, { API: mockAPI });

  // ImpSplen header is default on startup
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  await waitFor(() => expect(getByText('OK')).toBeInTheDocument());
  expect(getByText(apiErrors.startup)).toBeInTheDocument();
  // on startup error, do not change header image
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).not.toHaveBeenCalled();

  // dismiss error message
  fireEvent.click(getByText('OK'));

  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());
  expect(queryByText(apiErrors.startup)).not.toBeInTheDocument();
});

test('Show error on startup when loading the mod status fails', async () => {
  const vNr = 'versionNumber';
  mockVersion.mockResolvedValue(vNr);
  mockIsActive.mockRejectedValue(new Error('IsActiveError'));

  const { getByAltText, getByText, queryByText } = render(App, { API: mockAPI });
  // ImpSplen header is default on startup
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  await waitFor(() => expect(getByText('OK')).toBeInTheDocument());
  expect(getByText(apiErrors.startup)).toBeInTheDocument();

  // version loaded successfully
  expect(getByText(vNr)).toBeInTheDocument();
  // on startup error, do not change header image
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).toHaveBeenCalled();

  // dismiss error message
  fireEvent.click(getByText('OK'));

  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());
  expect(queryByText(apiErrors.startup)).not.toBeInTheDocument();
});

test('keeps ImpSplen header when mod is active', async () => {
  const vNr = 'versionNumber';
  mockVersion.mockResolvedValue(vNr);
  mockIsActive.mockResolvedValue(true);

  const { getByAltText, getByText } = render(App, { API: mockAPI });

  // ImpSplen header is default on startup
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  await waitFor(() => expect(getByText(vNr)).toBeInTheDocument());

  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).toHaveBeenCalled();

  // header stays on impsplen since mod is active
  expect(getByAltText(pageTitle)).toBeInTheDocument();
});

test('go through everything after successfull startup', async () => {
  const vNr = 'versionNumber';
  mockVersion.mockResolvedValue(vNr);
  mockIsActive.mockResolvedValueOnce(false);

  mockPlay.mockRejectedValueOnce(new Error('PlayError'));
  mockPlay.mockResolvedValueOnce(undefined);

  mockGoToWebsite.mockRejectedValueOnce(new Error('GoToWebsiteError'));
  mockGoToWebsite.mockResolvedValueOnce(undefined);

  mockUninstall.mockRejectedValueOnce(new Error('UninstallError'));
  mockUninstall.mockResolvedValueOnce(undefined);

  mockExit.mockRejectedValueOnce(new Error('ExitError'));
  mockExit.mockResolvedValueOnce(undefined);

  mockSwitch.mockRejectedValueOnce(new Error('SwitchError'));
  mockSwitch.mockResolvedValueOnce(undefined);
  mockIsActive.mockResolvedValueOnce(true);

  mockSwitch.mockRejectedValueOnce(new Error('SwitchError'));
  mockSwitch.mockResolvedValueOnce(undefined);
  mockIsActive.mockResolvedValueOnce(false);

  mockSwitch.mockResolvedValueOnce(undefined);
  mockIsActive.mockResolvedValueOnce(false);

  const { getByAltText, getByText, queryByText } = render(App, { API: mockAPI });
  // ImpSplen header is default on startup
  expect(getByAltText(pageTitle)).toBeInTheDocument();

  // startup
  await waitFor(() => expect(getByText(vNr)).toBeInTheDocument());
  expect(queryByText('OK')).not.toBeInTheDocument();

  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).toHaveBeenCalled();

  // Header is ETW since IS isn't active
  await waitFor(() => expect(getByAltText(etwTitle)).toBeInTheDocument());

  // --- PLAY ---

  // press Play button -> Error
  fireEvent.click(getByText('Play'));

  expect(mockPlay).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(getByText(apiErrors.play)).toBeInTheDocument());

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press Play button -> Success
  fireEvent.click(getByText('Play'));

  expect(mockPlay).toHaveBeenCalledTimes(2);
  expect(queryByText('OK')).not.toBeInTheDocument();

  // --- GO TO WEBSITE ---

  // press Website button -> Error
  fireEvent.click(getByText('Website'));

  expect(mockGoToWebsite).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(getByText(apiErrors.website)).toBeInTheDocument());

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press GoToWebsite button -> Success
  fireEvent.click(getByText('Website'));

  expect(mockGoToWebsite).toHaveBeenCalledTimes(2);
  expect(queryByText('OK')).not.toBeInTheDocument();

  // --- UNINSTALL ---

  // press Uninstall button -> Error
  fireEvent.click(getByText('Uninstall'));

  expect(mockUninstall).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(getByText(apiErrors.uninstall)).toBeInTheDocument());

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press GoToUninstall button -> Success
  fireEvent.click(getByText('Uninstall'));

  expect(mockUninstall).toHaveBeenCalledTimes(2);
  expect(queryByText('OK')).not.toBeInTheDocument();

  // --- EXIT ---

  // press Exit button -> Error
  fireEvent.click(getByText('Exit'));

  expect(mockExit).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(getByText(apiErrors.exit)).toBeInTheDocument());

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press GoToExit button -> Success
  fireEvent.click(getByText('Exit'));

  expect(mockExit).toHaveBeenCalledTimes(2);
  expect(queryByText('OK')).not.toBeInTheDocument();

  // --- SWITCH ---

  // press Switch button (to IS) -> Error + Header stays
  fireEvent.click(getByText('Switch'));

  expect(mockSwitch).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(getByText(apiErrors.switchToIS)).toBeInTheDocument());
  expect(getByAltText(etwTitle)).toBeInTheDocument();

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press Switch button (to IS) -> Success, switch to IS
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(getByAltText(pageTitle)).toBeInTheDocument());
  expect(queryByText('OK')).not.toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(2);
  expect(mockIsActive).toHaveBeenCalledTimes(2);

  // press Switch button (to ETW) -> Error + Header stays
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(getByText(apiErrors.switchToETW)).toBeInTheDocument());
  expect(getByAltText(pageTitle)).toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(3);

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press Switch button (to ETW) -> Success, switch to ETW
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(getByAltText(etwTitle)).toBeInTheDocument());
  expect(queryByText('OK')).not.toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(4);
  expect(mockIsActive).toHaveBeenCalledTimes(3);

  // press Switch button (to IS) -> No Error but isActive doesn't change
  fireEvent.click(getByText('Switch'));

  expect(getByAltText(etwTitle)).toBeInTheDocument();
  expect(queryByText('OK')).not.toBeInTheDocument();
  await waitFor(() => expect(mockSwitch).toHaveBeenCalledTimes(5));
  await waitFor(() => expect(mockIsActive).toHaveBeenCalledTimes(4));
});
