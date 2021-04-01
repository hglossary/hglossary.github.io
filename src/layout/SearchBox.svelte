<div class="search" data-screen={screen}>
  <div class="icon icon-search"></div>
  <div class="text">Tìm kiếm</div>
  <input class="input" type="text" class:active={$searchValue} bind:value={$searchValue} on:keydown={captureEsc}>
  <div class="close-search icon-close" class:hidden={!$searchValue} on:click={clearSearchInput}></div>
</div>

<script lang="ts">
  import {mobileSearchActive, searchValue} from '../share/store.js';

  export let screen: string;

  function clearSearchInput() {
    $searchValue = '';
    $mobileSearchActive = false;
  }

  function captureEsc(e) {
    if (e.key === 'Escape') {
      $searchValue = '';
    }
  }
</script>

<style lang="scss">
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
</style>
