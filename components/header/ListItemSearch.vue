<template>
  <li>
    <input v-model="query" type="search" class="input input-sm" :placeholder="$t('general.search')">
    <details :open="query !== ''">
      <summary class="hidden" />
      <ul v-if="query !== '' && items.length > 0" class="absolute menu bg-base-100 rounded-box z-5 w-52 p-2 shadow">
        <li v-for="item in items"><a :href="'/product/' + item.id + '.html'">{{ item.name }}</a></li>
      </ul>
      <div class="menu">
        <p v-if="query !== '' && items.length == 0" class="text-sm font-bold text-center bg-white px-3 py-3">
          {{ $t('catalog.no-product-found') }}
        </p>
      </div>
    </details>
  </li>
</template>

<script lang="ts" setup>
import { usePocketBase } from '~/util/pocketbase';

const pb = usePocketBase();
const items = ref([]);
const query = ref('');

watch(query, async () => {
  pb.autoCancellation(false)
  items.value = (await pb.collection('products').getList(1, 6, {
    filter: 'name~"' + query.value + '"'
  })).items;
});

</script>

<style></style>