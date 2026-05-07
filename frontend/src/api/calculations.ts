import { api } from './client'

export interface ForgeRequest {
    material: string
    shape: string
    dimension_a: number
    dimension_b: number
    initial_height: number
    final_height: number
    temperature: number
    friction_coeff: number
    deformation_speed: number
}

export interface ForgeResponse {
    forging_force: number
    forging_pressure: number
    work_done: number
    power: number
    deformation_speed: number
    strain_degree: number
    initial_volume: number
    final_volume: number
    workpiece_mass: number
    contact_area: number
    final_diameter: number
    final_side_a: number
    final_side_b: number
    height_reduction: number
}

export async function calculateForging(params: ForgeRequest): Promise<ForgeResponse> {
    const { data } = await api.post<ForgeResponse>('/forging', params)
    return data
}