<template>
  <section class="page">
    <ProductSeo v-if="item" :product="item" />
    <section class="grid grid-cols-8 gap-3 px-3 py-3 mx-auto max-w-6xl">
      <div class="col-span-8 md:col-span-3">
        <img class="aspect-square"
          src="https://ld-wp73.template-help.com/woocommerce/prod_26798/v1/wp-content/uploads/2019/11/Storytime-Snowman-Singing-Stuffed-Animal-With-Motion-2.jpg"
          alt="">
        <section class="grid grid-cols-6 gap-3 mt-3">
          <div class="col-span-2">
            <img class="aspect-square"
              src="https://ld-wp73.template-help.com/woocommerce/prod_26798/v1/wp-content/uploads/2019/11/Storytime-Snowman-Singing-Stuffed-Animal-With-Motion-2.jpg"
              alt="">
          </div>
          <div class="col-span-2">
            <img class="aspect-square"
              src="https://ld-wp73.template-help.com/woocommerce/prod_26798/v1/wp-content/uploads/2019/11/Storytime-Snowman-Singing-Stuffed-Animal-With-Motion-2.jpg"
              alt="">
          </div>
          <div class="col-span-2">
            <img class="aspect-square"
              src="https://ld-wp73.template-help.com/woocommerce/prod_26798/v1/wp-content/uploads/2019/11/Storytime-Snowman-Singing-Stuffed-Animal-With-Motion-2.jpg"
              alt="">
          </div>
        </section>
      </div>
      <div class="col-span-8 md:col-span-5">
        <h2 class="font-bold text-2xl"> {{ item.name }} </h2>
        <p class="text-sm">{{ item.description }}</p>

        <h3 v-if="item.price" class="font-bold text-4xl">
          <div class="glass px-3 py-3 mt-6 mb-6 md:flex md:justify-between ">
            <section class="text-center space-x-3">
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
              <button class="btn btn-warning hover:text-white">
                <Fa :icon="faStar" />
              </button>
            </section>
            <div class="mt-6 md:mt-2 text-right md:text-left">{{ item.price.toFixed(2) }} €</div>
          </div>
        </h3>

        <div class="form-control grid grid-cols-6 gap-3 text-right">
          <div class="col-span-6">
            <ProductCustomFields :fields="item.config?.fields" />
          </div>
          <div class="col-span-6">
            <section class="grid grid-cols-6 gap-3">
              <div class="col-span-6 md:col-span-1">
                <input type="number" min="1" placeholder="123"
                  class="input placeholder-black w-full bg-white rounded-lg text-center">
              </div>
              <div class="col-span-6 md:col-span-5">
                <button class="btn btn-secondary w-full" @click="addToCart(item.id)">
                  {{ $t('checkout.add-to-cart') }}
                </button>
              </div>
            </section>
          </div>
        </div>
      </div>
      <div class="col-span-8 mt-12">
        <div role="tablist" class="tabs tabs-bordered grid grid-cols-8">
          <input type="radio" name="my_tabs_1" role="tab" class="tab col-span-8 md:col-span-2"
            aria-label="Tab - Beschreibung">
          <div role="tabpanel" class="tab-content p-10">
            <section>
              <div
                class="gap-16 items-center py-8 px-4 mx-auto max-w-screen-xl lg:grid lg:grid-cols-2 lg:py-16 lg:px-6">
                <div class="font-light text-black">
                  <h2 class="mb-4 text-4xl tracking-tight font-extrabold text-white">We didn't reinvent the wheel</h2>
                  <p class="mb-4">We are strategists, designers and developers. Innovators and problem solvers. Small
                    enough to
                    be simple and quick, but big enough to deliver the scope you want at the pace you need. Small enough
                    to be
                    simple and quick, but big enough to deliver the scope you want at the pace you need.</p>
                  <p>We are strategists, designers and developers. Innovators and problem solvers. Small enough to be
                    simple and
                    quick.</p>
                </div>
                <div class="grid grid-cols-2 gap-4 mt-8">
                  <img class="w-full rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-2.png"
                    alt="office content 1">
                  <img class="mt-4 w-full lg:mt-10 rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-1.png"
                    alt="office content 2">
                </div>
              </div>
            </section>
            <section>
              <div
                class="gap-16 items-center py-8 px-4 mx-auto max-w-screen-xl lg:grid lg:grid-cols-2 lg:py-16 lg:px-6">
                <div class="grid grid-cols-2 gap-4 mt-8">
                  <img class="w-full rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-2.png"
                    alt="office content 1">
                  <img class="mt-4 w-full lg:mt-10 rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-1.png"
                    alt="office content 2">
                </div>
                <div class="font-light text-black">
                  <h2 class="mb-4 text-4xl tracking-tight font-extrabold text-white">We didn't reinvent the wheel</h2>
                  <p class="mb-4">We are strategists, designers and developers. Innovators and problem solvers. Small
                    enough to
                    be simple and quick, but big enough to deliver the scope you want at the pace you need. Small enough
                    to be
                    simple and quick, but big enough to deliver the scope you want at the pace you need.</p>
                  <p>We are strategists, designers and developers. Innovators and problem solvers. Small enough to be
                    simple and
                    quick.</p>
                </div>
              </div>
            </section>
            <section>
              <div
                class="gap-16 items-center py-8 px-4 mx-auto max-w-screen-xl lg:grid lg:grid-cols-2 lg:py-16 lg:px-6">
                <div class="font-light text-black">
                  <h2 class="mb-4 text-4xl tracking-tight font-extrabold text-white">We didn't reinvent the wheel</h2>
                  <p class="mb-4">We are strategists, designers and developers. Innovators and problem solvers. Small
                    enough to
                    be simple and quick, but big enough to deliver the scope you want at the pace you need. Small enough
                    to be
                    simple and quick, but big enough to deliver the scope you want at the pace you need.</p>
                  <p>We are strategists, designers and developers. Innovators and problem solvers. Small enough to be
                    simple and
                    quick.</p>
                </div>
                <div class="grid grid-cols-2 gap-4 mt-8">
                  <img class="w-full rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-2.png"
                    alt="office content 1">
                  <img class="mt-4 w-full lg:mt-10 rounded-lg"
                    src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/content/office-long-1.png"
                    alt="office content 2">
                </div>
              </div>
            </section>
          </div>

          <TabReviews />

          <input type="radio" name="my_tabs_1" role="tab" class="tab col-span-8 md:col-span-2" aria-label="Tab Simliar">
          <div role="tabpanel" class="tab-content p-10">Tab Simliar</div>

          <input type="radio" name="my_tabs_1" role="tab" class="tab col-span-8 md:col-span-2"
            aria-label="Tab Additional Product">
          <div role="tabpanel" class="tab-content p-10">Tab Additional Product</div>
        </div>
      </div>
    </section>
  </section>
</template>

<script lang="ts" setup>
import { faStar } from '@fortawesome/free-solid-svg-icons';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';
import { FontAwesomeIcon as Fa } from '@fortawesome/vue-fontawesome';
import { useMessagesStore } from '~/stores/messages';
import { useLocalStorage } from '@vueuse/core';

const storeMessages = useMessagesStore();
const route = useRoute();
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);
const item = ref({});
const cart = useLocalStorage('cart', [], {});

const addToCart = async function (id, qty = 1) {
  let found = false;
  if (typeof cart.value == "undefined") {
    cart.value = [];
  }

  const product = await pb.collection('products').getOne(id);

  cart.value.map((item) => {
    if (item.id == id) {
      found = true;
      storeMessages.add({
        message: 'Update ' + id
      })
    }
  });

  if (found) {
    cart.value.map((item) => {
      if (item.id == id) {
        item.qty += qty
      }
    });
  }
  else {
    storeMessages.add({
      message: 'new ID ' + id
    })
    cart.value.push({
      qty, id, product
    });
  }
}

const load = async () => {
  item.value = (await pb.collection('products').getFirstListItem('slug="' + route.params.slug.replace('.html', '') + '"'))
}

onMounted(() => {
  load();
});
</script>