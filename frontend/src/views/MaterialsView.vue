<template>
  <div class="space-y-8">
    <h2 class="text-2xl font-semibold">Матеріали та форми заготовок</h2>

    <!-- Матеріали -->
    <div class="bg-white p-6 rounded shadow space-y-4">
      <h3 class="text-lg font-semibold border-b pb-2">Матеріали</h3>

      <table class="w-full text-sm">
        <thead>
        <tr class="text-left text-gray-500 border-b">
          <th class="pb-2">Назва</th>
          <th class="pb-2">Ключ</th>
          <th class="pb-2">Густина (кг/м³)</th>
          <th class="pb-2">Межа плинності (МПа)</th>
          <th class="pb-2"></th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="m in store.materials" :key="m.id" class="border-b hover:bg-slate-50">
          <td class="py-2">{{ m.name }}</td>
          <td class="py-2 text-gray-500">{{ m.key }}</td>
          <td class="py-2">{{ m.density }}</td>
          <td class="py-2">{{ m.yield_strength }}</td>
          <td class="py-2 text-right">
            <button @click="removeMaterial(m.id)"
                    class="text-red-500 hover:text-red-700 text-xs px-2 py-1 rounded hover:bg-red-50">
              Видалити
            </button>
          </td>
        </tr>
        <tr v-if="store.materials.length === 0">
          <td colspan="5" class="py-4 text-center text-gray-400">Немає матеріалів</td>
        </tr>
        </tbody>
      </table>

      <!-- Форма додавання -->
      <form @submit.prevent="submitMaterial" class="grid grid-cols-2 md:grid-cols-4 gap-3 pt-2">
        <div>
          <label class="block text-xs text-gray-500 mb-1">Назва (укр.)</label>
          <input v-model="newMaterial.name" placeholder="Сталь 45"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div>
          <label class="block text-xs text-gray-500 mb-1">Ключ (англ.)</label>
          <input v-model="newMaterial.key" placeholder="steel_45"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div>
          <label class="block text-xs text-gray-500 mb-1">Густина (кг/м³)</label>
          <input v-model.number="newMaterial.density" type="number" min="0" step="0.1"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div>
          <label class="block text-xs text-gray-500 mb-1">Межа плинності (МПа)</label>
          <input v-model.number="newMaterial.yield_strength" type="number" min="0" step="1"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div class="col-span-2 md:col-span-4">
          <button type="submit"
                  class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            + Додати матеріал
          </button>
        </div>
      </form>
    </div>

    <!-- Форми заготовок -->
    <div class="bg-white p-6 rounded shadow space-y-4">
      <h3 class="text-lg font-semibold border-b pb-2">Форми заготовок</h3>

      <table class="w-full text-sm">
        <thead>
        <tr class="text-left text-gray-500 border-b">
          <th class="pb-2">Назва</th>
          <th class="pb-2">Ключ</th>
          <th class="pb-2"></th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="s in store.shapes" :key="s.id" class="border-b hover:bg-slate-50">
          <td class="py-2">{{ s.name }}</td>
          <td class="py-2 text-gray-500">{{ s.key }}</td>
          <td class="py-2 text-right">
            <button @click="removeShape(s.id)"
                    class="text-red-500 hover:text-red-700 text-xs px-2 py-1 rounded hover:bg-red-50">
              Видалити
            </button>
          </td>
        </tr>
        <tr v-if="store.shapes.length === 0">
          <td colspan="3" class="py-4 text-center text-gray-400">Немає форм</td>
        </tr>
        </tbody>
      </table>

      <form @submit.prevent="submitShape" class="flex gap-3 pt-2">
        <div class="flex-1">
          <label class="block text-xs text-gray-500 mb-1">Назва (укр.)</label>
          <input v-model="newShape.name" placeholder="Шестигранник"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div class="flex-1">
          <label class="block text-xs text-gray-500 mb-1">Ключ (англ.)</label>
          <input v-model="newShape.key" placeholder="hexagon"
                 class="w-full border rounded px-3 py-2 text-sm" required />
        </div>
        <div class="flex items-end">
          <button type="submit"
                  class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 text-sm">
            + Додати форму
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, onMounted } from 'vue'
import { useMaterialsStore } from '../stores/materials'

const store = useMaterialsStore()

const newMaterial = reactive({ name: '', key: '', density: 0, yield_strength: 0 })
const newShape = reactive({ name: '', key: '' })

onMounted(() => store.fetchAll())

async function submitMaterial() {
  await store.createMaterial({ ...newMaterial })
  Object.assign(newMaterial, { name: '', key: '', density: 0, yield_strength: 0 })
}

async function removeMaterial(id: string) {
  if (confirm('Видалити матеріал?')) await store.removeMaterial(id)
}

async function submitShape() {
  await store.createShape({ ...newShape })
  Object.assign(newShape, { name: '', key: '' })
}

async function removeShape(id: string) {
  if (confirm('Видалити форму?')) await store.removeShape(id)
}
</script>