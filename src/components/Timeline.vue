<template>
    <h1 class="text-4xl font-bold text-gray-900 dark:text-white">Timeline</h1>
    <div v-for="(round, index) in state.rounds" :key="index" class="my-8">
        <h2 class="text-2xl font-bold text-gray-500 dark:text-white">Round {{ index + 1 }}</h2>
        <ol class="items-center sm:flex flex-nowrap">
            <li class="relative mb-6 sm:mb-0" v-for="action in round.actions" :key="action.id">
                <div class="flex items-center">
                    <div
                        :class="{'bg-blue-100': action.player == 'P1', 'bg-red-100': action.player == 'P2', 'bg-gray-100': action.player == 'N/A'}"
                        class="z-10 flex items-center justify-center w-6 h-6 rounded-full ring-0 ring-white dark:bg-blue-900 sm:ring-8 dark:ring-gray-900 shrink-0">
                        <svg v-if="action.category == categories.SYSTEM" class="w-4 h-4 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.827 5.465-.435-2.324m.435 2.324a5.338 5.338 0 0 1 6.033 4.333l.331 1.769c.44 2.345 2.383 2.588 2.6 3.761.11.586.22 1.171-.31 1.271l-12.7 2.377c-.529.099-.639-.488-.749-1.074C5.813 16.73 7.538 15.8 7.1 13.455c-.219-1.169.218 1.162-.33-1.769a5.338 5.338 0 0 1 4.058-6.221Zm-7.046 4.41c.143-1.877.822-3.461 2.086-4.856m2.646 13.633a3.472 3.472 0 0 0 6.728-.777l.09-.5-6.818 1.277Z"/>
                        </svg>

                        <svg v-if="action.category == categories.NEUTRAL" class="w-4 h-4 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16h13M4 16l4-4m-4 4 4 4M20 8H7m13 0-4 4m4-4-4-4"/>
                        </svg>
                        <svg v-if="action.category == categories.DEFENSE" class="w-4 h-4 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path fill-rule="evenodd" d="M13.729 5.575c1.304-1.074 3.27-.146 3.27 1.544v9.762c0 1.69-1.966 2.618-3.27 1.544l-5.927-4.881a2 2 0 0 1 0-3.088l5.927-4.88Z" clip-rule="evenodd"/>
                        </svg>
                        <svg v-if="action.category == categories.OFFENSE" class="w-4 h-4 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="currentColor" viewBox="0 0 24 24">
                            <path fill-rule="evenodd" d="M10.271 5.575C8.967 4.501 7 5.43 7 7.12v9.762c0 1.69 1.967 2.618 3.271 1.544l5.927-4.881a2 2 0 0 0 0-3.088l-5.927-4.88Z" clip-rule="evenodd"/>
                        </svg>
                    </div>
                    <div class="hidden sm:flex w-full bg-gray-200 h-0.5 dark:bg-gray-700"></div>
                </div>
                <div class="mt-1 sm:pe-8">
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">{{ action.type }}</h3>
                    <time class="block mb-2 text-sm font-normal leading-none text-gray-400 dark:text-gray-500">{{ formatTimestamp(action.timestamp) }}</time>
                    <p class="text-base font-normal text-gray-500 dark:text-gray-400">
                        <span
                            :class="{ 'border-blue-400': action.player == 'P1', 'text-blue-400': action.player == 'P1', 'border-red-400': action.player == 'P2', 'text-red-400': action.player == 'P2'}"
                            class="border border-dashed rounded-lg px-2 text-xs font-semibold">
                                {{ action.player }}
                        </span>
                        <kbd class="px-2 py-1.5 text-xs font-semibold text-gray-800 bg-gray-100 border border-gray-200 rounded-lg dark:bg-gray-600 dark:text-gray-100 dark:border-gray-500">{{ action.hitContext }}</kbd>
                    </p>
                </div>
            </li>
        </ol>
    </div>
</template>

<script>
import { useGameStore } from '../store/gameStore';
import { formatTimestamp } from '../utils/formatTimestamp';
import { NEUTRAL, DEFENSE, OFFENSE, SYSTEM } from '../utils/categories';

export default {
    setup() {
        const gameStore = useGameStore();

        return {
            state: gameStore,
            categories: {
                NEUTRAL,
                DEFENSE,
                OFFENSE,
                SYSTEM,
            },
            formatTimestamp,
        };
    },
};
</script>
