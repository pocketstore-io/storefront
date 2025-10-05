<template>
  <div v-if="product" class="card shadow-xl bg-white">
    <figure>
      <a :href="'/' + locale + '/product/' + product.slug + '.html'">
        <img
            :src="
            'https://' +
            config.domains.pocketbase +
            '/api/files/' +
            product.collectionId +
            '/' +
            product.id +
            '/' +
            product.cover
          "
            alt="Shoes"
        />
      </a>
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ product.name }}</h2>
      <p class="text-ellipsis line-clamp-2">{{ product.description }}</p>
      <section class="badges space-x-3">
        <a :href="'/search?tag='+slugify(tag.name.toLowerCase())" v-for="tag in tags"
           class="badge badge-secondary mb-3">{{ tag.name }}</a>
      </section>
      <section>
        <button v-if="stock === null || stock.quantity === 0" class="btn btn-error">
          {{ $t('catalog.product.stock-empty') }}
        </button>
        <button v-if="stock !== null && stock.quantity > 0" class="btn btn-success">
          {{ $t('catalog.product.in-stock') }}: {{ stock.quantity }}
        </button>
      </section>
      <div class="card-actions join gap-0 flex justify-end">
        <section>
          <a v-if="product.price" :href="'/' + locale + '/product/' + product.slug + '.html'"
             class="btn btn-primary join-item">
            {{ product.price.toFixed(2) }} €
          </a>
          <CatalogAddToCart :product="product"/>
        </section>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {usePocketBase} from "~/util/pocketbase";
import config from '@/pocketstore.json'
import type {RecordModel} from "pocketbase";
import slugify from "slugify";

const pb = usePocketBase();

const i18n = useI18n();
const locale = i18n.locale;
const stock = ref({});
const tags = ref({});

const props = defineProps({
  identifier: {type: String, required: true},
});

const load = async () => {
  product.value = await pb
      .collection("products")
      .getFirstListItem('slug="' + props.identifier + '"', {
        expand: 'stock,tags',
      });
  stock.value = product.value.expand.stock;
  tags.value = product.value.expand.tags;
}

watch(() => props.identifier, (newVal, oldVal) => {
  load();
});

const product: Ref = ref({});
onMounted(async () => {
  pb.autoCancellation(false)
  load();
});
</script>