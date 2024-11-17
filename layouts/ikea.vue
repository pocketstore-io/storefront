<template>
  <div class="bg-white">
    <IkeaHeader />
    <slot />
    <div class="alert alert-warning mt-3 max-w-[90vw] mx-auto">
      Hallo Welt
    </div>
  </div>
</template>

<script lang="ts" setup>
import IkeaHeader from './ikea/Header.vue';
</script>