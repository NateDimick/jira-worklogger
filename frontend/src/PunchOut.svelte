<script lang="ts">
import type { main } from "./wailsjs/go/models";
import { config, issue, View, view } from "./stores"
import Back from "./Back.svelte";
import { ManualStopWork } from "./wailsjs/go/main/App";
import { configComplete } from "./config";


let issueKey: string = $issue.IssueKey
let comment: string
let workEndDt: string = localDt()

function localDt(): string  {
    let d = new Date()
    return `${d.getFullYear()}-${padNumber(d.getMonth() + 1)}-${padNumber(d.getDate())}T${padNumber(d.getHours())}:${padNumber(d.getMinutes())}`
}

function padNumber(n: number) {
    let ns: string = `${n}`
    if (ns.length == 1) {
        ns = `0${n}`
    }
    return ns
}

function closePunchOut() {
    view.set(View.Main)
}

function logWork() {
    let update = {IssueKey: issueKey, Comment: comment} as main.IssueUpdate
    ManualStopWork(update, workEndDt).then((success) => {
        console.log("maunal log work result:", success)
        if (success) { 
            issue.set({IssueKey: issueKey, Started: false, StartTime: new Date().toISOString()})
            closePunchOut() 
        }
    })
}

function inFuture(ts: string): boolean {
    let result = Date.parse(ts) > Date.now()
    return result
}
</script>

<Back/>
<div class="container">
    <h1>{issueKey}</h1>
    <label for="comment">Comment</label>
    <input type="text" name="comment" bind:value={comment}>
    <label for="workEnd">Work End Time</label>
    <input type="datetime-local" name="workEnd" bind:value={workEndDt}>
    <button on:click={logWork} disabled={!configComplete($config) || inFuture(workEndDt)}>
        Log Work Ending at {workEndDt}
    </button>
</div>

<style>
    .container {
        padding-top: 15vh;
    }
</style>