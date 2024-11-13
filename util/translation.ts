import PocketBase from 'pocketbase';

// Function to convert a key string to a nested object
function setNestedProperty(obj, key, value) {
  const keys = key.split('.');
  keys.reduce((acc, part, index) => {
    if (index === keys.length - 1) {
      acc[part] = value;
    } else {
      acc[part] = acc[part] || {};
    }
    return acc[part];
  }, obj);
}

export default async function fetchTranslations() {
  const pb = new PocketBase('https://admin.pocketstore.io');
  const translations = await pb.collection('translations').getFullList();
  
  const i18nData = translations.reduce((acc, record) => {
    const langCode = record.lang_code;
    const keyValues = record.key_value_pairs; // Assuming key_value_pairs is an object of keys and translations

    acc = acc[langCode] || {};
    
    for (const [key, value] of Object.entries(keyValues)) {
      setNestedProperty(acc[langCode], key, value);
    }
    
    return acc;
  }, {});

  return i18nData;
}