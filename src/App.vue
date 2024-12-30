<template>
  <div class="app-container">
    <NavBar />
    <!-- <ModaleTest /> -->

    <div class="p-4 pt-16 sm:ml-64">
      <RouterView />
    </div>
  </div>
</template>

<script>
import { onMounted } from 'vue';
import {
  initModals,
  initDropdowns,
} from 'flowbite';

import NavBar from './components/NavBar.vue';
// import ModaleTest from './components/ModaleTest.vue';


export default {
  components: { NavBar },
  setup() {
    onMounted(() => {
      initModals();
      initDropdowns();
    });
    return {};
  }
};


</script>

<style>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
