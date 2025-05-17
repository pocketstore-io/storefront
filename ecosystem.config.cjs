module.exports = {
    apps: [
        {
            name: "NuxtPocketStorefront",
            port: "4000",
            exec_mode: "cluster",
            instances: "4",
            script: ".output/server/index.mjs",
        },
    ],
};
