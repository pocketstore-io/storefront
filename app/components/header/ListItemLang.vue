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
import config from '~/pocketstore.json'
const { setLocale } = useI18n();
const lang = ref(config.language.fallback);

watch(lang, (value) => {
    setLocale(value);
});
</script>