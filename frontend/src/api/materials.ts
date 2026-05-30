import { api } from './client'

export interface CustomMaterial {
    id: string
    name: string
    key: string
    density: number
    yield_strength: number
}

export interface CustomShape {
    id: string
    name: string
    key: string
}

export async function getMaterials(): Promise<CustomMaterial[]> {
    const { data } = await api.get<CustomMaterial[]>('/materials')
    return data
}

export async function addMaterial(m: Omit<CustomMaterial, 'id'>): Promise<CustomMaterial> {
    const { data } = await api.post<CustomMaterial>('/materials', m)
    return data
}

export async function deleteMaterial(id: string): Promise<void> {
    await api.delete(`/materials/${id}`)
}

export async function getShapes(): Promise<CustomShape[]> {
    const { data } = await api.get<CustomShape[]>('/shapes')
    return data
}

export async function addShape(s: Omit<CustomShape, 'id'>): Promise<CustomShape> {
    const { data } = await api.post<CustomShape>('/shapes', s)
    return data
}

export async function deleteShape(id: string): Promise<void> {
    await api.delete(`/shapes/${id}`)
}