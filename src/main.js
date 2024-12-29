import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import { useGameStore } from './store/gameStore';

import './assets/styles.css'; // Inclut Tailwind

const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
const gameStore = useGameStore();
gameStore.initializeStore();

app.mount('#app');