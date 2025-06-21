<template>
  <label for="" class="floating-label">
    <span>{{ $t('checkout.label.country') }}</span>
    <select v-model="shipping.country" class="select select-neutral select-bordered w-full">
      <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
    </select>
  </label>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
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
const countries = ref([]);

onMounted(async () => {
    countries.value = await pb.collection("countries").getFullList(25);
});
</script>
