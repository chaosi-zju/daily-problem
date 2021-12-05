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
    let userid = app.globalData.userid
    util.request(wx, '/api/problem/common/daily?userid=' + userid, 'GET', {}, function (result) {
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
  onShareAppMessage(){
    return {
      title: '快来一起刷题吧，每天三道算法题，面试不是问题！',
      path: 'pages/index/index'
    }
  },
  onShareTimeline(){
    return {
      title: '快来一起刷题吧，每天三道算法题，面试不是问题！',
    }
  }
})