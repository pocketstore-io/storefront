<template>
  <div class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg">{{ $t('checkout.shipping-address') }}</h2>
    <select v-model="selectedShippingAddress" class="select w-full block">
      <option value="new-one">{{ $t('checkout.new-address') }}</option>
      <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
    </select>
    <div class="divider divider-primary">{{ $t('checkout.address') }}</div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.firstname') }}</label>
      <input v-model="shipping.name" type="text" class="input input-bordered input-primary border-2">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.surname') }}</label>
      <input v-model="shipping.surname" type="text" class="input input-primary border-2 input-bordered">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.street') }}</label>
      <input v-model="shipping.street" type="text" class="input input-primary border-2 input-bordered">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.number') }}</label>
      <input v-model="shipping.number" type="number" class="input input-primary border-2 input-bordered">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.zip') }}</label>
      <input v-model="shipping.zip" type="text" class="input input-primary border-2 input-bordered">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.city') }}</label>
      <input v-model="shipping.city" type="text" class="input input-primary border-2 input-bordered">
    </div>
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
      <select v-model="shipping.country" class="select select-bordered select-primary border-2">
        <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
      </select>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core'
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const addressess = ref([]);

const selectedShippingAddress = ref('new-one');
const shipping = useLocalStorage('shipping', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de',
}, {});
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const countries = reactive({
  'de': 'Deutschland'
});

onMounted(async () => {
  addressess.value = (await pb.collection('customer_addresses').getFullList(25));
  console.log(addressess.value);
  console.log('user: ', pb.authStore.model);
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