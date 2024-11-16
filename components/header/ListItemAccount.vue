<template>
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
</template>

<script lang="ts" setup>
import { faRightFromBracket, faUser } from '@fortawesome/free-solid-svg-icons';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const isLoggedIn = computed(() => {
  return pb.authStore.isValid == true;
});
</script>