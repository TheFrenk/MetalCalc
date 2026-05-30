<template>
  <div class="min-h-screen bg-gray-50 text-gray-900">

    <!-- Header -->
    <header class="bg-gradient-to-r from-slate-800 to-slate-700 text-white shadow-lg">
      <div class="max-w-5xl mx-auto px-6 py-4 flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="bg-blue-500 rounded-lg p-2">
            <svg class="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9 3H5a2 2 0 00-2 2v4m6-6h10a2 2 0 012 2v4M9 3v18m0 0h10a2 2 0 002-2V9M9 21H5a2 2 0 01-2-2V9m0 0h18" />
            </svg>
          </div>
          <div>
            <h1 class="text-lg font-bold leading-tight">MetalCalc</h1>
            <p class="text-xs text-slate-400">Розрахунок параметрів кування</p>
          </div>
        </div>

        <nav class="flex items-center gap-1">
          <RouterLink
              to="/"
              class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200"
              :class="$route.path === '/'
              ? 'bg-blue-500 text-white shadow'
              : 'text-slate-300 hover:bg-slate-600 hover:text-white'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 11h.01M12 11h.01M15 11h.01M4 19h16a2 2 0 002-2V7a2 2 0 00-2-2H4a2 2 0 00-2 2v10a2 2 0 002 2z" />
            </svg>
            Калькулятор
          </RouterLink>
          <RouterLink
              to="/materials"
              class="flex items-center gap-2 px-4 py-2 rounded-lg text-sm font-medium transition-all duration-200"
              :class="$route.path === '/materials'
              ? 'bg-blue-500 text-white shadow'
              : 'text-slate-300 hover:bg-slate-600 hover:text-white'"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z" />
            </svg>
            Матеріали і форми
          </RouterLink>
        </nav>
      </div>
    </header>

    <main class="max-w-5xl mx-auto px-6 py-8">
      <router-view />
    </main>

    <footer class="border-t border-gray-200 mt-12">
      <div class="max-w-5xl mx-auto px-6 py-4 flex items-center justify-between text-xs text-gray-400">
        <span>MetalCalc © {{ new Date().getFullYear() }}</span>
        <span>НТУ «ХПІ» — Кафедра КМІТ</span>
      </div>
    </footer>

  </div>
</template>