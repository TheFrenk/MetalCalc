import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { CustomMaterial, CustomShape } from '../api/materials'
import {
    getMaterials, addMaterial, deleteMaterial,
    getShapes, addShape, deleteShape,
} from '../api/materials'

export const useMaterialsStore = defineStore('materials', () => {
    const materials = ref<CustomMaterial[]>([])
    const shapes = ref<CustomShape[]>([])

    async function fetchAll() {
        materials.value = await getMaterials()
        shapes.value = await getShapes()
    }

    async function createMaterial(m: Omit<CustomMaterial, 'id'>) {
        const created = await addMaterial(m)
        materials.value.push(created)
    }

    async function removeMaterial(id: string) {
        await deleteMaterial(id)
        materials.value = materials.value.filter(m => m.id !== id)
    }

    async function createShape(s: Omit<CustomShape, 'id'>) {
        const created = await addShape(s)
        shapes.value.push(created)
    }

    async function removeShape(id: string) {
        await deleteShape(id)
        shapes.value = shapes.value.filter(s => s.id !== id)
    }

    return { materials, shapes, fetchAll, createMaterial, removeMaterial, createShape, removeShape }
})