import {Writable, writable} from 'svelte/store';
import data from '../_tmp/data.js';

export const tabTree = 'tree';
export const tabList = 'alpha';

export const selectedCategory: Writable<Category | undefined> = writable(undefined);
export const selectedEntry: Writable<Entry | undefined> = writable(undefined);
export const selectedTab = writable('tree');

export const {entries, mapEntries, categories} = buildData(data);
export const entryUrl = (key: string) => '/w/' + key;
export const categoryUrl = (key: string) => '/c/' + key;

export const searchValue = writable('');
export const searchValueNorm = writable('');

searchValue.subscribe((v) => {
  const norm = v.normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/ /g, '-');
  searchValueNorm.set(norm);
});

export type Term = {
  display: string
  key: string
  primary?: boolean
}

export type ContentEntry = {
  // raw: string
  html: string
}

export type Link = {
  display: string
  url: string
}

export type Entry = {
  display?: string
  key?: string
  primaryEntry?: Entry

  terms: Term[]
  categories: Category[]
  content: ContentEntry
  footnotes?: string[]
  relatedArticles?: Link[]
  relatedTerms?: Term[]
}

export type Category = {
  display: string
  key: string
  entries?: Entry[]
}

export function selectEntryByKey(key: string): Entry | undefined {
  const entry = mapEntries[key];
  selectedEntry.set(entry);
  return entry;
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

function buildData(data: Entry[]): {
  entries: Entry[]
  mapEntries: Record<string, Entry>
  categories: Category[]
} {
  const entries: Entry[] = [];

  const mapCats: Record<string, Category> = {};
  for (let entry of data) {
    const currentEntries: Entry[] = [];
    for (let term of entry.terms) {
      currentEntries.push({...entry, ...term});
    }

    // the first term is primary, all other terms link back to the first one
    currentEntries.slice(1).forEach((entry) => {
      entry.primaryEntry = currentEntries[0];
    });
    entries.push(...currentEntries);

    for (let catItem of entry.categories || []) {
      let cat = mapCats[catItem.key];
      if (!cat) {
        cat = mapCats[catItem.key] =
          Object.assign({}, catItem, {entries: []});
      }
      cat.entries.push(...currentEntries);
    }
  }

  const mapEntries: Record<string, Entry> = {};
  for (let entry of entries) {
    mapEntries[entry.key] = entry;
  }

  const categories: Category[] = [...Object.values(mapCats)];
  entries.sort((a, b) => less(a.key, b.key));
  categories.sort((a, b) => less(a.key, b.key));
  categories.forEach((cat) => {
    cat.entries.sort((a, b) => less(a.key, b.key));
  });
  // selectedCategory.set(categories[0]);
  return {entries, mapEntries, categories};
}

function less(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}
