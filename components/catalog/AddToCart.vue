<template>
  <button class="btn btn-secondary join-item" @click="addToCart(product.id)">
    {{ $t('checkout.add-to-cart') }}
  </button>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";
import {addToast} from "@/util/toast";
// TODO release date button, stock

const { product } = defineProps({
  product: {
    type: Object,
    requiered: true,
  },
});

const pb = usePocketBase();

const cart = useLocalStorage("cart", [], {});

const addToCart = async (id, qty = 1) => {
  if (!cart.value) {
    cart.value = [];
  }

  // Check if the product already exists in the cart
  const existingItem = cart.value.find((item) => item.id === id);
  addToast('Produkt hinzugefügt.', 'success')
};
</script>