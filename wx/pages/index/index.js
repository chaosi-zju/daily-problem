// index.js
const app = getApp()
const util = require('../../utils/util.js')

Page({
  data: {
    title: '',
    problems: []
  },
  onLoad() {
    let that = this
    util.request(wx, '/api/problem/common/daily', 'GET', {}, function (result) {
      that.setData({
        title: ' 提示：访问 chaosi-zju.com 体验更佳',
        problems: result
      })
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