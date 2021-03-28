import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

import {
  buttonWidth,
  heightRatio,
  imgBG,
  largeFontSize,
  smallFontFamily,
  smallFontSize,
  widthRatio,
} from './styles';

export const root = css`
  background: left top / cover ${imgBG};
  width: 100vw;
  height: 100vh;
  overflow: hidden;

  display: flex;
  flex-direction: column;
`;

export const heading = css`
  margin: 0 auto;
  padding: ${heightRatio * 58}px 0 ${heightRatio * 128}px;
  display: flex;
  align-items: center;
  justify-items: center;
  width: ${widthRatio * 1000}px;
  height: ${widthRatio * 374}px;

  & > img {
    width: ${widthRatio * 1000}px;
    height: auto;
    object-fit: contain;
  }
`;

export const buttonContainerNonIE = css`
  width: 100vw;
  flex: 1 1 0%;
  display: flex;
  justify-content: space-evenly;
  align-items: flex-start;
`;

export const buttonContainer = css`
  box-sizing: border-box;
  width: 100vw;
  flex: 1 1 0%;
  display: flex;
  justify-content: space-between;
  padding: 0 calc(100vw / 6 - 5 * ${buttonWidth}px / 6);
  align-items: flex-start;
`;

export const footer = css`
  padding-right: 0.25rem;
  margin-bottom: -0.25rem;

  text-align: right;
  color: black;

  font-family: ${smallFontFamily};

  & > .prefix {
    font-size: ${smallFontSize};
  }

  & > .version {
    font-size: ${largeFontSize};
  }
`;
