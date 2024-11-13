import settings from '~/configuration/settings.json';
export default {
    getSettingsValue: (key) => {
        return settings.find(item => item.key === key)?.value || ['setting not found'];
    },
    getSettingsAdditional: (key) => {
        return settings.find(item => item.key === key)?.additional || ['setting not found'];
    }
}