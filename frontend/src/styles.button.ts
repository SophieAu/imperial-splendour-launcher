import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

import { buttonTexture, heightRatio, largeFontFamily, largeFontSize, widthRatio } from './styles';

const buttonHeight = heightRatio * 100;
const buttonWidth = widthRatio * 290;

const linkHoverColor = '#5c4e3c';

export const root = css`
  background: no-repeat center/100% ${buttonTexture};
  font: normal ${largeFontSize} / ${buttonHeight}px ${largeFontFamily};

  text-align: center;
  text-decoration: none;
  width: ${buttonWidth}px;
  padding: 0;
  margin: 0;
  border: none;

  &:hover {
    color: ${linkHoverColor};
    cursor: pointer;
  }
`;
