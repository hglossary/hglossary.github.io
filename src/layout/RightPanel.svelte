<div class="panel">
  {#if $selectedEntry}
    <div class="title">
      <div class="primary">{$selectedEntry.display || ''}</div>
      {#if !primaryEntry}
        {#each listTerms($selectedEntry) as term,i}
          {#if i}
            <div class="circle">・</div>
          {/if}
          <div class="secondary">{term.display}</div>
        {/each}
      {/if}
    </div>
    <div class="content">
      {#if primaryEntry}
        Xem
        <a href={entryUrl(primaryEntry.key)} on:click|preventDefault={selectEntry(primaryEntry)}>
          {primaryEntry.display}</a>.
      {:else}
        {@html $selectedEntry.content.html}
      {/if}
    </div>
  {/if}
</div>

<script lang="ts">
  import type {Entry, Term} from '../share/store.js';
  import {entryUrl, selectedEntry} from '../share/store.js';

  $: primaryEntry = $selectedEntry?.primaryEntry;

  function selectEntry(entry: Entry) {
    return () => $selectedEntry = entry;
  }

  function listTerms(entry: Entry): Term[] {
    if (!entry) return [];
    const terms = entry.terms || [];
    return terms.filter((term) => term.key !== entry.key);
  }
</script>

<style lang="scss">
  .panel {
    flex: 3 3 0;
    padding: 30px 40px;
  }

  .title {
    margin-bottom: 24px;

    .primary {
      font-size: 24px;
      font-weight: bold;
    }

    .secondary {
      display: inline-block;
      font-size: 18px;
      font-weight: bold;
      color: #333;
    }

    .circle {
      display: inline-block;
      text-align: center;
      width: 20px;
      padding-left: 3px;
    }
  }

  a {
    color: var(--a-hover-color);
    text-decoration: underline var(--a-hover-color);
  }

  .content {
    line-height: 28px;
  }
</style>
