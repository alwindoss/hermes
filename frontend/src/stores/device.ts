import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { Events } from '@wailsio/runtime'
import { DeviceMessageReceived } from '../data/messages';

export const useDeviceStore = defineStore('device', () => {
    const deviceMsg = ref("");

    function initListeners() {
        Events.On(DeviceMessageReceived, (event) => {
        console.log("Event received:", event.data)
        deviceMsg.value = event.data
    })
    }


    return { deviceMsg, initListeners }
})