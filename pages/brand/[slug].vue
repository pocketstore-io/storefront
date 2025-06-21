<template>
  <section class="page bg-white px-3 py-3">
    <div class="grid grid-cols-6 gap-3">
      <div class="col-span-6 md:col-span-2">
        <img
          :src="
            url +
            '/api/files/' +
            item.collectionId +
            '/' +
            item.id +
            '/' +
            item.logo
          "
          alt=""
        />
      </div>
      <div class="col-span-6 md:col-span-4">
        <p class="text-sm" v-html="item.desc"></p>
      </div>
      <CatalogProductGrid identifier="bowling" />
    </div>
  </section>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";
import { usePocketBase, usePocketBaseUrl } from "~/util/pocketbase";

const pb = usePocketBase();
const url = usePocketBaseUrl();
const route = useRoute();
const item = ref({});
const seo = ref({});
const { t } = useI18n();

onMounted(async () => {
    item.value = await pb
        .collection("brands")
        .getFirstListItem('slug="' + route.params.slug + '"');

    useHead({
        title: "Brand: " + item.value.name + " - " + t("general.title"),
        meta: [
            {
                name: "description",
                content:
                    "Dies ist die Beschreibung für die Brand deiner Wahl. Aktuelle Brand " +
                    item.value.name,
            },
        ],
    });
});
</script>