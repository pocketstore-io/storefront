<template>
  <div class="grid grid-cols-6 gap-3">
    <div class="col-span-6 md:col-span-3">
      <FormLogin />
    </div>
    <div class="col-span-6">
     <LoginHero />
    </div>
  </div>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useRouter } from 'vue-router';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const router = useRouter();

onMounted(() => {
  if (pb.authStore.isValid) {
    // router.push('/')
  }
});
</script>
