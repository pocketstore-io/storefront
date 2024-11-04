<template>
  <section class="page bg-white">
    <section class="grid grid-cols-10 px-3 py-3 gap-3" v-if="loaded">
      <div class="col-span-10">
        <div id="headline" class="divider divider-primary font-bold text-lg">{{ $t('checkout.checkout') }}</div>
      </div>
      <div class="col-span-10 md:col-span-2">
        <button class="btn btn-block"
          :class="{ 'btn-secondary': checkoutStep == 'cart', 'btn-neutral': checkoutStep != 'cart' }">{{
            $t('checkout.cart') }}</button>
      </div>
      <div class="col-span-10 md:col-span-2">
        <button class="btn btn-block"
          :class="{ 'btn-primary': checkoutStep == 'customer', 'btn-neutral': checkoutStep != 'customer' }">{{
            $t('checkout.customer') }}</button>
      </div>
      <div class="col-span-10 md:col-span-2">
        <button class="btn btn-block"
          :class="{ 'btn-primary': checkoutStep == 'addresses', 'btn-neutral': checkoutStep != 'addresses' }">{{
            $t('checkout.addresses') }}</button>
      </div>
      <div class="col-span-10 md:col-span-2">
        <button class="btn btn-block"
          :class="{ 'btn-primary': checkoutStep == 'payment', 'btn-neutral': checkoutStep != 'payment' }">{{
            $t('checkout.payment-and-shipping') }}</button>
      </div>
      <div class="col-span-10 md:col-span-2">
        <button class="btn btn-block"
          :class="{ 'btn-primary': checkoutStep == 'confirm', 'btn-neutral': checkoutStep != 'confirm' }">{{
            $t('checkout.confirm') }}</button>
      </div>
    </section>

    <CheckoutCart />
    <CheckoutAddresses />
    <CheckoutCustomer />
    <CheckoutPayment />
    <CheckoutConfirm />
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { useBreadcrumbStore } from '~/stores/breadcrumb';

const storeBreadcrumb = useBreadcrumbStore();
const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});

const loaded = ref(false);

watch(checkoutStep, () => {
  storeBreadcrumb.clear();
  storeBreadcrumb.add({
    link: "checkout/" + checkoutStep.value,
    id: "checkout-" + checkoutStep.value,
  });
  document.getElementById('headline')?.scrollIntoView({
    behavior: "smooth"
  });
});

onMounted(() => {
  loaded.value = true;
});
</script>