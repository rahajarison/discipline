<template>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                <tr>
                    <th scope="col" class="px-6 py-3">
                        ID
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Date
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Replay ID
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Vidéo
                    </th>
                    <th scope="col" class="px-6 py-3">
                        Consulter
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" v-for="match in state.matches">
                    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                        {{ match.id }}
                    </th>
                    <td class="px-6 py-4">
                        {{ formatTimestamp(match.timestamp) }}
                    </td>
                    <td class="px-6 py-4">
                        XXXXXX
                    </td>
                    <td class="px-6 py-4">
                        YYYYYY
                    </td>
                    <td class="px-6 py-4 text-right">
                        <RouterLink :to="{ name: 'match-history', params: { matchId: match.id } }" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Choisir</RouterLink>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script>
// import { computed, ref } from 'vue';
import { useMatchStore } from '../store/matchStore';
import { formatTimestamp } from '../utils/formatTimestamp';

export default {
    setup() {
        const matchStore = useMatchStore();
        return {
            state: matchStore,
            formatTimestamp
        }
    }
}

</script>