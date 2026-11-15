<script setup lang="ts">
import { useDeviceStore } from '../stores/device';


const deviceStore = useDeviceStore()
</script>

<template>
    <h1>This is the home view</h1>
    <h1 class="text-3xl font-bold underline">
        Hello world!
    </h1>
    <button
        class="inline-block cursor-pointer rounded-md bg-gray-800 px-4 py-3 text-center text-sm font-semibold uppercase text-white transition duration-200 ease-in-out hover:bg-gray-900">
        Button
    </button>
    <button class="btn btn-primary">Button</button>
    <p>{{ deviceStore.deviceMsg }}</p>
</template>