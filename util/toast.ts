import { useLocalStorage } from "@vueuse/core"
import { v7 as uuidv7 } from 'uuid';

export const toasts = useLocalStorage('toasts', [], {});

export const add = (message,status) => {
  let id = uuidv7();
  toasts.value.push({
    id,
    message,
    status
  });
  setTimeout(() => {
    toasts.value = toasts.value.filter((item) => item.id != id)
  }, 5000)
}