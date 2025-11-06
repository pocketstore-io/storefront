<template>
  <div v-for="product in products" :key="product.id" class="col-span-6 md:col-span-2 py-3">
    <CatalogProductCard :identifier="product.slug"/>
  </div>
</template>

<script setup>
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();

const { identifier } = defineProps({
    identifier: { type: String, required: true },
});

const category = ref({});
const products = ref([]);

onMounted(async () => {
    category.value = await pb
        .collection("categories")
        .getFirstListItem('slug="' + identifier + '"');
    products.value = (
        await pb.collection("products").getList(1, 9, {
            filter: 'category="' + category.value.id + '"',
            sort: "-created",
        })
    ).items;
});
</script>
