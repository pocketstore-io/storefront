<template>
  <div>
    <button v-if="!locked" id="custom-stripe-button" @click="startCheckout" :disabled="loading">
      <Fa :icon="faStripeS" class="mr-2" />
      <span class="mt-2">{{ loading ? 'Redirecting...' : 'Buy Now' }}</span>
      <Fa :icon="faStripeS" class="ml-2" />
    </button>
    <button v-else class="btn btn-neutral w-full">{{$t('payment.stripe.locked')}}</button>
    <p v-if="error" class="error">{{ error.message }}</p>
  </div>
</template>

<script setup>
import { faStripeS } from '@fortawesome/free-brands-svg-icons';
import { useStripe } from '~/util/stripe';

const { locked } = defineProps({
  locked: {
    type: Boolean,
    required: false,
    default: ()=>{
      return false;
    }
  }
});

const { redirectToCheckout, error, loading } = useStripe();

// Replace with your actual price ID from Stripe's dashboard
const priceId = 'price_1QK0wTR5d3xw1mbRZXQKTEoM';

function startCheckout() {
  redirectToCheckout(priceId);
}
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