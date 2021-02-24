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
    background: #F7F9FC;
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .nav {
    flex: 0 0 48px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .body {
    flex: 1 1 0;
    display: flex;
    flex-direction: row;
    overflow: hidden;
    padding-top: 5px;
    margin-top: -5px;
  }
</style>
