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

<style>
  :global(html) {
    text-align: center;
    align-content: center;
    background-color: white;
    font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
    vertical-align: middle;
  }
  :global(body) {
    margin: 0;
  }
  :global(label) {
    display: block;
    margin: 1em;
  }
  :global(input) {
    display: block;
    margin: 1em auto;
  }
  :global(button) {
    background-color: forestgreen;
    color: whitesmoke;
    display: block;
    border-radius: 2em;
    border-color: forestgreen;
    padding: 1em 3em;
    margin: 1em auto;
  }
  :global(button:hover) {
    cursor: pointer;
  }
  :global(button:disabled) {
    background-color: maroon;
    border-color: maroon;
    cursor: default;
  }
  :global(.upper-left) {
    top: 0;
    left: 0;
    position: fixed;
    z-index: 3;
    background-color: white;
  }
  :global(.upper-left button) {
    padding: 1em;
    background-color: azure;
    border-color: black;
    color: black;
    margin: 0;
  }
  :global(.container) {
    padding-top: 25vh;
    height: 70vh;
    width: 100%;
    top: 0;
    position: fixed;
  }
</style>