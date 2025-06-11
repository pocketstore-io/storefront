<template>
  <div class="tabs tabs-border bg-white">
    <input
      type="radio"
      name="my_tabs_2"
      class="tab"
      @click="selected = 'description'"
      :checked="selected == 'description'"
      aria-label="Beschreibung"
    />
    <div class="tab-content bg-base-100 p-10">
      <CatalogTabDescription />
    </div>

    <input
      type="radio"
      name="my_tabs_2"
      class="tab"
      @click="selected = 'pictures'"
      :checked="selected == 'pictures'"
      aria-label="Fotos"
    />
    <div class="tab-content bg-base-100 p-10">
      <Carousel />
    </div>

    <input
      type="radio"
      name="my_tabs_2"
      class="tab"
      aria-label="Reviews"
      :checked="selected == 'reviews'"
      @click="selected = 'reviews'"
    />
    <div class="tab-content border-base-300 bg-base-100 p-10">
      <Review
        v-for="review in reviews"
        :key="review.id"
        :identifier="review.id"
      />
      <CatalogReviewForm />
    </div>

    <input
      type="radio"
      name="my_tabs_2"
      class="tab"
      @click="selected = 'similiar'"
      aria-label="Ähnliche Produkte"
      :checked="selected == 'similiar'"
    />
    <div class="tab-content border-base-300 bg-base-100 p-10">
      Tab content 3
    </div>

    <input
      type="radio"
      name="my_tabs_2"
      class="tab"
      @click="selected = 'equal'"
      aria-label="Dazu passende Produkte"
      :checked="selected == 'equal'"
    />
    <div class="tab-content border-base-300 bg-base-100 p-10">
      Tab content 3
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useRoute } from "vue-router";
import Review from "../catalog/Review.vue";
import { usePocketBase } from "~/util/pocketbase";
import { useLocalStorage } from "@vueuse/core";
import Carousel from "./tab/Carousel.vue";

const pb = usePocketBase();
const route = useRoute();

const props = defineProps({
  identifier: { type: String, required: true },
});

const reviews = ref([]);

const selected = ref("description");

onMounted(async () => {
  reviews.value = (
    await pb.collection("reviews").getList(1, 9, {
      filter: 'product="' + route.params.slug.replace(".html", "") + '"',
    })
  ).items;
});
</script>
