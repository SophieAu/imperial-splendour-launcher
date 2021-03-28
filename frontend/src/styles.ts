import background from './assets/background.png';
import textureBg from './assets/texture_bg.jpg';
import textureBtn from './assets/texture_btn.png';

export const imgBG = `url(${background})`;
export const buttonTexture = `url(${textureBtn})`;
export const modalBG = `url(${textureBg})`;

// window / bg image
// 1280 / 1920 = 2/3
// 800 / 1200 = 2/3
export const widthRatio = 2 / 3;
export const heightRatio = 2 / 3;

export const largeFontSize = '2rem';
export const largeFontFamily = 'IM FELL English SC';

export const smallFontSize = '1.25rem';
export const smallFontFamily = 'IM FELL English';

export const buttonHeight = heightRatio * 100;
export const buttonWidth = widthRatio * 290;

export const linkHoverColor = '#5c4e3c';
