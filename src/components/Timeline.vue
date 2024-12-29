<template>
    <div class="timeline-panel">
        <h2>Timeline</h2>
        <div v-for="(round, index) in state.rounds" :key="index">
            <h3>Round {{ index + 1 }}</h3>
            <ul>
                <li v-for="action in round.actions" :key="action.id">
                    [{{ action.hitContext }}] {{ action.type }} ({{ action.player }}) - {{ action.timestamp }}
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
import { useGameStore } from '../store/gameStore';

export default {
    setup() {
        const gameStore = useGameStore();

        return {
            state: gameStore,
        };
    },
};
</script>

<style>
.timeline-panel {
    flex-grow: 1;
    padding: 1rem;
    overflow-y: auto;
}
</style>