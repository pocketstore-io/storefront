<template>
  <div class="form-control">
    <label for="" class="label font-bold text-sm">{{ $t(props.label) }}</label>
    <input v-model="payment[props.var]" :type="type" class="input input-bordered input-primary border-2">
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';

const props = defineProps({
  var: { type: String, required: true },
  type: { type: String, required: true },
  label: { type: String, required: true },
});

const payment = useLocalStorage('payment', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de'
}, {});
</script>