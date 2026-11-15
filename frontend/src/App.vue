<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import Navigation from './components/Navigation.vue';
import { useDeviceStore } from './stores/device.js';


const deviceStore = useDeviceStore()

onMounted(() => {
  deviceStore.initListeners()
})
</script>

<template>
  <Navigation />
  <RouterView />
</template>
