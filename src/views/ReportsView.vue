<template>
    <h1 class="m-4 text-4xl font-bold text-gray-800 dark:text-white">Rapports</h1>
    <!-- {{ actionList }} -->
    <div class="relative overflow-x-auto">
        <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                <tr>
                    <th scope="col" class="px-6 py-3">
                        Round
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Joueur
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Catégorie
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Type
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Hit context
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Date
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700" v-for="action in actionList" :key="action.id">
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                        {{ action.round }}
                    </th>
                    <td class="px-6 py-4 " :class="{ 'text-blue-400': action.player == 'P1', 'text-red-400': action.player == 'P2'}">
                        {{ action.player }}
                    </td>
                    <td class="px-6 py-4">
                        {{ action.category }}
                    </td>
                    <td class="px-6 py-4">
                        {{ action.type }}
                    </td>
                    <!-- <td class="px-6 py-4 bg-red-500 text-white"> -->
                    <!-- <td class="px-6 py-4 bg-gray-500 text-white"> -->
                    <td class="px-6 py-4"
                        :class="{ 'bg-red-500 text-white': action.hitContext == hitContexts.ON_HIT, 'bg-gray-500 text-white': action.hitContext == hitContexts.ON_BLOCK, 'bg-red-800 text-white': action.hitContext == hitContexts.WHIFF }">
                        {{ action.hitContext }}
                    </td>
                    <td class="px-6 py-4">
                        {{ formatTimestamp(action.timestamp) }}
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
import { computed } from 'vue';
import { useGameStore } from '../store/gameStore';
import { formatTimestamp } from '../utils/formatTimestamp';
import { ON_HIT, ON_BLOCK, WHIFF, EVENT_NOTICEABLE } from '../utils/hitContexts';

export default {
    setup() {
        const gameStore = useGameStore();
        const actionList = computed(() => gameStore.flatActionList);

        return {
            state: gameStore,
            actionList,
            formatTimestamp,
            hitContexts: {
                ON_HIT,
                ON_BLOCK,
                WHIFF,
                EVENT_NOTICEABLE
            }
        }
    }
}
</script>