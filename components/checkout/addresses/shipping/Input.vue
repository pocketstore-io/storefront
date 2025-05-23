<template>
  <label class="floating-label">
    <span class="">{{ $t('checkout.label.' + props.var) }}</span>
    <input v-model="shipping[props.var]" :type="props.type" class="w-full input-neutral input-bordered input">
  </label>
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