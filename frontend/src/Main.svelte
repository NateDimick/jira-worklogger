<script lang="ts">
import { configComplete } from "./config";
import { config, issue, View, view } from "./stores";
import { AbortCurrentIssue, StartStopWorkTime } from "./wailsjs/go/main/App";
import type { main } from "./wailsjs/go/models";



let comment: string = ""
let workTime: string = ""
let intervalId
let issueKey: string = $issue.IssueKey || ""
let active: boolean = false
let startTime: number = Date.now()

$: {
    active = $issue.Started
    startTime = Date.parse($issue.StartTime)
}

function toggle() {
    let update = {IssueKey: issueKey, Comment: comment} as main.IssueUpdate
    console.log(`issue update: ${JSON.stringify(update)}`)
    StartStopWorkTime(update).then((result: main.StartStopResult) => {
        active = result.Success ? !active : active
        issue.set({IssueKey: issueKey, Started: active, StartTime: result.Time})
    })
}

function abort() {
    let d = new Date(startTime)
    AbortCurrentIssue().then((success) => {
        active = success ? !active : active
        issue.set({IssueKey: issueKey, Started: active, StartTime: d.toISOString()})
    })
}

function tick() {
    let now = Date.now()
    let elapsed = (now - startTime) / 1000
    let seconds = Math.floor(elapsed % 60)
    let minutes = Math.floor(elapsed % 3600 / 60)
    let hours = Math.floor(elapsed / 3600)
    workTime = `${hours} hours ${minutes} minutes ${seconds} seconds`
}

function punchOut() {
    view.set(View.PunchOut)
}

function openConfig() {
    view.set(View.Config)
}

$: {
    comment = active ? comment : ""
    if (active) {
    intervalId = setInterval(tick, 1000)
    } else {
    clearInterval(intervalId)
    }
}
</script>

<div class="upper-left">
    <button on:click={openConfig}>
        &#x2699;&#xFE0F; <!-- Gear -->
    </button>
</div>
<div class="container">
    <label for="issue">Jira Issue</label>
    <input type="text" name="issue" bind:value={issueKey} disabled={active}>
    {#if active }
        <label for="comment">Comment</label>
        <input type="text" name="comment" bind:value={comment}>
        <p>Current Session on {issueKey}</p>
        <p>{workTime}</p>
    {/if}
    <div class="btn-row">
        {#if active }
            <button on:click={abort}>
                Nevermind
            </button>
            <button on:click={toggle} disabled={!configComplete($config)}>
                Log Work Now
            </button>
            <button on:click={punchOut}>
                Pick End Time
            </button>
        {:else}
            <button on:click={toggle} disabled={issueKey === ""}>
                Start Work Now
            </button>
        {/if}
    </div>
</div>
  
<style>
  .btn-row {
    display: flex;
    flex-wrap: wrap;
    margin: 0 auto;
    text-align: center;
    width: fit-content;
  }
  .btn-row button {
    margin: 1em;
    width: 15em;
  }
</style>
  