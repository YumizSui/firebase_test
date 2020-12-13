module.exports = {
  devServer: {
    proxy: {
      "/api": {
        target: "http://localhost:9000",
        changeOrigin: true
      }
    },
    overlay: {
      warnings: true,
      errors: true
    }
  },
  publicPath: "/"
};
