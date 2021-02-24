<div class="panel">
  <div class="search">
    <div class="icon icon-search"></div>
    <div class="text">Tìm kiếm</div>
    <input class="input" type="text" class:active={$searchValue} bind:value={$searchValue}>
  </div>
  <div class="tabs">
    <div class="tab"
         class:active={$selectedTab === 'tree'}
         on:click={selectTab('tree')}>
      Danh mục
    </div>
    <div class="vline"></div>
    <div class="tab"
         class:active={$selectedTab === 'alpha'}
         on:click={selectTab('alpha')}>
      Alphabet
    </div>
  </div>
  <div class="list">
    <div class="list-body term-tree" class:active={$selectedTab==='tree'}>
      {#each categories as cat}
        <div class="item category">
          <div class="label" on:click|preventDefault={selectCategory(cat)}>
            {cat.display}
          </div>
          <div class="children">
            {#each cat.entries as entry}
              <div class="item term-item" class:hidden={isHiddenEntry(entry, $searchValueNorm)}>
                <a class="label" href={entryUrl(entry.key)}
                   class:active={$selectedEntry?.key === entry.key}
                   on:click|preventDefault={selectEntry(entry)}>
                  {entry.display}
                </a>
              </div>
            {/each}
          </div>
        </div>
      {/each}
    </div>
    <div class="list-body term-alpha" class:active={$selectedTab==='alpha'}>
      {#each entries as entry}
        <div class="item term-item" class:hidden={isHiddenEntry(entry, $searchValueNorm)}>
          <a class="label" href={entryUrl(entry.key)}
             class:active={$selectedEntry?.key === entry.key}
             on:click|preventDefault={selectEntry(entry)}>
            {entry.display}
          </a>
        </div>
      {/each}
    </div>
  </div>
</div>

<script lang="ts">
  import type {Category, Entry} from '../share/store.js';
  import {
    categories,
    entries,
    entryUrl,
    searchValue,
    searchValueNorm,
    selectedCategory,
    selectedEntry,
    selectedTab,
  } from '../share/store.js';

  function isSelected(selected: string, name: string) {
    return name === selected;
  }

  function selectTab(name: string) {
    return () => $selectedTab = name;
  }

  function selectCategory(cat: Category) {
    return () => $selectedCategory = cat;
  }

  function selectEntry(entry: Entry) {
    return () => $selectedEntry = entry;
  }

  function clearSearchInput() {
    $searchValue = '';
  }

  function isHiddenEntry(entry: Entry, searchValueNorm: string) {
    if (!searchValueNorm) return false;
    console.log('--', entry.key, searchValueNorm, entry.key.indexOf(searchValueNorm))
    return entry.key.indexOf(searchValueNorm) < 0;
  }
</script>

<style lang="scss">
  .panel {
    display: flex;
    flex-direction: column;
    flex: 1 1 0;
    border-right: solid 1px #aaa;
    font-size: 16px;
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

    .icon {
      width: 32px;
      height: 30px;
      margin-top: 2px;
      background-position: center;
      background-size: 20px;
      opacity: 60%;
    }

    .text {
      line-height: 32px;
      color: #444;
    }

    .input {
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      right: 0;
      width: 100%;
      display: block;
      padding: 0;
      border: none;
      text-align: center;
      font-size: 16px;
      background: #F8F6F4;
      opacity: 0;
      transition-duration: 0.2s;
      transition-property: opacity;

      &.active, &:focus {
        opacity: 1;
        outline: none;
        transition-duration: 0.1s;
      }
    }
  }

  .tabs {
    display: flex;
    line-height: 28px;
    height: 34px;
    border-top: solid 1px #aaa;
    border-bottom: solid 1px #aaa;

    .vline {
      flex: 0 0 1px;
      border-left: solid 1px #aaa;
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
    }
  }

  .list {
    padding: 16px 0;
    flex: 1 0 0;
    overflow: scroll;

    .label {
      display: block;
      padding-left: 18px;

      &.active {
        font-weight: bold;
      }
    }

    .children .label {
      padding-left: 30px;
    }

    .item {
      line-height: 24px;
    }

    .category:not(:first-child) {
      padding-top: 10px;
    }

    .category > .label {
      font-weight: bold;
    }

    .category > .label, .term-item {
      cursor: pointer;

      &:hover {
        background-color: var(--bg-hover-color);

        a {
          color: var(--a-hover-color);
        }
      }
    }
  }

  .list-body {
    display: none;
    flex-direction: column;

    &.active {
      display: flex;
    }
  }

</style>
