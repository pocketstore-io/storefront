<template>
  <div v-if="product" class="card shadow-xl bg-white">
    <figure>
      <a :href="'/' + locale + '/product/' + product.slug + '.html'">
        <img
          :src="
            'https://' +
            config.domain +
            '/api/files/' +
            product.collectionId +
            '/' +
            product.id +
            '/' +
            product.cover
          "
          alt="Shoes"
        />
      </a>
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ product.name }}</h2>
      <p class="text-ellipsis line-clamp-2">{{ product.description }}</p>
      <section class="w-full space-x-3 mb-3">
        <span v-for="tag in tags" class="badge badge-secondary text-sm">{{ tag }}</span>
      </section>
      <div class="card-actions join gap-0 justify-end">
        <a
          v-if="product.price"
          :href="'/' + locale + '/product/' + product.slug + '.html'"
          class="btn btn-primary join-item"
        >
          {{ product.price.toFixed(2) }} €
        </a>
        <ProductAddToCart :stock="stock" :qty="1" :item="product" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { usePocketBase } from "~/util/pocketbase";
import config from "@/pocketstore.json";
import { useLocalStorage } from "@vueuse/core";

const pb = usePocketBase();

const i18n = useI18n();
const locale = i18n.locale;
const stock = ref({});
const tags = ref([]);

const props = defineProps({
  identifier: { type: String, requiered: true },
});

onUpdated(() => {
  if (props.identifier != product.value.id) {
    load();
  }
});

const product = useLocalStorage("product-" + props.identifier + "", {}, {});
const date = useLocalStorage(
  "product-" + props.identifier + "-date",
  new Date().toLocaleDateString("de"),
  {}
);

function isEmpty(obj) {
  for (const prop in obj) {
    if (Object.hasOwn(obj, prop)) {
      return false;
    }
  }

  return true;
}

onMounted(async () => {
  if (
    product.value === null ||
    date.value != new Date().toLocaleDateString("de") ||
    isEmpty(product.value)
  ) {
    load();

    stock.value = await pb
      .collection("product_stocks")
      .getFirstListItem('product="' + product.value.id + '"');
  }
});

const load = async () => {
  product.value = await pb
    .collection("products")
    .getFirstListItem('slug="' + props.identifier + '"');

  tags.value = (
    await pb.collection("product_tags").getList(1, 1, {
      filter: 'product="' + product.value.id + '"',
    })
  ).items[0].tags.split(",");
};
</script>