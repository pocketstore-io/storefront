<template>
  <section class="header bg-[#1f6fed] px-3 py-3">
    <section class="flex justify-between">
      <section class="logo mt-3 text-white md:text-black md:mt-0">
        <a href="/" class="font-bold">{{ $t('general.title') }}</a>
      </section>
      <nav class="hidden md:block">
        <ul class="flex space-x-3">
          <li class="">
            <div>
              <div>
                <select v-model="currency" class="select select-sm">
                  <option value="euro">{{ $t('currency.euro') }}</option>
                  <option value="dollar">{{ $t('currency.dollar') }}</option>
                </select>
              </div>
            </div>
          </li>
          <li class="">
            <div>
              <div>
                <select v-model="lang" class="select select-sm">
                  <option value="de">{{ $t('lang.de') }}</option>
                  <option value="en">{{ $t('lang.en') }}</option>
                </select>
              </div>
            </div>
          </li>
          <li class="mt-1">
            <a href="/category/welcome">
              <Fa :icon="faLayerGroup" class="text-white mx-2 mt-1" />
            </a>
          </li>
          <li>
            <input v-model="query" type="search" class="input input-sm" :placeholder="$t('general.search')">
            <details :open="query !== ''">
              <summary class="hidden" />
              <ul
v-if="query !== '' && items.length > 0"
                class="absolute menu bg-base-100 rounded-box z-5 w-52 p-2 shadow">
                <li v-for="item in items"><a :href="'/product/' + item.id + '.html'">{{ item.name }}</a></li>
              </ul>
              <div class="menu">
                <p v-if="query !== '' && items.length == 0" class="text-sm font-bold text-center bg-white px-3 py-3">
                  {{ $t('catalog.no-product-found') }}
                </p>
              </div>
            </details>
          </li>
          <li class="mt-1">
            <a href="/checkout">
              <Fa :icon="faCartShopping" class="text-white mx-2 mt-1" />
              <span class="text-white">
                {{ cart.length }}
              </span>
            </a>
          </li>
          <li v-if="isLoggedIn" class="mt-1">
            <a href="/user/logout">
              <Fa :icon="faRightFromBracket" class="text-white mx-2 mt-1" />
            </a>
          </li>
          <li v-else class="mt-1">
            <a href="/login">
              <Fa :icon="faUser" class="text-white mx-2 mt-1" />
            </a>
          </li>
        </ul>
      </nav>
      <nav class="md:hidden flex items-center">
        <a href="/checkout">
          <Fa :icon="faCartShopping" class="text-white mx-2 mt-0" />
          <span class="text-white">
            {{ cart.length }}
          </span>
        </a>
        <button class="btn btn-primary" @click="open = !open">
          <span v-if="!open">
            <Fa :icon="faBars" size="2x" class="text-white" />
          </span>
          <span v-else>
            <Fa :icon="faTimes" size="2x" class="text-white" />
          </span>
        </button>
        <dialog id="my_modal_1" class="modal modal-top px-6">
          <div class="modal-box">
            <img src="https://cdn.jmse.cloud/jmse-logo.svg" alt="" class="w-full">
            <div class="divider divider-primary">Suche</div>
            <input v-model="query" type="search" class="input input-sm w-full input-primary my-3" placeholder="Search">
            <details :open="query !== ''" class="form-control">
              <summary class="hidden" />
              <ul
v-if="query !== '' && items.length > 0"
                class="absolute menu bg-base-100 rounded-box z-5 w-52 p-2 shadow">
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
                  <option value="euro">{{ $t('currency.euro') }}</option>
                  <option value="dollar">{{ $t('currency.dollar') }}</option>
                </select>
              </div>
              <div class="col-span-3">
                <a href="/category/welcome" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
              </div>
              <div class="col-span-3">
                <a href="/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
              </div>
              <div class="col-span-3">
                <a href="/category/welcome" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
              </div>
              <div class="col-span-3">
                <a href="/checkout/" class="btn btn-primary btn-block">{{ $t('checkout.checkout') }}</a>
              </div>
              <div class="col-span-3">
                <a href="/category/welcome" class="btn btn-primary btn-block">{{ $t('catalog.category') }}</a>
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
      </nav>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { faUser, faRightFromBracket, faCartShopping, faLayerGroup, faBars, faTimes } from '@fortawesome/free-solid-svg-icons'
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useLocalStorage } from '@vueuse/core';
import propertyManager from '~/util/settings'

const { setLocale } = useI18n()
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const cart = useLocalStorage('cart', [], {});
const items = ref([]);
const open = ref(false);
const query = ref('');
const lang = ref('de');
const currency = ref(propertyManager.getSettingsValue('currency'));

watch(lang, (value) => {
  setLocale(value);
});

watch(open, (value) => {
  if (value) {
    my_modal_1.showModal()
    document.activeElement.blur();
  }
});

watch(query, async () => {
  items.value = (await pb.collection('products').getList(1, 6, {
    filter: 'name~"' + query.value + '"'
  })).items;
});

const isLoggedIn = computed(() => {
  return pb.authStore.isValid == true;
});
</script>