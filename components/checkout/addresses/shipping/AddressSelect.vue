<template>
  <select v-model="selectedShippingAddress" class="select w-full block">
    <option value="new-one">{{ $t('checkout.new-address') }}</option>
    <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
  </select>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { usePocketBase } from '~/util/pocketbase';

const addressess = ref([]);
const selectedShippingAddress = useLocalStorage('selected-shipping', 'new-one', {});
const pb = usePocketBase();
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const shipping = useLocalStorage('shipping', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de',
}, {});

onMounted(async () => {
  pb.autoCancellation(false)
  addressess.value = (await pb.collection('customer_addresses').getFullList(25));
});

watch(selectedShippingAddress, (value) => {
  if (value == 'new-one') {
    shipping.value.name = '';
    shipping.value.surname = '';
    shipping.value.street = '';
    shipping.value.number = 1;
    shipping.value.city = '';
    shipping.value.zip = '';
    shipping.value.country = 'de';
  }
  else {
    addressess.value.map((item) => {
      if (item.id = value) {
        shipping.value.name = item.name.split(' ')[0];
        shipping.value.surname = item.name.split(' ')[1];
        shipping.value.street = item.street;
        shipping.value.number = item.number;
        shipping.value.city = item.city;
        shipping.value.zip = item.postcode;
        shipping.value.country = item.country;
      }
    })
  }
});

watch(shipping, () => {
  if (shipping.value.city == '') {
    valid.value.shipping = false;
    return
  }
  if (shipping.value.name == '') {
    valid.value.shipping = false;
    return
  }
  if (shipping.value.number < 1) {
    valid.value.shipping = false;
    return
  }
  if (shipping.value.street == '') {
    valid.value.shipping = false;
    return
  }
  if (shipping.value.surname == '') {
    valid.value.shipping = false;
    return
  }
  if (shipping.value.zip == '') {
    valid.value.shipping = false;
    return
  }
  valid.value.shipping = true;
}, { deep: true });

</script>

<style></style>