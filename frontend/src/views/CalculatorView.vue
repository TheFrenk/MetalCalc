<template>
  <div class="space-y-6">
    <h2 class="text-2xl font-semibold">Расчет параметров объемной штамповки / ковки</h2>

    <form @submit.prevent="onSubmit" class="bg-white p-6 rounded shadow space-y-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label class="block text-sm font-medium">Материал</label>
          <select v-model="form.material" class="mt-1 w-full border rounded px-3 py-2">
            <option value="steel">Сталь</option>
            <option value="aluminum">Алюминий</option>
            <option value="copper">Медь</option>
            <option value="titanium">Титан</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium">Форма заготовки</label>
          <select v-model="form.shape" class="mt-1 w-full border rounded px-3 py-2">
            <option value="cylinder">Цилиндр</option>
            <option value="rectangle">Прямоугольник</option>
          </select>
        </div>
        <div>
          <label class="block text-sm font-medium">
            {{ form.shape === 'cylinder' ? 'Диаметр D (мм)' : 'Сторона A (мм)' }}
          </label>
          <input v-model.number="form.dimension_a" type="number" step="1" min="1" class="mt-1 w-full border rounded px-3 py-2" />
        </div>
        <div v-if="form.shape === 'rectangle'">
          <label class="block text-sm font-medium">Сторона B (мм)</label>
          <input v-model.number="form.dimension_b" type="number" step="1" min="1" class="mt-1 w-full border rounded px-3 py-2" />
        </div>
        <div>
          <label class="block text-sm font-medium">Начальная высота H₀ (мм)</label>
          <input v-model.number="form.initial_height" type="number" step="1" min="1" class="mt-1 w-full border rounded px-3 py-2" />
        </div>
        <div>
          <label class="block text-sm font-medium">Конечная высота H₁ (мм)</label>
          <input v-model.number="form.final_height" type="number" step="0.1" min="1" class="mt-1 w-full border rounded px-3 py-2" />
        </div>
        <div>
          <label class="block text-sm font-medium">Температура (°C)</label>
          <input v-model.number="form.temperature" type="number" step="10" min="20" class="mt-1 w-full border rounded px-3 py-2" />
        </div>
        <div>
          <label class="block text-sm font-medium">Коэффициент трения μ</label>
          <input v-model.number="form.friction_coeff" type="number" step="0.05" min="0.05" max="0.8" class="mt-1 w-full border rounded px-3 py-2" />
          <p class="text-xs text-gray-500 mt-1">Горячая ковка: 0.3–0.5, холодная: 0.1–0.2</p>
        </div>
        <div class="md:col-span-2">
          <label class="block text-sm font-medium">Скорость деформации (мм/с)</label>
          <input v-model.number="form.deformation_speed" type="number" step="1" min="1" class="mt-1 w-full border rounded px-3 py-2" />
          <p class="text-xs text-gray-500 mt-1">Скорость движения пуансона / молота</p>
        </div>
      </div>
      <div class="flex gap-3">
        <button type="submit" :disabled="store.loading" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50">
          {{ store.loading ? 'Расчет...' : 'Рассчитать' }}
        </button>
        <button type="button" @click="onDownloadPDF" :disabled="!store.result || pdfLoading" class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50">
          {{ pdfLoading ? 'Генерация...' : 'Скачать PDF' }}
        </button>
      </div>
    </form>

    <div v-if="store.error" class="bg-red-50 text-red-700 p-4 rounded border border-red-200">
      {{ store.error }}
    </div>

    <div v-if="store.result" class="bg-white p-6 rounded shadow space-y-4">
      <h3 class="text-lg font-semibold border-b pb-2">Результаты расчета</h3>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Усилие ковки</p>
          <p class="text-lg font-bold">{{ fmt(store.result.forging_force) }} Н</p>
          <p class="text-xs text-gray-500">≈ {{ (store.result.forging_force / 1000).toFixed(1) }} кН</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Удельное давление</p>
          <p class="text-lg font-bold">{{ fmt(store.result.forging_pressure) }} МПа</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Работа деформации</p>
          <p class="text-lg font-bold">{{ fmt(store.result.work_done) }} Дж</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Мощность</p>
          <p class="text-lg font-bold">{{ fmt(store.result.power) }} кВт</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Скорость деформации ε̇</p>
          <p class="text-lg font-bold">{{ fmt(store.result.deformation_speed) }} с⁻¹</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Степень деформации</p>
          <p class="text-lg font-bold">{{ fmt(store.result.strain_degree) }}</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Масса заготовки</p>
          <p class="text-lg font-bold">{{ fmt(store.result.workpiece_mass) }} кг</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Площадь контакта</p>
          <p class="text-lg font-bold">{{ fmt(store.result.contact_area) }} мм²</p>
        </div>
      </div>

      <div class="border-t pt-3 text-sm space-y-1">
        <p><strong>Объем заготовки:</strong> {{ fmt(store.result.initial_volume) }} мм³</p>
        <p><strong>Обжатие по высоте:</strong> {{ fmt(store.result.height_reduction) }} мм</p>
        <p v-if="form.shape === 'cylinder'">
          <strong>Конечный диаметр:</strong> {{ fmt(store.result.final_diameter) }} мм
        </p>
        <p v-else>
          <strong>Конечные размеры:</strong> {{ fmt(store.result.final_side_a) }} × {{ fmt(store.result.final_side_b) }} мм
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {reactive, ref} from 'vue'
import { calculateForging, type ForgeRequest } from '../api/calculations'
import { useCalcStore } from '../stores/calculation'
import { downloadForgingPDF } from "../api/pdf.ts";

const store = useCalcStore()
const pdfLoading = ref(false)

const form = reactive<ForgeRequest>({
  material: 'steel',
  shape: 'cylinder',
  dimension_a: 80,
  dimension_b: 0,
  initial_height: 120,
  final_height: 60,
  temperature: 1100,
  friction_coeff: 0.4,
  deformation_speed: 50,
})

function fmt(n: number): string {
  return n.toFixed(2)
}

async function onSubmit() {
  store.loading = true
  store.error = ''
  try {
    const res = await calculateForging(form)
    store.setResult(res)
  } catch (e: any) {
    store.setError(e.response?.data?.error || 'Ошибка сервера')
  } finally {
    store.loading = false
  }
}

async function onDownloadPDF() {
  pdfLoading.value = true
  try {
    await downloadForgingPDF(form)
  } catch (e: any) {
    store.setError(e.response?.data?.error || 'Ошибка генерации PDF')
  } finally {
    pdfLoading.value = false
  }
}
</script>