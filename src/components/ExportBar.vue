<template>
    <div class="export-bar">
        <button class="btn-primary" @click="exportData('json')">Export JSON</button>
        <button class="btn-primary" @click="exportData('csv')">Export CSV</button>
    </div>
</template>

<script>
// import { useGameStore } from '../store/gameStore';

// export default {
//     setup() {
//         const gameStore = useGameStore();

//         const exportData = (format) => {
//             console.log(`Exporting data in ${format}`);
//             console.log(JSON.stringify(gameStore.rounds, null, 2))
//         };

//         return { exportData };
//     },
// };
</script>
