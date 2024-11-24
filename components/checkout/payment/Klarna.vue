<template>
  <div v-if="!locked" id="klarna-container" class="w-full mt-3" />
  <button v-else class="btn btn-neutral w-full mt-3">{{ $t('payment.klarna.locked') }}</button>
  <button v-if="locked" class="btn btn-secondary w-full mt-3" @click="resetLock()">
    <Fa :icon="faTrash" /> <span>Reset Klarna Payment</span>
    <Fa :icon="faTrash" />
  </button>
</template>

<script setup lang="ts">
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { useLocalStorage } from '@vueuse/core';
import { useRouter } from 'vue-router';
import { cartToKlarnaPayload } from '~/util/klarna';

const i18n = useI18n();

const { locked } = defineProps({
  locked: {
    type: Boolean,
    required: false,
    default: () => {
      return false;
    }
  }
});

const paymentMethodInfo = useLocalStorage('paymentMethodInfo', {}, {});

const router = useRouter();
useHead({
  script: [
    {
      async: true,
      onload: () => {
        klarnaAsyncCallback();
      },
      onerror: (err) => {
        console.error('Klarna SDK failed to load:', err);
      },
      src: 'https://x.klarnacdn.net/kp/lib/v1/api.js'
    }
  ]
});

const resetLock = function () {
  paymentMethodInfo.value.status = 'unlocked'; klarnaAsyncCallback()
}

const klarnaAsyncCallback = function () {
  window.Klarna.Payments.Buttons.init({
    client_id: findSettingByKey('klarna')['client-id'],
  }).load(
    {
      container: "#klarna-container",
      theme: "default",
      shape: "default",
      on_click: (authorize) => {
        authorize(
          { collect_shipping_address: true },
          cartToKlarnaPayload(),
          (result) => {
            if (result.session_id) {
              router.push('/' + i18n.locale + '/checkout/?klarna=' + result.session_id);
            }
          },
        );
      },
    },
    function load_callback(loadResult) {
      console.log('result', loadResult);
    },
  );
};

</script>
