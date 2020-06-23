const UglifyJsPlugin = require('uglifyjs-webpack-plugin');

module.exports = {
  configureWebpack: {
    optimization: {
      splitChunks: {
        chunks: 'async',
        maxSize: 300000,
        minChunks: 1,
        maxAsyncRequests: 6,
        maxInitialRequests: 4,
        automaticNameDelimiter: '~',
        cacheGroups: {
          defaultVendors: {
            test: /[\\/]node_modules[\\/]/,
            priority: -10
          },
          default: {
            minChunks: 2,
            priority: -20,
            reuseExistingChunk: true
          }
        }
      },
      minimizer: [
        new UglifyJsPlugin({
          parallel: 4,
        })
      ],
    },
  }
};