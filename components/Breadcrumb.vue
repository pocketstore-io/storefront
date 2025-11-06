<template>
  <div v-if="loaded" class="breadcrumbs text-sm px-3 py-3 bg-white text-black mb-3 rounded-lg xl:mx-auto">
    <ul>
      <li>
        <Fa :icon="faHouse" class="mr-3"/>
        <a href="/de" class="">{{ $t('breadcrumb.home') }}</a>
      </li>
      <li v-for="breadcrumb in breadcrumbs" :key="breadcrumb.id">
        <Fa :icon="['fas', breadcrumb.icon ?? 'question']" class="mr-3"/>
        <a :href="'/de/' + breadcrumb.link" class="">{{ $t('breadcrumb.' + breadcrumb.label) }}</a>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { faHouse } from "@fortawesome/free-solid-svg-icons";
import { breadcrumbs } from "@/util/breadcrumbs";

const loaded = ref(false);

onMounted(() => {
    loaded.value = true;
});
</script>