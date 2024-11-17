<script setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useBreadcrumbStore } from '~/stores/breadcrumb';

const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const category = ref({});
const products = ref([]);

onMounted(async () => {
    category.value = await pb.collection('categories').getFirstListItem('name="Welcome"', {
        expand: 'products',
        sort: 'products.created'
    });
    products.value = category.value.expand.products;
});

const storeBreadcrumb = useBreadcrumbStore();

storeBreadcrumb.clear();

</script>

<template>
    <section class="page">
        <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
            <div class="col-span-6">
                <div class="hero bg-base-200">
                    <div class="hero-content flex-col lg:flex-row">
                        <img src="https://img.daisyui.com/images/stock/photo-1635805737707-575885ab0820.webp"
                            class="max-w-sm rounded-lg shadow-2xl">
                        <div>
                            <h1 class="text-5xl font-bold">Box Office News!</h1>
                            <p data-testid="p-hero" class="py-6">
                                Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi
                                exercitationem quasi. In deleniti eaque aut repudiandae et a id nisi.
                            </p>
                            <button class="btn btn-primary">Get Started</button>
                        </div>
                    </div>
                </div>
            </div>
            <div v-for="product in products" class="col-span-6 md:col-span-2 py-3">
                <CatalogProductCard :identifier="product.id" />
            </div>
        </div>
    </section>
</template>