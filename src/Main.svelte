<div class="root">
  <div class="panel-left">
    <div class="header">
      <a href="/" on:click|preventDefault={()=>gotoUrl('/')}>NGÀNH DU LỊCH</a>
    </div>
    <LeftPanel/>
  </div>
  <div class="panel-right">
    <div class="panel-inner">
      <div class="header">
        <div class="item" class:active={$currentUrl==='/a/about'}>
          <a href="/a/about" on:click|preventDefault={()=>gotoUrl('/a/about')}>Giới thiệu</a>
        </div>
        <div class="item" class:active={$currentUrl==='/a/contact'}>
          <a href="/a/contact" on:click|preventDefault={()=>gotoUrl('/a/contact')}>Liên hệ</a>
        </div>
      </div>
      <RightPanel/>
    </div>
  </div>
</div>

<script>
  import LeftPanel from './layout/LeftPanel.svelte';
  import RightPanel from './layout/RightPanel.svelte';
  import {currentUrl, gotoUrl} from './share/store.js';
</script>

<style lang="scss">
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

    .panel-inner {

    }
  }
</style>
