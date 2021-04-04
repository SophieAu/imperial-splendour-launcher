import '@testing-library/jest-dom/extend-expect';

import { fireEvent, render, waitFor } from '@testing-library/svelte';

import App from './App.svelte';
import { apiErrors, etwTitle, newVersion, pageTitle, versionPrefix } from './strings';

const versionPingResolver = jest.fn();
const mockVersionPing = jest.fn(() => Promise.resolve({ json: () => versionPingResolver() }));
global.fetch = mockVersionPing;

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

xdescribe('On Startup', () => {
  test('show error and no content when loading the version fails', async () => {
    mockVersion.mockRejectedValue(new Error('VersionError'));
    versionPingResolver.mockRejectedValue(new Error('Oh Noes!'));
    mockIsActive.mockResolvedValue(true);

    const { getByText, queryByText } = render(App, { API: mockAPI });

    // No content while starting up
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    await waitFor(() => expect(queryByText('OK')).toBeInTheDocument());
    expect(queryByText(apiErrors.startup)).toBeInTheDocument();
    // on startup error, do not set content
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    expect(mockVersion).toHaveBeenCalled();
    expect(mockIsActive).not.toHaveBeenCalled();
    expect(mockVersionPing).not.toHaveBeenCalled();

    // dismiss error message
    fireEvent.click(getByText('OK'));

    await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());
    expect(queryByText(apiErrors.startup)).not.toBeInTheDocument();
  });

  test('show error and no content when loading the mod status fails', async () => {
    const vNr = 'versionNumber';
    mockVersion.mockResolvedValue(vNr);
    mockIsActive.mockRejectedValue(new Error('IsActiveError'));
    versionPingResolver.mockRejectedValue(new Error('Oh Noes!'));

    const { getByText, queryByText } = render(App, { API: mockAPI });

    // No content while starting up
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    await waitFor(() => expect(queryByText('OK')).toBeInTheDocument());
    expect(queryByText(apiErrors.startup)).toBeInTheDocument();

    // on startup error do not set version even if that did work
    expect(queryByText(vNr)).not.toBeInTheDocument();
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    expect(mockVersion).toHaveBeenCalled();
    expect(mockIsActive).toHaveBeenCalled();
    expect(mockVersionPing).not.toHaveBeenCalled();

    // dismiss error message
    fireEvent.click(getByText('OK'));

    await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());
    expect(queryByText(apiErrors.startup)).not.toBeInTheDocument();
  });

  test('ignore any errors when checking the newest version fails', async () => {
    const vNr = 'versionNumber';
    mockVersion.mockResolvedValue(vNr);
    mockIsActive.mockResolvedValue(true);
    versionPingResolver.mockRejectedValue(new Error('Oh Noes!'));

    const { queryByAltText, queryByText } = render(App, { API: mockAPI });

    // No content while starting up
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    await waitFor(() => expect(queryByText(vNr)).toBeInTheDocument());
    expect(queryByText('OK')).not.toBeInTheDocument(); // no error msg

    expect(mockVersion).toHaveBeenCalled();
    expect(mockIsActive).toHaveBeenCalled();
    expect(mockVersionPing).toHaveBeenCalledWith('https://imperialsplendour.com/version');

    // header turns into ImpSplen since it is active
    expect(queryByAltText(pageTitle)).toBeInTheDocument();
  });

  test('show upgrade message when there is a newer version', async () => {
    const vNr = '2.0';
    const vNrNew = '2.1';
    mockVersion.mockResolvedValue(vNr);
    mockIsActive.mockResolvedValue(true);
    versionPingResolver.mockResolvedValue(vNrNew);

    const { queryByAltText, getByText, queryByText } = render(App, { API: mockAPI });

    // No content while starting up
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    await waitFor(() => expect(queryByText('OK')).toBeInTheDocument());
    expect(queryByText(newVersion(vNrNew))).toBeInTheDocument();

    // show version and status in the background
    expect(queryByText(vNr)).toBeInTheDocument();
    expect(queryByAltText(pageTitle)).toBeInTheDocument();

    // dismiss error message
    fireEvent.click(getByText('OK'));

    await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());
    expect(queryByText(newVersion(vNrNew))).not.toBeInTheDocument();
  });

  test('do not show anything when current launcher version is the most recent one', async () => {
    const vNr = '2.0';
    const vNrNew = '2.0';
    mockVersion.mockResolvedValue(vNr);
    mockIsActive.mockResolvedValue(true);
    versionPingResolver.mockResolvedValue(vNrNew);

    const { queryByAltText, queryByText } = render(App, { API: mockAPI });

    // No content while starting up
    expect(queryByText(versionPrefix)).not.toBeInTheDocument();

    await waitFor(() => expect(queryByText(vNr)).toBeInTheDocument());
    expect(queryByAltText(pageTitle)).toBeInTheDocument();

    expect(mockVersionPing).toHaveBeenCalled();
    expect(queryByText(newVersion(vNrNew))).not.toBeInTheDocument();
  });

  test('show ImpSplen header when ImpSplen is active', async () => {
    mockVersion.mockResolvedValue('2.0');
    mockIsActive.mockResolvedValue(true);
    versionPingResolver.mockResolvedValue("doesn't matter");

    const { queryByAltText } = render(App, { API: mockAPI });

    await waitFor(() => expect(queryByAltText(pageTitle)).toBeInTheDocument());
  });

  test('show ETW header when ImpSplen is NOT active', async () => {
    mockVersion.mockResolvedValue('2.0');
    mockIsActive.mockResolvedValue(false);
    versionPingResolver.mockResolvedValue("doesn't matter");

    const { queryByAltText } = render(App, { API: mockAPI });

    await waitFor(() => expect(queryByAltText(etwTitle)).toBeInTheDocument());
  });
});

test('go through everything after successfull startup', async () => {
  mockVersion.mockResolvedValue('something');
  mockIsActive.mockResolvedValueOnce(false);

  mockPlay.mockRejectedValueOnce(new Error('PlayError'));
  mockPlay.mockResolvedValueOnce(undefined);

  mockGoToWebsite.mockRejectedValueOnce(new Error('WebsiteError'));
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

  const { getByText, queryByAltText, queryByText } = render(App, { API: mockAPI });
  // startup
  await waitFor(() => expect(queryByAltText(etwTitle)).toBeInTheDocument());
  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).toHaveBeenCalled();
  expect(mockVersionPing).toHaveBeenCalled();

  // --- PLAY ---

  // press Play button -> Error
  fireEvent.click(getByText('Play'));

  expect(mockPlay).toHaveBeenCalledTimes(1);
  await waitFor(() => expect(queryByText(apiErrors.play)).toBeInTheDocument());

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
  await waitFor(() => expect(queryByText(apiErrors.website)).toBeInTheDocument());

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
  await waitFor(() => expect(queryByText(apiErrors.uninstall)).toBeInTheDocument());

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
  await waitFor(() => expect(queryByText(apiErrors.unexpected)).toBeInTheDocument());

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
  await waitFor(() => expect(queryByText(apiErrors.switchToIS)).toBeInTheDocument());
  expect(queryByAltText(etwTitle)).toBeInTheDocument();

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press Switch button (to IS) -> Success, switch to IS
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(queryByAltText(pageTitle)).toBeInTheDocument());
  expect(queryByText('OK')).not.toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(2);
  expect(mockIsActive).toHaveBeenCalledTimes(2);

  // press Switch button (to ETW) -> Error + Header stays
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(queryByText(apiErrors.switchToETW)).toBeInTheDocument());
  expect(queryByAltText(pageTitle)).toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(3);

  // Dismiss Error Modal
  fireEvent.click(getByText('OK'));
  await waitFor(() => expect(queryByText('OK')).not.toBeInTheDocument());

  // press Switch button (to ETW) -> Success, switch to ETW
  fireEvent.click(getByText('Switch'));

  await waitFor(() => expect(queryByAltText(etwTitle)).toBeInTheDocument());
  expect(queryByText('OK')).not.toBeInTheDocument();
  expect(mockSwitch).toHaveBeenCalledTimes(4);
  expect(mockIsActive).toHaveBeenCalledTimes(3);

  // press Switch button (to IS) -> No Error but isActive doesn't change
  fireEvent.click(getByText('Switch'));

  expect(queryByAltText(etwTitle)).toBeInTheDocument();
  expect(queryByText('OK')).not.toBeInTheDocument();
  await waitFor(() => expect(mockSwitch).toHaveBeenCalledTimes(5));
  await waitFor(() => expect(mockIsActive).toHaveBeenCalledTimes(4));
});
