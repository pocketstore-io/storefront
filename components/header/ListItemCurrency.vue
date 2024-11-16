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
import propertyManager from '~/util/settings'

const currency = ref(propertyManager.getSettingsValue('currency'));
</script>