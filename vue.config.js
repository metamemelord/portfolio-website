const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const OptimizeCssAssetsPlugin = require('optimize-css-assets-webpack-plugin');

module.exports = {
  configureWebpack: {
    optimization: {
      splitChunks: {
        chunks: 'async',
        maxSize: 225000,
        minChunks: 1,
        maxAsyncRequests: Infinity,
        maxInitialRequests: Infinity,
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
          uglifyOptions: {
            keep_classnames: false,
            keep_fnames: false,
            output: {
              comments: false
            },
          },
          parallel: 4,
        })
      ],
    },
    plugins: [
      new OptimizeCssAssetsPlugin({
        assetNameRegExp: /\.optimize\.css\.scss$/g,
        cssProcessor: require('cssnano'),
        cssProcessorPluginOptions: {
          preset: ['default', { discardComments: { removeAll: true } }],
        },
        canPrint: true
      }),
    ],
  }
};
