<template>
  <section v-if="checkoutStep == 'confirm'" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3">
    <div class="col-span-6">
      {{ cart }} <br><br>
      {{ shippingMethod }}<br><br>
      {{ shipping }}<br><br>
      <div v-if="!same"> {{ payment }}</div>
      {{ paymentMethod }}<br><br>
    </div>
    <div class="col-span-6 flex justify-between py-3">
      <button class="btn btn-primary" @click="checkoutStep = 'payment'">Back to Payment</button>
      <button @click="confirmOrder()" class="btn btn-primary">{{ $t('checkout.order') }}</button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});
const cart = useLocalStorage('cart', [], {});

const paymentMethod = useLocalStorage('paymentMethod', 'vorkasse', {});
const paymentMethodInfo = useLocalStorage('paymentMethodInfo', {}, {});
const shippingMethod = useLocalStorage('shippingMethod', 'dhl', {});
const shippingMethodInfo = useLocalStorage('shippingMethodInfo', {}, {});

const same = useLocalStorage('same', true, {});
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const router = useRouter();
const shipping = useLocalStorage('shipping',
  {
    name: '',
    surname: '',
    street: '',
    number: 1,
    zip: '',
    city: '',
    country: 'de',
  },
  {});
const payment = useLocalStorage('payment', {
  name: '',
  surname: '',
  street: '',
  number: 1,
  zip: '',
  city: '',
  country: 'de'
}, {});

const confirmOrder = async () => {
  // console.log();
  let order = await pb.collection('orders').create({
    customer: pb.authStore.model?.id,
    cart: cart.value,
    payment_method: paymentMethod.value,
    payment_method_info: paymentMethodInfo.value,
    shipping_method_info: shippingMethodInfo.value,
    shipping_method: shippingMethod.value,
    payment_address: payment.value,
    shipping_address: shipping.value
  });
  router.push("/checkout/confirm?order=" + order.id);
  cart.value = [];
  checkoutStep.value = 'cart'
  paymentMethod.value = null;
  paymentMethodInfo.value = {};
  shippingMethod.value = null;
  shippingMethodInfo.value = {};
  payment.value = null;
  shipping.value = null;
}
</script>