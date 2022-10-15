<script lang="ts">
import { EventsOn } from "./wailsjs/runtime/runtime";

let error: string = ""

function closeError() {
    error = ""
}

EventsOn("error", (err: string) => {
    console.log(`got error: ${err}`)
    error = err
})
</script>

{#if error !== ""}
<div>
    <p>{error}</p>
    <button on:click={closeError}>
        Dismiss
    </button>
</div>
{/if}

<style>
div {
    position: fixed;
    background-color: red;
    color: whitesmoke;
    bottom: 0;
    width: 100%;
}
p {
    display: inline;
    margin: 0 auto;
}
button {
    display: inline;
    background-color: red;
    border-color: black;
    padding: 1em;
    right: 0;
    bottom: 0;
}
</style>
