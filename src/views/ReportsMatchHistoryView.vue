<template>
    <h1 class="m-4 text-4xl font-bold text-gray-800 dark:text-white">Rapports</h1>
    <!-- {{ actionList }} -->
    <!-- {{ filters }} -->
    
    <ReportsMatchHistoryFilterDropdown
        :options="roundOptions"
        placeholder="Filtrer par rounds"
        v-model="filters.round"
    />
    <ReportsMatchHistoryFilterDropdown
        :options="playerOptions"
        placeholder="Filtrer par joueur"
        v-model="filters.player"
    />
    <ReportsMatchHistoryFilterDropdown
        :options="categoryOptions"
        placeholder="Filtrer par catégorie"
        v-model="filters.category"
    />
    <ReportsMatchHistoryFilterDropdown
        :options="typeOptions"
        placeholder="Filtrer par type"
        v-model="filters.type"
    />
    <ReportsMatchHistoryFilterDropdown
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
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" v-for="action in filteredActionList" :key="action.id">
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
                        :class="{ 'bg-red-500 text-white': action.hitContext == hitContexts.ON_HIT, 'bg-gray-500 text-white': action.hitContext == hitContexts.ON_BLOCK, 'bg-red-800 text-white': action.hitContext == hitContexts.WHIFF, 'bg-gray-700 text-white': action.hitContext == hitContexts.TECHED }">
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
import { useRoute } from 'vue-router';
import { useMatchStore } from '../store/matchStore';

import ReportsMatchHistoryFilterDropdown from '../components/ReportsMatchHistoryFilterDropdown.vue';

import { formatTimestamp } from '../utils/formatTimestamp';
import { ON_HIT, ON_BLOCK, WHIFF, TECHED, EVENT_NOTICEABLE } from '../utils/hitContexts';
import { SYSTEM, NEUTRAL, DEFENSE, OFFENSE } from '../utils/categories';
import { EVENT, REVIEWER } from '../utils/types';
import ReportsHistoryFilterDropdown from '../components/ReportsMatchHistoryFilterDropdown.vue';


export default {
    components: { ReportsMatchHistoryFilterDropdown },
    setup() {
        const matchStore = useMatchStore();
        const route = useRoute();

        const filters = ref({
            round: '',
            player: '',
            category: '',
            type: '',
            hitContext: '',
        });

        // const actionList = computed(() => matchStore.flatActiveActionList);
        console.log(route.params);
        const match = matchStore.findMatch(route.params.matchId);
        // const flatActionList = computed(() => matchStore.getFlatActionList(match));
        const flatActionList = matchStore.getFlatActionList(match);
        const filteredActionList = computed(() => {
            return flatActionList.filter((action) => {
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
            ...match.rounds.map((_, index) => ({
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
            { value: TECHED, label: 'Teched' },
            { value: EVENT_NOTICEABLE, label: 'Evénement marquant' },
        ];

        return {
            state: matchStore,
            match,
            filters,
            filteredActionList,
            formatTimestamp,
            hitContexts: {
                ON_HIT,
                ON_BLOCK,
                WHIFF,
                TECHED,
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