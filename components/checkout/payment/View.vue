<template>
  <div class="col-span-6 md:col-span-3 space-y-3">
    <h2 class="font-bold text-lg">{{ $t('checkout.payment-method') }} {{ $t('general.select') }} </h2>
    <section v-for="(option) in optionsPayment" :key="option">
      <section class="flex px-6 py-3 border-2 border-base-300 rounded-lg">
        <input :id="'radio-' + option.code" v-model="paymentMethod"
          :disabled="paymentMethodInfo.status == 'locked'" type="radio" :value="option.code"
          name="paymentMethod" :class="{ '!border-primary !bg-primary': option.code == paymentMethod }"
          class="radio border-black bg-white border-2 mt-2 mr-3">
        <label class="label" :for="'radio-' + option.code">{{ $t('payment.methods.' + option.code)
          }}</label>
      </section>
      <div v-if="option.options && option.code == paymentMethod">
        <CheckoutPaymentPaypal v-if="option.code == 'paypal' && paymentMethod == 'paypal'"
          :locked="paymentMethodInfo.status == 'locked'" />
        <CheckoutPaymentStripe v-if="option.code == 'stripe' && paymentMethod == 'stripe'"
          :locked="paymentMethodInfo.status == 'locked'" class="mt-3" />
        <CheckoutPaymentKlarna v-if="option.code == 'klarna' && paymentMethod == 'klarna'"
          :locked="paymentMethodInfo.status == 'locked'" />
        <div v-for="(field, key) in option.options.fields">
          <label for="" class="label text-sm font-bold">{{ $t('payment.label.' + key) }}</label>
          <input v-model="paymentMethodInfo[field.name]" :type="field.type"
            class="input input-bordered bg-white border-2 w-full">
        </div>
      </div>
    </section>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const storePb = usePocketbaseStore();
const { url } = storeToRefs(storePb);
const pb = new PocketBase(url.value);
const paymentMethod = useLocalStorage('paymentMethod', 'vorkasse', {});
const paymentMethodInfo = useLocalStorage('paymentMethodInfo', {}, {});
const optionsPayment = ref([]);

onMounted(async () => {
	optionsPayment.value = await pb.collection('payment_methods').getFullList({
		filter: 'active=true',
		sort: 'sort'
	});
  });
</script>