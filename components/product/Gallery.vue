<template>
  <section class="grid grid-cols-6 gap-3 mt-3">
    <div class="col-span-6">
      <div class="divider divider-primary">Produkt Bilder</div>
    </div>
    <div v-for="pic in pictures" :key="pic.id" class="col-span-2">
      <img
        class="aspect-square"
        :src="
          url +
          '/api/files/' +
          pic.collectionId +
          '/' +
          pic.id +
          '/' +
          pic.media
        "
        alt=""
      />
    </div>
  </section>
</template>

<script lang="ts" setup>
import { usePocketBase, usePocketBaseUrl } from "~/util/pocketbase";

const props = defineProps({
    identifier: {
        type: String,
        required: true,
    },
});
const pb = usePocketBase();
const url = usePocketBaseUrl();
const pictures = ref([]);

watch(
    () => props.identifier,
    () => {
        load();
    },
);

const load = async () => {
    pictures.value = await pb.collection("product_pictures").getFullList(9, {
        filter: 'product.id="' + props.identifier + '"',
    });
};
onMounted(async () => {
    load();
});
</script>