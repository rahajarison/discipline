<template>
    <h1 class="text-4xl font-bold text-gray-800 dark:text-white">Analysis</h1>
    <div class="flex flex-col">
        <div class="grow">
            <Timeline />
        </div>
        <div class="my-8 basis-4/12">
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
import InputPanel from '../components/InputPanel.vue';
import Timeline from '../components/Timeline.vue';
import ExportBar from '../components/ExportBar.vue';
import KeyMapDisplay from '../components/KeyMapDisplay.vue';


export default {
    components: { InputPanel, Timeline, ExportBar, KeyMapDisplay },
}
</script>