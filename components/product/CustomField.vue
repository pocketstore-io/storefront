<template>
  <div v-if="props.field.type == 'text'">
    <div class="divider-secondary divider">
      {{
        $t(
          "product.label." +
            props.field.label.toLowerCase().replaceAll(" ", "-")
        )
      }}
    </div>
    <label class="floating-label">
      <span> {{ props.field.label }}</span>
      <input type="text" class="input input-secondary w-full" />
    </label>
    <p class="text-sm text-left mt-3">
      {{ props.field.desc }}
    </p>
  </div>
  <div v-if="props.field.type == 'select'">
    <div class="divider divider-secondary">
      {{
        $t(
          "product.label." +
            props.field.label.toLowerCase().replaceAll(" ", "-")
        )
      }}
    </div>
    <label class="floating-label">
      <span> {{ props.field.label }}</span>
      <select id="test" name="test" class="select w-full select-secondary">
        <option v-for="option in props.field.options" :key="option" :value="option">
          {{ option }}
        </option>
      </select></label
    >
    <p class="text-sm text-left mt-3">
      {{ props.field.desc }}
    </p>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
    field: { type: Object, required: true },
});
</script>