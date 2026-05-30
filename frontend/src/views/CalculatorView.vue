<template>
  <div class="space-y-6">
    <h2 class="text-2xl font-semibold">Розрахунок параметрів об'ємного кування</h2>

    <form @submit.prevent="onSubmit" class="bg-white p-6 rounded shadow space-y-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">

        <div>
          <label class="block text-sm font-medium">Матеріал</label>
          <select v-model="form.material" class="mt-1 w-full border rounded px-3 py-2">
            <option v-for="m in materialsStore.materials" :key="m.id" :value="m.key">
              {{ m.name }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium">Форма заготівлі</label>
          <select v-model="form.shape" class="mt-1 w-full border rounded px-3 py-2">
            <option v-for="s in materialsStore.shapes" :key="s.id" :value="s.key">
              {{ s.name }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-sm font-medium">
            {{ form.shape === 'cylinder' ? 'Діаметр D (мм)' : 'Сторона A (мм)' }}
          </label>
          <input v-model.number="form.dimension_a" type="number" step="1" min="1"
                 class="mt-1 w-full border rounded px-3 py-2" />
        </div>

        <div v-if="form.shape !== 'cylinder'">
          <label class="block text-sm font-medium">Сторона B (мм)</label>
          <input v-model.number="form.dimension_b" type="number" step="1" min="1"
                 class="mt-1 w-full border rounded px-3 py-2" />
        </div>

        <div>
          <label class="block text-sm font-medium">Початкова висота H₀ (мм)</label>
          <input v-model.number="form.initial_height" type="number" step="1" min="1"
                 class="mt-1 w-full border rounded px-3 py-2" />
        </div>

        <div>
          <label class="block text-sm font-medium">Кінцева висота H₁ (мм)</label>
          <input v-model.number="form.final_height" type="number" step="0.1" min="0.1"
                 class="mt-1 w-full border rounded px-3 py-2" />
        </div>

        <div>
          <label class="block text-sm font-medium">Температура (°C)</label>
          <input v-model.number="form.temperature" type="number" step="10" min="20"
                 class="mt-1 w-full border rounded px-3 py-2" />
        </div>

        <div>
          <label class="block text-sm font-medium">Коефіцієнт тертя μ</label>
          <input v-model.number="form.friction_coeff" type="number" step="0.05" min="0.05" max="0.8"
                 class="mt-1 w-full border rounded px-3 py-2" />
          <p class="text-xs text-gray-500 mt-1">Гаряче кування: 0.3–0.5, холодна: 0.1–0.2</p>
        </div>

        <div>
          <label class="block text-sm font-medium">Швидкість деформації (мм/с)</label>
          <input v-model.number="form.deformation_speed" type="number" step="1" min="1"
                 class="mt-1 w-full border rounded px-3 py-2" />
          <p class="text-xs text-gray-500 mt-1">Швидкість руху пуансону / молота</p>
        </div>

        <div>
          <label class="block text-sm font-medium">Кількість проходів</label>
          <input v-model.number="form.passes" type="number" step="1" min="1" max="20"
                 class="mt-1 w-full border rounded px-3 py-2" />
          <p class="text-xs text-gray-500 mt-1">
            Висота рівномірно розбивається на вказану кількість проходів
          </p>
        </div>

      </div>

      <div class="flex gap-3">
        <button type="submit" :disabled="store.loading"
                class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50">
          {{ store.loading ? 'Розрахунок...' : 'Розрахувати' }}
        </button>
        <button type="button" @click="onDownloadPDF" :disabled="!store.result || pdfLoading"
                class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 disabled:opacity-50">
          {{ pdfLoading ? 'Генерація...' : 'Завантажити PDF' }}
        </button>
      </div>
    </form>

    <div v-if="store.error" class="bg-red-50 text-red-700 p-4 rounded border border-red-200">
      {{ store.error }}
    </div>

    <div v-if="store.result" class="bg-white p-6 rounded shadow space-y-4">
      <div class="flex items-center justify-between border-b pb-2">
        <h3 class="text-lg font-semibold">Результати розрахунку</h3>
        <span class="text-sm text-gray-500 bg-slate-100 px-3 py-1 rounded-full">
          Проходів: {{ store.result.passes }}
        </span>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Зусилля кування (середнє)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.forging_force) }} Н</p>
          <p class="text-xs text-gray-500">≈ {{ (store.result.forging_force / 1000).toFixed(1) }} кН</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Питомий тиск (середній)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.forging_pressure) }} МПа</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Робота деформації (сумарна)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.work_done) }} Дж</p>
          <p class="text-xs text-gray-500">
            ≈ {{ (store.result.work_done / form.passes).toFixed(2) }} Дж / прохід
          </p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Потужність (сумарна)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.power) }} кВт</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Швидкість деформації ε̇ (середня)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.deformation_speed) }} с⁻¹</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Ступінь деформації (сумарний)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.strain_degree) }}</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Маса заготівлі</p>
          <p class="text-lg font-bold">{{ fmt(store.result.workpiece_mass) }} кг</p>
        </div>
        <div class="bg-slate-50 p-3 rounded">
          <p class="text-gray-600">Площа контакту (середня)</p>
          <p class="text-lg font-bold">{{ fmt(store.result.contact_area) }} мм²</p>
        </div>
      </div>

      <div class="border-t pt-3 text-sm space-y-1">
        <p><strong>Об'єм заготівлі:</strong> {{ fmt(store.result.initial_volume) }} мм³</p>
        <p><strong>Обтискання по висоті:</strong> {{ fmt(store.result.height_reduction) }} мм</p>
        <p v-if="form.shape === 'cylinder'">
          <strong>Кінцевий діаметр:</strong> {{ fmt(store.result.final_diameter) }} мм
        </p>
        <p v-else>
          <strong>Кінцеві розміри:</strong>
          {{ fmt(store.result.final_side_a) }} × {{ fmt(store.result.final_side_b) }} мм
        </p>
      </div>
    </div>
    <ForgingCharts
        v-if="store.result && store.result.pass_results?.length"
        :pass-results="store.result.pass_results"
        :passes="store.result.passes"
    />

  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import { calculateForging, type ForgeRequest } from '../api/calculations'
import { useCalcStore } from '../stores/calculation'
import { downloadForgingPDF } from '../api/pdf'
import { useMaterialsStore } from '../stores/materials'
import ForgingCharts from "./ForgingCharts.vue";

const store = useCalcStore()
const materialsStore = useMaterialsStore()
const pdfLoading = ref(false)

onMounted(() => materialsStore.fetchAll())

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
  passes: 1,
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
    store.setError(e.response?.data?.error || 'Помилка сервера')
  } finally {
    store.loading = false
  }
}

async function onDownloadPDF() {
  pdfLoading.value = true
  try {
    await downloadForgingPDF(form)
  } catch (e: any) {
    store.setError(e.response?.data?.error || 'Помилка генерації PDF')
  } finally {
    pdfLoading.value = false
  }
}
</script>