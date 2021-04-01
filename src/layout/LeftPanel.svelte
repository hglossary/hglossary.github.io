<div class="wrap">
  <div class="search">
    <div class="icon icon-search"></div>
    <div class="text">Tìm kiếm</div>
    <input class="input" type="text" class:active={$searchValue} bind:value={$searchValue} on:keydown={captureEsc}>
    <div class="close-search icon-close" class:hidden={!$searchValue} on:click={clearSearchInput}></div>
  </div>
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
    searchValue,
    searchValueNorm,
    selectedCategory,
    selectedEntry,
    selectedTab,
    selectEntry,
  } from '../share/store.js';
  import type {Category, Entry} from '../share/types.js';

  function isSelected(selected: string, name: string) {
    return name === selected;
  }

  function selectTab(name: string) {
    $selectedTab = name;
  }

  function selectCategory(cat: Category) {
    if ($selectedCategory !== cat) $selectedCategory = cat;
    else $selectedCategory = undefined;
  }

  function clearSearchInput() {
    $searchValue = '';
  }

  function captureEsc(e) {
    if (e.key === 'Escape') {
      $searchValue = '';
    }
  }

  function isHiddenEntry(entry: Entry, searchValueNorm: string) {
    if (!searchValueNorm) return false;
    return entry.key.indexOf(searchValueNorm) < 0;
  }
</script>

<style lang="scss">
  .wrap {
    display: flex;
    flex-direction: column;
    font-size: 16px;
    background-color: #F7F9FC;
    flex: 1 0 0;
  }

  .hidden {
    display: none;
  }

  .search {
    height: 32px;
    text-align: center;
    display: flex;
    justify-content: center;
    position: relative;
    margin: 5px 15px 5px;

    .icon {
      width: 32px;
      height: 30px;
      margin-top: 1px;
      background-position: center;
      background-size: 20px;
      opacity: 60%;
    }

    .close-search {
      cursor: pointer;
      position: absolute;
      top: 6px;
      right: 5px;
      width: 20px;
      height: 20px;
      border-radius: 10px;
      background-color: #ccc;
      background-position: center;
      background-size: 12px;
      opacity: 60%;
      transition: 0.2s;

      &:hover {
        opacity: 90%;
      }
    }

    .text {
      line-height: 30px;
      color: #888;
      font-size: 14px;
    }

    .input {
      position: absolute;
      top: 2px;
      bottom: 2px;
      width: 100%;
      display: block;
      padding: 0;
      border: solid 1px #ddd;
      border-radius: 40px;
      text-align: center;
      line-height: 30px;
      font-size: 14px;
      background: rgba(0, 0, 0, 0);
      transition-duration: 0.2s;
      transition-property: background-color, border-color;

      &.active, &:focus {
        outline: none;
        transition-duration: 0.1s;
        background: #fff;
        border-color: var(--outline-focus-color);
      }
    }
  }

  .tabs {
    display: flex;
    line-height: 28px;
    height: 34px;
    border-bottom: solid 1px var(--outline-color);
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
      border-bottom: solid 3px var(--a-hover-color);
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

</style>
