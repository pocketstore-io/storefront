<template>
  <div v-if="loaded" class="breadcrumbs text-sm px-3 py-3">
    <ul>
      <li>
        <a href="/de" class="">{{ $t('breadcrumb.home') }}</a>
      </li>
      <li v-for="breadcrumb in breadcrumbs" :key="breadcrumb.id">
        <a :href="'/de/' + breadcrumb.link" class="">{{ $t('breadcrumb.' + breadcrumb.id) }}</a>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';

const loaded = ref(false);
const breadcrumbs = useLocalStorage('breadcrumbs', [], {})

onMounted(() => {
  loaded.value = true;
});
</script>