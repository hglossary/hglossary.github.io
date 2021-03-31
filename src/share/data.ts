import data from '../_tmp/data.js';
import type {Category, Entry} from './types.js';
import {cmpStr} from './util.js';

const _data = buildData(data);
export const entries = _data.entries;
export const mapEntries = _data.mapEntries;
export const categories = _data.categories;

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
  entries.sort((a, b) => cmpStr(a.key, b.key));
  categories.sort((a, b) => cmpStr(a.key, b.key));
  categories.forEach((cat) => {
    cat.entries.sort((a, b) => cmpStr(a.key, b.key));
  });
  // selectedCategory.set(categories[0]);
  return {entries, mapEntries, categories};
}
