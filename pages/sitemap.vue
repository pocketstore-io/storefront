<template>
  <section class="page mx-auto max-w-6xl bg-white px-3 py-3">
    <section class="grid grid-cols-12 gap-3">
      <div class="col-span-12">
        <h2 class="font-bold text-xl">Produkte</h2>
      </div>
      <div
        v-for="item in products"
        :key="item.id"
        class="col-span-6 md:col-span-3"
      >
       <a :href="'/de/product/'+item.slug+'.html'">{{ item.name }}</a>
        <p class="text-sm">
          {{ item.desc }}
        </p>
      </div>
    </section>
    <section class="grid grid-cols-12 gap-3 mt-12">
      <div class="col-span-12">
        <h2 class="font-bold text-xl">Kategorien</h2>
      </div>
      <div
        v-for="item in categories"
        :key="item.id"
        class="col-span-6 md:col-span-3"
      >
       <a :href="'/de/category/'+item.slug+'.html'">{{ item.name }}</a>
        <p class="text-sm">
          {{ item.desc }}
        </p>
      </div>
    </section>
    <section class="grid grid-cols-12 gap-3 mt-12">
      <div class="col-span-12">
        <h2 class="font-bold text-xl">Brands</h2>
      </div>
      <div
        v-for="item in brands"
        :key="item.id"
        class="col-span-6 md:col-span-3"
      >
        <a :href="'/de/brand/'+item.slug">{{ item.name }}</a>
      </div>
    </section>
  </section>
</template>

<script setup lang="ts">
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const products = ref([]);
const categories = ref([]);
const brands = ref([]);

const load = async () => {
    products.value = await pb.collection("products").getFullList(100);
    categories.value = await pb.collection("categories").getFullList(100);
    brands.value = await pb.collection("brands").getFullList(100);
};

onMounted(() => {
    load();
});
</script>