module.exports = {
    siteUrl: 'http://localhost:3000',
    generateRobotsTxt: true,
    transform: async (config, path) => {
        // Transform function to modify the sitemap entry
        return {
            loc: path, // The URL of the page
            changefreq: 'daily',
            priority: 0.7,
            lastmod: new Date().toISOString(),
        };
    },
    additionalPaths: async (config) => [
        // Generate paths dynamically
        await getDynamicPathsFromAPI(),
    ],
};
