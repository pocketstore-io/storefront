<template>
  <div class="form-control">
    <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
    <select v-model="payment.country" class="select select-bordered select-primary border-2">
      <option v-for="(country, index) in countries" :value="index">{{ $t('country.' + country.code) }}</option>
    </select>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";

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