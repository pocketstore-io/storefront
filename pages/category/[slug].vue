<template>
  <section class="page">
    <div class="mx-auto max-w-6xl grid grid-cols-6">
      <div class="col-span-6 bg-white">
        <h2 class="font-bold text-lg px-3 py-6">{{ category.name }}</h2>
      </div>
      <div class="col-span-6 bg-white px-3 py-3">
        <section class="grid grid-cols-8 gap-3">
          <div class="col-span-4 md:col-span-2">
            <label class="floating-label">
              <span>Text Filter</span>
              <input
                type="text"
                v-model="query"
                class="w-full input input-bordered input-primary"
                id=""
              />
            </label>
          </div>
          <div class="col-span-4 md:col-span-2">
            <label class="floating-label">
              <span>Marke</span>
              <select
                v-model="brand"
                class="select select-bordered select-primary w-full"
              >
                <option value="all">Alle</option>
                <option v-for="sbrand in brands" :value="sbrand.id">{{ sbrand.name }}</option>
              </select>
            </label>
          </div>
        </section>
      </div>
    </div>
    <section v-if="filteredProducts.length > 0" class="grid grid-cols-6 gap-3">
      <div
        v-for="product in filteredProducts"
        class="col-span-6 md:col-span-2 py-3"
      >
        <CatalogProductCard :identifier="product.slug" />
      </div>
    </section>
    <section v-else>
      <p class="alert alert-warning text-sm font-bold block bg-white px-3 py-3 text-center mx-3 my-3">
        {{ t('category.empty') }}.
      </p>
    </section>
  </section>
</template>

<script setup>
import { useRoute } from "vue-router";
import { usePocketBase } from "~/util/pocketbase";

const {t} = useI18n()
const pb = usePocketBase();
const route = useRoute();

const categories = ref([]);
const category = ref({});
const products = ref([]);
const brands = ref([]);
const tags = ref('');

const menu = ref(true);
const query = ref("");
const brand = ref("all");
const selected = ref("all");

onMounted(async () => {
  pb.autoCancellation(false)
  await load();
  useHead({
    title: category.value.name + " - Kategorie",
    meta: [
      {
        name: "description",
        content: category.value.description,
      },
    ],
  });
});

const load = async () => {
  category.value = await pb
    .collection("categories")
    .getFirstListItem('slug="' + route.params.slug.replace(".html", "") + '"');

  categories.value = await pb.collection("categories").getFullList();
  brands.value = await pb.collection("brands").getFullList();
  tags.value = await pb.collection("product_tags").getFullList();

  products.value = await pb.collection("products").getFullList({
    filter: 'category="' + category.value.id + '"',
  });
};

const filteredProducts = computed(() => {
  return products.value
    .filter((item) =>
      item.name.toLowerCase().includes(query.value.toLowerCase())
    )
    .filter((item) => item.category == selected.value || selected.value == 'all')
    .filter((item)=> item.brand == brand.value || brand.value == 'all');
});

watch(query, load());
</script>
