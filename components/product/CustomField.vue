<template>
  <section class="field">
    <div v-if="field.type== 'text'" class="grid grid-cols-6">
      <label for="" class="label col-span-6 md:col-span-2 text-sm font-bold">
        Hallo Welt
      </label>
      <input type="text" class="input col-span-6 md:col-span-4" placeholder="Hallo Welt">
    </div>
  </section>
</template>

<script lang="ts" setup>
const { field } = defineProps({
  field: { type: Object, required: true }
});
</script>

<style>

</style>