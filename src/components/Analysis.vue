<template>
    <div class="flex flex-col">
        <div class="grow">
            <Timeline />
        </div>
        <div class="basis-4/12">
            <div class="flex flex-row">
                <div class="basis-1/3"><InputPanel /></div>
                <div class="basis-2/3"><KeyMapDisplay /></div>
            </div>
        </div>
        <div class="basis-1/12">
            <ExportBar />
        </div>
    </div>
</template>

<script>
import InputPanel from './InputPanel.vue';
import Timeline from './Timeline.vue';
import ExportBar from './ExportBar.vue';
import KeyMapDisplay from './KeyMapDisplay.vue';


export default {
    components: { InputPanel, Timeline, ExportBar, KeyMapDisplay },
}
</script>