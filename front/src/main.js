import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue';

import { useAnalysisStore } from './store/analysisStore';

// Views for router
import DashboardView from './views/DashboardView.vue';
import AnalysisView from './views/AnalysisView.vue';
import ObjectivesView from './views/ObjectivesView.vue';
import ReportsMatchListView from './views/ReportsMatchListView.vue';
import ReportsMatchHistoryView from './views/ReportsMatchHistoryView.vue';

// Inclusion of Tailwind
import './assets/styles.css';


const routes = [
    { path: '/', name: "dashboard", component: DashboardView },
    { path: '/analysis', name: 'analysis', component: AnalysisView },
    { path: '/objectives', name: 'objectives', component: ObjectivesView },
    { path: '/matches', name: 'match-list', component: ReportsMatchListView },
    { path: '/matches/:matchId/history', name: 'match-history', component: ReportsMatchHistoryView },
]

const app = createApp(App);
const pinia = createPinia();
const router = createRouter({
    history: createWebHistory(),
    routes,
})

app.use(pinia);
app.use(router);
const analysisStore = useAnalysisStore();
analysisStore.initializeStore();

app.mount('#app');