<script lang="ts">
    let ok: string = $state("loading...");

    fetch("/api/health").then((resp) => {
        resp.text().then((resp) => {
            ok = resp;
        });
    });
</script>

<div>
    <span>Status: {ok}</span>
</div>
