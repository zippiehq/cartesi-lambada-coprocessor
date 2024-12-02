module.exports = function override(config) {
    config.resolve.fallback = {
      ...config.resolve.fallback,
      "crypto": require.resolve("crypto-browserify"),
    };
    return config;
  };
