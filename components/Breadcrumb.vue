<template>
  <div v-if="loaded" class="breadcrumbs text-sm px-3 py-3 bg-white text-black mb-3 rounded-lg">
    <ul>
      <li>
        <Fa :icon="faHouse" class="mr-3"/>
        <a href="/de" class="">{{ $t('breadcrumb.home') }}</a>
      </li>
      <li v-for="breadcrumb in breadcrumbs" :key="breadcrumb.id">
        <a :href="'/de/' + breadcrumb.link" class="">{{ $t('breadcrumb.' + breadcrumb.id) }}</a>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { faHouse } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';

const loaded = ref(false);
const breadcrumbs = useLocalStorage('breadcrumbs', [], {})

onMounted(() => {
  loaded.value = true;
});
</script>