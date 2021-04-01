<div class={panelClassNames($currentRoute)}>
  <div class="inner">
    <svelte:component this={routeComponent($currentRoute)}/>
  </div>
</div>

<script lang="ts">
  import {routeComponent} from '../share/pages.js';
  import {currentRoute} from '../share/store.js';
  import type {Route} from '../share/types.js';

  function panelClassNames(r: Route): string {
    const cs = 'panel panel-right'.split(' ');
    cs.push(r.page + '-page');
    return cs.join(' ');
  }
</script>

<style lang="scss">
  .panel {
    border-radius: 4px 0 0 0;
    box-shadow: rgba(0, 0, 0, 0.2) 0 0 5px 0;
    background: #fff;
  }

  .inner {
    width: 100%;
    height: 100%;
    padding: 30px 40px;
    overflow-y: scroll;
  }
</style>
