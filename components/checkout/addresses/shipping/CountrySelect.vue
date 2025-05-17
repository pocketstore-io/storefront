<template>
  <div class="form-control">
    <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
    <select v-model="shipping.country" class="select select-bordered select-primary border-2">
      <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
    </select>
  </div>
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
