<template>
    <div class="w-1-3">
        <h1 class="">Action Input</h1>
        <ul>
            <li><span>A:</span> Jump In</li>
            <li><span>Z:</span> Dash In</li>
            <li><span>Enter:</span> Gain Round</li>
        </ul>
        <div class="controls">
            <button @click="undo">Undo</button>
            <button @click="redo">Redo</button>
        </div>
        <div class="active-player">
            <p>Active Player: {{ state.activePlayer }}</p>
            <button @click="togglePlayer"
                class="btn-primary">Switch
                Player</button>
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
            activePlayer: gameStore.activePlayer,
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