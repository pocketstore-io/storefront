<template>
  <section class="grid grid-cols-6 gap-3">
    <div class="col-span-6 md:col-span-2">
      <section class="card bg-white">
        <div class="card-body">
          <h2 class="font-bold text-xl">Welcome</h2>
          <ul class="list bg-base-100 rounded-box shadow-md">
            <li class="list-row">
              <div>
                <img
                  class="size-10 rounded-box"
                  src="/public/pocketstore-favicon.svg"
                />
              </div>
              <div>
                <div>Name</div>
                <div class="text-xs uppercase font-semibold opacity-60">
                  {{ $t("general.title") }}
                </div>
              </div>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faInfoCircle" />
              </button>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faUser" />
              </button>
            </li>
            <li class="list-row">
              <div>
                <img
                  class="size-10 rounded-box"
                  src="https://place-hold.it/300?text=EURO&fontsize=80"
                />
              </div>
              <div>
                <div>Währung</div>
                <div class="text-xs uppercase font-semibold opacity-60">
                  <span>
                    {{ currency.code }}
                  </span>
                  -
                  <span>
                    {{ currency.sign }}
                  </span>
                </div>
              </div>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faInfoCircle" />
              </button>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faUser" />
              </button>
            </li>
            <li class="list-row">
              <div>
                <img
                  class="size-10 rounded-box"
                  src="https://place-hold.it/300?text=DE&fontsize=100"
                />
              </div>
              <div>
                <div>Sprachen</div>
                <div class="text-xs uppercase font-semibold opacity-60">
                  {{ languages }}
                </div>
              </div>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faInfoCircle" />
              </button>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faUser" />
              </button>
            </li>
            <li class="list-row">
              <div>
                <img
                  class="size-10 rounded-box"
                  src="https://place-hold.it/300?text=DE&fontsize=100"
                />
              </div>
              <div>
                <div>Kunden</div>
                <div class="text-xs uppercase font-semibold opacity-60">
                  {{ customers.length }}
                </div>
              </div>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faInfoCircle" />
              </button>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faUser" />
              </button>
            </li>
            <li class="list-row">
              <div>
                <img
                  class="size-10 rounded-box"
                  src="https://place-hold.it/300?text=Orders&fontsize=60"
                />
              </div>
              <div>
                <div>Bestellungen</div>
                <div class="text-xs uppercase font-semibold opacity-60">
                  {{ customers.length }}
                </div>
              </div>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faInfoCircle" />
              </button>
              <button class="btn btn-square btn-ghost">
                <FontAwesomeIcon :icon="faUser" />
              </button>
            </li>
          </ul>
        </div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-2">
      <section class="card bg-white">
        <div class="card-body">
          <h2 class="font-bold text-xl">Produkte</h2>
          <AdminProductsTable />
        </div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-2">
      <section class="card bg-white">
        <div class="card-body">
          <h2 class="font-bold text-xl">Kategorien</h2>
          <AdminCategoryTable />
        </div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-3">
      <section class="card bg-white">
        <div class="card-body">
          <AdminChartOrders />
        </div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-3">
      <section class="card bg-white">
        <div class="card-body">
        <AdminChartRevenue />
        </div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-2">
      <section class="card bg-white">
        <div class="card-body">Reviews</div>
      </section>
    </div>
    <div class="col-span-6 md:col-span-2">
      <section class="card bg-white">
        <div class="card-body">Brands</div>
      </section>
    </div>
  </section>
</template>

<script setup lang="ts">
import { faInfoCircle, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { usePocketBase } from "~/util/pocketbase";

const pb = usePocketBase();
const t = useI18n();
const currency = ref({});
const translations = ref([]);
const products = ref([]);
const customers = ref([]);
const languages = ref([]);

const load = async () => {
    currency.value = await pb
        .collection("currencys")
        .getFirstListItem("default=true");
    translations.value = await pb.collection("translations").getFullList(10);
    customers.value = await pb.collection("customers").getFullList(10);
    languages.value = [...new Set(translations.value.map((item) => item.lang))];
};

onMounted(() => {
    load();
});

definePageMeta({
    layout: "admin",
});
</script>

