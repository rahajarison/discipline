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
            <p :class="{ 'text-blue-500': activePlayer === 'P1', 'text-red-500': activePlayer === 'P2' }">Active Player:
                {{ activePlayer }}</p>
            <button @click="togglePlayer" class="btn-primary">
                Switch Player
            </button>
        </div>
    </div>
</template>

<script>
import { computed, onMounted, onUnmounted } from 'vue';
import { useGameStore } from '../store/gameStore';

export default {
    setup() {
        const gameStore = useGameStore();
        const activePlayer = computed(() => gameStore.activePlayer)

        const handleKeydown = (event) => {
            switch (event.key) {
                case 'a':
                    gameStore.addAction({ type: 'Jump In' });
                    break;
                case 'z':
                    gameStore.addAction({ type: 'Dash In' });
                    break;
                case 'Enter':
                    gameStore.markRoundWinner(gameStore.activePlayer);
                    break;
                case 'Tab':
                    event.preventDefault();
                    gameStore.togglePlayer();
                    break;
            }
        };

        onMounted(() => {
            window.addEventListener('keydown', handleKeydown);
        });

        onUnmounted(() => {
            window.removeEventListener('keydown', handleKeydown);
        });

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