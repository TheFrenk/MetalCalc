# MetalCalc — Metal Forging Calculator

> 🇷🇺 [Русская версия ниже](#metalcalc--калькулятор-ковки-металла)

---

## Overview

**MetalCalc** is a web application for engineering calculations in metal forging processes. It helps technologists and engineers quickly determine key parameters of plastic deformation.

### Calculated Parameters

| Parameter | Description |
|-----------|-------------|
| Forging Force | Force required to deform the workpiece |
| Specific Pressure | Contact pressure per unit area |
| Deformation Work | Energy consumed during the forging process |
| Power | Required equipment power |
| Strain Rate ε̇ | Rate of plastic deformation |
| Degree of Deformation | Relative reduction in workpiece dimensions |
| Workpiece Mass | Mass of the initial billet |
| Contact Area | Area of tool-workpiece contact |
| Workpiece Volume | Volume of the initial billet |
| Height Reduction | Absolute reduction in height |
| Final Diameter | Resulting diameter after forging |

---

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Backend | Go |
| Frontend | Vue.js + TypeScript |
| Styling | CSS |

---

## Project Structure

```
MetalCalc/
├── backend/        # Go REST API server
├── frontend/       # Vue.js + TypeScript application
└── .gitignore
```

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.20+
- [Node.js](https://nodejs.org/) 18+ and npm

### Backend

```bash
cd backend
go run main.go
```

The API server will start at `http://localhost:8080` (default).

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The application will be available at `http://localhost:5173` (default Vite port).

---

## Usage

1. Open the application in your browser.
2. Enter the initial workpiece parameters (dimensions, material).
3. The calculator will instantly compute all derived forging parameters.
4. Use the results for process planning and equipment selection.

---

## License

This project is open source. See [LICENSE](LICENSE) for details.

---
---

# MetalCalc — Калькулятор ковки металла

> 🇬🇧 [English version above](#metalcalc--metal-forging-calculator)

---

## Описание

**MetalCalc** — веб-приложение для инженерных расчётов в процессах ковки металла. Помогает технологам и инженерам быстро определить ключевые параметры пластической деформации.

### Рассчитываемые параметры

| Параметр | Описание |
|----------|----------|
| Усилие ковки | Сила, необходимая для деформирования заготовки |
| Удельное давление | Контактное давление на единицу площади |
| Работа деформации | Энергия, затрачиваемая в процессе ковки |
| Мощность | Требуемая мощность оборудования |
| Скорость деформации ε̇ | Скорость пластической деформации |
| Степень деформации | Относительное уменьшение размеров заготовки |
| Масса заготовки | Масса исходной заготовки |
| Площадь контакта | Площадь контакта инструмента с заготовкой |
| Объём заготовки | Объём исходной заготовки |
| Обжатие по высоте | Абсолютное уменьшение высоты |
| Конечный диаметр | Диаметр заготовки после ковки |

---

## Технологический стек

| Уровень | Технология |
|---------|-----------|
| Бэкенд | Go |
| Фронтенд | Vue.js + TypeScript |
| Стили | CSS |

---

## Структура проекта

```
MetalCalc/
├── backend/        # REST API сервер на Go
├── frontend/       # Приложение на Vue.js + TypeScript
└── .gitignore
```

---

## Начало работы

### Требования

- [Go](https://go.dev/) 1.20+
- [Node.js](https://nodejs.org/) 18+ и npm

### Запуск бэкенда

```bash
cd backend
go run main.go
```

API-сервер запустится по адресу `http://localhost:8080` (по умолчанию).

### Запуск фронтенда

```bash
cd frontend
npm install
npm run dev
```

Приложение будет доступно по адресу `http://localhost:5173` (порт Vite по умолчанию).

---

## Использование

1. Откройте приложение в браузере.
2. Введите исходные параметры заготовки (размеры, материал).
3. Калькулятор мгновенно вычислит все производные параметры ковки.
4. Используйте результаты для планирования технологического процесса и выбора оборудования.

---

## Лицензия

Проект с открытым исходным кодом. Подробнее в файле [LICENSE](LICENSE).
