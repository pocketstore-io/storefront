<template>
  <div class="col-span-6 md:col-span-3">
    <h2 class="font-bold text-lg">{{ $t('checkout.shipping-address') }}</h2>
    <AddressSelect />
    <div class="divider divider-primary">{{ $t('checkout.address') }}</div>
    <Input var="name" type="text" />
    <Input var="surname" type="text" />
    <Input var="street" type="text" />
    <Input var="number" type="number" />
    <Input var="zip" type="number" />
    <Input var="city" type="text" />
    <CountrySelect />
  </div>
</template>

<script lang="ts" setup>
import AddressSelect from './shipping/AddressSelect.vue';
import CountrySelect from './shipping/CountrySelect.vue';
import Input from './shipping/Input.vue';
</script>