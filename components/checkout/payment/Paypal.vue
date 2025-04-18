<template>
  <div>
    <div id="paypal-button-container" class="mt-3" />
    <button v-if="locked" class="btn btn-secondary w-full mt-3" @click="resetLock()">
      <Fa :icon="faTrash" /> <span>Reset PayPal Payment</span>
      <Fa :icon="faTrash" />
    </button>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from '@vueuse/core';
import { loadScript } from "@paypal/paypal-js";
import { cartToPurchaseUnits } from '~/util/paypal';
import findSettingByKey from '~/util/settings';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

const { locked } = defineProps({
  locked: {
    type: Boolean,
    required: false,
    default: () => {
      return false;
    }
  }
});

const paymentMethod = useLocalStorage('paymentMethod', 'vorkasse', {});
const checkoutStep = useLocalStorage('checkoutStep', 'cart', {});

const resetLock = function () {

}

const initPaypal = () => {
  loadScript({ "client-id": findSettingByKey('paypal')['client-id'], currency: findSettingByKey('paypal')['currency'] })
    .then((paypal) => {
      paypal.Buttons({
        createOrder: function (data, actions) {
          // Directly create an order with the amount, no server needed
          return actions.order.create({
            purchase_units: cartToPurchaseUnits()
          });
        },
        onApprove: function (data, actions) {
          return actions.order.capture().then(function (details) {
            alert('Transaction completed by ' + details.payer.name.given_name);
            console.log(details); // You can log or handle the transaction details here
          });
          // TODO store details for order
        },
        onError: function (err) {
          console.error('An error occurred during the transaction', err);
        }
      }).render('#paypal-button-container');
    })
    .catch((err) => {
      console.error("failed to load the PayPal JS SDK script", err);
    });
}

onMounted(() => {
  if (paymentMethod.value == 'paypal' && checkoutStep.value == 'payment') {
    initPaypal();
  }
});

watch(paymentMethod, (value) => {
  if (value === 'paypal') {
    initPaypal();
  }
});

watch(checkoutStep, (value) => {
  if (paymentMethod.value == 'paypal' && value == 'payment') {
    initPaypal();
  }
});
</script>