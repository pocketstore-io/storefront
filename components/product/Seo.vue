<script lang="ts" setup>
import { usePocketBase } from "~/util/pocketbase";
import { defineProps } from "vue";

const props = defineProps({
    product: { type: Object, required: true },
});

const pb = usePocketBase();
const loadSeo = async () => {
    if (props.product) {
        const productSeo = await pb
            .collection("product_seo")
            .getFirstListItem('product="' + props.product.id + '"');
        useHead({
            title: props.product.name + " - Produkt Ansicht",
            meta: [
                {
                    name: "description",
                    content: productSeo.desc,
                },
            ],
        });
    }
};

onUpdated(() => {
    loadSeo();
});

onMounted(async () => {
    loadSeo();
});
</script>
