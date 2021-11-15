module.exports = {
    title: 'chaosi · 笔记',
    description: '温故而知新',
    head: [
        ['link', {rel: 'icon', href: 'http://chaosi-zju.com/favicon.ico'}]
    ],
    base: '/',
    themeConfig: {
        nav: [
            {text: '刷新', link: 'http://note.chaosi-zju.com/api/note/update'},
            {text: '每日做题', link: 'http://chaosi-zju.com'},
            {text: 'Github', link: 'http://github.com/chaosi-zju/daily-problem'}
        ],
        docsRepo: 'chaosi-zju/daily-problem',
        docsDir: '/',
        lastUpdated: 'Last Updated',
    },
    plugins: ['permalink-pinyin', ['autobar', {'pinyinNav': true}], 'rpurl'],
    chainWebpack: (config, isServer) => {
        const inlineLimit = 10000
        config.module.rule('images')
            .test(/\.(png|jpe?g|gif|webp)(\?.*)?$/)
            .use('url-loader')
            .loader('url-loader')
            .options({
                limit: inlineLimit,
                name: `assets/img/[name].[hash:8].[ext]`
            })
    }
}
