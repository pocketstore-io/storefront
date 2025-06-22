<template>
  <div>
    <div class="divider divider-primary mt-6">Formular</div>
    <section class="grid grid-cols-12 gap-3">
      <div class="col-span-6 md:col-span-2 py-3">
        <select
          v-model="language"
          class="select select-bordered select-primary w-full"
        >
          <option value="all">Alle</option>
          <option value="de">Deutsch</option>
          <option value="en">English</option>
        </select>
      </div>
      <div class="col-span-6 md:col-span-2 py-3">
        <label class="floating-label">
          <span>Wonach suchst du ?</span>
          <input type="text" v-model="query" class="input input-md" />
        </label>
      </div>
    </section>
    <div class="divider divider-primary mt-6">Wörter</div>
    <section class="grid grid-cols-12 gap-3">
      <div
        v-for="translation in translationsWords"
        :key="translation.id"
        class="col-span-6 md:col-span-3"
      >
        <section class="card bg-white">
          <div class="card-body">
            <p class="text-sm">
              <b>Key: </b>
              <span>{{ translation.key }}</span>
            </p>
            <p class="text-sm">
              <b>Language: </b>
              <span>{{ translation.lang }}</span>
            </p>
            <p class="text-sm">
              <b>Übersetzt: </b>
              <span>{{ translation.translated }}</span>
            </p>
            <button
              @click="
                modalEdit = translation.id;
                open = true;
              "
              class="btn btn-secondary btn-block btn-sm mt-3"
            >
              <FontAwesomeIcon :icon="faEdit" />
              <span>Bearbeiten</span>
            </button>
          </div>
        </section>
      </div>
    </section>
    <div class="divider divider-primary mt-6">Längere Sätze</div>
    <section class="grid grid-cols-12 gap-3 mt-3">
      <div
        v-for="translation in translationsSentence"
        :key="translation.id"
        class="col-span-6 md:col-span-3"
      >
        <section class="card bg-white">
          <div class="card-body">
            <p class="text-sm">
              <b>Key: </b>
              <span>{{ translation.key }}</span>
            </p>
            <p class="text-sm">
              <b>Lang: </b>
              <span>{{ translation.lang }}</span>
            </p>
            <p class="text-sm">
              <b>Übersetzt: </b>
            </p>
            <p class="">{{ translation.translated }}</p>
            <button
              @click="
                modalEdit = translation.id;
                open = true;
              "
              class="btn btn-secondary btn-block btn-sm mt-3"
            >
              <FontAwesomeIcon :icon="faEdit" />
              <span> Bearbeiten </span>
            </button>
          </div>
        </section>
      </div>
    </section>
    <div class="offcanvas-modal-edit">
      <AdminModalEdit />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faEdit, faTrash } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useLocalStorage } from "@vueuse/core";
import { useRouter } from "vue-router";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const router = useRouter();

const translations = ref([]);

const modalEdit = useLocalStorage("modal-edit", "", {});
const open = useLocalStorage("modal-edit-open", false, {});

const language = ref("all");
const query = ref("");

const load = async () => {
  pb.autoCancellation(false);
  translations.value = await pb.collection("translations").getFullList(100);
};

const translationsWords = computed(() => {
  return translations.value
    .filter((item) => item.type == "words")
    .filter(
      (item) =>
        item.translated.toLowerCase().includes(query.value.toLowerCase()) ||
        query.value == "" ||
        item.key.toLowerCase().includes(query.value.toLowerCase()) ||
        query.value == ""
    )
    .filter((item) => item.lang == language.value || language.value == "all")
    .sort(function (a, b) {
      if (a.translated < b.translated) {
        return -1;
      }
      if (a.translated > b.translated) {
        return 1;
      }
      return 0;
    });
});

const translationsSentence = computed(() => {
  return translations.value
    .filter((item) => item.type == "sentence")
    .filter(
      (item) =>
        item.translated.toLowerCase().includes(query.value.toLowerCase()) ||
        query.value == "" ||
        item.key.toLowerCase().includes(query.value.toLowerCase()) ||
        query.value == ""
    )
    .filter((item) => item.lang == language.value || language.value == "all")
    .sort(function (a, b) {
      if (a.translated < b.translated) {
        return -1;
      }
      if (a.translated > b.translated) {
        return 1;
      }
      return 0;
    });
});

onMounted(() => {
  load();

  open.value =false;
  modalEdit.value = '';

  pb.collection("translations").subscribe(
    "*",
    () => {
      load();
    },
    {}
  );
});

definePageMeta({
  layout: "admin",
});
</script>