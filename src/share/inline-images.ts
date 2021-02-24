import * as iconSearch from './icon-search.js';

let style = '';

export const getStyle = () => style;

const mapType = {
  png: 'image/png',
  svg: 'image/svg+xml',
};

type ImageData = { type: string, data: string }

function buildStyle(className: string, {type, data}: ImageData) {
  const dataType = mapType[type];
  if (!dataType) throw new Error(`unknown data type: ${type}`);

  style += `
.${className} {
  background-repeat: no-repeat;
  background-image: url("data:${dataType};base64,${btoa(data)}")
}`;
}

export function setupStyle() {
  const _$style = document.createElement('style');
  _$style.innerHTML = style;
  document.head.append(_$style);
}

export const classIconSearch = 'icon-search';
buildStyle(classIconSearch, iconSearch);
