<template>
  <li class="">
    <div>
      <div>
        <select v-model="lang" class="select select-sm">
          <option value="de">{{ $t('lang.de') }}</option>
          <option value="en">{{ $t('lang.en') }}</option>
        </select>
      </div>
    </div>
  </li>
</template>

<script lang="ts" setup>
import propertyManager from '~/util/settings';

const { setLocale } = useI18n()
const lang = ref(propertyManager.getSettingsValue('language'));

watch(lang, (value) => {
  setLocale(value);
});
</script>