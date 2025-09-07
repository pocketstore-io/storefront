<template>
  <section class="grid grid-cols-6 gap-3">
    <div class="col-span-6 px-3 py-3 bg-red-400">
      <label class="floating-label">
        <span>Wonach suchst du ?</span>
        <input type="search" v-model="query" class="input w-full"/>
      </label>
    </div>
    <div v-for="item in filtered" class="col-span-6 xs:col-span-3 md:col-span-2">
      <ProductCard :identifier="item.slug"/>
    </div>
  </section>
</template>

<script setup lang="ts">
import ProductCard from "@/components/catalog/ProductCard.vue";
import {ref, computed, watch} from 'vue'
import {addBreadcrumb} from '@/util/breadcrumbs'
import {usePocketBase} from '@/util/pocketbase'
import type PocketBase from 'pocketbase'
import {useRoute} from 'vue-router'

const {t} = useI18n();
const query = ref('');
const pb: PocketBase = usePocketBase();
const products = ref([]);
const tag = ref({});
const filtered = computed(() => {
  return products.value.filter(item => {
    const hasQuery = query.value && query.value.trim() !== '';
    const hasTag = tag.value && tag.value.trim() !== '';

    const matchesName = !hasQuery || item.name.toLowerCase().includes(query.value.toLowerCase());
    const matchesTag = !hasTag || item.tags.some(t => t.toLowerCase() === tag.value.toLowerCase());

    return matchesName && matchesTag;
  });
});
const route = useRoute();

const load = async () => {
  pb.autoCancellation(false);
  products.value = (await pb.collection('products').getList(1, 10, {
    sort: '-created',
  })).items.slice(0, 6);
}

watch(query, load);

onMounted(async () => {
  tag.value = route.query.tag;
  addBreadcrumb({
    label: 'search',
    icon: 'magnifying-glass',
    link: 'search',
  });

  load();

  useHead({
    title: 'Suche - ' + t('general.title'),
    meta: [
      {
        name: 'description',
        content: t('general.description'),
      }
    ]
  });
});
</script>