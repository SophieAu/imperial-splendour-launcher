import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

import {
  heightRatio,
  imgBG,
  largeFontSize,
  smallFontFamily,
  smallFontSize,
  widthRatio,
  windowWidth,
} from './styles';

export const heading = css`
  margin: 0 auto;
  padding: ${heightRatio * 58}px 0 ${heightRatio * 128}px;
  display: flex;
  align-items: center;
  justify-items: center;

  & > img {
    width: ${widthRatio * 1000}px;
    height: ${widthRatio * 374}px;
    object-fit: contain;
  }
`;

export const headingImg = css``;

export const root = css`
  background: center / contain ${imgBG};
  height: ${windowWidth / 1.6}px;
  overflow: hidden;

  display: flex;
  flex-direction: column;
`;

export const buttonContainer = css`
  flex: 1;
  display: flex;
  justify-content: space-evenly;
  align-items: flex-start;
`;

export const footer = css`
  text-align: right;
  margin-bottom: -0.25rem;
  padding-right: 0.25rem;
  font-family: ${smallFontFamily};

  & > .prefix {
    font-size: ${smallFontSize};
  }

  & > .version {
    font-size: ${largeFontSize};
  }
`;
