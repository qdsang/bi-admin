'use strict'
const path = require('path')
function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  devServer: {
    disableHostCheck: true,
    open: true,
    overlay: {
        warnings: false,
        errors: true
    },
    proxy: {
      '^/api': {
        target: 'http://localhost:8000',
        changeOrigin: true,// 如果接口跨域，需要进行这个参数配置
        //ws: true, // proxy websockets
        //pathRewrite方法重写url
        pathRewrite: {
          '^/api': '/v2' 
        }
      }
    }
  },
  configureWebpack: {
      //    @路径走src文件夹
      resolve: {
          alias: {
              '@': resolve('src')
          }
      }
  },
  publicPath: '/',
  outputDir: 'docs',
  productionSourceMap: false,
  chainWebpack(config) {
    // set svg-sprite-loader
    config.module
      .rule('svg')
      .exclude.add(resolve('src/icons'))
      .end()
    config.module
      .rule('icons')
      .test(/\.svg$/)
      .include.add(resolve('src/icons'))
      .end()
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
      .end()
    config
      .optimization.splitChunks({
        chunks: 'all',
        cacheGroups: {
          libs: {
            name: 'chunk-libs',
            test: /[\\/]node_modules[\\/]/,
            priority: 10,
            chunks: 'initial' // only package third parties that are initially dependent
          },
          elementUI: {
            name: 'chunk-elementUI', // split elementUI into a single package
            priority: 20, // the weight needs to be larger than libs and app or it will be packaged into libs or app
            test: /[\\/]node_modules[\\/]_?element-ui(.*)/ // in order to adapt to cnpm
          },
          ECharts: {
            name: 'chunk-ECharts',
            priority: 20,
            test: /[\\/]node_modules[\\/]_?echarts(.*)/ // in order to adapt to cnpm
          }
        }
      })
    // config.optimization.runtimeChunk('single')
  }
}
