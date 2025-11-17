<template>
  <li class="">
    <div>
      <div>
        <select v-model="currency" class="select select-sm">
          <option value="euro">{{ $t('checkout.currency.euro') }}</option>
          <option value="dollar">{{ $t('checkout.currency.dollar') }}</option>
        </select>
      </div>
    </div>
  </li>
</template>

<script lang="ts" setup>
import findSettingByKey from "~/utils/settings";
const currency = ref(findSettingByKey("currency"));
</script>