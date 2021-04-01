<div class="root" class:route-active={$routeActive}>
  <div class="panel-left">
    <div class="header" class:search-active={$mobileSearchActive}>
      <a href="/" on:click|preventDefault={()=>gotoUrl('/')}>NGÀNH DU LỊCH</a>
      <div class="icon icon-back" data-screen="mobile"
           on:click|preventDefault={()=>back()}></div>
      <div class="icon icon-search" data-screen="mobile"
           on:click|preventDefault={()=>toggleSearchActive()}></div>
      <SearchBox screen="mobile"/>
    </div>
    <div class="panel-scroll" data-screen="mobile">
      <RightPanel screen="mobile"/>
    </div>
    <LeftPanel/>
  </div>
  <div class="panel-right" data-screen="desktop">
    <div class="panel-inner">
      <div class="header">
        <div class="item" class:active={$currentUrl==='/about'}>
          <a href="/about" on:click|preventDefault={()=>gotoUrl('/about')}>Giới thiệu</a>
        </div>
        <div class="item" class:active={$currentUrl==='/contact'}>
          <a href="/contact" on:click|preventDefault={()=>gotoUrl('/contact')}>Liên hệ</a>
        </div>
      </div>
      <RightPanel screen="desktop"/>
    </div>
  </div>
</div>

<script>
  import LeftPanel from './layout/LeftPanel.svelte';
  import RightPanel from './layout/RightPanel.svelte';
  import SearchBox from './layout/SearchBox.svelte';
  import {currentUrl, gotoUrl, mobileSearchActive, routeActive, searchValue} from './share/store.js';
  import {querySel} from './share/util.js';

  function toggleSearchActive() {
    $mobileSearchActive = !$mobileSearchActive;
    $routeActive = false;
    $searchValue = '';
    setTimeout(() => {
      querySel('.search[data-screen=mobile] .input').focus();
    }, 100);
  }

  function back() {
    $routeActive = false;
    if ($searchValue !== '') {
      $mobileSearchActive = true;
    }
  }
</script>

<style lang="scss">
  .root {
    background: #F7F9FC;
    height: 100%;
    display: flex;
  }

  .panel-left {
    .header {
      text-align: center;

      a {
        font-family: Rockwell, serif;
        font-size: 20px;
      }
    }
  }

  @media (min-width: 769px) {
    .root {
      background: #F7F9FC;
      height: 100%;
      display: flex;
    }

    .nav {
      flex: 0 0 48px;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
    }

    .panel-left {
      flex: 1 0 160px;
      display: flex;
      flex-direction: column;

      .header {
        padding: 16px 40px 0;
        text-align: center;

        a {
          font-family: Rockwell, serif;
          font-size: 20px;
        }
      }
    }

    .panel-right {
      flex: 6 6 0;
      overflow-y: scroll;
      padding-left: 5px;
      margin-left: -5px;

      .header {
        display: flex;
        justify-content: flex-end;

        .item {
          font-size: 16px;
          line-height: 48px;
          padding-right: 24px;
        }

        .item.active a {
          color: var(--a-hover-color);
          text-decoration: underline;
        }
      }
    }
  }

  @media (max-width: 768px) {
    .root {
      flex-direction: column;
    }

    .panel-left {
      display: flex;
      flex: 1 0 0;
      flex-direction: column;

      .header {
        padding: 10px 10px;
        border-bottom: solid 1px var(--outline-color);
        box-shadow: rgb(0 0 0 / 4%) 0 2px 3px 0;
        z-index: 10;
        position: relative;
      }
    }

    .panel-right {
      display: none;
    }

    .icon {
      width: 32px;
      height: 30px;
      margin-top: 1px;
      background-position: center;
      background-size: 20px;
      opacity: 60%;
    }

    .icon-search {
      position: absolute;
      right: 10px;
      top: 6px;
      cursor: pointer;
    }

    .icon-back {
      position: absolute;
      left: 10px;
      top: 6px;
      cursor: pointer;
      display: none;
    }

    .route-active {
      .icon-back {
        display: block;
      }
    }

    .panel-scroll {
      display: none;
      overflow-y: scroll;
      flex: 1 0 0;
    }

    .route-active {
      .panel-scroll {
        display: block;
      }
    }
  }
</style>
