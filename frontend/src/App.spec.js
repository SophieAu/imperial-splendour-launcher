import '@testing-library/jest-dom/extend-expect';

import { render, waitFor } from '@testing-library/svelte';
import App from './App';

describe('derp', () => {
  const mockVersion = jest.fn();

  const mockIsActive = jest.fn();
  const mockPlay = jest.fn();
  const mockSwitch = jest.fn();
  const mockGoToWebsite = jest.fn();
  const mockUninstall = jest.fn();
  const mockExit = jest.fn();

  const mockAPI = {
    Version: async () => mockVersion(),
    IsActive: async () => mockIsActive(),
    Play: async () => mockPlay(),
    Switch: async () => mockSwitch(),
    GoToWebsite: async () => mockGoToWebsite(),
    Uninstall: async () => mockUninstall(),
    Exit: async () => mockExit(),
  };

  test('shows proper heading when rendered', async () => {
    mockVersion.mockReturnValue('flerpaferp');

    const { getByAltText, getByText } = render(App, { API: mockAPI });

    await waitFor(() => expect(getByText('flerpaferp')).toBeInTheDocument());

    expect(mockVersion).toHaveBeenCalled();

    expect(getByAltText('Imperial Splendour: Rise of the Republic')).toBeInTheDocument();
  });
});
