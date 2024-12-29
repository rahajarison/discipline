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
import { keyMap, getHitContext } from '../utils/keyMap';

export default {
    setup() {
        const gameStore = useGameStore();
        const activePlayer = computed(() => gameStore.activePlayer);


        const handleKeydown = (event) => {
            const keyAction = keyMap[event.key];
            if (!keyAction) return;
            console.log("Key action detected:", keyAction);

            const hitContext = keyAction.hitContexts ? getHitContext(event, keyAction.hitContexts) : null;
            console.log("Determined hitContext:", hitContext);

            if (keyAction.action) {
                gameStore[keyAction.action]();
            } else {
                gameStore.addAction({
                    type: keyAction.type,
                    hitContext,
                });
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