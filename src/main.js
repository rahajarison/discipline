import { createApp } from 'vue';
import { createPinia } from 'pinia';
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue';

import { useAnalysisStore } from './store/analysisStore';

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
    history: createWebHistory(),
    routes,
})

app.use(pinia);
app.use(router);
const analysisStore = useAnalysisStore();
analysisStore.initializeStore();

app.mount('#app');