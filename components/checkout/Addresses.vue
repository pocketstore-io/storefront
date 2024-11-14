<template>
  <section v-if="checkoutStep == 'addresses'" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3">
    <div class="col-span-6 md:col-span-3">
      <h2 class="font-bold text-lg">{{ $t('checkout.shipping-address') }}</h2>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.firstname') }}</label>
        <input v-model="shipping.name" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.surname') }}</label>
        <input v-model="shipping.surname" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.street') }}</label>
        <input v-model="shipping.street" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.number') }}</label>
        <input v-model="shipping.number" type="number" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.zip') }}</label>
        <input v-model="shipping.zip" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.city') }}</label>
        <input v-model="shipping.city" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
        <select v-model="shipping.country" class="select select-bordered">
          <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
        </select>
      </div>
    </div>
    <div v-if="!same" class="col-span-6 md:col-span-3">
      <h2 class="font-bold text-lg mb-3">Payment</h2>
      <section class="alert alert-secondary">
        <input v-model="same" type="checkbox" class="checkbox">
        <span>{{ $t('checkout.same') }}</span>
      </section>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.firstname') }}</label>
        <input v-model="payment.name" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.surname') }}</label>
        <input v-model="payment.surname" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.street') }}</label>
        <input v-model="payment.street" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.number') }}</label>
        <input v-model="payment.number" type="number" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.zip') }}</label>
        <input v-model="payment.zip" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.city') }}</label>
        <input v-model="payment.city" type="text" class="input input-bordered">
      </div>
      <div class="form-control">
        <label for="" class="label font-bold text-sm">{{ $t('checkout.label.country') }}</label>
        <select v-model="payment.country" class="select select-bordered">
          <option v-for="(country, index) in countries" :value="index">{{ country }}</option>
        </select>
      </div>
    </div>
    <div v-else class="col-span-6 md:col-span-3">
      <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment-address') }}</h2>
      <section class="alert alert-secondary">
        <input v-model="same" type="checkbox" class="checkbox">
        <span>{{ $t('checkout.same') }}</span>
      </section>
    </div>
    <div class="col-span-6 flex justify-between py-3">
      <button class="btn btn-primary" @click="checkoutStep = 'customer'">{{ $t('back.to.customer') }}</button>
      <button class="btn btn-primary" :disabled="(!valid.payment && !same) || !valid.shipping"
        @click="checkoutStep = 'payment'"><span>{{ $t('contiue.to.payment') }}</span>
        <Fa :icon="faChevronCircleRight" />
      </button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { faChevronCircleRight } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const same = useLocalStorage('same', true, {});
const shipping = useLocalStorage('shipping', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de',
}, {});
const payment = useLocalStorage('payment', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de'
}, {});
const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});
const countries = reactive({
  'de': 'Deutschland'
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