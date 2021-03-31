import type {Writable} from 'svelte/store';
import {writable} from 'svelte/store';
import {mapEntries} from './data.js';
import {parseRoute} from './pages.js';
import type {Category, Entry, Route, Term} from './types.js';

export const kTabTree = 'tree';
export const kTabList = 'alpha';
export const currentUrl: Writable<string> = writable(location.pathname);
export const currentRoute: Writable<Route> = writable(parseRoute(location.pathname));

export const selectedCategory: Writable<Category | undefined> = writable(undefined);
export const selectedEntry: Writable<Entry | undefined> = writable(undefined);
export const selectedTab = writable(kTabTree);

export const entryUrl = (key: string) => '/w/' + key;    // word
export const categoryUrl = (key: string) => '/c/' + key; // category
export const metaUrl = (key: string) => '/a/' + key;     // about

export const searchValue = writable('');
export const searchValueNorm = writable('');

window.addEventListener('popstate', (e) => {
  const route = e.state as Route;
  if (!route) return;

  currentUrl.set(route?.url || '/');
  currentRoute.set(route);

  console.log('<- url', route?.url, route);
});

currentRoute.subscribe((route) => {
  const {page, key} = route;
  if (page === 'entry') {
    const entry = mapEntries[key];
    selectedEntry.set(entry);
  } else {
    selectedEntry.set(null);
  }
});

searchValue.subscribe((v) => {
  const norm = v.normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/ /g, '-');
  searchValueNorm.set(norm);
});

export function selectEntry(key: string) {
  gotoUrl(entryUrl(key));
}

export function randomRelatedTerms(entry): Term[] | undefined {
  if (!entry || !entry.relatedTerms || !entry.relatedTerms.length) {
    return;
  }

  let relatedTerms = entry.relatedTerms;
  const n = entry.relatedTerms.length;
  if (n > 3) {
    const picked = {};
    for (let i = 0; i < 3; i++) {
      let x;
      do {
        x = Math.floor(Math.random() * n);
      } while (picked[x]);
      picked[x] = 1;
    }
    relatedTerms = Object.keys(picked).sort().map((i) => entry.relatedTerms[i]);
  }
  return relatedTerms;
}

export function gotoUrl(url: string) {
  const u = new URL(url, location.href);

  // external url
  if (u.origin !== location.origin) {
    window.open(url, '_blank');
    return true;
  }

  const pathname = u.pathname;
  const route = parseRoute(pathname);
  currentUrl.set(pathname);
  currentRoute.set(route);

  if (pathname !== location.pathname) {
    history.pushState(route, 'Ngành du lịch', route.url);
  }
  console.log('-> url', route.url, `(${url})`);
}
