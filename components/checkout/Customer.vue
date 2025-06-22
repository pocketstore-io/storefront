<template>
  <section
    v-if="checkoutStep == 'customer'"
    class="page bg-red-400 mx-auto max-w-6xl"
  >
    <div class="grid grid-cols-6">
      <div class="col-span-6 hero bg-base-200 py-32">
        <div class="hero-content flex-col lg:flex-row-reverse">
          <CustomerHeroContent />
          <div class="card bg-base-100 w-full max-w-sm shrink-0 shadow-2xl">
            <CheckoutCustomerForm
              v-if="!pb.authStore.isValid || pb.authStore.isSuperuser"
            />
            <section v-else class="card-body">
              <p class="text-sm text-center font-bold block">
                Weiter als Kunde
              </p>
              <button
                class="btn btn-sm btn-secondary"
                @click="checkoutStep = 'addresses'"
              >
                {{ pb.authStore.record?.name }}
              </button>
            </section>
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
import CustomerHeroContent from "./customer/HeroContent.vue";
import { usePocketBase } from "~/util/pocketbase";

const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const pb = usePocketBase();

onMounted(() => {
  if (pb.authStore.isValid && checkoutStep.value == "customer") {
    checkoutStep.value = "addresses";
  }
});
</script>
