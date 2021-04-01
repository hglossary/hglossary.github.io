<div class="text panel-left">
  <a href="/" on:click|preventDefault={()=>gotoUrl('/')}>NGÀNH DU LỊCH</a>
</div>
<div class="right panel-right">
  <div class="item" class:active={$currentUrl==='/a/about'}>
    <a href="/a/about" on:click|preventDefault={()=>gotoUrl('/a/about')}>Giới thiệu</a>
  </div>
  <div class="item" class:active={$currentUrl==='/a/contact'}>
    <a href="/a/contact" on:click|preventDefault={()=>gotoUrl('/a/contact')}>Liên hệ</a>
  </div>
</div>

<script lang="ts">
  import {currentUrl, gotoUrl} from '../share/store.js';
</script>

<style lang="scss">
  .text {
    padding: 16px 40px 0;
    text-align: center;

    a {
      font-family: Rockwell, serif;
      font-size: 20px;
    }
  }

  .right {
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
</style>
