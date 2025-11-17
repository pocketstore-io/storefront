<template>
  <div>
    <input
      type="checkbox"
      v-model="modalAdd"
      id="my_modal_7"
      class="modal-toggle"
    />
    <div class="modal" role="dialog">
      <div class="modal-box">
        <h3 class="text-lg font-bold">Hello!</h3>
        <form @submit.prevent="update()" class="space-y-3 mt-3">
          <label class="floating-label">
            <span>Sprache</span>
            <select
              v-model="item.lang"
              class="select select-bordered select-primary w-full"
            >
              <option value="de">Deutsch</option>
              <option value="en">English</option>
            </select>
          </label>
          <label class="floating-label">
            <span>Selector</span>
            <input
              type="text"
              v-model="item.key"
              class="input input-bordered input-primary w-full"
            /> </label
          ><label class="floating-label">
            <span>Übersetzung</span>
            <input
              type="text"
              v-model="item.translated"
              class="input input-bordered input-primary w-full"
            />
          </label>
          <section class="actions flex justify-between">
            <button
              @click="close()"
              class="btn btn-sm btn-warning"
              type="button"
            >
              schließen
            </button>
            <button type="submit" class="btn btn-primary btn-sm">
              Speichern
            </button>
          </section>
        </form>
      </div>
      <label class="modal-backdrop" @click="close()">Close</label>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/utils/pocketbase";

const modalAdd = useLocalStorage("modal-add", false, {});
const item = ref({
    translated: "",
    type: "words",
    key: "",
    lang: "de",
});
const pb = usePocketBase();

const update = async () => {
    await pb.collection("translations").create(item.value);
    modalAdd.value = false;
    item.value = {
        translated: "",
        key: "",
        lang: "",
        type: "words",
    };
};

const close = async () => {
    modalAdd.value = false;
};
</script>
