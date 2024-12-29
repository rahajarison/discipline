<template>
  <div class="app-container">
    <KeyMapDisplay />
    <InputPanel />
    <Timeline />
    <ExportBar />
  </div>
</template>

<script>
import InputPanel from './components/InputPanel.vue';
import Timeline from './components/Timeline.vue';
import ExportBar from './components/ExportBar.vue';
import KeyMapDisplay from './components/KeyMapDisplay.vue';

export default {
  components: { InputPanel, Timeline, ExportBar, KeyMapDisplay },
};
</script>

<style>
.app-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
}
</style>
