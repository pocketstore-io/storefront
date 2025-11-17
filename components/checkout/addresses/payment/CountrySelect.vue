<template>
  <label for="" class="floating-label">
    <span>{{ $t('checkout.label.country') }}</span>
    <select v-model="payment.country" class="select select-secondary select-bordered w-full">
      <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
    </select>
  </label>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/utils/pocketbase";

const pb = usePocketBase();

const countries = ref([]);
const payment = useLocalStorage(
    "payment",
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

onMounted(async () => {
    countries.value = await pb.collection("countries").getFullList(25);
});
</script>