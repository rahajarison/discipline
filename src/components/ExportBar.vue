<template>
    <div class="export-bar">
        <button @click="exportData('json')">Export JSON</button>
        <button @click="exportData('csv')">Export CSV</button>
    </div>
</template>

<script>
import { useGameStore } from '../store/gameStore';

export default {
    setup() {
        const gameStore = useGameStore();

        const exportData = (format) => {
            console.log(`Exporting data in ${format}`);
            // Ajouter logique d'export
        };

        return { exportData };
    },
};
</script>

<style>
.export-bar {
    position: fixed;
    bottom: 0;
    width: 100%;
    padding: 1rem;
    background: #f8f8f8;
    border-top: 1px solid #ccc;
    display: flex;
    justify-content: center;
}
</style>
