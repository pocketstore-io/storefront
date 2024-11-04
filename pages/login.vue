<template>
  <div class="grid grid-cols-6 gap-3">
    <div class="col-span-6 bg-white">
      <div class="hero bg-base-200 py-24">
        <div class="hero-content text-center">
          <div class="max-w-lg">
            <h1 class="text-5xl font-bold">Herzlich Willkommen</h1>
            <p class="py-6">in deinem Online Shop bei PocketStore.io.</p>
            <button class="btn btn-primary">Get Started</button>
          </div>
        </div>
      </div>
    </div>
    <div class="col-span-6 md:col-span-2 bg-white">
      <FormLogin />
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
    router.push('/')
  }
});
</script>
