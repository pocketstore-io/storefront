<template>
  <select v-model="selectedPaymentAddress" class="select w-full block">
    <option value="new-one">{{ $t('checkout.new-address') }}</option>
    <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
  </select>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { usePocketBase } from '~/util/pocketbase';
import { select,validation } from '~/usecase/checkout/payment';

const same = useLocalStorage('same', true, {});
const addressess = ref([]);
const pb = usePocketBase();
const selectedPaymentAddress = useLocalStorage('selected-payment', 'new-one', {});
const selectedShippingAddress = useLocalStorage('selected-payment', 'new-one', {});
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const payment = useLocalStorage('payment', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de'
}, {});

onMounted(async () => {
  addressess.value = (await pb.collection('customer_addresses').getFullList(25));
});

watch(selectedPaymentAddress, (value) => {
  select(value,payment,addressess,same,selectedShippingAddress)
});

watch(payment, () => {
  validation(valid,same,payment);
}, { deep: true });
</script>