<div class="wrap-left">
  <SearchBox screen="desktop"/>
  <div class="tabs">
    <div class="tab"
         class:active={$selectedTab === kTabTree}
         on:click={()=>selectTab(kTabTree)}>
      Danh mục
    </div>
    <div class="vline"></div>
    <div class="tab"
         class:active={$selectedTab === kTabList}
         on:click={()=>selectTab(kTabList)}>
      Alphabet
    </div>
  </div>
  <div class="list">
    <div class="list-body term-tree" class:active={$selectedTab===kTabTree}>
      {#each categories as cat}
        <div class="item category">
          <div class="label" on:click|preventDefault={()=>selectCategory(cat)}>
            {cat.display}
          </div>
        </div>
        <div class="items" class:hidden={$selectedCategory?.key!==cat.key && !$searchValue}>
          {#each cat.entries as entry}
            <div class="item term-item" class:hidden={isHiddenEntry(entry, $searchValueNorm)}>
              <a class="label" href={entryUrl(entry.key)}
                 class:active={$selectedEntry?.key === entry.key}
                 on:click|preventDefault={()=>selectEntry(entry.key)}>
                {entry.display}
              </a>
            </div>
          {/each}
        </div>
      {/each}
    </div>
    <div class="list-body term-alpha" class:active={$selectedTab===kTabList}>
      {#each entries as entry}
        <div class="item term-item" class:hidden={isHiddenEntry(entry, $searchValueNorm)}>
          <a class="label" href={entryUrl(entry.key)}
             class:active={$selectedEntry?.key === entry.key}
             on:click|preventDefault={()=>selectEntry(entry.key)}>
            {entry.display}
          </a>
        </div>
      {/each}
    </div>
  </div>
</div>

<script lang="ts">
  import {categories, entries} from '../share/data.js';
  import {
    entryUrl,
    kTabList,
    kTabTree,
    routeActive,
    searchValue,
    searchValueNorm,
    selectedCategory,
    selectedEntry,
    selectedTab,
    selectEntry,
  } from '../share/store.js';
  import type {Category, Entry} from '../share/types.js';
  import SearchBox from './SearchBox.svelte';

  function selectTab(name: string) {
    $selectedTab = name;
    $routeActive = false;
  }

  function selectCategory(cat: Category) {
    if ($selectedCategory !== cat) $selectedCategory = cat;
    else $selectedCategory = undefined;
  }

  function isHiddenEntry(entry: Entry, searchValueNorm: string) {
    if (!searchValueNorm) return false;
    return entry.key.indexOf(searchValueNorm) < 0;
  }
</script>

<style lang="scss">
  .wrap-left {
    display: flex;
    flex-direction: column;
    font-size: 16px;
    background-color: var(--bg-cyan-color);
    flex: 1 0 0;
  }

  .hidden {
    display: none;
  }

  .tabs {
    display: flex;
    line-height: 28px;
    height: 34px;
    box-shadow: rgba(0, 0, 0, 0.04) 0 2px 3px 0;

    .vline {
      flex: 0 0 1px;
      height: 20px;
      margin-top: 6px;
      border-left: solid 1px var(--outline-color);
    }
  }

  .tab {
    flex: 1 1 0;
    text-align: center;
    cursor: pointer;

    &:hover {
      color: #1768B0;
    }

    &.active {
      font-weight: bold;
      color: var(--a-hover-color);
    }
  }

  .list {
    flex: 1 0 0;
    overflow-y: scroll;
    overflow-x: auto;

    .label {
      display: block;
      padding-left: 18px;
    }

    .item {
      line-height: 24px;
    }

    .category:not(:first-child) {
      margin-top: 10px;
    }

    .category .label {
      font-weight: bold;
    }

    .term-item, .category {
      .label {
        cursor: pointer;

        &:hover {
          background-color: var(--bg-hover-color);
          color: var(--a-hover-color);
        }

        &.active {
          background-color: var(--a-hover-color);
          color: #fff;
          font-weight: bold;
        }
      }
    }

    .term-tree .term-item .label {
      padding-left: 30px;
    }

    .term-tree {
      .category {

      }

      .items {
        flex: 0 1 auto;
      }
    }
  }

  .list-body {
    display: none;
    flex-direction: column;
    padding: 15px 0 25px;

    &.active {
      display: flex;
    }
  }

  @media (min-width: 768px) {
    .header {
      display: none;
    }

    .tabs {
      height: 34px;
      border-bottom: solid 1px var(--outline-color);
    }

    .tab.active {
      border-bottom: solid 3px var(--a-hover-color);
    }
  }

  @media (max-width: 768px) {
    .wrap-left {
      flex-direction: column-reverse;
    }

    .search {
      display: none;
    }

    .tabs {
      height: 40px;
      border-top: solid 1px var(--outline-color);

      .vline {
        margin-top: 10px;
      }
    }

    .tab {
      border-top: solid 3px rgba(0, 0, 0, 0);
    }

    .tab.active {
      border-top: solid 3px var(--a-hover-color);
    }
  }
</style>
