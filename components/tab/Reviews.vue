<template>
  <input
type="radio" name="my_tabs_1" role="tab" class="tab col-span-8 md:col-span-2" aria-label="Tab - Reviews"
    checked="checked" >
  <div role="tabpanel" class="tab-content my-6">
    <Review v-for="review in reviews" :identifier="review.id" />
    <CatalogReviewForm />
  </div>
</template>

<script lang="ts" setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useRoute } from 'vue-router';
import Review from '../catalog/Review.vue';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const route = useRoute();

const reviews = ref([]);

onMounted(async () => {
  reviews.value = (await pb.collection('reviews').getList(1, 9, {
    filter: 'product="' + route.params.slug.replace('.html', '') + '"'
  })).items;
});
</script>
