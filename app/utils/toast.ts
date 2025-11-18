import { ref } from "vue";
import { useLocalStorage } from "@vueuse/core";
import { v7 as uuidv7 } from "uuid";

export type ToastMessage = {
    message: string;
    id: string;
    timer: number;
    status: string;
};

// Correctly type the ref as an array of ToastMessage
export const messages: Ref<ToastMessage[]> = useLocalStorage(
    "messages",
    [],
    {},
);

export const addToast = (message: string, status: string) => {
    const id = uuidv7();
    const unixNowPlus5 = Math.floor(Date.now() / 1000) + 5;
    messages.value.push({ id, message, status, timer: unixNowPlus5 }); // Don't reassign; just push
    setTimeout(() => {
        messages.value = messages.value.filter((item) => item.id !== id);
    }, 5000);
};
