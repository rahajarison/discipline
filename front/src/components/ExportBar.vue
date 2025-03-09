<template>
    <div class="export-bar">
        <!-- <button class="btn-primary" @click="exportData('json')">Export JSON</button> -->
        <button class="btn-primary" @click="downloadMatch(flatActionList)">Export CSV</button>
    </div>
</template>

<script>
import { useMatchStore } from '../store/matchStore';
import { useRoute } from 'vue-router';

export default {
    setup() {
        const matchStore = useMatchStore();
        const route = useRoute();
        const matchId = route.params.matchId;

        const match = matchStore.findMatch(matchId);
        const flatActionList = matchStore.getFlatActionList(match);

        const csvmaker = function (data) {
            if (!data || !data.length) {
                console.error('Aucune donnée à exporter.');
                return;
            }

            // Obtenir les clés pour les en-têtes du CSV
            const headers = Object.keys(data[0]);

            // Créer les lignes CSV
            const csvContent = [
                headers.join(','), // Première ligne : en-têtes
                ...data.map((row) => headers.map((header) => JSON.stringify(row[header] || '')).join(',')) // Contenu
            ].join('\n');

            return csvContent;
        }

        const downloadFile = function (data) {
            const blob = new Blob([data], { type: 'text/csv;charset=utf-8;' });
            const url = URL.createObjectURL(blob);
            const a = document.createElement('a');
            
            // Set the URL and download attribute of the anchor tag
            a.href = url;
            a.download = `discipline_match_${matchId}.csv`;
            
            a.style.visibility = 'hidden';
            // Trigger the download by clicking the anchor tag
            a.click();
        }

        const downloadMatch = function (data) {
            console.log(JSON.stringify(data));
            return downloadFile(csvmaker(data));
        }

        return {
            flatActionList,
            downloadMatch
        };
    },
};
</script>
