<template>
  <section class="w-full bg-base-200 px-3 py-3">
    Review {{ identifier }}
  </section>
</template>

<script lang="ts" setup>
const { identifier } = defineProps({
    identifier: { type: String, required: true },
});
</script>