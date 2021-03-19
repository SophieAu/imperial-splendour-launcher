import { css } from '@emotion/css/dist/emotion-css.umd.min.js';

import background from './assets/background.png';
import textureBtn from './assets/texture_btn.png';
import textureBg from './assets/texture_bg.jpg';

export const root = css`
  --img-bg: url(${background});
  --button-texture: url(${textureBtn});
  --modal-bg: url(${textureBg});
`;
