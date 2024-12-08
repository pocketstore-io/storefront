module.exports = {
	globDirectory: '.output',
	swDest: 'public/sw.js',
	runtimeCaching: [
		{
			urlPattern: ({ request }) => request.destination === 'image',
			handler: 'CacheFirst',
			options: {
				cacheName: 'images',
				expiration: {
					maxEntries: 50,
					maxAgeSeconds: 30 * 24 * 60 * 60, // 30 Tage
				},
			},
		},
	],
};