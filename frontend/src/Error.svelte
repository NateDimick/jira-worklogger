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
        &#10062; <!-- X -->
    </button>
</div>
{/if}

<style>
div {
    position: fixed;
    z-index: 10;
    background-color: red;
    color: whitesmoke;
    bottom: 0;
}
</style>
