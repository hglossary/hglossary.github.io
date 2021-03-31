export const querySel = (sel) => document.querySelector(sel);
export const queryAll = (sel) => document.querySelectorAll(sel);

export function cmpStr(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}
