import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createMemoryHistory, createRouter } from 'vue-router'

import App from './App.vue';

import { useGameStore } from './store/gameStore';

// Views for router
import DashboardView from './views/DashboardView.vue';
import AnalysisView from './views/AnalysisView.vue';
import ObjectivesView from './views/ObjectivesView.vue';
import ReportsView from './views/ReportsView.vue';

// Inclusion of Tailwind
import './assets/styles.css';


const routes = [
    { path: '/', component: DashboardView },
    { path: '/analysis', component: AnalysisView },
    { path: '/objectives', component: ObjectivesView },
    { path: '/reports', component: ReportsView },
]


const app = createApp(App);
const pinia = createPinia();
const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

app.use(pinia);
app.use(router);
const gameStore = useGameStore();
gameStore.initializeStore();

app.mount('#app');