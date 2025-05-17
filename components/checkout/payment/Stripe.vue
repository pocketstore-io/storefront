<template>
  <div>
    <button v-if="!locked" id="custom-stripe-button" :disabled="loading" @click="startCheckout">
      <Fa :icon="faStripeS" class="mr-2" />
      <span class="mt-2">{{ loading ? 'Weiterleitung...' : 'Jetzt kaufen' }}</span>
      <Fa :icon="faStripeS" class="ml-2" />
    </button>
    <button v-else class="btn btn-neutral w-full">{{ $t('payment.stripe.locked') }}</button>
    <p v-if="error" class="error">{{ error.message }}</p>
    <button v-if="locked" class="btn btn-secondary w-full mt-3" @click="resetLock()">
      <Fa :icon="faTrash" /> <span>Reset Stripe Payment</span>
      <Fa :icon="faTrash" />
    </button>
  </div>
</template>

<script setup>
// TODO get config by option
import { faStripeS } from "@fortawesome/free-brands-svg-icons";
import { useLocalStorage } from "@vueuse/core";
import { useStripe } from "~/util/stripe";
import { faTrash } from "@fortawesome/free-solid-svg-icons";

const { locked } = defineProps({
    locked: {
        type: Boolean,
        required: false,
        default: () => {
            return false;
        },
    },
});

const paymentMethodInfo = useLocalStorage("paymentMethodInfo", {}, {});

const { redirectToCheckout, error, loading } = useStripe();

// Replace with your actual price ID from Stripe's dashboard
const priceId = "price_1QK0wTR5d3xw1mbRZXQKTEoM";

function startCheckout() {
    redirectToCheckout(priceId);
}

const resetLock = function () {
    paymentMethodInfo.value.status = "unlocked";
};
</script>

<style>
#custom-stripe-button {
  background-color: #ff69b4;
  /* Pink background */
  color: white;
  /* Text color */
  padding: 12px 24px;
  width: 100%;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

#custom-stripe-button:hover {
  background-color: #ff85c0;
  /* Lighter pink on hover */
}
</style>