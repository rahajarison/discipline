<template>
    <div class="w-1-3">
        <h1>Action Input</h1>
        <ul>
            <li><span>A:</span> Jump In</li>
            <li><span>Z:</span> Dash In</li>
            <li><span>Enter:</span> Gain Round</li>
        </ul>
        <div class="controls">
            <button @click="undo" class="btn-primary">Undo</button>
            <button @click="redo" class="btn-primary">Redo</button>
        </div>
        <div class="active-player">
            <p>Active Player: {{ activePlayer }}</p>
            <button  @click="togglePlayer" class="btn-primary">
                Switch Player
            </button>
        </div>
    </div>
</template>

<script>
import { computed } from 'vue';
import { useGameStore } from '../store/gameStore';

export default {
    setup() {
        const gameStore = useGameStore();

        const activePlayer = computed(() => gameStore.activePlayer)

        return {
            state: gameStore,
            activePlayer,
            togglePlayer: gameStore.togglePlayer,
            undo: gameStore.undo,
            redo: gameStore.redo,
        };
    },
};
</script>

<style>
/* .input-panel {
    width: 30%;
    border-right: 1px solid #ccc;
    padding: 1rem;
} */
</style>