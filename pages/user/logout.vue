<template>
  <div />
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { usePocketBase } from '~/util/pocketbase';

const router = useRouter();
const pb = usePocketBase();

onMounted(() => {
  pb.authStore.clear();
  router.push('/');
});
</script>
