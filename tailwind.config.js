import config from './pocketstore.json';

module.exports = {
    plugins: [require('daisyui')],
    daisyui: config.daisyui,
    content: [
        "./components/**/*.{js,vue,ts}",
        "./layouts/**/*.vue",
        "./pages/**/*.vue",
        "./plugins/**/*.{js,ts}",
        "./app.vue",
        "./error.vue",
    ],
};
