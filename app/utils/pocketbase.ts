import PocketBase from "pocketbase";
import config from "../pocketstore.json";

const url = "https://" + config.domains.pocketbase;
const pb = new PocketBase(url);

export const usePocketBaseUrl = () => {
    return url;
};

export const usePocketBase = () => {
    return pb;
};
