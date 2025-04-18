<template>
  <section v-if="checkoutStep == 'cart' && loaded" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl">
    <div class="col-span-6 md:col-span-4">
      <table v-if="cart.length > 0" class="w-full">
        <CartHeader />
        <tbody>
          <CartItem :item="item" :index="index" v-for="(item, index) in cart" :key="item.id" />
        </tbody>
      </table>
      <section v-else class="px-3 py-3">
        <section v-if="loaded" class="alert alert-warning text-white">
          <p class="block text-center">
            {{ $t('checkout.cart.empty') }}
          </p>
        </section>
      </section>
    </div>

    <div v-if="cart.length > 0" class="col-span-6 md:col-span-2">
      <CustomCartTotal />
    </div>
    <div v-if="cart.length > 0" class="col-span-6 flex justify-end px-3 py-3">
      <CartNext />
    </div>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import CartItem from './cart/Item.vue'
import CartNext from './cart/Next.vue'
import CartHeader from './cart/Header.vue'

const cart = useLocalStorage("cart", [], {});
const valid = useLocalStorage(
  "checkout-valid",
  { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false },
  {}
);
const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const loaded = ref(false);

onMounted(() => {
  loaded.value = true;

  if (cart.value.length > 0) {
    valid.value.cart = true;
  }
});
</script>