import {Writable, writable} from 'svelte/store';
import data from '../_tmp/data.js';

export const tabTree = 'tree';
export const tabList = 'alpha';

export const selectedCategory: Writable<Category | null> = writable(null);
export const selectedEntry: Writable<Entry | null> = writable(null);
export const selectedTab = writable('tree');

export const {entries, categories} = buildData(data);
export const entryUrl = (key: string) => '/w/' + key;
export const categoryUrl = (key: string) => '/c/' + key;

export const searchValue = writable('');
export const searchValueNorm = writable('');

searchValue.subscribe((v) => {
  const norm = v.normalize("NFD")
    .replace(/[\u0300-\u036f]/g, "")
    .replace(/ /g, '-');
  searchValueNorm.set(norm);

  console.log('norm', norm);
});

export type Term = {
  display: string
  key: string
  primary?: boolean
}

export type ContentEntry = {
  raw: string
  html: string
}

export type Entry = {
  display?: string
  key?: string
  primaryEntry?: Entry
  terms: Term[]
  categories: Category[]
  content: ContentEntry
  comment: string
}

export type Category = {
  display: string
  key: string
  entries?: Entry[]
}

function buildData(data: Entry[]): {
  entries: Entry[]
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

  const categories: Category[] = [...Object.values(mapCats)];
  entries.sort((a, b) => less(a.key, b.key));
  categories.sort((a, b) => less(a.key, b.key));
  categories.forEach((cat) => {
    cat.entries.sort((a, b) => less(a.key, b.key));
  });

  selectedCategory.set(categories[0]);
  return {entries, categories};
}

function less(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

