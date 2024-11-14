<template>
  <section v-if="checkoutStep == 'payment'" class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3">
    <div class="col-span-6 md:col-span-3 space-y-3">
      <h2 class="font-bold text-lg">{{ $t('checkout.payment-method') }} {{ $t('general.select') }} </h2>
      <section v-for="(option) in optionsPayment" :key="option">
        <section class="alert alert-neutral flex">
          <input :id="'radio-' + option.code" v-model="paymentMethod" :disabled="paymentMethodInfo.status == 'locked'"
            type="radio" :value="option.code" name="paymentMethod" class="radio mt-2 mr-3">
          <label class="label" :for="'radio-' + option.code">{{ $t('payment.methods.' + option.code) }}</label>
        </section>
        <div v-if="option.options && option.code == paymentMethod">
          <CheckoutPaymentPaypal v-if="option.code == 'paypal' && paymentMethod == 'paypal'"
            :locked="paymentMethodInfo.status == 'locked'" />
          <CheckoutPaymentStripe v-if="option.code == 'stripe' && paymentMethod == 'stripe'"
            :locked="paymentMethodInfo.status == 'locked'" class="mt-3" />
          <CheckoutPaymentKlarna v-if="option.code == 'klarna' && paymentMethod == 'klarna'"
            :locked="paymentMethodInfo.status == 'locked'" />
          <div v-for="(field, key) in option.options.fields">
            <label for="" class="label">{{ $t('payment.label.' + key) }}</label>
            <input v-model="paymentMethodInfo[field.name]" :type="field.type"
              class="input input-bordered input-primary w-full">
          </div>
        </div>
      </section>
    </div>
    <!--| Split up Payment & Shipping by Component Level |-->
    <div class="col-span-6 md:col-span-3 space-y-3">
      <h2 class="font-bold text-lg">{{ $t('checkout.shipping-method') }} {{ $t('general.select') }}</h2>
      <section v-for="(option) in optionsShipping" :key="option" class="bg-base-200 rounded-lg py-6 px-3 w-full">
        <div class="flex justify-between items-center">
          <section class="input-group flex">
            <input v-model="shippingMethod" type="radio" :value="option.code" name="shippingMethod"
              class="ml-auto radio mt-2 mr-3">
            <label for="" class="label">{{ $t('shipping.methods.' + option.code) }}</label>
          </section>
          <div class="price-tag">
            <b>{{ option.price.toFixed(2) }} €</b>
          </div>
        </div>
      </section>
      <section v-if="shippingMethod == 'click-and-collect'">
        <label for="" class="label font-bold text-lg">{{ $t('general.stores') }}</label>
        <select v-model="selectedStore" class="select w-full bg-base-200">
          <option v-for="store in stores" :value="store.id">
            {{ store.street }} {{ store.number }}, {{ store.zip }} {{ store.city }}
          </option>
        </select>
      </section>
    </div>
    <div class="col-span-6 flex justify-between py-3">
      <button class="btn btn-primary" @click="checkoutStep = 'addresses'">{{ $t('checkout.back.addresses') }}</button>
      <button class="btn btn-primary" @click="checkoutStep = 'confirm'"><span>{{ $t('checkout.continue.confirm') }}</span>
        <Fa :icon="faChevronCircleRight" />
      </button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome';
import { faChevronCircleRight } from '@fortawesome/free-solid-svg-icons';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useRoute } from 'vue-router';

const route = useRoute();
const storePb = usePocketbaseStore();
const { url } = storeToRefs(storePb);
const pb = new PocketBase(url.value);

const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});
const paymentMethod = useLocalStorage('paymentMethod', 'vorkasse', {});
const shippingMethod = useLocalStorage('shippingMethod', 'click-and-collect', {});
const paymentMethodInfo = useLocalStorage('paymentMethodInfo', {}, {});
const shippingMethodInfo = useLocalStorage('shippingMethodInfo', {}, {});

const stores = ref([]);
const optionsPayment = ref([]);
const optionsShipping = ref([]);
const optionsCountry = ref([]);
const selectedStore = ref(null);

watch(selectedStore, (value) => {
  shippingMethodInfo.value = {
    store: value
  }
});

watch(shippingMethod, (value) => {
  if (value != 'click-and-collect') {
    shippingMethodInfo.value = {}
  }
});

onMounted(async () => {
  stores.value = await pb.collection('stores').getFullList();
  optionsPayment.value = await pb.collection('payment_methods').getFullList({
    filter: 'active=true',
    sort: 'sort'
  });
  optionsShipping.value = await pb.collection('shipping_methods').getFullList({
    filter: 'active=true'
  });
  optionsCountry.value = await pb.collection('countries').getFullList();
  if (shippingMethod.value === 'click-and-collect') {
    shippingMethodInfo.value = {
      store: stores.value[0].id
    }
  }
  else {
    shippingMethodInfo.value = {}
  }

  if (route.query.stripe && route.query.stripe !== "") {
    paymentMethod.value = 'stripe';
    paymentMethodInfo.value.stripe = route.query.stripe;
    paymentMethodInfo.value.status = 'locked';
  }
  if (route.query.klarna && route.query.klarna !== "") {
    paymentMethod.value = 'klarna';
    paymentMethodInfo.value.klarna = route.query.klarna;
    paymentMethodInfo.value.status = 'locked';
  }
});
</script>