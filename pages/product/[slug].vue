<template>
  <section class="page bg-gray-400 px-3 py-3">
    <ProductSeo v-if="item" :product="item" />
    <section class="grid grid-cols-8 gap-3 px-3 py-3 mx-auto max-w-6xl">
      <div class="col-span-8 md:col-span-3">
        <img
          class="aspect-square"
          :src=" url +
          '/api/files/' +
          item.collectionId +
          '/' +
          item.id +
          '/' +
          item.cover"
          alt=""
        />
        <ProductGallery :identifier="item.id" />
      </div>
      <div class="col-span-8 md:col-span-5">
        <h2 class="font-bold text-2xl">{{ item.name }}</h2>
        <p class="text-sm">{{ item.description }}</p>

        <h3 v-if="item.price" class="font-bold text-4xl">
          <div class="glass px-3 py-3 mt-6 mb-6 md:flex md:justify-between">
            <section class="text-center space-x-3">
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
            </section>
            <div class="mt-6 md:mt-2 text-right md:text-left">
              {{ item.price.toFixed(2) }} €
            </div>
          </div>
        </h3>

        <div class="form-control grid grid-cols-6 gap-3 text-right">
          <div class="col-span-6">
            <ProductCustomFields :fields="item.config?.fields" />
          </div>
          <div class="col-span-6">
            <section class="grid grid-cols-6 gap-3">
              <div class="col-span-6 md:col-span-1">
                <input
                  type="number"
                  min="1"
                  placeholder="123"
                  class="input placeholder-black w-full bg-white rounded-lg text-center"
                />
              </div>
              <div class="col-span-6 md:col-span-5">
                <button
                  class="btn btn-secondary w-full"
                  @click="addToCart(item.id)"
                >
                  {{ $t("checkout.add-to-cart") }}
                </button>
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
import { faStar } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase, usePocketBaseUrl } from "~/util/pocketbase";

const route = useRoute();
const pb = usePocketBase();
const url = usePocketBaseUrl();
const item = ref({});
const cart = useLocalStorage("cart", [], {});

const addToCart = async function (id, qty = 1) {
    let found = false;
    if (typeof cart.value == "undefined") {
        cart.value = [];
    }

    const product = await pb.collection("products").getOne(id);

    cart.value.map((item) => {
        if (item.id == id) {
            found = true;
        }
    });

    if (found) {
        cart.value.map((item) => {
            if (item.id == id) {
                item.qty += qty;
            }
        });
    } else {
        // TODO cart message
        cart.value.push({
            qty,
            id,
            product,
        });
    }
};

const load = async () => {
    item.value = await pb
        .collection("products")
        .getFirstListItem(
            'slug="' + route.params.slug.replace(".html", "") + '"',
        );
};

onMounted(() => {
    load();
});
</script>