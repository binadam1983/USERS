module.exports = {
	globDirectory: 'static/',
	globPatterns: [
		'**/*.{js,json,css,html}'
	],
	swDest: 'static/sw.js',
	ignoreURLParametersMatching: [
		/^utm_/,
		/^fbclid$/
	]
};