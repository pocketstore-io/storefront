<template>
  <div v-if="product" class="card shadow-xl bg-white">
    <figure>
      <a :href="'/' + locale + '/product/' + product.slug + '.html'">
        <img
          :src="
            'https://' +
            config.domain +
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
      <section>
        <button v-if="stock === null || stock.quantity === 0" class="btn btn-error">
          {{$t('catalog.product.stock-empty')}}
        </button>
        <button v-if="stock !== null && stock.quantity > 0" class="btn btn-success">
          {{$t('catalog.product.stock')}}: {{ stock.quantity }}
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

const pb = usePocketBase();

const i18n = useI18n();
const locale = i18n.locale;
const stock = ref({});

const {identifier} = defineProps({
  identifier: {type: String, requiered: true},
});

const product = ref({});
onMounted(async () => {
  pb.autoCancellation(false)
  product.value = await pb
      .collection("products")
      .getFirstListItem('slug="' + identifier + '"');
  stock.value = (await pb
      .collection("product_stocks")
      .getList(1, 1, {
        "filter": 'product="' + identifier + '"'
      })).items[0] ?? null;
});
</script>