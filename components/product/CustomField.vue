<template>
  <div v-if="field.type == 'text'" class="col-span-2">
    <div class="divider divider-primary">
      {{ $t('product.config.' + field.label.replaceAll(' ', '-')) }}
    </div>
    <p class="text-sm text-left">
      {{ field.desc }}
    </p>
    <input type="text" class="input input-primary bg-base-200 placeholder-black w-full mt-3" placeholder="Hallo Welt">
  </div>
  <div v-if="field.type == 'select'" class="col-span-2">
    <div class="divider divider-primary">
      {{ $t('product.config.' + field.label.replaceAll(' ', '-')) }}
    </div>
    <p class="text-sm text-left">
      {{ field.desc }}
    </p>
    <select name="test" id="test" class="select w-full select-primary mt-3 bg-base-200">
      <option value="test" v-for="option in field.options">{{ option }}</option>
    </select>
  </div>
</template>

<script lang="ts" setup>
const { field } = defineProps({
  field: { type: Object, required: true }
});
</script>

<style></style>