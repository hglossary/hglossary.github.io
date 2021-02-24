<div class="root">
  <div class="nav">
    <Navigator/>
  </div>
  <div class="body">
    <LeftPanel/>
    <RightPanel/>
  </div>
</div>

<script>
  import LeftPanel from './layout/LeftPanel.svelte';
  import Navigator from './layout/Navigator.svelte';
  import RightPanel from './layout/RightPanel.svelte';
</script>

<style lang="scss">
  .root {
    background: var(--bg-color);
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .nav {
    flex: 0 0 40px;
    border-bottom: solid 1px #aaaaaa;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .body {
    flex: 1 0 auto;
    display: flex;
    flex-direction: row;
  }
</style>
