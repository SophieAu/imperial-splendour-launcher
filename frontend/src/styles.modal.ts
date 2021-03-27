import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

import { modalBG, smallFontFamily, smallFontSize } from './styles';

export const overlay = css`
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  display: flex;
  align-items: center;
  justify-content: center;

  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(2px);
`;

export const container = css`
  width: 33vw;
  background: ${modalBG};

  border: 0.25rem solid rgb(77, 50, 50);
  border-radius: 0.375rem;
  padding: 0.75rem;

  display: flex;
  flex-direction: column;
`;

export const button = css`
  align-self: flex-end;
`;

export const message = css`
  font: normal ${smallFontSize} ${smallFontFamily};

  margin: 0;
  padding: 0.5rem 0.5rem 0.75rem;
`;
