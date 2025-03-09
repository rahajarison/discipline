<template>
    <h1 class="m-4 text-4xl font-bold text-gray-800 dark:text-white">Rapports</h1>
    <div class="relative overflow-x-auto">
        <ul>
            <li>Total actions: {{ flatActionList.length }}</li>
            <li>Total actions neutral: {{ stats.totalNeutralActions }}</li>
            <li>Total actions offense: {{ stats.totalOffenseActions }}</li>
            <li>Total actions défense: {{ stats.totalDefenseActions }}</li>
            <li>P1 actions neutral: {{ stats.P1.neutralActions }}</li>
            <li>P1 actions offense: {{ stats.P1.offenseActions }}</li>
            <li>P1 actions défense: {{ stats.P1.defenseActions }}</li>
            <li>P2 actions neutral: {{ stats.P2.neutralActions }}</li>
            <li>P2 actions offense: {{ stats.P2.offenseActions }}</li>
            <li>P2 actions défense: {{ stats.P2.defenseActions }}</li>
        </ul>
    </div>
</template>

<script>
import { useRoute } from 'vue-router';
import { useMatchStore } from '../store/matchStore';

import { ON_HIT, ON_BLOCK, WHIFF, TECHED, EVENT_NOTICEABLE } from '../utils/hitContexts';
import { SYSTEM, NEUTRAL, DEFENSE, OFFENSE } from '../utils/categories';
import { EVENT, REVIEWER } from '../utils/types';
import { P1 } from '../utils/players';


export default {
    setup() {
        const matchStore = useMatchStore();
        const route = useRoute();

        console.log(route.params);
        const match = matchStore.findMatch(route.params.matchId);
        const flatActionList = matchStore.getFlatActionList(match);

        const stats =  {
            P1: {
                neutralActions: 0,
                offenseActions: 0,
                defenseActions: 0,
            },
            P2: {
                neutralActions: 0,
                offenseActions: 0,
                defenseActions: 0,
            },
            totalNeutralActions: 0,
            totalOffenseActions: 0,
            totalDefenseActions: 0,
        }

        for (let i = 0; i < flatActionList.length; i++) {
            const action = flatActionList[i];

            if (action.type != EVENT && action.type != REVIEWER) {
                if (action.category == NEUTRAL) {
                    stats[action.player].neutralActions++;
                    stats.totalNeutralActions++;
                } else if (action.category == OFFENSE) {
                    stats[action.player].offenseActions++;
                    stats.totalOffenseActions++;
                } else if (action.category == DEFENSE) {
                    stats[action.player].defenseActions++;
                    stats.totalDefenseActions++;
                }
            }
        }

        return {
            state: matchStore,
            match,
            flatActionList,
            hitContexts: {
                ON_HIT,
                ON_BLOCK,
                WHIFF,
                TECHED,
                EVENT_NOTICEABLE
            },
            stats,
        }
    }
}
</script>