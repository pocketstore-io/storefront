<template>
  <div v-if="!same" class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment') }}</h2>
    <PaymentSelect />
    <div class="divider divider-primary">{{ $t('checkout.addresses') }}</div>
    <RadioSame />
    <div class="space-y-6">
      <PaymentInput var="name" type="text" />
      <PaymentInput var="surname" type="text" />
      <PaymentInput var="street" type="text" />
      <PaymentInput var="number" type="number" />
      <PaymentInput var="zip" type="number" />
      <PaymentInput var="city" type="text" />
      <CountrySelect />
    </div>
  </div>
  <div v-else class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg mb-3">{{ $t('checkout.payment-address') }}</h2>
    <CheckboxSame />
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import PaymentInput from "./payment/Input.vue";
import CountrySelect from "./payment/CountrySelect.vue";
import RadioSame from "./payment/RadioSame.vue";
import PaymentSelect from "./payment/AddressSelect.vue";
import CheckboxSame from "./payment/CheckboxSame.vue";

const same = useLocalStorage("same", true, {});
</script>