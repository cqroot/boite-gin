module.exports = {
  productionSourceMap: false,
  assetsDir: "static",
  publicPath: process.env.NODE_ENV === "production" ? "/ui/" : "/",
  devServer: {
    proxy: {
      "/api": {
        target: "http://localhost:8088",
      },
    },
  },
};
