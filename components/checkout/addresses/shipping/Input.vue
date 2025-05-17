<template>
  <div class="form-control">
    <label for="" class="label font-bold text-sm">{{ $t('checkout.label.' + props.var) }}</label>
    <input v-model="shipping[props.var]" :type="props.type" class="input input-bordered input-primary border-2">
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";

const props = defineProps({
    var: { type: String, required: true },
    type: { type: String, required: true },
});

const shipping = useLocalStorage(
    "shipping",
    {
        name: "",
        surname: "",
        street: "",
        number: 1,
        zip: "",
        city: "",
        country: "de",
    },
    {},
);
</script>