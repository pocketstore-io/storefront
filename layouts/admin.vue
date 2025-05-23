<template>
    <main>
        <aside id="default-sidebar"
            class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
            aria-label="Sidebar">
            <div class="h-full px-3 py-4 overflow-y-auto bg-green-400">
                <ul class="font-medium ">
                    <li>
                        <h5 class="font-bold text-xl block text-center mb-3">Hallo</h5>
                    </li>
                </ul>
                <ul class="">
                    <li>
                        <h5 class="font-bold">Dashboard</h5>
                    </li>
                    <li>
                        <a href="#" class="flex items-center p-2 group flex justify-between">
                            <span class="ms-3 text-sm">Overview</span>
                            <Fa :icon="faLink" class="w-5" color="black" />
                        </a>
                    </li>
                    <li>
                        <div class="divider divier-primary my-0"></div>
                    </li>
                    <li>
                        <h5 class="font-bold">Katalog</h5>
                    </li>
                    <li>
                        <a href="/admin/products" class="flex items-center p-2 group flex justify-between">
                            <span class="ms-3 text-sm">Produkte</span>
                            <Fa :icon="faLink" class="w-5" color="black" />
                        </a>
                    </li>
                    <li>
                        <a href="#" class="flex items-center p-2 group flex justify-between">
                            <span class="ms-3 text-sm">Kategorien</span>
                            <Fa :icon="faLink" class="w-5" color="black" />
                        </a>
                    </li>
                    <li>
                        <div class="divider divier-primary my-0"></div>
                    </li>
                    <li>
                        <h5 class="font-bold">Hallo</h5>
                    </li>
                    <li>
                        <a href="#" class="flex items-center p-2 group flex justify-between">
                            <span class="ms-3 text-sm">Dashboard</span>
                            <Fa :icon="faLink" class="w-5" color="black" />
                        </a>
                    </li>
                </ul>
            </div>
        </aside>

        <div class="p-4 sm:ml-64 bg-red-400 min-h-screen">
            <slot></slot>
        </div>
    </main>
</template>

<script setup lang="ts">
import { faAngleDoubleDown, faAngleDoubleRight, faAngleDown, faArrowsToDot, faDotCircle, faEdit, faLink, faListDots, faMinus, faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome';
</script>