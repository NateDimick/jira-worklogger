<script lang="ts">
import { UpdateConfig } from "./wailsjs/go/main/App";
import { config, View, view } from "./stores";
import type { main } from "./wailsjs/go/models";
import Back from "./Back.svelte"

let token: string = $config.Token
let server: string = $config.Server

function closeConfig() {
    view.set(View.Main)
}

function updateConfig() {
    let update = {Token: token, Server: server} as main.Config
    UpdateConfig(update).then((success: boolean) => {
        if (success) {
            config.set(update)
            closeConfig()
        }
    })
}
</script>

<Back/>
<div class="container">
    <h1>Config</h1>
    <label for="token">Personal Access Token</label>
    <input type="password" name="token" bind:value={token}>
    <label for="server">Jira Domain</label>
    <input type="text" name="server" bind:value={server}>
    <button on:click={updateConfig}>
        Save
    </button>
</div>

<style>
.container {
    padding-top: 15vh;
}
</style>