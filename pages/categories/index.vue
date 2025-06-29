<template>
  <section class="mx-auto max-w-6xl px-3 py-3">
    <div class="glass px-3 py-3 rounded-2xl">
      <section class="grid grid-cols-6 gap-3">
        <div class="col-span-6 md:col-span-2 mx-6 my-6">
          <img
            class="rounded-2xl w-[100vh]"
            src="https://place-hold.it/1000x1000"
            alt=""
          />
        </div>
        <div class="col-span-6 md:col-span-4">
          <h2 class="font-bold text-3xl text-white mt-12">
            Hier findest du alle Kategorien von Geek Wear
          </h2>
          <p class="my-6 text-white">
            Lorem ipsum dolor sit amet consect adipisicing elit. Possimus magnam
            voluptatum cupiditate veritatis in accusamus quisquam voluptatum
            cupiditate veritatis in accusamus quisquam.
          </p>
          <section class="grid grid-cols-6 gap-3 text-white">
            <div
              v-for="category in computedCategories"
              class="col-span-6 md:col-span-3 lg:col-span-2"
            >
              <FontAwesomeIcon :icon="faCheckCircle" />
              <span class="ml-3">{{ category.name }}</span>
            </div>
          </section>
          <section class="cta mt-6">
            <a href="/" class="link link-hover text-secondary">
              <span class="mr-3">See our Search</span>
              <FontAwesomeIcon :icon="faAnglesDown" />
            </a>
          </section>
        </div>
      </section>
    </div>
    <section class="grid grid-cols-6 gap-6 my-12">
      <div class="col-span-6">
        <label class="floating-label">
          <span>Wonach suchst du ?</span>
          <input
            type="text"
            v-model="query"
            class="input input-bordered input-primary w-full"
          />
        </label>
      </div>
      <div v-for="category in filtered" class="col-span-6 md:col-span-2">
        <section class="bg-white rounded-xl px-3 py-3">
          <img
            :src="'https://place-hold.it/380x120?text=' + category.slug"
            alt=""
          />
          <NuxtLink
            :to="localePath('/category/' + category.slug + '.html')"
            class="font-bold text-lg"
            >{{ category.name }}</NuxtLink
          >
          <p class="text-sm overflow-hidden text-ellipsis h-4">
            Hallo Welt Beschreibung direkt aus der Hölle. Hallo Welt Beschreibung
            direkt aus der Hölle
          </p>
        </section>
      </div>
    </section>
  </section>
</template>

<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faCheckCircle, faAnglesDown } from "@fortawesome/free-solid-svg-icons";
import { watch } from "vue";
import { usePocketBase } from "@/util/pocketbase";
const localePath = useLocalePath();
const { t } = useI18n();

const pb = usePocketBase();
const categories = ref([]);
const query = ref("");

const computedCategories = computed(() => {
    return categories.value.slice(0, 9);
});

watch(query, () => {
    load();
});

const load = async () => {
    categories.value = (
        await pb.collection("categories").getList(1, 30)
    ).items.sort(function (a, b) {
        if (a.name < b.name) {
            return -1;
        }
        if (a.name > b.name) {
            return 1;
        }
        return 0;
    });
};

const filtered = computed(() => {
    return categories.value.filter((item) =>
        item.name.toLowerCase().includes(query.value.toLowerCase()),
    );
});

onMounted(async () => {
    load();

    useHead({
        title: "Kategorien | " + t("general.title"),
    });
});
</script>