<template>
    <div class="p-4 bg-gray-50 border border-gray-200 rounded-lg">
        <h2 class="text-xl font-bold text-gray-700 dark:text-white">Informations</h2>
        <div class="w-1-3">
            <div class="active-player">
                <p :class="{ 'text-blue-600': activePlayer === 'P1', 'text-red-600': activePlayer === 'P2' }">
                    Joueur actif:
                    <span
                        :class="{ 'border-blue-400': activePlayer == 'P1', 'border-red-400': activePlayer == 'P2'}"
                        class="border border-dashed rounded-lg px-2 text-xs font-semibold">
                            {{ activePlayer }}
                    </span>
                </p>
                <button @click="togglePlayer" class="btn-primary">
                    Switch Player
                </button>
            </div>
            <div class="controls">
                <button @click="undo" :disabled="!state.canUndo" class="btn-primary">Undo</button>
                <button @click="redo" :disabled="!state.canRedo" class="btn-primary">Redo</button>
            </div>
        </div>
    </div>
</template>

<script>
import { computed, onMounted, onUnmounted } from 'vue';
import { useAnalysisStore } from '../store/analysisStore';
import { keyMap, getHitContext } from '../utils/keymap';

export default {
    setup() {
        const analysisStore = useAnalysisStore();
        const activePlayer = computed(() => analysisStore.activePlayer);

        const handleKeydown = (event) => {
            const keyAction = keyMap[event.key];
            if (!keyAction) return;

            // Block keys that are mapped
            //  Prevent default behaviors such as select all the text with ctrl A
            const blockedKeys = Object.keys(keyMap);
            if ((event.ctrlKey || event.altKey || event.metaKey) && blockedKeys.includes(event.key)) {
                event.preventDefault();
            } else if (event.key == "Tab") {
                event.preventDefault();
            }

            console.log("Key action detected:", keyAction);

            const hitContext = keyAction.hitContexts ? getHitContext(event, keyAction.hitContexts) : null;
            console.log("Determined hitContext:", hitContext);

            if (keyAction.action) {
                analysisStore[keyAction.action]();
            } else {
                analysisStore.addAction({
                    type: keyAction.type,
                    hitContext,
                    category: keyAction.category
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
            state: analysisStore,
            activePlayer,
            togglePlayer: analysisStore.togglePlayer,
            undo: analysisStore.undo,
            redo: analysisStore.redo,
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