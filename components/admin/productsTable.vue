<template>
  <ul class="list bg-gray-400 rounded-box shadow-md">
    <li v-for="product in products" class="list-row">
      <div>
        <img
          class="size-10 rounded-box"
          :src="
            url +
            '/api/files/' +
            product.collectionId +
            '/' +
            product.id +
            '/' +
            product.cover
          "
        />
      </div>
      <div>
        <div>{{ product.name }}</div>
        <div class="text-xs uppercase font-semibold opacity-60">
          {{ product.created }}
        </div>
      </div>
      <button class="btn btn-square btn-ghost">
        <FontAwesomeIcon :icon="faInfoCircle" />
      </button>
      <button class="btn btn-square btn-ghost">
        <FontAwesomeIcon :icon="faLink" />
      </button>
    </li>
  </ul>
</template>

<script setup lang="ts">
import { usePocketBase, usePocketBaseUrl } from "~/utils/pocketbase";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
    faEdit,
    faInfoCircle,
    faLink,
} from "@fortawesome/free-solid-svg-icons";

const pb = usePocketBase();
const url = usePocketBaseUrl();
const products = ref([]);

const load = async () => {
    products.value = (
        await pb.collection("products").getList(1, 5, {
            sort: "created",
        })
    ).items;
};
onMounted(() => {
    load();
});
</script>