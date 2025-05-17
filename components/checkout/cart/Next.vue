<template>
  <button class="btn btn-primary" :disabled="!valid.cart" @click="checkoutStep = 'customer'">
    <span>{{ $t('checkout.continue.customer') }}</span>
    <Fa :icon="faChevronCircleRight" />
  </button>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { faChevronCircleRight } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon as Fa } from "@fortawesome/vue-fontawesome";

const cart = useLocalStorage("cart", [], {});
const valid = useLocalStorage(
    "checkout-valid",
    {
        cart: false,
        addresses: false,
        payment: false,
        shipping: false,
        confirm: false,
        customer: false,
    },
    {},
);
const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
</script>