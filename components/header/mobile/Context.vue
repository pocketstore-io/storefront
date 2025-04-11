<template>
  <dialog id="modal_mobile_menu" class="modal modal-top px-6">
    <div class="modal-box">
      <img src="https://cdn.jmse.cloud/jmse-logo.svg" alt="" class="w-full">
      <div class="divider divider-primary">Suche</div>
      <input v-model="query" type="search" class="input input-sm w-full input-primary my-3" placeholder="Search">
      <details :open="query !== ''" class="form-control">
        <summary class="hidden" />
        <ul v-if="query !== '' && items.length > 0" class="absolute menu bg-base-100 rounded-box z-5 w-52 p-2 shadow">
          <li v-for="item in items" :key="item.id"><a :href="'/product/' + item.id + '.html'">{{ item.name }}</a>
          </li>
        </ul>
        <div class="menu">
          <p v-if="query !== '' && items.length == 0" class="text-sm font-bold text-center bg-white px-3 py-3">
            {{ $t('catalog.no-product-found') }}
          </p>
        </div>
      </details>
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
          <a href="/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
        </div>
        <div class="col-span-3">
          <a href="/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
        </div>
        <div class="col-span-3">
          <a href="/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
        </div>
        <div class="col-span-3">
          <a href="/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
        </div>
        <div class="col-span-3">
          <a href="/category/bowling" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
        </div>
        <div class="col-span-3">
          <a href="/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
        </div>
      </div>
      <div class="modal-action">
        <form method="dialog">
          <!-- if there is a button in form, it will close the modal -->
          <button class="btn">{{ $t('general.close') }}</button>
        </form>
      </div>
    </div>
  </dialog>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
const lang = useLocalStorage('lang', 'de', {});
const currency = useLocalStorage('currency', 'euro', {});
const items = ref([]);
const query = ref('');
</script>
