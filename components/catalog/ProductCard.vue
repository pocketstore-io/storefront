<template>
  <div v-if="product" class="card shadow-xl">
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
        <button class="btn btn-secondary join-item" @click="addToCart(product.id)">
          {{ $t('checkout.add-to-cart') }}
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useLocalStorage } from '@vueuse/core';
import { useMessagesStore } from '~/stores/messages';
const i18n = useI18n();
const locale = i18n.locale;

const storeMessages = useMessagesStore();

const { identifier } = defineProps({
  identifier: { type: String, requiered: true }
});

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const cart = useLocalStorage('cart', [], {});
const product = ref({});

const addToCart = async function (id, qty = 1) {
  let found = false;
  if (typeof cart.value == "undefined") {
    cart.value = [];
  }

  const product = await pb.collection('products').getOne(id);

  cart.value.map((item) => {
    if (item.id == id) {
      found = true;
      storeMessages.add({
        message: 'Update ' + id
      })
    }
  });

  if (found) {
    cart.value.map((item) => {
      if (item.id == id) {
        item.qty += qty
      }
    });
  }
  else {
    storeMessages.add({
      message: 'new ID ' + id
    })
    cart.value.push({
      qty, id, product
    });
  }
}

const load = async function () {
  product.value = (await pb.collection('products').getOne(identifier ?? ''));
}

onMounted(() => {
  load();
});
</script>