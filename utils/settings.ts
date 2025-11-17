import data from "~/pocketstore.json";

export default function findSettingByKey(key) {
    // Check if the settings object exists and contains the key
    if (data.settings && data.settings[key]) {
        return data.settings[key];
    } else {
        console.log(`Key "${key}" not found in settings.`);
        return undefined;
    }
}
