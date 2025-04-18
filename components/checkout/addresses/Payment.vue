<template>
  <div v-if="!same" class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment') }}</h2>
    <select v-model="selectedPaymentAddress" class="select w-full block">
      <option value="new-one">{{ $t('checkout.new-address') }}</option>
      <option v-for="address in addressess" :key="address.id" :value="address.id">{{ address.name }}</option>
    </select>
    <div class="divider divider-primary">{{ $t('checkout.addresses') }}</div>
    <section class="alert alert-neutral input-primary border-2">
      <input v-model="same" type="checkbox" class="checkbox checkbox-primary border-2 bg-white">
      <span>{{ $t('checkout.same') }}</span>
    </section>
    <PaymentStreet var="name" type="text" label="checkout.label.firstname" />
    <PaymentStreet var="surname" type="text" label="checkout.label.surname" />
    <PaymentStreet var="street" type="text" label="checkout.label.street" />
    <PaymentStreet var="number" type="number"  label="checkout.label.number" />
    <PaymentStreet var="zip" type="number" label="checkout.label.zip" />
    <PaymentStreet var="city" type="text" label="checkout.label.city" />
    <div class="form-control">
      <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
      <select v-model="payment.country" class="select select-bordered select-primary border-2">
        <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
      </select>
    </div>
  </div>
  <div v-else class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment-address') }}</h2>
    <section class="alert border-primary border-2">
      <input v-model="same" type="checkbox" class="checkbox">
      <span>{{ $t('checkout.same') }}</span>
    </section>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core'
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import PaymentFirstName from './input/PaymentFirstName.vue';
import PaymentSurName from './input/PaymentSurName.vue';
import PaymentStreet from './input/PaymentStreet.vue';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const selectedPaymentAddress = ref('new-one');
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const same = useLocalStorage('same', true, {});
const addressess = ref([]);
const selectedShippingAddress = ref('new-one');
const countries = reactive({
  'de': 'Deutschland'
});

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
  if (value == selectedShippingAddress.value) {
    same.value = true;
  }
  else if (value == 'new-one') {
    payment.value.name = '';
    payment.value.surname = '';
    payment.value.street = '';
    payment.value.number = 1;
    payment.value.city = '';
    payment.value.zip = '';
    payment.value.country = 'de';
  }
  else {
    addressess.value.map((item) => {
      if (item.id = value) {
        payment.value.name = item.name.split(' ')[0];
        payment.value.surname = item.name.split(' ')[1];
        payment.value.street = item.street;
        payment.value.number = item.number;
        payment.value.city = item.city;
        payment.value.zip = item.postcode;
        payment.value.country = item.country;
      }
    })
  }
});

watch(payment, () => {
  if (!same.value) {
    if (payment.value.city == '') {
      valid.value.payment = false;
      return
    }
    if (payment.value.name == '') {
      valid.value.payment = false;
      return
    }
    if (payment.value.number < 1) {
      valid.value.payment = false;
      return
    }
    if (payment.value.street == '') {
      valid.value.payment = false;
      return
    }
    if (payment.value.surname == '') {
      valid.value.payment = false;
      return
    }
    if (payment.value.zip == '') {
      valid.value.payment = false;
      return
    }
  }
  valid.value.payment = true;
}, { deep: true });
</script>