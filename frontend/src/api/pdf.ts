import { api } from './client.ts'
import type { ForgeRequest } from "./calculations.ts";

export async function downloadForgingPDF(params: ForgeRequest) {
    const response = await api.post('/forging/pdf', params, {
        responseType: 'blob',
    })

    const blob = new Blob([response.data], { type: 'application/pdf' })
    const url = window.URL.createObjectURL(blob)
    const link = document.createElement('a')
    link.href = url
    link.download = 'forging.pdf'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
}