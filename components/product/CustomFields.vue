<template>
  <section class="">
    <section class="field" v-for="item in fields">
      <ProductCustomField :field="item" />
    </section>
    Hallo 123
  </section>
</template>

<script lang="ts" setup>

const { fields } = defineProps({
  fields: { type: Array, required: true }
});
</script>

<style></style>