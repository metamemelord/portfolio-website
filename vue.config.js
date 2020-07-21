const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const PreloadWebpackPlugin = require('@vue/preload-webpack-plugin');

module.exports = {
  configureWebpack: {
    plugins: [
      new PreloadWebpackPlugin({
        rel: 'preload',
        as(entry) {
          if (/\.css$/.test(entry)) return 'style';
          if (/\.woff$/.test(entry)) return 'font';
          return 'script';
        }
      })
    ],
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
