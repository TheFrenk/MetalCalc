<template>
  <div class="bg-white p-6 rounded shadow space-y-8">
    <h3 class="text-lg font-semibold border-b pb-2">Графіки</h3>

    <div>
      <h4 class="text-sm font-medium text-gray-600 mb-3">Зусилля кування vs Висота заготовки</h4>
      <div class="relative h-64">
        <svg class="w-full h-full" :viewBox="`0 0 ${W} ${H}`" preserveAspectRatio="xMidYMid meet">
          <line v-for="i in 5" :key="'hy'+i"
                :x1="PAD" :y1="PAD + (i-1) * innerH / 4"
                :x2="W - PAD" :y2="PAD + (i-1) * innerH / 4"
                stroke="#e5e7eb" stroke-width="1" />
          <line v-for="i in 5" :key="'vx'+i"
                :x1="PAD + (i-1) * innerW / 4" :y1="PAD"
                :x2="PAD + (i-1) * innerW / 4" :y2="H - PAD"
                stroke="#e5e7eb" stroke-width="1" />

          <polyline
              :points="forceHeightPoints"
              fill="none" stroke="#2563eb" stroke-width="2.5"
              stroke-linejoin="round" stroke-linecap="round" />

          <circle v-for="(pt, i) in forceHeightData" :key="i"
                  :cx="pt.cx" :cy="pt.cy" r="4"
                  fill="white" stroke="#2563eb" stroke-width="2" >
            <title>H={{ pt.h.toFixed(1) }} мм, F={{ (pt.f/1000).toFixed(1) }} кН</title>
          </circle>

          <line :x1="PAD" :y1="PAD" :x2="PAD" :y2="H-PAD" stroke="#6b7280" stroke-width="1.5"/>
          <line :x1="PAD" :y1="H-PAD" :x2="W-PAD" :y2="H-PAD" stroke="#6b7280" stroke-width="1.5"/>

          <text v-for="(pt, i) in forceHeightData" :key="'lx'+i"
                :x="pt.cx" :y="H-PAD+16" text-anchor="middle"
                font-size="10" fill="#6b7280">
            {{ pt.h.toFixed(0) }}
          </text>

          <text v-for="i in 5" :key="'ly'+i"
                :x="PAD-6" :y="PAD + (4-(i-1)) * innerH / 4 + 4"
                text-anchor="end" font-size="10" fill="#6b7280">
            {{ yForceLabel(i-1) }}
          </text>

          <text :x="W/2" :y="H-2" text-anchor="middle" font-size="11" fill="#374151">
            Висота (мм)
          </text>
          <text :x="12" :y="H/2" text-anchor="middle" font-size="11" fill="#374151"
                :transform="`rotate(-90, 12, ${H/2})`">
            Зусилля (кН)
          </text>
        </svg>
      </div>
    </div>

    <div v-if="passes > 1">
      <h4 class="text-sm font-medium text-gray-600 mb-3">Зусилля по проходах</h4>
      <div class="relative h-64">
        <svg class="w-full h-full" :viewBox="`0 0 ${W} ${H}`" preserveAspectRatio="xMidYMid meet">
          <line v-for="i in 5" :key="'bgy'+i"
                :x1="PAD" :y1="PAD + (i-1) * innerH / 4"
                :x2="W - PAD" :y2="PAD + (i-1) * innerH / 4"
                stroke="#e5e7eb" stroke-width="1" />

          <g v-for="(bar, i) in barData" :key="i">
            <rect
                :x="bar.x" :y="bar.y"
                :width="bar.w" :height="bar.barH"
                fill="#2563eb" rx="3" opacity="0.85">
              <title>Прохід {{ bar.pass }}: {{ (bar.force/1000).toFixed(1) }} кН</title>
            </rect>
            <text :x="bar.x + bar.w/2" :y="bar.y - 4"
                  text-anchor="middle" font-size="10" fill="#1d4ed8" font-weight="600">
              {{ (bar.force/1000).toFixed(0) }}кН
            </text>
            <text
                :x="bar.x + bar.w/2"
                :y="H - PAD + 8"
                text-anchor="end"
                font-size="10"
                fill="#374151"
                :transform="`rotate(-45, ${bar.x + bar.w/2}, ${H - PAD + 8})`">
              Прохід {{ bar.pass }}
            </text>
          </g>

          <line :x1="PAD" :y1="PAD" :x2="PAD" :y2="H-PAD" stroke="#6b7280" stroke-width="1.5"/>
          <line :x1="PAD" :y1="H-PAD" :x2="W-PAD" :y2="H-PAD" stroke="#6b7280" stroke-width="1.5"/>

          <text v-for="i in 5" :key="'bly'+i"
                :x="PAD-6" :y="PAD + (4-(i-1)) * innerH / 4 + 4"
                text-anchor="end" font-size="10" fill="#6b7280">
            {{ yBarLabel(i-1) }}
          </text>

          <text :x="W/2" :y="H-2" text-anchor="middle" font-size="11" fill="#374151">
            Прохід
          </text>
          <text :x="12" :y="H/2" text-anchor="middle" font-size="11" fill="#374151"
                :transform="`rotate(-90, 12, ${H/2})`">
            Зусилля (кН)
          </text>
        </svg>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import type { PassResult } from '../api/calculations'

const props = defineProps<{
  passResults: PassResult[]
  passes: number
}>()

const W = 500
const H = 220
const PAD = 45
const innerW = W - PAD * 2
const innerH = H - PAD * 2

const forceHeightData = computed(() => {
  if (!props.passResults.length) return []

  const points: { h: number; f: number; cx: number; cy: number }[] = []
  const allH = [
    ...props.passResults.map(p => ({ h: p.height_start, f: p.force })),
    { h: props.passResults.at(-1)!.height_end, f: props.passResults.at(-1)!.force },
  ]

  const maxH = Math.max(...allH.map(p => p.h))
  const minH = Math.min(...allH.map(p => p.h))
  const maxF = Math.max(...allH.map(p => p.f))

  allH.forEach(p => {
    const cx = PAD + ((maxH - p.h) / (maxH - minH || 1)) * innerW
    const cy = PAD + (1 - p.f / maxF) * innerH
    points.push({ h: p.h, f: p.f, cx, cy })
  })
  return points
})

const forceHeightPoints = computed(() =>
    forceHeightData.value.map(p => `${p.cx},${p.cy}`).join(' ')
)

function yForceLabel(i: number) {
  if (!props.passResults.length) return ''
  const maxF = Math.max(...props.passResults.map(p => p.force))
  return ((maxF / 4 * i) / 1000).toFixed(0) + 'к'
}

const barData = computed(() => {
  if (!props.passResults.length) return []
  const maxF = Math.max(...props.passResults.map(p => p.force))
  const n = props.passResults.length
  const gap = 10
  const totalGap = gap * (n + 1)
  const barW = (innerW - totalGap) / n

  return props.passResults.map((p, i) => {
    const barH = (p.force / maxF) * innerH
    return {
      pass: p.pass,
      force: p.force,
      x: PAD + gap + i * (barW + gap),
      y: PAD + innerH - barH,
      w: barW,
      barH,
    }
  })
})

function yBarLabel(i: number) {
  if (!props.passResults.length) return ''
  const maxF = Math.max(...props.passResults.map(p => p.force))
  return ((maxF / 4 * i) / 1000).toFixed(0) + 'к'
}
</script>