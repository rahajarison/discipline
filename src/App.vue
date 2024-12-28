<template>
  <div class="app-container">
    <InputPanel />
    <Timeline />
    <ExportBar />
  </div>
</template>

<script>
import InputPanel from './components/InputPanel.vue';
import Timeline from './components/Timeline.vue';
import ExportBar from './components/ExportBar.vue';

export default {
  components: { InputPanel, Timeline, ExportBar },
};
</script>

<style>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
