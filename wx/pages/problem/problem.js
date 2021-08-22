// pages/problem.js
const app = getApp()

Page({
  data: {
    markmode: 0, // choose markdown plugin: 0.wemark; 1.towxml 
    article: {},
  },
  onLoad: function (options) {
    this.setData({markmode: options.idx % 2})
    let p = app.globalData.problems[options.idx]
    let content = "## " + p.name + "\n[OJ链接](" + p.link + ")\n\n" + p.content + "\n\n### 解答\n" + p.result
    if (this.data.markmode != 0) {
      content = app.towxml(content, 'markdown', {})
    }
    this.setData({article: content})
  }
})