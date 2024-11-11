// src/composables/useStripe.js
import { ref } from 'vue';
import { loadStripe } from '@stripe/stripe-js';

// Replace with your Stripe publishable key
const stripePromise = loadStripe('pk_test_51QJlDLR5d3xw1mbRyaq24na2qlIiPpPjYfnq9m6TL2PSmhjvt2DFh0DHK6aUGaw5M4Mp81wlCOhpqW2jcvZfi3YQ000eWIgGiJ');

export function useStripe() {
  const error = ref(null);
  const loading = ref(false);

  const redirectToCheckout = async (priceId) => {
    loading.value = true;
    error.value = null;

    try {
      const stripe = await stripePromise;
      return await stripe.redirectToCheckout({
        lineItems: [{ price: priceId, quantity: 1 }],
        mode: 'payment',
        successUrl: window.location.origin + '/checkout/?stripe={CHECKOUT_SESSION_ID}',  // Replace with your success URL
        cancelUrl: window.location.origin + '/checkout',   // Replace with your cancel URL
      });
    } catch (err) {
      error.value = err;
    } finally {
      loading.value = false;
    }
  };

  return {
    error,
    loading,
    redirectToCheckout,
  };
}
