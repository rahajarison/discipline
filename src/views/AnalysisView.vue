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
            <RouterLink to="/reports">
                <button type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center me-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    <svg class="w-3.5 h-3.5 me-2" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 18 21">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m8-2h3m-3 3h3m-4 3v6m4-3H8M19 4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1ZM8 12v6h8v-6H8Z"/>
                    </svg>
                    Voir rapport
                </button>
            </RouterLink>
            <RouterLink to="/analysis">
                <button type="button" class="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-center inline-flex items-center me-2 text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700">
                    Nouvelle analyse
                    <svg class="w-4 h-4 ms-2 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7 7V5"/>
                    </svg>

                </button>
            </RouterLink>
            <!-- <ExportBar /> -->
        </div>
    </div>
</template>

<script>
import InputPanel from '../components/InputPanel.vue';
import Timeline from '../components/Timeline.vue';
// import ExportBar from '../components/ExportBar.vue';
import KeyMapDisplay from '../components/KeyMapDisplay.vue';


export default {
    components: { InputPanel, Timeline, KeyMapDisplay },
}
</script>