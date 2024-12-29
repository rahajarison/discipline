<template>
    <div>
        <h1 class="text-4xl font-bold text-gray-900 dark:text-white">Timeline</h1>
        <div v-for="(round, index) in state.rounds" :key="index" class="my-8">
            <h2 class="text-2xl font-bold text-gray-500 dark:text-white">Round {{ index + 1 }}</h2>
            <!-- <ul> -->
                <ol class="items-center sm:flex flex-nowrap">
                    <li class="relative mb-6 sm:mb-0" v-for="action in round.actions" :key="action.id">
                        <div class="flex items-center">
                            <div class="z-10 flex items-center justify-center w-6 h-6 bg-blue-100 rounded-full ring-0 ring-white dark:bg-blue-900 sm:ring-8 dark:ring-gray-900 shrink-0">
                                <svg class="w-2.5 h-2.5 text-blue-800 dark:text-blue-300" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                                    <path d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z"/>
                                </svg>
                            </div>
                            <div class="hidden sm:flex w-full bg-gray-200 h-0.5 dark:bg-gray-700"></div>
                        </div>
                        <div class="mt-3 sm:pe-8">
                            <h3 class="text-lg font-semibold text-gray-900 dark:text-white">{{ action.type }}</h3>
                            <time class="block mb-2 text-sm font-normal leading-none text-gray-400 dark:text-gray-500">{{ formatTimestamp(action.timestamp) }}</time>
                            <p class="text-base font-normal text-gray-500 dark:text-gray-400">
                                <kbd class="px-2 py-1.5 text-xs font-semibold text-gray-800 bg-gray-100 border border-gray-200 rounded-lg dark:bg-gray-600 dark:text-gray-100 dark:border-gray-500">{{ action.hitContext }}</kbd>
                                ({{ action.player }})
                            </p>
                        </div>
                    </li>
                </ol>
                <!-- <li v-for="action in round.actions" :key="action.id">
                    [{{ action.hitContext }}] {{ action.type }} ({{ action.player }}) - {{ formatTimestamp(action.timestamp) }}
                </li> -->
            <!-- </ul> -->
        </div>
    </div>
</template>

<script>
import { useGameStore } from '../store/gameStore';
import { formatTimestamp } from '../utils/formatTimestamp';

export default {
    setup() {
        const gameStore = useGameStore();

        return {
            state: gameStore,
            formatTimestamp,
        };
    },
};
</script>
