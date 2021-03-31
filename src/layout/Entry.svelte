{#if $selectedEntry}
  <div class="title">
    <div class="primary">{$selectedEntry.display || ''}</div>
    {#if !primaryEntry}
      {#each listTerms($selectedEntry) as term, i}
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
      <a href={entryUrl(primaryEntry.key)}
         on:click|preventDefault={()=>selectEntry(primaryEntry.key)}>
        {primaryEntry.display}</a>.
    {:else}
      {@html $selectedEntry.content.html}
    {/if}
  </div>
{/if}
{#if !primaryEntry}
  {#if relatedTerms}
    <div class="extra related-terms">
      <div class="header">Thuật ngữ liên quan</div>
      {#each relatedTerms as item,i}
        <div class="item">
          <a href={entryUrl(item.key)}
             on:click|preventDefault={()=>selectEntry(item.key)}>
            {item.display}</a>
        </div>
      {/each}
    </div>
  {/if}
  {#if $selectedEntry?.relatedArticles}
    <div class="extra related-articles">
      <div class="header">Bài viết liên quan</div>
      {#each $selectedEntry.relatedArticles as item,i}
        <div class="item">
          <a href={item.url} target="_blank">{item.display}</a>
        </div>
      {/each}
    </div>
  {/if}
  {#if $selectedEntry?.footnotes}
    <div class="footnotes-border"></div>
    <ol class="footnotes">
      {#each $selectedEntry.footnotes as item,i}
        <li class="note" id={'note-'+(i+1)}>{item}</li>
      {/each}
    </ol>
  {/if}
{/if}

<script lang="ts">
  import {entryUrl, gotoUrl, randomRelatedTerms, selectedEntry, selectEntry} from '../share/store.js';
  import type {Entry, Term} from '../share/types.js';

  $: primaryEntry = $selectedEntry?.primaryEntry;
  $: relatedTerms = randomRelatedTerms($selectedEntry);
  $: {
    $selectedEntry;
    setTimeout(() => {
      const link$Items = document.querySelectorAll('.content a[data-ilink]');
      for (let item: Element of link$Items) {
        item.addEventListener('click', (e) => {
          e.preventDefault();
          const a = e.target as HTMLAnchorElement;
          gotoUrl(a.href);
        });
      }
    });
  }

  function listTerms(entry: Entry): Term[] {
    if (!entry) return [];
    const terms = entry.terms || [];
    return terms.filter((term) => term.key !== entry.key);
  }
</script>

<style lang="scss">
  .panel {
    flex: 6 6 0;
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

  .content {
    line-height: 28px;
  }

  .footnotes-border {
    margin: 60px 0 20px 0;
    height: 1px;
    width: 100px;
    background: #aaa;
  }

  .footnotes {
    margin: 0;
    padding: 0 0 0 30px;
  }

  .content + .extra {
    margin-top: 80px;
  }

  .extra {
    margin-top: 40px;

    .header {
      font-size: 18px;
      font-weight: bold;
    }
  }
</style>
