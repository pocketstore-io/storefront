import PocketBase, {type RecordModel} from "pocketbase";
import config from "../pocketstore.json";

const url = "https://" + config.domains.pocketbase;
const pb = new PocketBase(url);

export const getMediaUrl = (media: RecordModel, field: String) => {
    return url + '/api/files/' + media.collectionId + '/' + media.id + '/' + media[field];
}

export const usePocketBaseUrl = () => {
    return url;
};

export const usePocketBase = () => {
    return pb;
};
