<template>
    <section class="alert alert-neutral input-primary border-2">
      <input v-model="same" type="checkbox" class="checkbox checkbox-primary border-2 bg-white">
      <span>{{ $t('checkout.same') }}</span>
    </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const same = useLocalStorage("same", true, {});
</script>
