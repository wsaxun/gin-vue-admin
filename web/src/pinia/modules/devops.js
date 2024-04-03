import { defineStore } from 'pinia'
import { ref } from 'vue'

export const piniaTools = defineStore("devops",() => {
    const toolTypeID = ref(0)
    function setToolTypeID (id) {
        toolTypeID.value = id
    }
    return {
        toolTypeID,
        setToolTypeID,
    }
})