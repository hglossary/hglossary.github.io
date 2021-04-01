<div class={panelClassNames($currentRoute)}>
  <svelte:component this={routeComponent($currentRoute)}/>
</div>

<script lang="ts">
  import {routeComponent} from '../share/pages.js';
  import {currentRoute} from '../share/store.js';
  import type {Route} from '../share/types.js';

  function panelClassNames(r: Route): string {
    const cs = ['wrap'];
    cs.push(r.page + '-page');
    return cs.join(' ');
  }
</script>

<style lang="scss">
  .wrap {
    border-radius: 4px 0 0 0;
    box-shadow: rgba(0, 0, 0, 0.2) 0 0 5px 0;
    background: #fff;
    padding: 30px 40px;
    min-height: calc(100vh - 48px);
  }
</style>
