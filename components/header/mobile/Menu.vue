<template>
  <div class="divider divider-primary">Menu</div>
  <div class="md:hidden grid grid-cols-6 gap-3">
    <div class="col-span-3">
      <select v-model="lang" class="select select-primary w-full">
        <option value="de">{{ $t('lang.de') }}</option>
        <option value="en">{{ $t('lang.en') }}</option>
      </select>
    </div>
    <div class="col-span-3">
      <select v-model="currency" class="select select-primary w-full">
        <option value="€">{{ $t('checkout.currency.euro') }}</option>
        <option value="$">{{ $t('checkout.currency.dollar') }}</option>
      </select>
    </div>
    <div class="col-span-3">
      <a href="/de/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
    </div>
    <div class="col-span-3">
      <a href="/de/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
    </div>
    <div class="col-span-3">
      <a href="/de/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
    </div>
    <div class="col-span-3">
      <a href="/de/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
    </div>
    <div class="col-span-3">
      <a href="/de/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
    </div>
    <div class="col-span-3">
      <a href="/de/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
const lang = useLocalStorage("lang", "de", {});
const currency = useLocalStorage("currency", "euro", {});
</script>
