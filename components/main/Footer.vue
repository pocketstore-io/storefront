<template>
  <section class="bg-white">
    <div class="grid grid-cols-7 px-3 py-3 text-center mx-auto max-w-5xl">
      <div class="col-span-7 md:col-span-2">
        <NuxtLink :to="'/impressum'" class="text-blue-400">Impressum</NuxtLink>
        & <a href="/data-privacy" class="text-blue-400">Datenschutz</a>
      </div>
      <div class="col-span-7 md:col-span-2">Made with
        <a :href="'https://' + url.value" target="_blank" rel="noopener noreferrer">
          <Fa :icon="faHeart" class="text-red-400" />
        </a>
        by Jonathan
      </div>
      <div class="col-span-7 md:col-span-3">Copyright by <a class="text-blue-600"
          href="https://www.jonathan-martz.de">Jonathan Martz Softwareentwicklung</a></div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { usePocketbaseStore } from '~/stores/pocketbase';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome'
import { faHeart } from '@fortawesome/free-solid-svg-icons';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
</script>
