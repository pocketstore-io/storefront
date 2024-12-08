<script setup>
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { useBreadcrumbStore } from '~/stores/breadcrumb';

const store = usePocketbaseStore();
const storeBreadcrumb = useBreadcrumbStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const category = ref({});
const products = ref([]);

onMounted(async () => {
  category.value = await pb.collection('categories').getFirstListItem('slug="bowling"', {
    expand: 'products',
    sort: '-products.created'
  });
  products.value = category.value.expand.products;
  storeBreadcrumb.clear();
  storeBreadcrumb.add({
    label: 'Category View',
    link: 'category/bowling',
    id: 'category-view'
  });
});

</script>

<template>
  <section class="page">
    <div class="mx-auto max-w-6xl grid grid-cols-6 gap-3">
      <div class="col-span-6">
        <h2 class="font-bold text-lg px-3 py-6">{{ category.name }} </h2>
      </div>
      <div class="col-span-6">
        <div class="drawer">
          <input id="my-drawer" type="checkbox" class="drawer-toggle" >
          <div class="drawer-content">
            <!-- Page content here -->
            <label for="my-drawer" class="btn btn-primary drawer-button">Open drawer</label>
          </div>
          <div class="drawer-side z-10">
            <label for="my-drawer" aria-label="close sidebar" class="drawer-overlay"/>
            <section class="menu bg-base-200 text-base-content min-h-full w-96 max-w-full p-4">
              <section class="grid grid-cols-6">
                <div class="col-span-6 bg-red-400 px-3 py-3 flex justify-between">
                  <span>Filter</span>
                  <label for="my-drawer" class="btn btn-primary drawer-button">Close drawer</label>
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Suche</label>
                  <input type="text" class="input block input-bordered input-primary col-span-6 md:col-span-4">
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Sprache</label>
                  <select class="select block select-bordered select-primary col-span-6 md:col-span-4">
                    <option value="de">German</option>
                  </select>
                </div>
                <div class="col-span-6 mt-3">
                  <label for="" class="label font-bold text-sm">Kategorien:</label>
                  <section class="grid grid-cols-6 gap-3">
                    <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-sm btn-block flex justify-between">
                        <input type="checkbox" class="checkbox checkbox-sm checkbox-primary">
                        <span>Hallo</span>
                        <span class="invisible"/>
                      </button>
                    </div>
                    <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-sm btn-block flex justify-between">
                        <input type="checkbox" class="checkbox checkbox-sm checkbox-primary">
                        <span>Hallo</span>
                        <span class="invisible"/>
                      </button>
                    </div>
                    <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-sm btn-block flex justify-between">
                        <input type="checkbox" class="checkbox checkbox-sm checkbox-primary">
                        <span>Hallo</span>
                        <span class="invisible"/>
                      </button>
                    </div>
                    <div class="col-span-6 md:col-span-3">
                      <button class="btn btn-sm btn-block flex justify-between">
                        <input type="checkbox" class="checkbox checkbox-sm checkbox-primary">
                        <span>Hallo</span>
                        <span class="invisible"/>
                      </button>
                    </div>
                  </section>
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Suche</label>
                  <input type="text" class="input block input-bordered input-primary col-span-6 md:col-span-4">
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Sprache</label>
                  <select class="select block select-bordered select-primary col-span-6 md:col-span-4">
                    <option value="de">German</option>
                  </select>
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Suche</label>
                  <input type="text" class="input block input-bordered input-primary col-span-6 md:col-span-4">
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Sprache</label>
                  <select class="select block select-bordered select-primary col-span-6 md:col-span-4">
                    <option value="de">German</option>
                  </select>
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Suche</label>
                  <input type="text" class="input block input-bordered input-primary col-span-6 md:col-span-4">
                </div>
                <div class="form-control col-span-6 grid grid-cols-6 mt-3">
                  <label for="" class="label col-span-6 md:col-span-2">Sprache</label>
                  <select class="select block select-bordered select-primary col-span-6 md:col-span-4">
                    <option value="de">German</option>
                  </select>
                </div>
              </section>
            </section>
          </div>
        </div>
      </div>
      <div v-for="product in products" class="col-span-6 md:col-span-2 py-3">
        <CatalogProductCard :identifier="product.id" />
      </div>
    </div>
  </section>
</template>