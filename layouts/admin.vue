<template>
  <main>
    <aside
      id="default-sidebar"
      class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
      aria-label="Sidebar"
    >
      <div class="h-full px-3 py-4 overflow-y-auto bg-green-400">
        <ul class="font-medium">
          <li>
            <h5 class="font-bold text-xl block text-center mb-3">Hallo</h5>
          </li>
        </ul>
        <ul class="">
          <li>
            <h5 class="font-bold">Dashboard</h5>
          </li>
          <li>
            <a
              href="/de/admin"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Übersicht</span>
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
            <a
              href="/de/admin/brands"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Produkte</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="/de/admin/brands"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Kategorien</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="/de/admin/brands"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Marken</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <div class="divider divier-primary my-0"></div>
          </li>
          <li>
            <h5 class="font-bold">Kasse</h5>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Bestellungen</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">shipping_methods</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Reviews</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">payment_methods</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">tax</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <div class="divider divier-primary my-0"></div>
          </li>
          <li>
            <h5 class="font-bold">Einstellungen</h5>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Länder</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Währungen</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Läden</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Übersetzungen</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <div class="divider divier-primary my-0"></div>
          </li>
          <li>
            <h5 class="font-bold">CMS</h5>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Blog</span>
              <Fa :icon="faLink" class="w-5" color="black" />
            </a>
          </li>
          <li>
            <a
              href="#"
              class="flex items-center p-2 group flex justify-between"
            >
              <span class="ms-3 text-sm">Reviews</span>
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
import "~/main.css";
import {
    faAngleDoubleDown,
    faAngleDoubleRight,
    faAngleDown,
    faArrowsToDot,
    faDotCircle,
    faEdit,
    faLink,
    faListDots,
    faMinus,
    faPlus,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";
</script>