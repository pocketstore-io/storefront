import { v4 as uuidv4 } from "uuid";

export const breadcrumbs: Ref = ref([]);

type Breadcrumb = {
    id: string;
    label: string;
    icon: string | null; // You may use a specific icon type if you use a UI library
    link: string;
};

export const addBreadcrumb = (item: Breadcrumb) => {
    breadcrumbs.value.push(item);
};
