import { createRouter, createWebHistory } from 'vue-router'
import CalculatorView from '../views/CalculatorView.vue'
import MaterialsView from '../views/MaterialsView.vue'

const routes = [
    { path: '/',          component: CalculatorView },
    { path: '/materials', component: MaterialsView },
]

export default createRouter({
    history: createWebHistory(),
    routes,
})