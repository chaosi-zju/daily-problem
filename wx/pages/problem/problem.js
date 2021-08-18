// pages/problem.js
const app = getApp()

Page({
  data: {
    article: {}
  },
  onLoad: function (options) {
    let p = app.globalData.problems[options.idx]
    let content = "## " + p.name + "\n[OJ链接](" + p.link + ")\n\n" + p.content + "\n\n### 解答\n" + p.result
    this.setData({
      article: content
    })
  }
})