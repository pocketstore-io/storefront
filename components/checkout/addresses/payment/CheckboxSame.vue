<template>
  <section class="alert border-primary border-2">
    <input v-model="same" type="checkbox" class="checkbox">
    <span>{{ $t('checkout.same') }}</span>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const same = useLocalStorage("same", true, {});
</script>
