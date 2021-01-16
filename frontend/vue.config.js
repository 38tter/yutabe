module.exports = {
  devSever: {
    proxy: {
      "/api": {
        target: "http://localhost:8081",
      },
    },
  },
}
