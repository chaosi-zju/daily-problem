// pages/problem.js
const app = getApp()
var WxParse = require('../wxParse/wxParse.js');

Page({
  data: {
    article: {}
  },
  onLoad: function (options) {
    let p = app.globalData.problems[options.idx]
    let content = "## " + p.name + "<br><br>[OJ链接](" + p.link + ")<br><br> " + p.content + "<br><br> ### 解答<br>" + p.result
    WxParse.wxParse('article', 'md', content, this, 5);
  }
})