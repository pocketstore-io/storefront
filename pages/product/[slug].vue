<template>
  <section class="page bg-gray-400 px-3 py-3">
    <ProductSeo v-if="item" :product="item" />
    <section class="grid grid-cols-8 gap-3 px-3 py-3 mx-auto max-w-6xl">
      <div class="col-span-8 md:col-span-3">
        <img
          class="w-full"
          :src="
            url +
            '/api/files/' +
            item.collectionId +
            '/' +
            item.id +
            '/' +
            item.cover
          "
          alt=""
        />
        <ProductGallery :identifier="item.id" />
      </div>
      <div class="col-span-8 md:col-span-5">
        <h2 class="font-bold text-2xl">{{ item.name }}</h2>
        <p class="text-sm">{{ item.description }}</p>

        <review-stars :item="item" />
        <div class="form-control grid grid-cols-6 gap-3 text-right">
          <div class="col-span-6">
            <ProductCustomFields :fields="item.config?.fields" />
          </div>
          <div class="col-span-6">
            <section class="grid grid-cols-6 gap-3">
              <section class="col-span-6">
                <ProductStock :quantity="stock.quantity" />
              </section>
              <div class="col-span-6 md:col-span-1">
                <input
                  type="number"
                  min="1"
                  :value="qty"
                  :max="stock.quantity"
                  class="input placeholder-black w-full bg-white rounded-lg text-center"
                />
              </div>
              <div class="col-span-6 md:col-span-5">
                <ProductAddToCart
                  :stock="stock"
                  :qty="qty"
                  :item="item"
                />
              </div>
            </section>
          </div>
        </div>
      </div>
      <div class="col-span-8 mt-12">
        <CatalogProductTabs :identifier="item.id" />
      </div>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { usePocketBase, usePocketBaseUrl } from "~/util/pocketbase";

const route = useRoute();
const pb = usePocketBase();
const url = usePocketBaseUrl();
const item = ref({});
const qty = ref(1);
const stock = ref({});

const load = async () => {
  item.value = await pb
    .collection("products")
    .getFirstListItem('slug="' + route.params.slug.replace(".html", "") + '"');

  stock.value = (
    await pb.collection("product_stocks").getList(1, 1, {
      filter: 'product="' + item.value.id + '"',
    })
  ).items[0];
};

onMounted(() => {
  load();
});
</script>