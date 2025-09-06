<template>
  <section
    v-if="checkoutStep == 'addresses'"
    class="grid grid-cols-6 gap-3 mx-auto max-w-6xl px-3 py-3"
  >
    <AddressesShipping />
    <AddressesPayment />
    <div class="col-span-6 flex justify-between py-3">
      <button class="btn btn-secondary" @click="checkoutStep = 'customer'">
        {{ $t("checkout.back.customer") }}
      </button>
      <button
        class="btn btn-primary"
        :disabled="(!valid.payment && !same) || !valid.shipping"
        @click="checkoutStep = 'payment'"
      >
        <span>{{ $t("checkout.continue.addresses") }}</span>
        <Fa :icon="faChevronCircleRight" />
      </button>
    </div>
  </section>
</template>

<script lang="ts" setup>
import { faChevronCircleRight } from "@fortawesome/free-solid-svg-icons";
import { useLocalStorage } from "@vueuse/core";
import AddressesPayment from "./addresses/Payment.vue";
import AddressesShipping from "./addresses/Shipping.vue";
const { t } = useI18n();

const checkoutStep = useLocalStorage("checkoutStep", "cart", {});
const valid = useLocalStorage(
  "checkout-valid",
  {
    cart: false,
    addresses: false,
    payment: false,
    shipping: false,
    confirm: false,
    customer: false,
  },
  {}
);
const same = useLocalStorage("same", true, {});

onMounted(() => {
  useHead({
    title: "Checkout Adressen - " + t("general.title"),
  });
});
</script>