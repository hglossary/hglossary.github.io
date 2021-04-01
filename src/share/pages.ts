import Component from 'svelte/types/compiler/compile/Component.js';
// @ts-ignore
import About from '../layout/About.svelte';
// @ts-ignore
import Contact from '../layout/Contact.svelte';
// @ts-ignore
import Entry from '../layout/Entry.svelte';
import {mapEntries} from './data.js';
import type {Route} from './types.js';

export const kDefPageMeta = 'about';
export const mapPageMeta = {
  about: About,
  contact: Contact,
};

export function routeComponent(r: Route): Component {
  if (r.page === 'entry') return Entry;
  return mapPageMeta[r.key] || mapPageMeta[kDefPageMeta];
}

export function validateMetaKey(key: string): string {
  return mapPageMeta[key] ? key : kDefPageMeta;
}

const reUrl = /^\/([w])\/([0-9A-Za-z-]+)$/;
const routeHome: Route = {url: '/', page: 'meta', key: 'home'};

export function parseRoute(path: string, hash?: string): Route {
  const parts = reUrl.exec(path);
  if (!parts) {
    const key = path.slice(1);
    if (mapPageMeta[key]) return {url: path, page: 'meta', key};
    return routeHome;
  }

  const page = parts[1], key = parts[2];
  if (page === 'w' && mapEntries[key]) {
    const title = `Ngành du lịch - ${mapEntries[key].display}`;
    return {url: path, page: 'entry', key, title};
  }
  return routeHome;
}
