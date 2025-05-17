<template>
  <select v-model="selectedShippingAddress" class="select w-full block">
    <option value="new-one">{{ $t('checkout.new-address') }}</option>
    <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
  </select>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { select, validation } from "~/usecase/checkout/shipping";
import { usePocketBase } from "~/util/pocketbase";

const addressess = ref([]);
const selectedShippingAddress = useLocalStorage(
    "selected-shipping",
    "new-one",
    {},
);
const pb = usePocketBase();
const valid = useLocalStorage(
    "checkout-valid",
    {
        cart: false,
        addresses: false,
        payment: false,
        shipping: false,
        confirm: false,
        customer: false,
    },
    {},
);
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

onMounted(async () => {
    pb.autoCancellation(false);
    addressess.value = await pb
        .collection("customer_addresses")
        .getFullList(25);
});

watch(selectedShippingAddress, (value) => {
    select(value, shipping);
});

watch(
    shipping,
    () => {
        validation(shipping, valid);
    },
    { deep: true },
);
</script>
