<script>
import { view, View } from "./stores";

function back() {
    view.set(View.Main)
}
</script>

<div class="upper-left">
    <button on:click={back}>
        &#x2B05;&#xFE0F; <!-- Back arrow -->
    </button>
</div>