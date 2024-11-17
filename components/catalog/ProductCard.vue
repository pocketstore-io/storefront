<template>
  <div v-if="product" class="card shadow-xl bg-white">
    <figure>
      <a :href="'/' + locale + '/product/' + product.slug + '.html'">
        <img src="https://img.daisyui.com/images/stock/photo-1606107557195-0e29a4b5b4aa.webp" alt="Shoes">
      </a>
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ product.name }}</h2>
      <p class="text-ellipsis line-clamp-2">{{ product.description }}</p>
      <div class="card-actions join gap-0 justify-end">
        <a v-if="product.price" :href="'/' + locale + '/product/' + product.slug + '.html'"
          class="btn btn-primary join-item">
          {{ product.price.toFixed(2) }} €
        </a>
        <CatalogAddToCart :product="product" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const i18n = useI18n();
const locale = i18n.locale;

const { identifier } = defineProps({
  identifier: { type: String, requiered: true }
});

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const product = ref({});

onMounted(async () => {
  product.value = (await pb.collection('products').getOne(identifier ?? ''));
});
</script>