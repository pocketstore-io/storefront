import { createConfigForNuxt } from "@nuxt/eslint-config/flat";

export default createConfigForNuxt({}).overrideRules({
    "vue/multi-word-component-names": "off",
    "vue/no-multiple-template-root": "off",
    "vue/valid-v-for": "off",
    "vue/require-v-for-key": "off",
    "vue/require-default-prop": "off",
    "vue/first-attribute-linebreak": "off",
});
