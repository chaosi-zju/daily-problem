// pages/problem.js
const app = getApp()
const util = require('../../utils/util.js')

Page({
  data: {
    userid: 0,
    idx: 0,
    markmode: 0, // choose markdown plugin: 0.wemark; 1.towxml 
    article: {},
    finished: false
  },
  onLoad: function (options) {
    this.setData({
      idx: parseInt(options.idx),
      markmode: options.idx % 2
    })
    let p = app.globalData.problems[options.idx]
    let content = "## " + p.name + "\n[OJ链接](" + p.link + ")\n\n" + p.content + "\n\n### 解答\n" + p.result
    if (this.data.markmode != 0) {
      content = app.towxml(content, 'markdown', {})
    }
    this.setData({article: content})
    this.setData({userid: app.globalData.userid})
  },
  finishProblem() {
    let p = app.globalData.problems[this.data.idx]
    let url = '/api/user/problem/finish?problem_id=' + p.ID + '&userid=' + app.globalData.userid
    util.request(wx, url, 'GET', {}, function (result) {
      this.setData({
        finished: true
      })
    })
  },
  preProblem() {
    let len = app.globalData.problems.length
    let idx = (this.data.idx + len - 1) % len
    wx.redirectTo({
      url: '../problem/problem?idx=' + idx
    })
  },
  nextProblem() {
    let idx = (this.data.idx + 1) % app.globalData.problems.length
    wx.redirectTo({
      url: '../problem/problem?idx=' + idx
    })
  },
  onShareAppMessage() {
    return {
      title: '快来一起刷题吧，每天三道算法题，面试不是问题！',
      path: 'pages/index/index'
    }
  },
  onShareTimeline() {
    return {
      title: '快来一起刷题吧，每天三道算法题，面试不是问题！',
    }
  }
})