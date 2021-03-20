import '@testing-library/jest-dom/extend-expect';

import { render, waitFor } from '@testing-library/svelte';
import App from './App';

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

test('shows proper heading when rendered', async () => {
  mockVersion.mockResolvedValue('flerpaferp');
  mockIsActive.mockRejectedValue(new Error('Async error'));

  const { getByAltText, getByText } = render(App, { API: mockAPI });

  await waitFor(() => expect(getByText('flerpaferp')).toBeInTheDocument());
  await waitFor(() => expect(getByText('OK')).toBeInTheDocument());

  expect(mockVersion).toHaveBeenCalled();
  expect(mockIsActive).toHaveBeenCalled();

  expect(getByAltText('Imperial Splendour: Rise of the Republic')).toBeInTheDocument();
  expect(getByText('OK')).toBeInTheDocument();
});

xtest('show error message on startup when getting the version fails', async () => {
  const { getByAltText, getByText } = render(App, { API: mockAPI });
});
xtest('show error message on startup when getting the status fails', async () => {
  const { getByAltText, getByText } = render(App, { API: mockAPI });
});
xtest('show imp splen logo when IS is active on startup', async () => {
  const { getByAltText, getByText } = render(App, { API: mockAPI });
});
xtest('show imp splen logo when ETW is active on startup', async () => {
  const { getByAltText, getByText } = render(App, { API: mockAPI });
});

// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
// test('', async () => {});
