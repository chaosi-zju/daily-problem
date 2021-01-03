
module.exports = {
    productionSourceMap: false,

    outputDir: '../website',

    devServer: {
        port: 4001,
        proxy: {
            '/api': {
                target: `http://localhost:5001/api`,
                changeOrigin: true,
                pathRewrite: {
                    '^/api': ''
                }
            }
        }
    }
}
