<template>
  <section class="">
    <section class="card py-3 px-3 bg-white mb-3">
      <div class="card-body px-3 py-3 m-0 grid grid-cols-4 gap-3">
        <ProductCustomField v-for="item in fields" :key="item.id" :field="item" />
      </div>
    </section>
  </section>
</template>

<script lang="ts" setup>

const { fields } = defineProps({
  fields: { type: Array, required: true }
});
</script>

<style></style>