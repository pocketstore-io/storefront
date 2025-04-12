import PocketBase from 'pocketbase';
import config from '../pocketstore.json';

const pb = new PocketBase('https://' + config.domain);

export const usePocketBase = () => {
	return pb;
};
