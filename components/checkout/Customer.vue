<template>
  <section v-if="checkoutStep == 'customer'" class="page bg-red-400 mx-auto max-w-6xl">
    <div class="grid grid-cols-6">
      <div class="col-span-6 hero bg-base-200 py-32">
        <div class="hero-content flex-col lg:flex-row-reverse">
          <div class="text-center lg:text-left" v-if="!pb.authStore.isValid">
            <h1 class="text-5xl font-bold">Login now!</h1>
            <p class="py-6">
              Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
              quasi. In deleniti eaque aut repudiandae et a id nisi.
            </p>
          </div>
          <div class="text-center lg:text-left" v-else>
            <h1 class="text-2xl font-bold">Deine Vorteile als Kunde</h1>
            <p class="py-6">
              Provident cupiditate voluptatem et in. Quaerat fugiat ut assumenda excepturi exercitationem
              quasi. In deleniti eaque aut repudiandae et a id nisi.
            </p>
          </div>
          <div class="card bg-base-100 w-full max-w-sm shrink-0 shadow-2xl">
            <form class="card-body" v-if="!pb.authStore.isValid">
              <div class="form-control">
                <label class="label">
                  <span class="label-text">Email</span>
                </label>
                <input type="email" v-model="email" placeholder="email" class="input input-bordered" required />
              </div>
              <div class="form-control">
                <label class="label">
                  <span class="label-text">Password</span>
                </label>
                <input type="password" v-model="password" placeholder="password" class="input input-bordered"
                  required />
                <label class="label">
                  <a href="#" class="label-text-alt link link-hover">Forgot password?</a>
                </label>
              </div>
              <div class="form-control mt-6">
                <button class="btn btn-primary" @click.prevent="login()">Login</button>
              </div>
            </form>
            <section v-else class="card-body">
              <p class="text-sm text-center font-bold block">
                Weiter als Kunde
              </p>
              <button @click="checkoutStep='addresses'" class="btn btn-sm btn-secondary">{{pb.authStore.model?.name}}</button>
            </section>
          </div>
        </div>
      </div>
      <div class="col-span-6 flex justify-between py-3">
        <button class="btn btn-primary" @click="checkoutStep = 'cart'">Back to Cart</button>
        <button class="btn btn-primary" :disabled="!pb.authStore.isValid"
          @click="checkoutStep = 'addresses'"><span>Continue</span>
          <Fa :icon="faChevronCircleRight" />
        </button>
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { faChevronCircleRight } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';
import PocketBase from 'pocketbase';
import { usePocketbaseStore } from '~/stores/pocketbase';

const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});
const valid = useLocalStorage('checkout-valid', { cart: false, addresses: false, payment: false, shipping: false, confirm: false, customer: false }, {});
const store = usePocketbaseStore();
const { url } = storeToRefs(store);
const pb = new PocketBase(url.value);

const email = ref('');
const password = ref('');
const login = async () => {
  await pb.collection('customer').authWithPassword(email.value, password.value);

  if (pb.authStore.isValid) {
    checkoutStep.value = 'addresses';
  }
}

onMounted(() => {
  if (pb.authStore.isValid && checkoutStep.value == 'customer') {
    checkoutStep.value = 'addresses';
  }
});
</script>
