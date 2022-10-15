<script lang="ts">
import Config from "./Config.svelte";
import Error from "./Error.svelte";
import Main from "./Main.svelte";
import PunchOut from "./PunchOut.svelte";
import { config, issue, View, view } from "./stores";
import { GetConfig, GetCurrentIssue } from "./wailsjs/go/main/App";

async function load() {
  let i = await GetCurrentIssue()
  issue.set(i)
  let c = await GetConfig()
  config.set(c)
}

let promise = load()
</script>

{#await promise}
  <p>Loading...</p>
{:then result} 
  {#if $view === View.Main}
  <Main/>
  {:else if $view === View.Config}
  <Config/>
  {:else if $view === View.PunchOut}
  <PunchOut/>
  {/if}
{/await}
<Error/>
