<template>
  <div class="divider divider-primary">Suche</div>
  <input v-model="query" type="search" class="input input-sm w-full input-primary my-3" placeholder="Search">
  <details :open="query !== ''" class="form-control">
    <summary class="hidden" />
    <ul v-if="query !== '' && items.length > 0" class="absolute menu bg-base-100 rounded-box z-5 w-52 p-2 shadow">
      <li v-for="item in items" :key="item.id"><a :href="'/product/' + item.id + '.html'">{{ item.name }}</a>
      </li>
    </ul>
    <div class="menu">
      <p v-if="query !== '' && items.length == 0" class="text-sm font-bold text-center bg-white px-3 py-3">
        {{ $t('catalog.no-product-found') }}
      </p>
    </div>
  </details>
</template>

<script lang="ts" setup>
const items = ref([]);
const query = ref("");
</script>
