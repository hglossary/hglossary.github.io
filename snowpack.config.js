/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
    'src': {
      url: '/',
    },
    'static': {
      url: '/',
      static: true,
    },
  },
  plugins: [
    '@snowpack/plugin-sass',
    ['@snowpack/plugin-svelte', {}],
  ],
  packageOptions: {
    source: 'remote',
    types: true,
    knownEntrypoints: [],
  },
  devOptions: {},
  buildOptions: {
    baseUrl: "/hindex/",
    out: "docs",
  },
  optimize: {
    bundle: true,
    minify: true,
    target: 'es2018',
  },
};
