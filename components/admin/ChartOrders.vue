<template>
  <LineChart :chartData="testData" />
</template>

<script setup lang="ts">
import { LineChart } from "vue-chart-3";
import { Chart, registerables } from "chart.js";
import { usePocketBase } from "~/utils/pocketbase";

Chart.register(...registerables);

const orders = ref([]);
const days = ref([]);
const groups = ref({});
const pb = usePocketBase();

function getWeekBounds(date = new Date()) {
    // Copy date so we don't mutate the original
    const d = new Date(date);
    const day = d.getDay(); // 0 (Sun) to 6 (Sat)
    // Calculate distance to Monday (1)
    const diffToMonday = (day === 0 ? -6 : 1) - day;
    const weekStart = new Date(d);
    weekStart.setDate(d.getDate() + diffToMonday);
    weekStart.setHours(0, 0, 0, 0);

    const weekEnd = new Date(weekStart);
    weekEnd.setDate(weekStart.getDate() + 6);
    weekEnd.setHours(23, 59, 59, 999);

    return { weekStart, weekEnd };
}

const load = async () => {
    const { weekStart, weekEnd } = getWeekBounds();
    pb.autoCancellation(false);
    orders.value = (
        await pb.collection("orders").getList(1, 100, {
            sort: "created",
            filter: `created >= "${weekStart.toLocaleDateString(
                "de",
            )} 00:00:00" && created <= "${weekEnd.toLocaleDateString(
                "de",
            )} 23:59:59"`,
        })
    ).items;

    orders.value = orders.value.map((item) => {
        if (!groups.value[new Date(item.created).toLocaleDateString("de")]) {
            groups.value[new Date(item.created).toLocaleDateString("de")] = [];
        }
        groups.value[new Date(item.created).toLocaleDateString("de")].push(
            item,
        );
    });

    days.value = [0, 0, 0, 0, 0, 0, 0];
    for (let i = 0; i < Object.entries(groups.value).length; i++) {
        days.value[i] = Object.values(groups.value)[i].length;
    }
};

const generateLabels = () => {
    const days = [];
    const today = new Date();
    const day = today.getDay(); // Sunday - 0, Monday - 1, ..., Saturday - 6
    const diffToMonday = (day === 0 ? -6 : 1) - day;
    const monday = new Date(today);
    const locale = "de";
    monday.setDate(today.getDate() + diffToMonday);

    for (let i = 0; i < 7; i++) {
        const d = new Date(monday);
        d.setDate(monday.getDate() + i);
        const dayName = d.toLocaleDateString(locale, {
            weekday: "long",
        });
        const localDate = d.toLocaleDateString(locale, {
            year: "numeric",
            month: "2-digit",
            day: "2-digit",
        });
        days.push(`${dayName} ${localDate}`);
    }
    return days;
};

const testData = computed(() => {
    return {
        labels: generateLabels(),
        datasets: [
            {
                label: "Orders",
                data: days.value,
                fill: true,
                borderColor: "red",
                backgroundColor: "#1EADEF",
                tension: 0.5,
            },
        ],
    };
});

onMounted(() => {
    load();
});
</script>