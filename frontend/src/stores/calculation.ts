import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { ForgeResponse } from '../api/calculations'

export const useCalcStore = defineStore('calculation', () => {
    const result = ref<ForgeResponse | null>(null)
    const loading = ref(false)
    const error = ref('')

    function setResult(data: ForgeResponse) {
        result.value = data
        error.value = ''
    }

    function setError(msg: string) {
        error.value = msg
        result.value = null
    }

    return { result, loading, error, setResult, setError }
})