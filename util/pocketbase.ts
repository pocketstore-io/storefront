import PocketBase from 'pocketbase';
import config from '../pocketstore.json';

const url = 'https://' + config.domain;
const pb = new PocketBase(url);

export const usePocketBaseUrl = () => {
	return url;
};

export const usePocketBase = () => {
	return pb;
};
