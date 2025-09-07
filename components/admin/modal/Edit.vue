<template>
  <div>
    <input
      type="checkbox"
      v-model="open"
      id="my_modal_7"
      class="modal-toggle"
    />
    <div class="modal" role="dialog">
      <div class="modal-box">
        <h3 class="text-lg font-bold">Hello!</h3>
        <form @submit.prevent="save()" class="space-y-3 mt-3">
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
              Aktualisieren
            </button>
          </section>
        </form>
      </div>
      <label class="modal-backdrop" @click="close()">Close</label>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { addToast } from "~/util/toast";
import { useLocalStorage } from "@vueuse/core";
import { usePocketBase } from "~/util/pocketbase";

const modalEdit = useLocalStorage("modal-edit", "", {});
const open = useLocalStorage("modal-edit-open", false, {});
const item = ref({});
const pb = usePocketBase();
const toasts = useLocalStorage("toasts", [], {});

const load = async () => {
  item.value = await pb.collection("translations").getOne(modalEdit.value);
};

const close = () => {
  open.value = false;
};

const save = async () => {
  try {
    await pb.collection("translations").update(item.value.id, item.value);
  } catch (e) {
    if (e.status === 400) {
      addToast("Übersetzung schon vorhanden", "warning");
    }
  }

  open.value = false;
  modalEdit.value = "";
};

watch(
  () => modalEdit,
  () => {
    load();
  },
  { deep: true }
);

onMounted(async () => {
  toasts.value = [];
});
</script>
