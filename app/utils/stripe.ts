// src/composables/useStripe.js
import { ref } from "vue";
import { loadStripe } from "@stripe/stripe-js";
import findSettingByKey from "./settings";

// Replace with your Stripe publishable key
const stripePromise = loadStripe(findSettingByKey("stripe")["client-id"]);

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
                mode: "payment",
                successUrl:
                    window.location.origin +
                    "/checkout/?stripe={CHECKOUT_SESSION_ID}", // Replace with your success URL
                cancelUrl: window.location.origin + "/checkout", // Replace with your cancel URL
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
