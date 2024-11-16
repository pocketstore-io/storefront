<template>
  <a href="/checkout">
    <Fa :icon="faCartShopping" class="text-white mx-2 mt-0" />
    <span class="text-white">
      {{ cart.length }}
    </span>
  </a>
</template>

<script lang="ts" setup>
import { faCartShopping } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';
const cart = useLocalStorage('cart', [], {});
</script>