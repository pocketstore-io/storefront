import settings from '~/configuration/settings.json';
export default {
    getSettings: (key) => {
        return settings.find(item => item.key === key)?.additional || ['setting not found'];
    }
}