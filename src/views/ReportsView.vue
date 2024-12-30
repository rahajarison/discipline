<template>
    <h1 class="m-4 text-4xl font-bold text-gray-800 dark:text-white">Rapports</h1>
    <!-- {{ actionList }} -->
    {{ filters }}
    
    <ReportsFilterDropdown
        :options="roundOptions"
        placeholder="Filtrer par rounds"
        v-model="filters.round"
    />
    <ReportsFilterDropdown
        :options="playerOptions"
        placeholder="Filtrer par joueur"
        v-model="filters.player"
    />
    <ReportsFilterDropdown
        :options="categoryOptions"
        placeholder="Filtrer par catégorie"
        v-model="filters.category"
    />
    <ReportsFilterDropdown
        :options="typeOptions"
        placeholder="Filtrer par type"
        v-model="filters.type"
    />
    <ReportsFilterDropdown
        :options="hitContextOptions"
        placeholder="Filtrer par hit context"
        v-model="filters.hitContext"
    />

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
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700" v-for="action in filteredActionList" :key="action.id">
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
import { computed, ref } from 'vue';
import { useGameStore } from '../store/gameStore';

import ReportsFilterDropdown from '../components/ReportsFilterDropdown.vue';

import { formatTimestamp } from '../utils/formatTimestamp';
import { ON_HIT, ON_BLOCK, WHIFF, EVENT_NOTICEABLE } from '../utils/hitContexts';
import { SYSTEM, NEUTRAL, DEFENSE, OFFENSE } from '../utils/categories';
import { EVENT, REVIEWER } from '../utils/types';

export default {
    components: { ReportsFilterDropdown },
    setup() {
        const gameStore = useGameStore();
        const filters = ref({
            round: '',
            player: '',
            category: '',
            type: '',
            hitContext: '',
        });

        const actionList = computed(() => gameStore.flatActionList);
        const filteredActionList = computed(() => {
            return gameStore.flatActionList.filter((action) => {
                return (
                (!filters.value.round || action.round === Number(filters.value.round)) &&
                (!filters.value.player || action.player === filters.value.player) &&
                (!filters.value.category || action.category === filters.value.category) &&
                (!filters.value.type || action.type === filters.value.type) &&
                (!filters.value.hitContext || action.hitContext === filters.value.hitContext)
                );
            });
        });

        // Filters options
        const roundOptions = computed(() => [
            { value: '', label: 'Tous les rounds' },
            ...gameStore.rounds.map((_, index) => ({
                value: "" + (index + 1),
                label: `Round ${index + 1}`,
            })),
        ]);
        const playerOptions = [
            { value: '', label: 'Tous les joueurs' },
            { value: 'P1', label: 'Joueur 1' },
            { value: 'P2', label: 'Joueur 2' },
        ];
        const categoryOptions = [
            { value: '', label: 'Toutes les categories' },
            { value: NEUTRAL, label: 'Neutral' },
            { value: OFFENSE, label: 'Offense' },
            { value: DEFENSE, label: 'Defense' },
            { value: SYSTEM, label: 'Système' },
        ];
        const typeOptions = [
            { value: '', label: 'Tous les types' },
            { value: EVENT, label: 'Event' },
            { value: REVIEWER, label: 'Reviewer' },
        ];
        const hitContextOptions = [
            { value: '', label: 'Tous les hit contexts' },
            { value: ON_HIT, label: 'On Hit' },
            { value: ON_BLOCK, label: 'On Block' },
            { value: WHIFF, label: 'Whiff' },
            { value: EVENT_NOTICEABLE, label: 'Evénement marquant' },
        ];

        return {
            state: gameStore,
            filters,
            actionList,
            filteredActionList,
            formatTimestamp,
            hitContexts: {
                ON_HIT,
                ON_BLOCK,
                WHIFF,
                EVENT_NOTICEABLE
            },
            roundOptions,
            playerOptions,
            categoryOptions,
            typeOptions,
            hitContextOptions,
        }
    }
}
</script>