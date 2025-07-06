<template>
  <section
    v-if="checkoutStep == 'customer'"
    class="page bg-red-400 mx-auto max-w-6xl"
  >
    <div class="grid grid-cols-6">
      <div class="col-span-6 md:col-span-2">
        <div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
          <div class="sm:mx-auto sm:w-full sm:max-w-sm">
            <img
              class="mx-auto h-10 w-auto"
              src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=indigo&amp;shade=600"
              alt="Your Company"
            />
            <h2
              class="mt-10 text-center text-2xl/9 font-bold tracking-tight text-gray-900"
            >
              Sign in to your account
            </h2>
          </div>
          <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <form class="space-y-6" @submit.prevent="login">
              <div>
                <label
                  for="email"
                  class="block text-sm/6 font-medium text-gray-900"
                  >Email address</label
                >
                <div class="mt-2">
                  <input
                    v-model="email"
                    type="email"
                    name="email"
                    id="email"
                    class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                  />
                </div>
              </div>
              <div>
                <div class="flex items-center justify-between">
                  <label
                    for="password"
                    class="block text-sm/6 font-medium text-gray-900"
                    >Password</label
                  >
                  <div class="text-sm">
                    <a
                      href="#"
                      class="font-semibold text-indigo-600 hover:text-indigo-500"
                      >Forgot password?</a
                    >
                  </div>
                </div>
                <div class="mt-2">
                  <input
                    type="password"
                    name="password"
                    id="password"
                    v-model="password"
                    class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                  />
                </div>
              </div>
              <div>
                <button
                  type="submit"
                  class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                >
                  Sign in
                </button>
              </div>
            </form>
            <p class="mt-10 text-center text-sm/6 text-gray-500">
              Not a member?
              <a
                href="#"
                class="font-semibold text-indigo-600 hover:text-indigo-500"
                >Start a 14 day free trial</a
              >
            </p>
          </div>
        </div>
      </div>
      <div class="col-span-6 md:col-span-4">
        <div class="hero bg-gray-400 my-12">
          <div class="hero-content text-center">
            <div class="max-w-md">
              <h1 class="text-3xl font-bold">Login - Done already</h1>
              <p class="py-6">
                Provident cupiditate voluptatem et in. Quaerat fugiat ut
                assumenda excepturi exercitationem quasi. In deleniti eaque aut
                repudiandae et a id nisi.
              </p>
              <button class="btn btn-primary">Get Started</button>
            </div>
          </div>
        </div>
        <div class="divider divider-primary"></div>
        <div class="hero bg-gray-400 my-12">
          <div class="hero-content text-center">
            <div class="max-w-md">
              <h1 class="text-3xl font-bold">Register - Forever</h1>
              <p class="py-6">
                Provident cupiditate voluptatem et in. Quaerat fugiat ut
                assumenda excepturi exercitationem quasi. In deleniti eaque aut
                repudiandae et a id nisi.
              </p>
              <button class="btn btn-primary">Get Started</button>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-6 flex justify-between py-3">
        <button class="btn ml-3 btn-secondary" @click="checkoutStep = 'cart'">
          {{ $t("checkout.back.cart") }}
        </button>
        <button
          class="btn mr-3 btn-primary"
          v-if="pb.authStore.isValid || !pb.authStore.isSuperuser"
          @click="checkoutStep = 'addresses'"
        >
          <span>{{ $t("checkout.continue.addresses") }}</span>
          <Fa :icon="faChevronCircleRight" />
        </button>
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { faChevronCircleRight } from "@fortawesome/free-solid-svg-icons";
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";

const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const pb = usePocketBase();

const email = ref("");
const password = ref("");

const login = async () => {
  await pb
    .collection("customers")
    .authWithPassword(email.value, password.value);
  if (pb.authStore.isValid) {
    checkoutStep.value = "addresses";
  }
};

onMounted(() => {
  if (pb.authStore.isValid && checkoutStep.value == "customer") {
    checkoutStep.value = "addresses";
  }
});
</script>
