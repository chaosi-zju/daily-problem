// index.js
const app = getApp()
const util = require('../../utils/util.js')

Page({
  data: {
    problems: []
  },
  onLoad() {
    let that = this
    util.request(wx, '/api/problem/common/daily', 'GET', {}, function (result) {
      that.setData({problems: result})
      app.globalData.problems = result
    })
  },
  // 事件处理函数
  viewProblem(e) {
    let idx = e.currentTarget.dataset.index
    wx.navigateTo({
      url: '../problem/problem?idx=' + idx
    })
  },
})