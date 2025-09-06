<template>
  <section class="mx-auto max-w-6xl bg-white px-3 py-3 rounded-lg">
    <section class="grid grid-cols-6 gap-3">
      <div class="col-span-6">
        <h3 class="font-bold text-xl mb-3">Unsere Läden und deren Status</h3>
        <p class="block text-sm text-center">
          Dies sind alle psysischen Lände von uns an dem Button kannst du erkennen
          ob click und collect verfügbar ist.
        </p>
      </div>
      <div
        v-for="store in stores"
        :key="store.id"
        class="col-span-6 md:col-span-2"
      >
        <section class="card bg-gray-400">
          <figure>
            <img src="https://place-hold.it/500x270" alt="Shoes" />
          </figure>
          <div class="card-body pb-2">
            <div class="">
              <div class="font-bold">{{ store.name }}</div>
              <div class="text-xs mb-1 mt-1">
                {{ store.street }} {{ store.number }}
              </div>
              <div class="text-xs mb-1">{{ store.zip }} {{ store.city }}</div>
            </div>
          </div>
          <section class="card-actions px-3 pb-3">
            <button v-if="store.clickandcollect" class="btn btn-success btn-block">
              Click & Collect: verfügbar
            </button>
            <button v-else class="btn btn-warning btn-block">
              Click & Collect: nicht verfügbar
            </button>
          </section>
        </section>
      </div>
    </section>
  </section>
</template>


<script setup lang="ts">
import {
  faBagShopping,
  faBan,
  faEdit,
  faInfoCircle,
  faPlay,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const stores = ref([]);

const load = async () => {
  stores.value = await pb.collection("stores").getFullList(50);
};

onMounted(() => {
  load();
});
</script>